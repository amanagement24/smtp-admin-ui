package endpoints

import (
	"net/http"

	"github.com/dgb9/smtp-admin/internal/service"
	"github.com/dgb9/smtp-admin/internal/ui"
)

type Controller struct {
	svc     *service.Svc
	session *ui.AppSession
}

func NewController(svc *service.Svc, session *ui.AppSession) *Controller {
	return &Controller{
		svc:     svc,
		session: session,
	}
}

func checkLoggedIn(w http.ResponseWriter, session *ui.SessionStore, checkAdmin bool) bool {
	res := true

	if !session.IsLoggedIn() {
		errMsg := ""
		if session.Expired {
			errMsg = "session expired"
			session.Expired = false
		}

		ui.RenderLogin(w, ui.ViewLogin{
			ViewHeader: ui.ViewHeader{
				Login: "",
				Error: errMsg,
			},
			Login: "",
		})

		res = false
	}

	if res && checkAdmin && !session.IsAdmin() {
		ui.RenderLogin(w, ui.ViewLogin{
			ViewHeader: ui.ViewHeader{
				Login: "",
				Error: "this functionality is for admin users only",
			},
			Login: "",
		})

		res = false
	}

	return res
}

func getMap(selected []string) map[string]bool {
	res := make(map[string]bool)

	for _, domain := range selected {
		res[domain] = true
	}

	return res
}
