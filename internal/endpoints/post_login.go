package endpoints

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgb9/smtp-admin/internal/service"
	"github.com/dgb9/smtp-admin/internal/ui"
)

func (c *Controller) PostLogin(w http.ResponseWriter, r *http.Request) {
	session := c.session.GetSession(r, w)
	session.Expired = false

	login := r.FormValue("login")
	password := r.FormValue("password")

	err := c.processPostLogin(w, session, login, password)
	if err != nil {
		ui.RenderLogin(w, ui.ViewLogin{
			ViewHeader: ui.ViewHeader{
				Login: login,
				Error: err.Error(),
			},
			Login: login,
		})
	}

}

func (c *Controller) processPostLogin(w http.ResponseWriter, session *ui.SessionStore, login string, password string) error {
	err := checkLoginParameters(login)
	if err != nil {
		return err
	}

	svc := c.svc
	user, err := svc.GetUserByLogin(login)
	if err != nil {
		return err
	}

	if user == nil {
		return fmt.Errorf("User %s not found", login)
	}

	// check password
	hashedPassword := service.HashPassword(password)
	if hashedPassword != user.Password {
		return fmt.Errorf("Invalid credentials")
	}

	session.User = user

	if user.Admin {
		domains, err := svc.GetDomains()
		if err != nil {
			return err
		}
		session.Domains = domains
		ui.RenderDomains(w, session.NewViewDomains(""))
	} else {
		ui.RenderChPass(w, session.NewViewChPass("", *user, false))
	}

	return nil
}

func checkLoginParameters(login string) error {
	login = strings.TrimSpace(login)
	if len(login) == 0 {
		return errors.New("please provide the login")
	}

	return nil
}
