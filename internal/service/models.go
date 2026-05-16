package service

type Domain struct {
	DomainID      string
	Name          string
	CatchAll      bool
	CatchallLogin string
}

type User struct {
	UserID   string
	DomainID string
	Login    string
	Password string
	Admin    bool
}

type Mailbox struct {
	MailboxID       string
	UserID          string
	Name            string
	FlagNonExistent bool
	FlagNoInferiors bool
	FlagNoSelect    bool
	FlagMarked      bool
	FlagSubscribed  bool
	FlagRemote      bool
	FlagArchive     bool
	FlagDrafts      bool
	FlagFlagged     bool
	FlagJunk        bool
	FlagSent        bool
	FlagTrash       bool
	FlagImportant   bool
}

// Flags returns the names of the IMAP flags that are set on the mailbox.
func (m Mailbox) Flags() []string {
	var f []string
	for _, fl := range []struct {
		set  bool
		name string
	}{
		{m.FlagNonExistent, "non-existent"},
		{m.FlagNoInferiors, "no-inferiors"},
		{m.FlagNoSelect, "no-select"},
		{m.FlagMarked, "marked"},
		{m.FlagSubscribed, "subscribed"},
		{m.FlagRemote, "remote"},
		{m.FlagArchive, "archive"},
		{m.FlagDrafts, "drafts"},
		{m.FlagFlagged, "flagged"},
		{m.FlagJunk, "junk"},
		{m.FlagSent, "sent"},
		{m.FlagTrash, "trash"},
		{m.FlagImportant, "important"},
	} {
		if fl.set {
			f = append(f, fl.name)
		}
	}
	return f
}

type UserSession struct {
	SessionID   string
	UserID      string
	Token       string
	SessionData string
	CreatedDate string
	Expired     bool
}
