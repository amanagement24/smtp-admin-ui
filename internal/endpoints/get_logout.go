package endpoints

import (
	"net/http"

	"github.com/dgb9/smtp-admin/internal/ui"
)

func (c *Controller) GetLogout(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)

	// clear the session
	session.Clear()

	ui.RenderLogin(w, session.NewViewLogin("", ""))
}
