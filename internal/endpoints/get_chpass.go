package endpoints

import (
	"net/http"

	"github.com/dgb9/smtp-admin/internal/ui"
)

func (c *Controller) GetChpass(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)
	if !checkLoggedIn(w, session, true) {
		return
	}

	id := r.URL.Query().Get("id")

	for _, u := range session.ViewDomainUsers {
		if u.UserID == id {
			ui.RenderChPass(w, session.NewViewChPass("", u, false))
			return
		}
	}

	if session.CurrentDomain != nil {
		ui.RenderViewDomain(w, session.NewViewDomain("user not found"))
		return
	}

	ui.RenderDomains(w, session.NewViewDomains("user not found"))
}
