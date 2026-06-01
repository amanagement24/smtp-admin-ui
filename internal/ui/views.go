package ui

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/dgb9/smtp-admin/internal/service"
)

//go:embed templates
var templateFiles embed.FS

//go:embed static
var staticFiles embed.FS

// StaticHandler serves embedded static assets (CSS, etc.) under /static/.
func StaticHandler() http.Handler {
	return http.FileServer(http.FS(staticFiles))
}

type ViewHeader struct {
	Login string
	Error string
	Admin bool
}

func (v ViewHeader) IsLoggedIn() bool {
	return len(v.Login) > 0
}

func (v ViewHeader) IsAdmin() bool {
	return v.Admin
}

type ViewLogin struct {
	ViewHeader
	Login string
}

func RenderLogin(w http.ResponseWriter, viewLogin ViewLogin) {
	t := template.Must(template.ParseFS(templateFiles, "templates/header.html", "templates/login.html"))
	if err := t.ExecuteTemplate(w, "header", viewLogin); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type ViewDomains struct {
	ViewHeader
	Domains  []service.Domain
	Selected map[string]bool
}

func RenderDomains(w http.ResponseWriter, v ViewDomains) {
	t := template.Must(template.ParseFS(templateFiles, "templates/header.html", "templates/domains.html"))
	if err := t.ExecuteTemplate(w, "header", v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type ViewChPass struct {
	ViewHeader
	User    service.User
	Success bool
}

func RenderChPass(w http.ResponseWriter, v ViewChPass) {
	t := template.Must(template.ParseFS(templateFiles, "templates/header.html", "templates/chpass.html"))
	if err := t.ExecuteTemplate(w, "header", v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type ViewCdDomains struct {
	ViewHeader
	Domains []service.Domain
}

func RenderCdDomains(w http.ResponseWriter, v ViewCdDomains) {
	t := template.Must(template.ParseFS(templateFiles, "templates/header.html", "templates/cddomains.html"))
	if err := t.ExecuteTemplate(w, "header", v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type ViewCdUsers struct {
	ViewHeader
	Domain service.Domain
	Users  []service.User
}

func RenderCdUsers(w http.ResponseWriter, v ViewCdUsers) {
	t := template.Must(template.ParseFS(templateFiles, "templates/header.html", "templates/cdusers.html"))
	if err := t.ExecuteTemplate(w, "header", v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type ViewDomain struct {
	ViewHeader
	Domain   service.Domain
	Users    []service.User
	Selected map[string]bool
}

func RenderViewDomain(w http.ResponseWriter, v ViewDomain) {
	t := template.Must(template.ParseFS(templateFiles, "templates/header.html", "templates/viewdomain.html"))
	if err := t.ExecuteTemplate(w, "header", v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type ViewEditDomain struct {
	ViewHeader
	Name          string
	DomainID      string
	Adding        bool
	CatchAll      bool
	CatchAllLogin string
}

func RenderEditDomain(w http.ResponseWriter, v ViewEditDomain) {
	t := template.Must(template.ParseFS(templateFiles, "templates/header.html", "templates/editdomain.html"))
	if err := t.ExecuteTemplate(w, "header", v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type ViewEditUser struct {
	ViewHeader
	User      service.User
	Mailboxes []service.Mailbox
	Adding    bool
}

func RenderEditUser(w http.ResponseWriter, v ViewEditUser) {
	t := template.Must(template.ParseFS(templateFiles, "templates/header.html", "templates/edituser.html"))
	if err := t.ExecuteTemplate(w, "header", v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
