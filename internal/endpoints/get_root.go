package endpoints

import (
	"net/http"

	"github.com/dgb9/smtp-admin/internal/ui"
)

func (c *Controller) GetRoot(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)

	if session.IsLoggedIn() {
		ui.RenderDomains(w, session.NewViewDomains(""))
	} else {
		errMsg := ""
		if session.Expired {
			errMsg = "session expired"
			session.Expired = false
		}
		ui.RenderLogin(w, session.NewViewLogin(errMsg, ""))
	}
}
