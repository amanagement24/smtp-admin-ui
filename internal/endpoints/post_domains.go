package endpoints

import (
	"errors"
	"net/http"

	"github.com/dgb9/smtp-admin/internal/service"
	"github.com/dgb9/smtp-admin/internal/ui"
	"github.com/google/uuid"
)

func (c *Controller) PostDomains(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)

	if !checkLoggedIn(w, session, true) {
		return
	}

	_ = r.ParseForm()

	add := r.FormValue("add")
	del := r.FormValue("delete")

	selected := r.Form["selected"]
	session.SelectedDomains = selected

	var err error

	if add != "" {
		err = c.processAddDomain(w, session)

		if err == nil {
			return
		}
	} else if del != "" {
		err = c.processDeleteDomains(w, session)

		if err == nil {
			return
		}
	}

	message := ""
	if err != nil {
		message = err.Error()
	}

	ui.RenderDomains(w, session.NewViewDomains(message))
}

func (c *Controller) processDeleteDomains(w http.ResponseWriter, session *ui.SessionStore) error {
	if len(session.SelectedDomains) == 0 {
		return errors.New("please select at least one domain")
	}

	selected := getMap(session.SelectedDomains)

	var domains []service.Domain
	for _, d := range session.Domains {
		if selected[d.DomainID] {
			domains = append(domains, d)
		}
	}

	session.CdDomains = domains

	ui.RenderCdDomains(w, session.NewViewCdDomains(""))

	return nil
}

func (c *Controller) processAddDomain(w http.ResponseWriter, session *ui.SessionStore) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	ui.RenderEditDomain(w, session.NewViewEditDomain("", "", id.String(), true, false, ""))

	return nil
}
