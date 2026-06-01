package endpoints

import (
	"net/http"

	"github.com/dgb9/smtp-admin/internal/ui"
)

func (c *Controller) PostCdDomains(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)
	if !checkLoggedIn(w, session, true) {
		return
	}

	_ = r.ParseForm()

	if r.FormValue("cancel") != "" {
		ui.RenderDomains(w, session.NewViewDomains(""))
		return
	}

	if r.FormValue("submit") != "" {
		ids := make([]string, 0, len(session.CdDomains))
		for _, d := range session.CdDomains {
			ids = append(ids, d.DomainID)
		}

		if err := c.svc.DeleteDomains(ids); err != nil {
			ui.RenderCdDomains(w, session.NewViewCdDomains(err.Error()))
			return
		}

		domains, err := c.svc.GetDomains()
		if err != nil {
			ui.RenderCdDomains(w, session.NewViewCdDomains(err.Error()))
			return
		}

		session.Domains = domains
		session.SelectedDomains = nil
		session.CdDomains = nil

		ui.RenderDomains(w, session.NewViewDomains(""))
	}
}
