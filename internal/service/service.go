package service

import (
	"crypto/sha512"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type Svc struct {
	db *sql.DB
}

func (s Svc) GetUserByLogin(login string) (*User, error) {
	at := strings.LastIndex(login, "@")
	if at <= 0 || at == len(login)-1 {
		return nil, nil
	}
	localPart := login[:at]
	domainName := login[at+1:]

	const q = `
		SELECT u.user_id, u.domain_id, u.login, u.password, u.admin_ind
		FROM mailbox_user u
		JOIN domain d ON d.domain_id = u.domain_id
		WHERE u.login = ? AND d.name = ?`

	var u User
	var adminInd string
	err := s.db.QueryRow(q, localPart, domainName).
		Scan(&u.UserID, &u.DomainID, &u.Login, &u.Password, &adminInd)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	u.Admin = adminInd == "Y"
	return &u, nil
}

func (s Svc) GetDomains() ([]Domain, error) {
	const q = `
		SELECT domain_id, name, catchall_ind, catchall_login
		FROM domain
		ORDER BY name`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var domains []Domain
	for rows.Next() {
		var d Domain
		var catchallInd string
		if err := rows.Scan(&d.DomainID, &d.Name, &catchallInd, &d.CatchallLogin); err != nil {
			return nil, err
		}
		d.CatchAll = catchallInd == "Y"
		domains = append(domains, d)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return domains, nil
}

func (s Svc) GetDomain(id string) (*Domain, error) {
	const q = `
		SELECT domain_id, name, catchall_ind, catchall_login
		FROM domain
		WHERE domain_id = ?`

	var d Domain
	var catchallInd string
	err := s.db.QueryRow(q, id).Scan(&d.DomainID, &d.Name, &catchallInd, &d.CatchallLogin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	d.CatchAll = catchallInd == "Y"
	return &d, nil
}

func (s Svc) AddDomain(id string, domain string) error {
	domain = strings.ToLower(strings.TrimSpace(domain))
	if domain == "" {
		return errors.New("domain must be filled out")
	}
	if !strings.Contains(domain, ".") {
		return errors.New("domain must contain at least one dot")
	}

	var existing string
	err := s.db.QueryRow(`SELECT domain_id FROM domain WHERE name = ?`, domain).Scan(&existing)
	if err == nil {
		return fmt.Errorf("domain %s already exists", domain)
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	_, err = s.db.Exec(
		`INSERT INTO domain (domain_id, name, catchall_ind, catchall_login) VALUES (?, ?, 'N', '')`,
		id, domain,
	)
	return err
}

func (s Svc) UpdateDomain(id, domain string, catchAll bool, catchAllLogin string) error {
	domain = strings.ToLower(strings.TrimSpace(domain))
	if domain == "" {
		return errors.New("domain must be filled out")
	}
	if !strings.Contains(domain, ".") {
		return errors.New("domain must contain at least one dot")
	}

	var existing string
	err := s.db.QueryRow(
		`SELECT domain_id FROM domain WHERE name = ? AND domain_id <> ?`,
		domain, id,
	).Scan(&existing)
	if err == nil {
		return fmt.Errorf("domain %s already exists", domain)
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	catchAllInd := "N"
	if catchAll {
		if strings.TrimSpace(catchAllLogin) == "" {
			return errors.New("catchall login must be filled out")
		}
		var userID string
		err := s.db.QueryRow(
			`SELECT user_id FROM mailbox_user WHERE login = ? AND domain_id = ?`,
			catchAllLogin, id,
		).Scan(&userID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("catchall login %s does not exist under this domain", catchAllLogin)
			}
			return err
		}
		catchAllInd = "Y"
	} else {
		catchAllLogin = ""
	}

	_, err = s.db.Exec(
		`UPDATE domain SET name = ?, catchall_ind = ?, catchall_login = ? WHERE domain_id = ?`,
		domain, catchAllInd, catchAllLogin, id,
	)
	return err
}

func (s Svc) DeleteDomains(ids []string) error {
	if len(ids) == 0 {
		return nil
	}

	placeholders := strings.Repeat("?,", len(ids))
	placeholders = placeholders[:len(placeholders)-1]
	args := make([]any, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	_, err := s.db.Exec(`DELETE FROM domain WHERE domain_id IN (`+placeholders+`)`, args...)
	if err != nil {
		return fmt.Errorf("failed to delete: there are dependent items (%w)", err)
	}
	return nil
}

func (s Svc) GetUsersByDomain(domain string) ([]User, error) {
	const q = `
		SELECT u.user_id, u.domain_id, u.login, u.password, u.admin_ind
		FROM mailbox_user u
		JOIN domain d ON d.domain_id = u.domain_id
		WHERE d.name = ?
		ORDER BY u.login`

	rows, err := s.db.Query(q, domain)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		var adminInd string
		if err := rows.Scan(&u.UserID, &u.DomainID, &u.Login, &u.Password, &adminInd); err != nil {
			return nil, err
		}
		u.Admin = adminInd == "Y"
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (s Svc) AddUser(id, domainID, login string, admin bool) error {
	login = strings.TrimSpace(login)
	if login == "" {
		return errors.New("login must be filled out")
	}

	var existing string
	err := s.db.QueryRow(
		`SELECT user_id FROM mailbox_user WHERE login = ? AND domain_id = ?`,
		login, domainID,
	).Scan(&existing)
	if err == nil {
		return fmt.Errorf("login %s already exists in this domain", login)
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	adminInd := "N"
	if admin {
		adminInd = "Y"
	}

	_, err = s.db.Exec(
		`INSERT INTO mailbox_user (user_id, domain_id, login, password, admin_ind) VALUES (?, ?, ?, '', ?)`,
		id, domainID, login, adminInd,
	)
	return err
}

func (s Svc) UpdateUser(id, domainID, login string, admin bool) error {
	login = strings.TrimSpace(login)
	if login == "" {
		return errors.New("login must be filled out")
	}

	var existing string
	err := s.db.QueryRow(
		`SELECT user_id FROM mailbox_user WHERE login = ? AND domain_id = ? AND user_id <> ?`,
		login, domainID, id,
	).Scan(&existing)
	if err == nil {
		return fmt.Errorf("login %s already exists in this domain", login)
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	adminInd := "N"
	if admin {
		adminInd = "Y"
	}

	_, err = s.db.Exec(
		`UPDATE mailbox_user SET login = ?, admin_ind = ? WHERE user_id = ?`,
		login, adminInd, id,
	)
	return err
}

func (s Svc) UpdatePassword(id, hashedPassword string) error {
	_, err := s.db.Exec(
		`UPDATE mailbox_user SET password = ? WHERE user_id = ?`,
		hashedPassword, id,
	)
	return err
}

func (s Svc) DeleteUsers(ids []string) error {
	if len(ids) == 0 {
		return nil
	}

	placeholders := strings.Repeat("?,", len(ids))
	placeholders = placeholders[:len(placeholders)-1]
	args := make([]any, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	_, err := s.db.Exec(`DELETE FROM mailbox_user WHERE user_id IN (`+placeholders+`)`, args...)
	if err != nil {
		return fmt.Errorf("failed to delete: there are dependent items (%w)", err)
	}
	return nil
}

func (s Svc) GetMailboxes(userID string) ([]Mailbox, error) {
	const q = `
		SELECT mailbox_id, user_id, name,
			flag_non_existent, flag_no_inferiors, flag_no_select, flag_marked,
			flag_subscribed, flag_remote, flag_archive, flag_drafts,
			flag_flagged, flag_junk, flag_sent, flag_trash, flag_important
		FROM mailbox
		WHERE user_id = ?
		ORDER BY name`

	rows, err := s.db.Query(q, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mailboxes []Mailbox
	for rows.Next() {
		var m Mailbox
		var nonExistent, noInferiors, noSelect, marked, subscribed, remote, archive, drafts, flagged, junk, sent, trash, important string
		if err := rows.Scan(
			&m.MailboxID, &m.UserID, &m.Name,
			&nonExistent, &noInferiors, &noSelect, &marked,
			&subscribed, &remote, &archive, &drafts,
			&flagged, &junk, &sent, &trash, &important,
		); err != nil {
			return nil, err
		}
		m.FlagNonExistent = nonExistent == "Y"
		m.FlagNoInferiors = noInferiors == "Y"
		m.FlagNoSelect = noSelect == "Y"
		m.FlagMarked = marked == "Y"
		m.FlagSubscribed = subscribed == "Y"
		m.FlagRemote = remote == "Y"
		m.FlagArchive = archive == "Y"
		m.FlagDrafts = drafts == "Y"
		m.FlagFlagged = flagged == "Y"
		m.FlagJunk = junk == "Y"
		m.FlagSent = sent == "Y"
		m.FlagTrash = trash == "Y"
		m.FlagImportant = important == "Y"
		mailboxes = append(mailboxes, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return mailboxes, nil
}

func (s Svc) AddMailbox(m Mailbox) error {
	m.Name = strings.TrimSpace(m.Name)
	if m.Name == "" {
		return errors.New("getmalname must be filled out")
	}
	if strings.EqualFold(m.Name, "inbox") {
		m.Name = "INBOX"
	}

	var existing string
	err := s.db.QueryRow(
		`SELECT mailbox_id FROM mailbox WHERE user_id = ? AND name = ?`,
		m.UserID, m.Name,
	).Scan(&existing)
	if err == nil {
		return fmt.Errorf("mailbox %s already exists for this user", m.Name)
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	_, err = s.db.Exec(
		`INSERT INTO mailbox (
			mailbox_id, user_id, name,
			flag_non_existent, flag_no_inferiors, flag_no_select, flag_marked,
			flag_subscribed, flag_remote, flag_archive, flag_drafts,
			flag_flagged, flag_junk, flag_sent, flag_trash, flag_important
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		m.MailboxID, m.UserID, m.Name,
		boolToInd(m.FlagNonExistent), boolToInd(m.FlagNoInferiors), boolToInd(m.FlagNoSelect), boolToInd(m.FlagMarked),
		boolToInd(m.FlagSubscribed), boolToInd(m.FlagRemote), boolToInd(m.FlagArchive), boolToInd(m.FlagDrafts),
		boolToInd(m.FlagFlagged), boolToInd(m.FlagJunk), boolToInd(m.FlagSent), boolToInd(m.FlagTrash), boolToInd(m.FlagImportant),
	)
	return err
}

func (s Svc) UpdateMailbox(m Mailbox) error {
	m.Name = strings.TrimSpace(m.Name)
	if m.Name == "" {
		return errors.New("name must be filled out")
	}
	if strings.EqualFold(m.Name, "inbox") {
		m.Name = "INBOX"
	}

	var existing string
	err := s.db.QueryRow(
		`SELECT mailbox_id FROM mailbox WHERE user_id = ? AND name = ? AND mailbox_id <> ?`,
		m.UserID, m.Name, m.MailboxID,
	).Scan(&existing)
	if err == nil {
		return fmt.Errorf("mailbox %s already exists for this user", m.Name)
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	_, err = s.db.Exec(
		`UPDATE mailbox SET
			name = ?,
			flag_non_existent = ?, flag_no_inferiors = ?, flag_no_select = ?, flag_marked = ?,
			flag_subscribed = ?, flag_remote = ?, flag_archive = ?, flag_drafts = ?,
			flag_flagged = ?, flag_junk = ?, flag_sent = ?, flag_trash = ?, flag_important = ?
		WHERE mailbox_id = ?`,
		m.Name,
		boolToInd(m.FlagNonExistent), boolToInd(m.FlagNoInferiors), boolToInd(m.FlagNoSelect), boolToInd(m.FlagMarked),
		boolToInd(m.FlagSubscribed), boolToInd(m.FlagRemote), boolToInd(m.FlagArchive), boolToInd(m.FlagDrafts),
		boolToInd(m.FlagFlagged), boolToInd(m.FlagJunk), boolToInd(m.FlagSent), boolToInd(m.FlagTrash), boolToInd(m.FlagImportant),
		m.MailboxID,
	)
	return err
}

func boolToInd(b bool) string {
	if b {
		return "Y"
	}
	return "N"
}

func NewService(db *sql.DB) *Svc {
	return &Svc{db: db}
}

// HashPassword hashes a password using SHA512 and returns the hex string
func HashPassword(password string) string {
	hash := sha512.Sum512([]byte(password))
	return fmt.Sprintf("%x", hash)
}
