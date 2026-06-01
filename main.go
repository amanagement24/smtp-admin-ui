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
	Context        string   `json:"context"`
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
	session := ui.NewSessionStore(cfg.SessionTimeout, cfg.Context)
	ctrl := endpoints.NewController(svc, session)

	ctx := cfg.Context
	mux := http.NewServeMux()
	mux.Handle("GET "+ctx+"/static/", http.StripPrefix(ctx, ui.StaticHandler()))
	mux.HandleFunc("GET "+ctx+"/", ctrl.GetRoot)
	mux.HandleFunc("POST "+ctx+"/login", ctrl.PostLogin)
	mux.HandleFunc("GET "+ctx+"/logout", ctrl.GetLogout)
	mux.HandleFunc("GET "+ctx+"/editdomain", ctrl.GetEditDomain)
	mux.HandleFunc("POST "+ctx+"/editdomain", ctrl.PostEditDomain)
	mux.HandleFunc("GET "+ctx+"/edituser", ctrl.GetEditUser)
	mux.HandleFunc("POST "+ctx+"/edituser", ctrl.PostEditUser)
	mux.HandleFunc("GET "+ctx+"/chpass", ctrl.GetChpass)
	mux.HandleFunc("POST "+ctx+"/chpass", ctrl.PostChpass)
	mux.HandleFunc("GET "+ctx+"/domains", ctrl.GetDomains)
	mux.HandleFunc("POST "+ctx+"/domains", ctrl.PostDomains)
	mux.HandleFunc("POST "+ctx+"/cddomains", ctrl.PostCdDomains)
	mux.HandleFunc("POST "+ctx+"/cdusers", ctrl.PostCdUsers)
	mux.HandleFunc("GET "+ctx+"/viewdomain", ctrl.GetViewDomain)
	mux.HandleFunc("POST "+ctx+"/viewdomain", ctrl.PostViewDomain)

	log.Printf("listening on %s", cfg.Address)
	if err := http.ListenAndServe(cfg.Address, mux); err != nil {
		log.Fatalf("server: %v", err)
	}
}
