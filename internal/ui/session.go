package ui

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/dgb9/smtp-admin/internal/service"
)

type SessionStore struct {
	Context         string
	Created         time.Time
	Expired         bool
	User            *service.User
	Domains         []service.Domain
	SelectedDomains []string
	CdDomains       []service.Domain
	CurrentDomain   *service.Domain
	ViewDomainUsers []service.User
	SelectedUsers   []string
	CdUsers         []service.User
}

func (s *SessionStore) IsAdmin() bool {
	return s.IsLoggedIn() && s.User.Admin
}

func (s *SessionStore) IsLoggedIn() bool {
	return s.User != nil
}

func (s *SessionStore) GetLogin() string {
	res := ""

	if s.User != nil {
		res = s.User.Login
	}

	return res
}

func (s *SessionStore) Clear() {
	s.Expired = false
	s.User = nil
	s.Domains = nil
	s.SelectedDomains = nil
	s.CdDomains = nil
	s.CurrentDomain = nil
	s.ViewDomainUsers = nil
	s.SelectedUsers = nil
	s.CdUsers = nil
}

type AppSession struct {
	sessionMap map[string]*SessionStore
	timeout    time.Duration
	context    string
}

func NewSessionStore(sessionTimeoutSeconds int, context string) *AppSession {
	return &AppSession{
		sessionMap: make(map[string]*SessionStore),
		timeout:    time.Duration(sessionTimeoutSeconds) * time.Second,
		context:    context,
	}
}

func (a *AppSession) Get(key string) *SessionStore {
	return a.sessionMap[key]
}

func (a *AppSession) Contains(key string) bool {
	_, ok := a.sessionMap[key]
	return ok
}

func (a *AppSession) GetSession(r *http.Request, w http.ResponseWriter) *SessionStore {
	cookie, err := r.Cookie("smtpadmin")
	if err == nil {
		if store, ok := a.sessionMap[cookie.Value]; ok {
			if a.timeout > 0 && time.Since(store.Created) > a.timeout {
				store.Clear()
				store.Created = time.Now()
				store.Expired = true
			}
			return store
		}
	}

	store := &SessionStore{Created: time.Now(), Context: a.context}
	key := newSessionKey()
	http.SetCookie(w, &http.Cookie{
		Name:     "smtpadmin",
		Value:    key,
		HttpOnly: true,
	})
	a.sessionMap[key] = store
	return store
}

func newSessionKey() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
