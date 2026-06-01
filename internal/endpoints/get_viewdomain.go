package endpoints

import (
	"errors"
	"net/http"

	"github.com/dgb9/smtp-admin/internal/ui"
)

func (c *Controller) GetViewDomain(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)
	if !checkLoggedIn(w, session, true) {
		return
	}

	id := r.URL.Query().Get("id")

	domain, err := c.svc.GetDomain(id)
	if err == nil && domain == nil {
		err = errors.New("domain not found")
	}
	if err != nil {
		ui.RenderDomains(w, session.NewViewDomains(err.Error()))
		return
	}

	users, err := c.svc.GetUsersByDomain(domain.Name)
	if err != nil {
		ui.RenderDomains(w, session.NewViewDomains(err.Error()))
		return
	}

	session.CurrentDomain = domain
	session.ViewDomainUsers = users
	session.SelectedUsers = nil

	ui.RenderViewDomain(w, session.NewViewDomain(""))
}
