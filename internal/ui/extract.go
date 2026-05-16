package ui

import (
	"net/http"
)

type PostLogin struct {
	Login    string
	Password string
}

type PostChpass struct {
	NewPassword    string
	RepeatPassword string
	Submit         bool
}

type PostDomains struct {
	Selected []string
	Add      bool
	Delete   bool
}

type PostViewDomain struct {
	Selected []string
	Action   string
}

type PostEditDomain struct {
	ID            string
	Name          string
	CatchAll      bool
	CatchAllLogin string
	Action        string
}

type PostCdDomains struct {
	Confirmed bool
}

type PostCdUsers struct {
	Confirmed bool
}

type PostEditUser struct {
	ID       string
	DomainID string
	Login    string
	Admin    bool
	Action   string
}

func ExtractLogin(r *http.Request) PostLogin {
	return PostLogin{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}
}

func ExtractChpass(r *http.Request) PostChpass {
	return PostChpass{
		NewPassword:    r.FormValue("new_password"),
		RepeatPassword: r.FormValue("repeat_password"),
		Submit:         r.FormValue("submit") != "",
	}
}

func ExtractDomains(r *http.Request) PostDomains {
	_ = r.ParseForm()
	return PostDomains{
		Selected: r.Form["selected"],
		Add:      r.FormValue("add") != "",
		Delete:   r.FormValue("delete") != "",
	}
}

func ExtractViewDomain(r *http.Request) PostViewDomain {
	_ = r.ParseForm()
	return PostViewDomain{
		Selected: r.Form["selected"],
		Action:   r.FormValue("action"),
	}
}

func ExtractEditDomain(r *http.Request) PostEditDomain {
	return PostEditDomain{
		ID:            r.FormValue("id"),
		Name:          r.FormValue("name"),
		CatchAll:      r.FormValue("catchall") != "",
		CatchAllLogin: r.FormValue("catchall_login"),
		Action:        r.FormValue("action"),
	}
}

func ExtractCdDomains(r *http.Request) PostCdDomains {
	return PostCdDomains{
		Confirmed: r.FormValue("submit") != "",
	}
}

func ExtractCdUsers(r *http.Request) PostCdUsers {
	return PostCdUsers{
		Confirmed: r.FormValue("submit") != "",
	}
}

func ExtractEditUser(r *http.Request) PostEditUser {
	return PostEditUser{
		ID:       r.FormValue("id"),
		DomainID: r.FormValue("domain_id"),
		Login:    r.FormValue("login"),
		Admin:    r.FormValue("admin") != "",
		Action:   r.FormValue("action"),
	}
}
