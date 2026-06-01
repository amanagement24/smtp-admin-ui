package endpoints

import (
	"net/http"

	"github.com/dgb9/smtp-admin/internal/service"
	"github.com/dgb9/smtp-admin/internal/ui"
)

func (c *Controller) PostEditDomain(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)
	if !checkLoggedIn(w, session, true) {
		return
	}

	if r.FormValue("cancel") != "" {
		ui.RenderDomains(w, session.NewViewDomains(""))
		return
	}

	adding := r.FormValue("adding") == "true"
	id := r.FormValue("id")

	domain := r.FormValue("domain")
	catchAll := r.FormValue("catchall") != ""
	catchAllUser := r.FormValue("catchall_login")

	var err error
	if adding {
		err = c.svc.AddDomain(id, domain)

		if err == nil {
			session.Domains = append(session.Domains, service.Domain{
				DomainID:      id,
				Name:          domain,
				CatchAll:      false,
				CatchallLogin: "",
			})

			ui.RenderDomains(w, session.NewViewDomains(""))
			return
		}
	} else {
		err = c.svc.UpdateDomain(id, domain, catchAll, catchAllUser)

		if err == nil {
			var domains []service.Domain
			domains, err = c.svc.GetDomains()

			if err == nil {
				session.Domains = domains
				ui.RenderDomains(w, session.NewViewDomains(""))
				return
			}
		}
	}

	message := ""
	if err != nil {
		message = err.Error()
	}

	ui.RenderEditDomain(w, session.NewViewEditDomain(message, domain, id, adding, catchAll, catchAllUser))
}
