package api

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
)

var (
	store *sessions.FilesystemStore
)

const AUTH_SESSION = "auth-session"

func Init(secret string) error {
	store = sessions.NewFilesystemStore("", []byte(secret))
	gob.Register(map[string]interface{}{})
	return nil
}
