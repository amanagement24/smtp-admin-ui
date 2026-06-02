package endpoints

import (
	"net/http"

	"github.com/dgb9/smtp-admin/internal/service"
	"github.com/dgb9/smtp-admin/internal/ui"
	"github.com/google/uuid"
)

func (c *Controller) GetEditMailbox(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)
	if !checkLoggedIn(w, session, true) {
		return
	}

	id := r.URL.Query().Get("id")

	if id != "" {
		mailbox, err := c.svc.GetMailbox(id)
		if err != nil || mailbox == nil {
			ui.RenderDomains(w, session.NewViewDomains("mailbox not found"))
			return
		}
		userLogin := findUserLogin(session, mailbox.UserID)
		ui.RenderEditMailbox(w, session.NewViewEditMailbox("", *mailbox, false, userLogin))
		return
	}

	userID := r.URL.Query().Get("user_id")
	mid, err := uuid.NewV7()
	if err != nil {
		ui.RenderDomains(w, session.NewViewDomains("failed to generate mailbox id"))
		return
	}
	mailbox := service.Mailbox{
		MailboxID: mid.String(),
		UserID:    userID,
	}
	userLogin := findUserLogin(session, userID)
	ui.RenderEditMailbox(w, session.NewViewEditMailbox("", mailbox, true, userLogin))
}

func findUserLogin(session *ui.SessionStore, userID string) string {
	for _, u := range session.ViewDomainUsers {
		if u.UserID == userID {
			return u.Login
		}
	}
	return ""
}
