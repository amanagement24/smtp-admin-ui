package endpoints

import (
	"net/http"

	"github.com/dgb9/smtp-admin/internal/ui"
)

func (c *Controller) GetDomains(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)
	if !checkLoggedIn(w, session, true) {
		return
	}

	ui.RenderDomains(w, session.NewViewDomains(""))
}
