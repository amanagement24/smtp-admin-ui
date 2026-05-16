package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgb9/smtp-admin/internal/endpoints"
	"github.com/dgb9/smtp-admin/internal/service"
	"github.com/dgb9/smtp-admin/internal/ui"
	_ "github.com/go-sql-driver/mysql"
)

type DbConfig struct {
	Machine  string `json:"machine"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ConfigData struct {
	Address        string   `json:"address"`
	SessionTimeout int      `json:"sessionTimeout"`
	Db             DbConfig `json:"db"`
}

func main() {
	cfgPath := os.Getenv("CONFIG_FILE")
	if cfgPath == "" {
		cfgPath = "/config/config.json"
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Fatalf("read config: %v", err)
	}

	var cfg ConfigData
	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("parse config: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Db.Login, cfg.Db.Password, cfg.Db.Machine, cfg.Db.Port, cfg.Db.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("db ping: %v", err)
	}

	svc := service.NewService(db)
	session := ui.NewSessionStore(cfg.SessionTimeout)
	ctrl := endpoints.NewController(svc, session)

	mux := http.NewServeMux()
	mux.Handle("GET /static/", ui.StaticHandler())
	mux.HandleFunc("GET /", ctrl.GetRoot)
	mux.HandleFunc("POST /login", ctrl.PostLogin)
	mux.HandleFunc("GET /logout", ctrl.GetLogout)
	mux.HandleFunc("GET /editdomain", ctrl.GetEditDomain)
	mux.HandleFunc("POST /editdomain", ctrl.PostEditDomain)
	mux.HandleFunc("GET /edituser", ctrl.GetEditUser)
	mux.HandleFunc("POST /edituser", ctrl.PostEditUser)
	mux.HandleFunc("GET /chpass", ctrl.GetChpass)
	mux.HandleFunc("POST /chpass", ctrl.PostChpass)
	mux.HandleFunc("GET /domains", ctrl.GetDomains)
	mux.HandleFunc("POST /domains", ctrl.PostDomains)
	mux.HandleFunc("POST /cddomains", ctrl.PostCdDomains)
	mux.HandleFunc("POST /cdusers", ctrl.PostCdUsers)
	mux.HandleFunc("GET /viewdomain", ctrl.GetViewDomain)
	mux.HandleFunc("POST /viewdomain", ctrl.PostViewDomain)

	log.Printf("listening on %s", cfg.Address)
	if err := http.ListenAndServe(cfg.Address, mux); err != nil {
		log.Fatalf("server: %v", err)
	}
}
