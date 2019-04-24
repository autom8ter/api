package api

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
)

func (c *Config) CallbackHandler(loggedIn string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(c.Scopes) == 0 {
			c.Scopes = []string{"openid", "profile", "email"}
		}
		conf := &oauth2.Config{
			ClientID:     c.ClientId,
			ClientSecret: c.ClientSecret,
			RedirectURL:  c.Redirect,
			Scopes:       c.Scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://" + c.Domain + "/authorize",
				TokenURL: "https://" + c.Domain + "/oauth/token",
			},
		}

		if c.Audience == "" {
			c.Audience = "https://" + c.Domain + "/userinfo"
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
		resp, err := client.Get(c.Audience)
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
		http.Redirect(w, r, loggedIn, http.StatusSeeOther)
	}

}

func (c *Config) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(c.Scopes) == 0 {
			c.Scopes = []string{"openid", "profile", "email"}
		}
		conf := &oauth2.Config{
			ClientID:     c.ClientId,
			ClientSecret: c.ClientSecret,
			RedirectURL:  c.Redirect,
			Scopes:       c.Scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://" + c.Domain + "/authorize",
				TokenURL: "https://" + c.Domain + "/oauth/token",
			},
		}

		if c.Audience == "" {
			c.Audience = "https://" + c.Domain + "/userinfo"
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

		audience := oauth2.SetAuthURLParam("audience", c.Audience)
		url := conf.AuthCodeURL(state, audience)

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func (c *Config) LogoutHandler(returnTo string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Url *url.URL
		Url, err := url.Parse("https://" + c.Domain)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Url.Path += "/v2/logout"
		parameters := url.Values{}
		parameters.Add("returnTo", returnTo)
		parameters.Add("client_id", c.ClientId)
		Url.RawQuery = parameters.Encode()

		http.Redirect(w, r, Url.String(), http.StatusTemporaryRedirect)
	}
}
