package endpoints

import (
	"errors"
	"net/http"

	"github.com/dgb9/smtp-admin/internal/service"
	"github.com/dgb9/smtp-admin/internal/ui"
	"github.com/google/uuid"
)

func (c *Controller) PostViewDomain(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)
	if !checkLoggedIn(w, session, true) {
		return
	}

	_ = r.ParseForm()

	selected := r.Form["selected"]
	session.SelectedUsers = selected

	if r.FormValue("back") != "" {
		ui.RenderDomains(w, session.NewViewDomains(""))
		return
	}

	domain := session.CurrentDomain
	if domain == nil {
		ui.RenderDomains(w, session.NewViewDomains("domain not found in session"))
		return
	}

	if r.FormValue("add") != "" {
		id, err := uuid.NewV7()
		if err != nil {
			renderViewDomain(w, session, err.Error())
			return
		}
		ui.RenderEditUser(w, session.NewViewEditUser("", service.User{UserID: id.String(), DomainID: domain.DomainID}, nil, true))
		return
	}

	if r.FormValue("delete") != "" {
		if len(selected) == 0 {
			renderViewDomain(w, session, errors.New("please select at least one user").Error())
			return
		}

		sel := getMap(selected)
		var users []service.User
		for _, u := range session.ViewDomainUsers {
			if sel[u.UserID] {
				users = append(users, u)
			}
		}

		session.CdUsers = users

		ui.RenderCdUsers(w, session.NewViewCdUsers(""))
		return
	}
}

func renderViewDomain(w http.ResponseWriter, session *ui.SessionStore, errMsg string) {
	ui.RenderViewDomain(w, session.NewViewDomain(errMsg))
}
