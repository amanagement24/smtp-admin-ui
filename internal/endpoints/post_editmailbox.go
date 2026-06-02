package endpoints

import (
	"net/http"

	"github.com/dgb9/smtp-admin/internal/service"
	"github.com/dgb9/smtp-admin/internal/ui"
)

func (c *Controller) PostEditMailbox(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)
	if !checkLoggedIn(w, session, true) {
		return
	}

	_ = r.ParseForm()

	userID := r.FormValue("user_id")
	mailboxID := r.FormValue("id")
	adding := r.FormValue("adding") == "true"

	var currentUser *service.User
	for i, u := range session.ViewDomainUsers {
		if u.UserID == userID {
			currentUser = &session.ViewDomainUsers[i]
			break
		}
	}

	renderEditUser := func(errMsg string) {
		if currentUser == nil {
			ui.RenderDomains(w, session.NewViewDomains(errMsg))
			return
		}
		mailboxes, _ := c.svc.GetMailboxes(userID)
		ui.RenderEditUser(w, session.NewViewEditUser(errMsg, *currentUser, mailboxes, false))
	}

	if r.FormValue("cancel") != "" {
		renderEditUser("")
		return
	}

	mailbox := service.Mailbox{
		MailboxID:       mailboxID,
		UserID:          userID,
		Name:            r.FormValue("name"),
		FlagNonExistent: r.FormValue("flag_non_existent") != "",
		FlagNoInferiors: r.FormValue("flag_no_inferiors") != "",
		FlagNoSelect:    r.FormValue("flag_no_select") != "",
		FlagMarked:      r.FormValue("flag_marked") != "",
		FlagSubscribed:  r.FormValue("flag_subscribed") != "",
		FlagRemote:      r.FormValue("flag_remote") != "",
		FlagArchive:     r.FormValue("flag_archive") != "",
		FlagDrafts:      r.FormValue("flag_drafts") != "",
		FlagFlagged:     r.FormValue("flag_flagged") != "",
		FlagJunk:        r.FormValue("flag_junk") != "",
		FlagSent:        r.FormValue("flag_sent") != "",
		FlagTrash:       r.FormValue("flag_trash") != "",
		FlagImportant:   r.FormValue("flag_important") != "",
	}

	var err error
	if adding {
		err = c.svc.AddMailbox(mailbox)
	} else {
		err = c.svc.UpdateMailbox(mailbox)
	}

	if err != nil {
		userLogin := ""
		if currentUser != nil {
			userLogin = currentUser.Login
		}
		ui.RenderEditMailbox(w, session.NewViewEditMailbox(err.Error(), mailbox, adding, userLogin))
		return
	}

	renderEditUser("")
}
