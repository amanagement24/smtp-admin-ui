package endpoints

import (
	"net/http"

	"github.com/dgb9/smtp-admin/internal/service"
	"github.com/dgb9/smtp-admin/internal/ui"
	"github.com/google/uuid"
)

func (c *Controller) PostEditUser(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)
	if !checkLoggedIn(w, session, true) {
		return
	}

	_ = r.ParseForm()

	domain := session.CurrentDomain

	if r.FormValue("cancel") != "" {
		ui.RenderViewDomain(w, session.NewViewDomain(""))
		return
	}

	adding := r.FormValue("adding") == "true"
	id := r.FormValue("id")
	domainID := r.FormValue("domain_id")
	login := r.FormValue("login")
	admin := r.FormValue("admin") != ""
	if !admin && session.User != nil && id == session.User.UserID {
		ui.RenderEditUser(w, session.NewViewEditUser(
			"current user cannot remove own admin flag",
			service.User{UserID: id, DomainID: domainID, Login: login, Admin: true},
			nil,
			adding,
		))
		return
	}

	var err error
	if adding {
		err = c.svc.AddUser(id, domainID, login, admin)
		if err == nil {
			if mid, idErr := uuid.NewV7(); idErr != nil {
				err = idErr
			} else {
				err = c.svc.AddMailbox(service.Mailbox{
					MailboxID: mid.String(),
					UserID:    id,
					Name:      "INBOX",
				})
			}
		}
		if err == nil {
			session.ViewDomainUsers = append(session.ViewDomainUsers, service.User{
				UserID:   id,
				DomainID: domainID,
				Login:    login,
				Admin:    admin,
			})
			session.SelectedUsers = nil
			ui.RenderViewDomain(w, session.NewViewDomain(""))
			return
		}
	} else {
		err = c.svc.UpdateUser(id, domainID, login, admin)
		if err == nil {
			users, fetchErr := c.svc.GetUsersByDomain(domain.Name)
			if fetchErr == nil {
				session.ViewDomainUsers = users
				session.SelectedUsers = nil
				ui.RenderViewDomain(w, session.NewViewDomain(""))
				return
			}
			err = fetchErr
		}
	}

	ui.RenderEditUser(w, session.NewViewEditUser(
		err.Error(),
		service.User{UserID: id, DomainID: domainID, Login: login, Admin: admin},
		nil,
		adding,
	))
}
