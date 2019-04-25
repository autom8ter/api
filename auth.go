//go:generate godocdown -o README.md

package api

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
)

type RouterPaths struct {
	HomePath     string
	LoggedInPath string
	LoginPath    string
	LogoutPath   string
	CallbackPath string
	HomeURL      string
}

var DEFAULT_OAUTH_REDIRECT = "http://localhost:8080/callback"
var DEFAULT_OAUTH_SCOPES = []Scope{Scope_OPENID, Scope_PROFILE, Scope_EMAIL}

func (a *Auth) defaultIfEmpty() {
	if len(a.Scopes) == 0 {
		a.Scopes = DEFAULT_OAUTH_SCOPES
	}
	if a.Audience == "" {
		a.Audience = a.UserInfoURL()
	}
	if a.Redirect == "" {
		a.Redirect = DEFAULT_OAUTH_REDIRECT
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

func (a *Auth) callbackHandler(c *RouterPaths) http.HandlerFunc {
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
			Scopes:       NormalizeScopes(a.Scopes...),
			Endpoint: oauth2.Endpoint{
				AuthURL:  a.AuthURL(),
				TokenURL: a.TokenURL(),
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
		session.Values["user"] = userinfo
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to logged in page
		http.Redirect(w, r, c.LoggedInPath, http.StatusSeeOther)
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
			Scopes:       NormalizeScopes(a.Scopes...),
			Endpoint: oauth2.Endpoint{
				AuthURL:  a.AuthURL(),
				TokenURL: a.TokenURL(),
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

func (a *Auth) logoutHandler(paths *RouterPaths) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.defaultIfEmpty()
		if err := a.validate(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		u, err := a.LogoutURL(paths)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	}
}

func (a *Auth) RequireLogin(paths *RouterPaths, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, AUTH_SESSION)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, ok := session.Values["user"]; !ok {
			http.Redirect(w, r, paths.LoginPath, http.StatusSeeOther)
		} else {
			next(w, r)
		}
	}
}

func (a *Auth) Router(c *RouterPaths, home, loggedIn http.HandlerFunc) *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc(c.LogoutPath, a.logoutHandler(c))
	m.HandleFunc(c.LoginPath, a.loginHandler())
	m.HandleFunc(c.CallbackPath, a.callbackHandler(c))
	m.HandleFunc(c.HomePath, home)
	m.HandleFunc(c.LoggedInPath, a.RequireLogin(c, loggedIn))
	return m
}

func (a *Auth) ListenAndServe(addr string, c *RouterPaths, home, loggedIn http.HandlerFunc) error {
	return http.ListenAndServe(addr, a.Router(c, home, loggedIn))
}

func (c *Auth) AuthURL() string {
	return "https://" + c.Domain + "/authorize"
}

func (c *Auth) LogoutURL(r *RouterPaths) (string, error) {
	var Url *url.URL
	Url, err := url.Parse("https://" + c.Domain)
	if err != nil {
		return "", err
	}

	Url.Path += "/v2/logout"
	parameters := url.Values{}
	parameters.Add("returnTo", r.HomeURL)
	parameters.Add("client_id", c.ClientId)
	Url.RawQuery = parameters.Encode()

	return Url.String(), nil
}

func (c *Auth) TokenURL() string {
	return "https://" + c.Domain + "/oauth/token"
}

func (c *Auth) UserInfoURL() string {
	return "https://" + c.Domain + "/userinfo"
}

func (c *Auth) UsersURL() string {
	return "https://" + c.Domain + "/api/v2/users"
}

func (c *Auth) SearchUsersURL() string {
	return "https://" + c.Domain + "/api/v2/users-by-email"
}

func (c *Auth) RolesURL() string {
	return "https://" + c.Domain + "/api/v2/roles"
}

func (c *Auth) LogsURL() string {
	return "https://" + c.Domain + "/api/v2/logs"
}

func (c *Auth) GrantsURL() string {
	return "https://" + c.Domain + "/api/v2/grants"
}

func (c *Auth) StatsURL() string {
	return "https://" + c.Domain + "/api/v2/stats/active-users"
}

func (c *Auth) ClientsURL() string {
	return "https://" + c.Domain + "/api/v2/clients"
}

func (c *Auth) ClientGrantsURL() string {
	return "https://" + c.Domain + "/api/v2/client-grants"
}

func (c *Auth) ConnectionsURL() string {
	return "https://" + c.Domain + "/api/v2/connections"
}

func (c *Auth) CustomDomainsURL() string {
	return "https://" + c.Domain + "/api/v2/custom-domains"
}

func (c *Auth) RulesURL() string {
	return "https://" + c.Domain + "/api/v2/rules"
}

func (c *Auth) TenantURL() string {
	return "https://" + c.Domain + "/api/v2/tenants/settings"
}

func (c *Auth) EmailURL() string {
	return "https://" + c.Domain + "/api/v2/emails/provider"
}

func (c *Auth) EmailTemplateURL(name string) string {
	return "https://" + c.Domain + "/api/v2/email-templates/" + name
}

func (c *Auth) DeviceCredentials(name string) string {
	return "https://" + c.Domain + "/api/v2/device-credentials"
}

func (c *Auth) JWKSURL() string {
	return "https://" + c.Domain + "/.well-known/jwks.json"
}

func (c *Auth) GetJWKS() (*Jwks, error) {
	resp, err := HTTPClient.Get(c.JWKSURL())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var jwks = &Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return nil, err
	}
	return jwks, nil
}

func (c *Jwks) GetCert(token *jwt.Token) (string, error) {
	var cert string
	for k, _ := range c.Keys {
		if token.Header["kid"] == c.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + c.Keys[k].X5C[0] + "\n-----END CERTIFICATE-----"
		}
	}
	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}
	return cert, nil
}
