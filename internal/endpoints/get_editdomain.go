package endpoints

import (
	"net/http"

	"github.com/dgb9/smtp-admin/internal/ui"
)

func (c *Controller) GetEditDomain(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)
	if !checkLoggedIn(w, session, true) {
		return
	}

	id := r.URL.Query().Get("id")

	for _, d := range session.Domains {
		if d.DomainID == id {
			ui.RenderEditDomain(w, session.NewViewEditDomain("", d.Name, d.DomainID, false, d.CatchAll, d.CatchallLogin))
			return
		}
	}

	ui.RenderDomains(w, session.NewViewDomains("domain not found"))
}
