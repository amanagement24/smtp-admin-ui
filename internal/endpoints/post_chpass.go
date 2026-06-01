package endpoints

import (
	"net/http"

	"github.com/dgb9/smtp-admin/internal/service"
	"github.com/dgb9/smtp-admin/internal/ui"
)

func (c *Controller) PostChpass(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)
	if !checkLoggedIn(w, session, false) {
		return
	}

	_ = r.ParseForm()

	if r.FormValue("cancel") != "" {
		if session.IsAdmin() {
			ui.RenderViewDomain(w, session.NewViewDomain(""))
		} else {
			ui.RenderChPass(w, session.NewViewChPass("", *session.User, false))
		}
		return
	}

	id := r.FormValue("id")

	var user *service.User
	if session.IsAdmin() {
		for i := range session.ViewDomainUsers {
			if session.ViewDomainUsers[i].UserID == id {
				user = &session.ViewDomainUsers[i]
				break
			}
		}
	} else {
		if id != session.User.UserID {
			ui.RenderChPass(w, session.NewViewChPass("not allowed to change another user's password", *session.User, false))
			return
		}
		user = session.User
	}

	if user == nil {
		ui.RenderViewDomain(w, session.NewViewDomain("user not found"))
		return
	}

	newPassword := r.FormValue("new_password")
	repeatPassword := r.FormValue("repeat_password")

	var errMsg string
	switch {
	case newPassword == "" || repeatPassword == "":
		errMsg = "both password and repeat must be filled out"
	case newPassword != repeatPassword:
		errMsg = "passwords do not match"
	default:
		if err := c.svc.UpdatePassword(id, service.HashPassword(newPassword)); err != nil {
			errMsg = err.Error()
		}
	}

	if errMsg != "" {
		ui.RenderChPass(w, session.NewViewChPass(errMsg, *user, false))
		return
	}

	if session.IsAdmin() {
		ui.RenderViewDomain(w, session.NewViewDomain(""))
	} else {
		ui.RenderChPass(w, session.NewViewChPass("", *user, true))
	}
}
