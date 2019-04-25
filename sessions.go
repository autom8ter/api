package api

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"os"
)

var (
	store *sessions.CookieStore
)

var AUTH_SESSION = "auth-session"

func InitSessions(secret string) error {
	if secret == "" {
		secret = os.Getenv("SECRET")
	}
	store = sessions.NewCookieStore([]byte(secret))
	gob.Register(map[string]interface{}{})
	return nil
}
