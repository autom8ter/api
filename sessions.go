package api

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"os"
)

var (
	store *sessions.FilesystemStore
)

var AUTH_SESSION = "auth-session"
var SESSIONS_PATH = os.Getenv("PWD") + "/sessions"

func InitSessions(secret string) error {
	_ = os.MkdirAll(SESSIONS_PATH, 0755)
	if secret == "" {
		secret = os.Getenv("SECRET")
	}
	store = sessions.NewFilesystemStore(SESSIONS_PATH, []byte(secret))
	gob.Register(map[string]interface{}{})
	return nil
}
