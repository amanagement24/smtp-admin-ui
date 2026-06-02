package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/dgb9/smtp-admin/internal/endpoints"
	"github.com/dgb9/smtp-admin/internal/service"
	"github.com/dgb9/smtp-admin/internal/ui"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/natefinch/lumberjack.v2"
)

type DbConfig struct {
	Machine  string `json:"machine"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LogConfig struct {
	FileLogEnabled bool   `json:"fileLogEnabled"`
	FileName       string `json:"fileName"`
	MaxSize        int    `json:"maxSize"`
	MaxBackups     int    `json:"maxBackups"`
	MaxDays        int    `json:"maxDays"`
	Compress       bool   `json:"compress"`
}

type ConfigData struct {
	Address        string    `json:"address"`
	Context        string    `json:"context"`
	SessionTimeout int       `json:"sessionTimeout"`
	Db             DbConfig  `json:"db"`
	Log            LogConfig `json:"log"`
}

func main() {
	cfgPath := os.Getenv("CONFIG_FILE")
	cfgPath = strings.TrimSpace(cfgPath)
	if len(cfgPath) == 0 {
		slog.Error("CONFIG_FILE environment variable not set")
		os.Exit(1)
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		slog.Error("read config", "err", err)
		os.Exit(1)
	}

	var cfg ConfigData
	if err := json.Unmarshal(data, &cfg); err != nil {
		slog.Error("parse config", "err", err)
		os.Exit(1)
	}

	var logOut io.Writer = os.Stdout
	if cfg.Log.FileLogEnabled {
		lj := &lumberjack.Logger{
			Filename:   cfg.Log.FileName,
			MaxSize:    cfg.Log.MaxSize,
			MaxBackups: cfg.Log.MaxBackups,
			MaxAge:     cfg.Log.MaxDays,
			Compress:   cfg.Log.Compress,
		}
		logOut = io.MultiWriter(os.Stdout, lj)
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(logOut, nil)))

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Db.Login, cfg.Db.Password, cfg.Db.Machine, cfg.Db.Port, cfg.Db.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		slog.Error("open db", "err", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		slog.Error("db ping", "err", err)
		os.Exit(1)
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
	mux.HandleFunc("GET "+ctx+"/editmailbox", ctrl.GetEditMailbox)
	mux.HandleFunc("POST "+ctx+"/editmailbox", ctrl.PostEditMailbox)
	mux.HandleFunc("GET "+ctx+"/chpass", ctrl.GetChpass)
	mux.HandleFunc("POST "+ctx+"/chpass", ctrl.PostChpass)
	mux.HandleFunc("GET "+ctx+"/domains", ctrl.GetDomains)
	mux.HandleFunc("POST "+ctx+"/domains", ctrl.PostDomains)
	mux.HandleFunc("POST "+ctx+"/cddomains", ctrl.PostCdDomains)
	mux.HandleFunc("POST "+ctx+"/cdusers", ctrl.PostCdUsers)
	mux.HandleFunc("GET "+ctx+"/viewdomain", ctrl.GetViewDomain)
	mux.HandleFunc("POST "+ctx+"/viewdomain", ctrl.PostViewDomain)

	slog.Info("listening", "addr", cfg.Address)
	if err := http.ListenAndServe(cfg.Address, mux); err != nil {
		slog.Error("server", "err", err)
		os.Exit(1)
	}
}
