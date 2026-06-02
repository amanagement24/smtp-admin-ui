package ui

import "github.com/dgb9/smtp-admin/internal/service"

func (s *SessionStore) NewViewLogin(errMsg, loginValue string) ViewLogin {
	return ViewLogin{
		ViewHeader: ViewHeader{Context: s.Context, Error: errMsg},
		Login:      loginValue,
	}
}

func sliceToMap(keys []string) map[string]bool {
	m := make(map[string]bool, len(keys))
	for _, k := range keys {
		m[k] = true
	}
	return m
}

func (s *SessionStore) NewViewDomains(errMsg string) ViewDomains {
	return ViewDomains{
		ViewHeader: ViewHeader{Login: s.GetLogin(), Admin: s.IsAdmin(), Error: errMsg, Context: s.Context},
		Domains:    s.Domains,
		Selected:   sliceToMap(s.SelectedDomains),
	}
}

func (s *SessionStore) NewViewDomain(errMsg string) ViewDomain {
	d := service.Domain{}
	if s.CurrentDomain != nil {
		d = *s.CurrentDomain
	}
	return ViewDomain{
		ViewHeader: ViewHeader{Login: s.GetLogin(), Admin: s.IsAdmin(), Error: errMsg, Context: s.Context},
		Domain:     d,
		Users:      s.ViewDomainUsers,
		Selected:   sliceToMap(s.SelectedUsers),
	}
}

func (s *SessionStore) NewViewCdDomains(errMsg string) ViewCdDomains {
	return ViewCdDomains{
		ViewHeader: ViewHeader{Login: s.GetLogin(), Admin: s.IsAdmin(), Error: errMsg, Context: s.Context},
		Domains:    s.CdDomains,
	}
}

func (s *SessionStore) NewViewCdUsers(errMsg string) ViewCdUsers {
	d := service.Domain{}
	if s.CurrentDomain != nil {
		d = *s.CurrentDomain
	}
	return ViewCdUsers{
		ViewHeader: ViewHeader{Login: s.GetLogin(), Admin: s.IsAdmin(), Error: errMsg, Context: s.Context},
		Domain:     d,
		Users:      s.CdUsers,
	}
}

func (s *SessionStore) NewViewEditDomain(errMsg, name, domainID string, adding, catchAll bool, catchAllLogin string) ViewEditDomain {
	return ViewEditDomain{
		ViewHeader:    ViewHeader{Login: s.GetLogin(), Admin: s.IsAdmin(), Error: errMsg},
		Name:          name,
		DomainID:      domainID,
		Adding:        adding,
		CatchAll:      catchAll,
		CatchAllLogin: catchAllLogin,
	}
}

func (s *SessionStore) NewViewEditUser(errMsg string, user service.User, mailboxes []service.Mailbox, adding bool) ViewEditUser {
	return ViewEditUser{
		ViewHeader: ViewHeader{Login: s.GetLogin(), Admin: s.IsAdmin(), Error: errMsg, Context: s.Context},
		User:       user,
		Mailboxes:  mailboxes,
		Adding:     adding,
	}
}

func (s *SessionStore) NewViewEditMailbox(errMsg string, mailbox service.Mailbox, adding bool, userLogin string) ViewEditMailbox {
	return ViewEditMailbox{
		ViewHeader: ViewHeader{Login: s.GetLogin(), Admin: s.IsAdmin(), Error: errMsg, Context: s.Context},
		Mailbox:    mailbox,
		Adding:     adding,
		UserLogin:  userLogin,
	}
}

func (s *SessionStore) NewViewChPass(errMsg string, user service.User, success bool) ViewChPass {
	return ViewChPass{
		ViewHeader: ViewHeader{Login: s.GetLogin(), Admin: s.IsAdmin(), Error: errMsg, Context: s.Context},
		User:       user,
		Success:    success,
	}
}
