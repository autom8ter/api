package api

import (
	"encoding/json"
	oauth3 "github.com/autom8ter/oauth2"
	"github.com/autom8ter/oauth2/callback"
	"github.com/autom8ter/objectify"
	"github.com/pkg/errors"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/auth0/go-jwt-middleware"
	"github.com/autom8ter/oauth2/middleware"

	sessions2 "github.com/autom8ter/oauth2/sessions"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
)

func init() {
	if strings.Contains(os.Getenv("DEBUG"), "t") {
		Debug = true
	}
	if strings.Contains(os.Getenv("DEBUG"), "T") {
		Debug = true
	}
}

var Util = objectify.Default()
var Debug bool
var AUTH_SESSION_NAME = "autom8ter"

func (p *UserInfo) Validate() error {
	if p.Name == "" {
		return errors.New("Validate: UserInfo requires a name")
	}
	return Util.Validate(p)
}

func (p *UserInfo) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *UserInfo) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}

func (p *UserMetadata) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *UserMetadata) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}

func (p *AppMetadata) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *AppMetadata) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}

func (p *Tokens) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *Tokens) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}

func (p *Paths) JSONString() string {
	return string(Util.MarshalJSON(p))
}

func (p *Paths) Render(tmpl *template.Template, w io.Writer) error {
	return Util.RenderHTML(tmpl, p, w)
}

func DefaultPaths() *Paths {
	return &Paths{
		Home:              "/",
		Dashboard:         "/dashboard",
		Logout:            "/logout",
		Callback:          "/callback",
		Login:             "/login",
		LoggedOutReturnTo: "http://localhost:8080",
	}
}

func (a *Auth0) AsMux(paths *Paths) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(paths.Callback, a.HandleOAuthCallback(paths))
	mux.HandleFunc(paths.Logout, a.HandleOAuthLogout(paths))
	mux.HandleFunc(paths.Login, a.HandleOAuthLogin())
	return mux
}

func NewAuth0(domain string, clientID string, clientSecret string, redirectURL string, resourceURL string, scopes ...string) (*Auth0, error) {
	a := &Auth0{
		Domain:       domain,
		ClientId:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Redirect:     redirectURL,
		ResourceUrl:  resourceURL,
	}
	if a.Domain == "" {
		return nil, errors.New("empty domain")
	}
	if a.ClientId == "" {
		return nil, errors.New("empty client id")
	}
	if a.ClientSecret == "" {
		return nil, errors.New("empty client secret")
	}
	if a.Redirect == "" {
		return nil, errors.New("empty redirect")
	}
	if a.Scopes == nil {
		a.Scopes = []string{"userid", "profile", "email"}
	}
	if a.ResourceUrl == "" {
		a.ResourceUrl = oauth3.UserInfoEndpoint(a.Domain)
	}
	return a, nil
}

func (o *Auth0) OAuthConfig() *oauth2.Config {
	return oauth3.NewOauth2Config(o.ClientId, o.ClientSecret, o.Domain, o.Redirect, o.Scopes...)
}

func (o *Auth0) Validate() {
	if err := Util.Validate(o); err != nil {
		Util.Entry().Fatalln(err.Error())
	}
}

func (o *Auth0) OAuthLoginURL(randomstate string) string {
	return o.OAuthConfig().AuthCodeURL(randomstate, oauth2.SetAuthURLParam("audience", o.ResourceUrl))
}

func (o *Auth0) HandleOAuthCallback(p *Paths) http.HandlerFunc {
	return callback.HandleOAuthCallback(o.OAuthConfig(), AUTH_SESSION_NAME, oauth3.UserInfoEndpoint(o.Domain), p.Dashboard)
}

func (o *Auth0) HandleOAuthLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Generate random state
		state := Util.RandomString(32)

		session, err := sessions2.Session.Get(r, "state")
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
		http.Redirect(w, r, o.OAuthLoginURL(state), http.StatusTemporaryRedirect)
	}
}

func (o *Auth0) HandleOAuthLogout(p *Paths) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := oauth3.LogoutEndpoint(o.ClientId, o.Domain, p.LoggedOutReturnTo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		Debugf("logging out and returning to: %s\n", p.LoggedOutReturnTo)
		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	}
}

func (o *Auth0) GetUserInfo(r *http.Request) (*UserInfo, error) {
	obj, err := sessions2.GetResourceFromSession(oauth3.UserInfoEndpoint(o.Domain), AUTH_SESSION_NAME, r)
	if err != nil {
		return nil, err
	}
	bits := Util.MarshalJSON(obj)
	Debugf("got resource from session: %s", bits)
	if len(bits) == 0 {
		Fatalf("failed to marshal resource from session")
	}
	p := &UserInfo{}
	err = json.Unmarshal(bits, p)
	if err != nil {
		return p, err
	}
	Debugf("got user info from resource: %s", p)

	return p, nil
}

func (o *Auth0) GetTokens(r *http.Request) (*Tokens, error) {
	access, err := sessions2.GetAccessTokenFromSession(AUTH_SESSION_NAME, r)
	if err != nil {
		return nil, err
	}
	id, err := sessions2.GetIDTokenFromSession(AUTH_SESSION_NAME, r)
	if err != nil {
		return nil, err
	}
	ref, err := sessions2.GetRefreshTokenFromSession(AUTH_SESSION_NAME, r)
	if err != nil {
		return nil, err
	}
	return &Tokens{
		Access: &AccessToken{
			Token: access,
		},
		Id: &IDToken{
			Token: id,
		},
		Refresh: &RefreshToken{
			Token: ref,
		},
	}, nil
}

func (o *Auth0) SecureFunc(redirect string, handler http.HandlerFunc) http.HandlerFunc {
	return o.Secure(redirect, handler)
}

func (o *Auth0) Secure(redirect string, handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prof, err := o.GetUserInfo(r)
		Debugf("got profile: %s", Util.MarshalJSON(prof))
		if err != nil {
			Debugf("redirect: error")
			http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
		}
		if prof == nil {
			Debugf("redirect: empty profile")
			http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
		}
		if prof.Name == "" {
			Debugf("redirect: empty name")
			http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
		}
		if err := Util.Validate(prof); err != nil {
			Debugf("redirect: validation error")
			http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
		}
		handler.ServeHTTP(w, r)
	}
}

func (a *Auth0) SecureJWTFunc(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.Auth0Middleware().Handler(handler).ServeHTTP(w, r)
	}
}

func (a *Auth0) SecureJWT(handler http.Handler) http.Handler {
	return a.Auth0Middleware().Handler(handler)
}

func (a *Auth0) SecureJWTWithRedirectFunc(redirect string, handler http.HandlerFunc) http.HandlerFunc {
	middle := a.Auth0MiddlewareWithRedirect(redirect)
	return func(w http.ResponseWriter, r *http.Request) {
		middle.Handler(handler).ServeHTTP(w, r)
	}
}

func (a *Auth0) SecureJWTWithRedirect(redirect string, handler http.Handler) http.Handler {
	middle := a.Auth0MiddlewareWithRedirect(redirect)
	return middle.Handler(handler)
}

func (a *Auth0) Auth0Middleware() *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'aud' claim
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(a.ResourceUrl, false)
			if !checkAud {
				return token, errors.New("Invalid audience.")
			}
			// Verify 'iss' claim
			iss := "https://" + a.Domain + "/"
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer.")
			}

			cert, err := oauth3.GetPemCert(a.Domain, token)
			if err != nil {
				return "", err
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		Debug:         Debug,
		SigningMethod: jwt.SigningMethodRS256,
	})
}

func (a *Auth0) Auth0MiddlewareWithRedirect(redirect string) *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'aud' claim
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(a.ResourceUrl, false)
			if !checkAud {
				return token, errors.New("Invalid audience.")
			}
			// Verify 'iss' claim
			iss := "https://" + a.Domain + "/"
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer.")
			}

			cert, err := oauth3.GetPemCert(a.Domain, token)
			if err != nil {
				panic(err.Error())
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		Debug:         Debug,
		SigningMethod: jwt.SigningMethodRS256,
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
		},
	})
}

func (o *Auth0) TokenHasScope(scope string, token string) bool {
	return middleware.TokenHasScope(scope, token)
}

func (o *Auth0) TokenHasAudience(scope string, token string) bool {
	return middleware.TokenHasAudience(scope, token)
}

func Debugf(format string, args ...interface{}) {
	if Debug {
		Util.Entry().Debugf(format, args...)
	}
}

func Fatalf(format string, args ...interface{}) {
	if Debug {
		Util.Entry().Fatalf(format, args...)
	}
}

func Warnf(format string, args ...interface{}) {
	if Debug {
		Util.Entry().Warnf(format, args...)
	}
}
