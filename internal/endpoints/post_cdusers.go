package endpoints

import (
	"fmt"
	"net/http"

	"github.com/dgb9/smtp-admin/internal/service"
	"github.com/dgb9/smtp-admin/internal/ui"
)

func (c *Controller) PostCdUsers(w http.ResponseWriter, r *http.Request) {
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

	if r.FormValue("submit") != "" {
		ids := make([]string, 0, len(session.CdUsers))
		for _, u := range session.CdUsers {
			ids = append(ids, u.UserID)
		}

		if domain.CatchAll {
			for _, u := range session.CdUsers {
				if u.Login == domain.CatchallLogin {
					errMsg := fmt.Sprintf("cannot delete %s: it is the catch-all login for this domain", u.Login)
					ui.RenderCdUsers(w, session.NewViewCdUsers(errMsg))
					return
				}
			}
		}

		if err := c.svc.DeleteUsers(ids); err != nil {
			ui.RenderCdUsers(w, session.NewViewCdUsers(err.Error()))
			return
		}

		deleted := getMap(ids)
		var remaining []service.User
		for _, u := range session.ViewDomainUsers {
			if !deleted[u.UserID] {
				remaining = append(remaining, u)
			}
		}
		session.ViewDomainUsers = remaining
		session.SelectedUsers = nil
		session.CdUsers = nil

		ui.RenderViewDomain(w, session.NewViewDomain(""))
	}
}
