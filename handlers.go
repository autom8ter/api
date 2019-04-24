package api

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/Masterminds/sprig"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var LOGIN_PATH = "/login"
var LOGOUT_PATH = "/logout"
var CALLBACK_PATH = "/callback"
var HOMEURL = "http://localhost:8080"
var DEFAULT_REDIRECT = "http://localhost:8080/callback"
var DEFAULT_SCOPES = []string{"openid", "profile", "email"}

func (a *Auth) defaultIfEmpty() {
	if len(a.Scopes) == 0 {
		a.Scopes = DEFAULT_SCOPES
	}
	if a.Audience == "" {
		a.Audience = "https://" + a.Domain + "/userinfo"
	}
	if a.Redirect == "" {
		a.Redirect = DEFAULT_REDIRECT
	}

}
func (a *Auth) validate() error {
	if len(a.Scopes) == 0 {
		return errors.New("validation error: empty scope")
	}
	if a.Redirect == "" {
		return errors.New("validation error: empty redirect")
	}
	if a.ClientId == "" {
		return errors.New("validation error: empty client id")
	}
	if a.Domain == "" {
		return errors.New("validation error: empty domain")
	}
	if a.ClientSecret == "" {
		return errors.New("validation error: empty client secret")
	}
	if a.Audience == "" {
		return errors.New("validation error: empty audience")
	}
	return nil
}

func (a *Auth) callbackHandler(loggedInPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.defaultIfEmpty()
		if err := a.validate(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		conf := &oauth2.Config{
			ClientID:     a.ClientId,
			ClientSecret: a.ClientSecret,
			RedirectURL:  a.Redirect,
			Scopes:       a.Scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://" + a.Domain + "/authorize",
				TokenURL: "https://" + a.Domain + "/oauth/token",
			},
		}
		state := r.URL.Query().Get("state")
		session, err := store.Get(r, "state")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if state != session.Values["state"] {
			http.Error(w, "Invalid state parameter", http.StatusInternalServerError)
			return
		}

		code := r.URL.Query().Get("code")

		token, err := conf.Exchange(r.Context(), code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Getting now the userInfo
		client := conf.Client(r.Context(), token)
		resp, err := client.Get(a.Audience)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		var userinfo map[string]interface{}
		if err = json.NewDecoder(resp.Body).Decode(&userinfo); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session, err = store.Get(r, AUTH_SESSION)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Values["id_token"] = token.Extra("id_token")
		session.Values["refresh_token"] = token.RefreshToken
		session.Values["access_token"] = token.AccessToken
		session.Values["userinfo"] = userinfo
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to logged in page
		http.Redirect(w, r, loggedInPath, http.StatusSeeOther)
	}

}

func (a *Auth) loginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.defaultIfEmpty()
		if err := a.validate(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		conf := &oauth2.Config{
			ClientID:     a.ClientId,
			ClientSecret: a.ClientSecret,
			RedirectURL:  a.Redirect,
			Scopes:       a.Scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://" + a.Domain + "/authorize",
				TokenURL: "https://" + a.Domain + "/oauth/token",
			},
		}

		// Generate random state
		b := make([]byte, 32)
		rand.Read(b)
		state := base64.StdEncoding.EncodeToString(b)

		session, err := store.Get(r, "state")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["state"] = state
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		audience := oauth2.SetAuthURLParam("audience", a.Audience)
		u := conf.AuthCodeURL(state, audience)

		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	}
}

func (a *Auth) logoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.defaultIfEmpty()
		if err := a.validate(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var Url *url.URL
		Url, err := url.Parse("https://" + a.Domain)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Url.Path += "/v2/logout"
		parameters := url.Values{}
		parameters.Add("returnTo", HOMEURL)
		parameters.Add("client_id", a.ClientId)
		Url.RawQuery = parameters.Encode()

		http.Redirect(w, r, Url.String(), http.StatusTemporaryRedirect)
	}
}

func (a *Auth) RequireLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, AUTH_SESSION)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, ok := session.Values["userinfo"]; !ok {
			http.Redirect(w, r, LOGIN_PATH, http.StatusSeeOther)
		} else {
			next(w, r)
		}
	}
}

func (a *Auth) Mux(loggedInPath string) *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc(LOGOUT_PATH, a.logoutHandler())
	m.HandleFunc(LOGIN_PATH, a.loginHandler())
	m.HandleFunc(CALLBACK_PATH, a.callbackHandler(loggedInPath))
	return m
}

func FuncMap() template.FuncMap {
	m := make(map[string]interface{})
	for k, v := range sprig.GenericFuncMap() {
		m[k] = v
	}
	return m
}

func ServeTemplate(name string, sessVal string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, AUTH_SESSION)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bits, err := ioutil.ReadFile(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bitstring := string(bits)
		if strings.Contains(bitstring, "{{") {
			templ, err := template.New(name).Funcs(FuncMap()).Parse(string(bits))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = templ.Execute(w, session.Values[sessVal])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		io.WriteString(w, bitstring)
		return
	}

}
