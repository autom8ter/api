package api

import (
	"context"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/autom8ter/objectify"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"html/template"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func init() {
	Context = context.WithValue(context.TODO(), "env", os.Environ())
	Util = objectify.Default()
	httpClient = http.DefaultClient

}

var (
	Util                   *objectify.Handler
	Context                context.Context
	httpClient             *http.Client
	store                  *sessions.CookieStore
	AUTH_SESSION           = "auth-session"
	DEFAULT_OAUTH_REDIRECT = "http://localhost:8080/callback"
	DEFAULT_OAUTH_SCOPES   = []Scope{Scope_OPENID, Scope_PROFILE, Scope_EMAIL}
)

type ClientSet struct {
	Utility  UtilityServiceClient
	Contact  ContactServiceClient
	Payment  PaymentServiceClient
	Resource ResourceServiceClient
}

func NewClientSet(conn *grpc.ClientConn) *ClientSet {
	return &ClientSet{
		Utility:  NewUtilityServiceClient(conn),
		Contact:  NewContactServiceClient(conn),
		Resource: NewResourceServiceClient(conn),
	}
}

func ENVFromContext(ctx context.Context) []string {
	return ctx.Value("env").([]string)
}

func SearchUSPhoneNumbersURL(account string) string {
	return fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/AvailablePhoneNumbers/US/Local.json", account)
}

func IncomingPhoneNumbersURL(account string) string {
	return fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/IncomingPhoneNumbers.json", account)
}

func ChatServiceURL() string {
	return "https://chat.twilio.com/v2/Services"
}

func (s Scope) Normalize() string {
	switch s {
	case Scope_PROFILE:
		return "profile"
	case Scope_EMAIL:
		return "email"
	case Scope_OPENID:
		return "openid"
	case Scope_READ_USERS:
		return "read:users"
	case Scope_READ_USER_IDP_TOKENS:
		return "read:user_idp_tokens"
	case Scope_CREATE_USERS:
		return "create:users"
	case Scope_READ_STATS:
		return "read:stats"
	case Scope_CREATE_EMAIL_TEMPLATES:
		return "create:email_templates"
	case Scope_UPDATE_EMAIL_TEMPLATES:
		return "update:email_templates"
	case Scope_READ_EMAIL_TEMPLATES:
		return "read:email_templates"
	case Scope_CREATE_RULES:
		return "create:rules"
	case Scope_DELETE_RULES:
		return "delete:rules"
	case Scope_UPDATE_RULES:
		return "update:rules"
	case Scope_READ_RULES:
		return "read:rules"
	case Scope_CREATE_ROLES:
		return "create:rules"
	case Scope_DELETE_ROLES:
		return "delete:rules"
	case Scope_UPDATE_ROLES:
		return "update:rules"
	case Scope_READ_ROLES:
		return "read:rules"
	case Scope_READ_LOGS:
		return "read:logs"
	default:
		return ""
	}
}

func NormalizeScopes(scopes ...Scope) []string {
	s := []string{}
	for _, scope := range scopes {
		s = append(s, scope.Normalize())
	}
	return s
}

func FuncMap() template.FuncMap {
	m := make(map[string]interface{})
	for k, v := range sprig.GenericFuncMap() {
		m[k] = v
	}
	return m
}

func NewTemplateFromFile(filename string) (*Template, error) {
	bits, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Template{
		Text: string(bits),
		Name: filename,
	}, nil
}

func (m *Template) IsTemplate() bool {
	return strings.Contains(m.Text, "{{")
}

func (m *Template) AsTemplate() (*template.Template, error) {
	return (template.New(m.Name).Funcs(FuncMap())).Parse(m.Text)
}

func (m *Template) RenderBytes(w io.Writer, bits *Bytes) error {
	templ, err := template.New(m.Name).Funcs(FuncMap()).Parse(m.Text)
	if err != nil {
		return err
	}
	return templ.Execute(w, bits.Bits)
}

func (m *Template) Render(w io.Writer, data interface{}) error {
	templ, err := template.New(m.Name).Funcs(FuncMap()).Parse(m.Text)
	if err != nil {
		return err
	}
	return templ.Execute(w, data)
}

func (t *Template) RenderUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, AUTH_SESSION)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if t.IsTemplate() {
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			obj := session.Values["user"]
			bits, err := json.Marshal(obj)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			u := &User{}
			err = json.Unmarshal(bits, u)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = t.RenderBytes(w, AsBytes(u))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		io.WriteString(w, t.Text)
		return
	}
}

func RenderFile(name string, data []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bits, err := ioutil.ReadFile(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bitstring := string(bits)
		if strings.Contains(bitstring, "{{") {
			templ, err := template.New("").Funcs(FuncMap()).Parse(string(bits))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = templ.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		_, err = io.WriteString(w, bitstring)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

}

func WriteFile(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bits, err := ioutil.ReadFile(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bitstring := string(bits)
		_, err = io.WriteString(w, bitstring)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

}

func (h HTTPMethod) Normalize() string {
	if h == HTTPMethod_POST {
		return "POST"
	} else {
		return "GET"
	}
}

func (h *HTTPRequest) Do(token *Token) (*Bytes, error) {
	u, err := url.Parse(h.Url)
	if err != nil {
		return nil, err
	}
	r := &http.Request{
		Method: h.Method.Normalize(),
		URL:    u,
	}
	if token.TokenType == "" {
		token.TokenType = "Bearer"
	}
	if token.AccessToken != "" {
		r.Header.Set("Authorization", token.TokenType+" "+token.AccessToken)
	}

	if len(h.Form) != 0 {
		for k, v := range h.Form {
			r.Form.Set(k, v)
		}
	}
	resp, err := httpClient.Do(r)
	if err != nil {
		return nil, Util.WrapErr(err, resp.Status)
	}
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, Util.WrapErr(err, "reading http response body")
	}
	return &Bytes{
		Bits: bits,
	}, nil

}

func NewBytes() *Bytes {
	return &Bytes{
		Bits: []byte{},
	}
}

func BytesFromString(str string) *Bytes {
	return &Bytes{
		Bits: []byte(str),
	}
}

func BytesFromBytes(bits []byte) *Bytes {
	return &Bytes{
		Bits: bits,
	}
}

func BytesFromReader(r io.Reader) (*Bytes, error) {
	b := NewBytes()
	_, err := io.Copy(b, r)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func BytesFromFile(fileName string) (*Bytes, error) {
	bits, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Bytes{
		Bits: bits,
	}, nil
}

func (m *Bytes) Read(p []byte) (n int, err error) {
	before := len(m.Bits)
	m.Bits = append(m.Bits, p...)
	after := len(m.Bits)
	return after - before, nil
}

func (m *Bytes) Write(p []byte) (n int, err error) {
	before := len(m.Bits)
	m.Bits = append(m.Bits, p...)
	after := len(m.Bits)
	return after - before, nil
}

func (m *Bytes) YAML() []byte {
	return Util.MarshalYAML(m.Bits)
}

func (m *Bytes) JSON() []byte {
	return Util.MarshalJSON(m.Bits)
}

func (m *Bytes) XML() []byte {
	return Util.MarshalXML(m.Bits)
}

func (m *Bytes) Length() int {
	return len(m.Bits)
}

func (m *Bytes) Compile(w io.Writer, t *Template) error {
	return t.Render(w, m)
}

func (m *Bytes) Contains(str string) bool {
	return strings.Contains(string(m.Bits), str)
}

func (m *Bytes) BitString() string {
	return string(m.Bits)
}

func (m *Bytes) CompileHTTP(t *Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := m.Compile(w, t); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (m *Bytes) WriteString() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, m.BitString())
		return
	}
}

func (m *Bytes) UnMarshalJSON(obj interface{}) error {
	return json.Unmarshal(m.JSON(), obj)
}

func (m *Bytes) UnMarshalProto(obj interface{}) error {
	return proto.Unmarshal(m.Bits, obj.(proto.Message))
}

func (m *Bytes) Clear() {
	m.Bits = []byte{}
}

func AsBytes(obj interface{}) *Bytes {
	return &Bytes{
		Bits: Util.MarshalJSON(obj),
	}
}

type RouterPaths struct {
	HomePath     string
	LoggedInPath string
	LoginPath    string
	LogoutPath   string
	CallbackPath string
	HomeURL      string
}

func (a *Auth) defaultIfEmpty() {
	if len(a.Scopes) == 0 {
		a.Scopes = DEFAULT_OAUTH_SCOPES
	}
	if a.Redirect == "" {
		a.Redirect = DEFAULT_OAUTH_REDIRECT
	}
}
func (a *Auth) validate() error {
	if len(a.Scopes) == 0 {
		return Util.NewError("validation error: empty scope")
	}
	if a.Redirect == "" {
		return Util.NewError("validation error: empty redirect")
	}
	if a.ClientId == "" {
		return Util.NewError("validation error: empty client id")
	}
	if a.Domain == "" {
		return Util.NewError("validation error: empty domain")
	}
	if a.ClientSecret == "" {
		return Util.NewError("validation error: empty client secret")
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

		conf := a.OAuth2Config()

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
		resp, err := client.Get(URL_USER_INFOURL.Normalize(a.Domain))
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
		sesh, err := GetAuthSession(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		TokenAuthSessionValues(sesh, token)
		AuthSessionValues(sesh, "user", userinfo)
		SaveSession(w, r)
		// Redirect to logged in page
		http.Redirect(w, r, c.LoggedInPath, http.StatusSeeOther)
	}

}

func TokenAuthSessionValues(session *sessions.Session, token *oauth2.Token) {
	session.Values["id_token"] = token.Extra("id_token")
	session.Values["refresh_token"] = token.RefreshToken
	session.Values["access_token"] = token.AccessToken
	session.Values["token_type"] = token.TokenType
	session.Values["expiry"] = token.Expiry.String()
}

func AuthSessionValues(session *sessions.Session, key string, data map[string]interface{}) {
	session.Values[key] = data
}

func TokenFromAuthSession(session *sessions.Session) (*Token, error) {
	return &Token{
		AccessToken:  session.Values["access_token"].(string),
		TokenType:    session.Values["token_type"].(string),
		RefreshToken: session.Values["refresh_token"].(string),
		Expiry:       session.Values["expiry"].(string),
		IdToken:      session.Values["id_token"].(string),
	}, nil

}
func TokenFromOAuthToken(tok *oauth2.Token) *Token {
	return &Token{
		AccessToken:  tok.AccessToken,
		TokenType:    tok.TokenType,
		RefreshToken: tok.RefreshToken,
		Expiry:       tok.Expiry.String(),
		IdToken:      tok.Extra("id_token").(string),
	}
}

func SaveSession(w http.ResponseWriter, r *http.Request) {
	if err := sessions.Save(r, w); err != nil {
		http.Error(w, Util.WrapErr(err, "saving session").Error(), http.StatusInternalServerError)
	}
}

func UserFromSession(session *sessions.Session) (*User, error) {
	return session.Values["user"].(*User), nil
}

func FromSession(key string, session *sessions.Session) interface{} {
	return session.Values[key]
}

func GetAuthSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, AUTH_SESSION)
}

func (a *Auth) loginHandler(audience string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.defaultIfEmpty()
		if err := a.validate(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		conf := a.OAuth2Config()

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

		audience := oauth2.SetAuthURLParam("audience", URL_USER_INFOURL.Normalize(a.Domain))
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

func (a *Auth) ToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "auth", a)
}

func AuthFromContext(ctx context.Context) *Auth {
	return ctx.Value("auth").(*Auth)
}

func TokenFromContext(ctx context.Context) *Token {
	return ctx.Value("token").(*Token)
}

func (t *Token) ToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "token", t)
}

func (a *Auth) RequireLogin(paths *RouterPaths, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, AUTH_SESSION)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, ok := session.Values["access_token"]; !ok {
			http.Redirect(w, r, paths.LoginPath, http.StatusSeeOther)
		} else {
			next(w, r)
		}
	}
}

func (a *Auth) OAuth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     a.ClientId,
		ClientSecret: a.ClientSecret,
		RedirectURL:  a.Redirect,
		Scopes:       NormalizeScopes(a.Scopes...),
		Endpoint: oauth2.Endpoint{
			AuthURL:  URL_AUTHORIZEURL.Normalize(a.Domain),
			TokenURL: URL_TOKENURL.Normalize(a.Domain),
		},
	}
}

func (a *Auth) Router(c *RouterPaths, home, loggedIn http.HandlerFunc) *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc(c.LogoutPath, a.logoutHandler(c))
	m.HandleFunc(c.LoginPath, a.loginHandler(URL_USER_INFOURL.Normalize(a.Domain)))
	m.HandleFunc(c.CallbackPath, a.callbackHandler(c))
	m.HandleFunc(c.HomePath, home)
	m.HandleFunc(c.LoggedInPath, a.RequireLogin(c, loggedIn))
	return m
}

func (a *Auth) ListenAndServe(addr string, c *RouterPaths, home, loggedIn http.HandlerFunc) error {
	return http.ListenAndServe(addr, a.Router(c, home, loggedIn))
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

func (u URL) Normalize(domain string) string {
	switch u {
	case URL_TOKENURL:
		return tokenURL(domain)
	case URL_USER_INFOURL:
		return userInfoURL(domain)
	case URL_USERSURL:
		return usersURL(domain)
	case URL_AUTHORIZEURL:
		return authURL(domain)
	case URL_SEARCH_USERSURL:
		return searchUsersURL(domain)
	case URL_ROLESURL:
		return rolesURL(domain)
	case URL_LOGSURL:
		return logsURL(domain)
	case URL_GRANTSURL:
		return grantsURL(domain)
	case URL_STATSURL:
		return statsURL(domain)
	case URL_CLIENTSURL:
		return clientsURL(domain)
	case URL_CONNECTIONSURL:
		return connectionsURL(domain)
	case URL_RULESURL:
		return rulesURL(domain)
	case URL_TENANTSURL:
		return tenantsURL(domain)
	case URL_JWKSURL:
		return jWKSURL(domain)
	case URL_DEVICEURL:
		return deviceCredentials(domain)
	case URL_EMAILURL:
		return emailURL(domain)
	case URL_EMAIL_TEMPLATEURL:
		return emailTemplateURL(domain)
	default:
		return ""
	}
}

func tokenURL(domain string) string {
	return "https://" + domain + "/oauth/token"
}

func userInfoURL(domain string) string {
	return "https://" + domain + "/userinfo"
}

func usersURL(domain string) string {
	return "https://" + domain + "/api/v2/users"
}

func authURL(domain string) string {
	return "https://" + domain + "/authorize"
}

func searchUsersURL(domain string) string {
	return "https://" + domain + "/api/v2/users-by-email"
}

func rolesURL(domain string) string {
	return "https://" + domain + "/api/v2/roles"
}

func logsURL(domain string) string {
	return "https://" + domain + "/api/v2/logs"
}

func grantsURL(domain string) string {
	return "https://" + domain + "/api/v2/grants"
}

func statsURL(domain string) string {
	return "https://" + domain + "/api/v2/stats/active-users"
}

func clientsURL(domain string) string {

	return "https://" + domain + "/api/v2/clients"
}

func clientGrantsURL(domain string) string {
	return "https://" + domain + "/api/v2/client-grants"
}

func connectionsURL(domain string) string {
	return "https://" + domain + "/api/v2/connections"
}

func customDomainsURL(domain string) string {
	return "https://" + domain + "/api/v2/custom-domains"
}

func rulesURL(domain string) string {
	return "https://" + domain + "/api/v2/rules"
}

func tenantsURL(domain string) string {
	return "https://" + domain + "/api/v2/tenants/settings"
}

func emailURL(domain string) string {
	return "https://" + domain + "/api/v2/emails/provider"
}

func emailTemplateURL(domain string) string {
	return "https://" + domain + "/api/v2/email-templates"
}

func deviceCredentials(domain string) string {
	return "https://" + domain + "/api/v2/device-credentials"
}

func jWKSURL(domain string) string {
	return "https://" + domain + "/.well-known/jwks.json"
}

func (b *Bytes) UnmarshalProto(obj interface{}) error {
	return proto.Unmarshal(b.Bits, obj.(proto.Message))
}

func (b *Bytes) UnmarshalJSON(obj interface{}) error {
	return json.Unmarshal(b.Bits, obj)
}

func (c *Jwks) TokenCert(token *jwt.Token) (string, error) {
	var cert string
	for k, _ := range c.Keys {
		if token.Header["kid"] == c.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + c.Keys[k].X5C[0] + "\n-----END CERTIFICATE-----"
		}
	}
	if cert == "" {
		err := Util.NewError("Unable to find appropriate key.")
		return cert, err
	}
	return cert, nil
}

func SecretFromEnv() *Secret {
	return &Secret{
		Text: os.Getenv("SECRET"),
	}
}

func (s *Secret) InitSessions() error {
	store = sessions.NewCookieStore([]byte(s.Text))
	gob.Register(map[string]interface{}{})
	return nil
}

func (a *Auth) Token(ctx context.Context, code string) (*Token, error) {
	token, err := a.OAuth2Config().Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return TokenFromOAuthToken(token), nil
}

func (t *Token) ResourceRequest(domain string, method HTTPMethod, u URL, form map[string]string, body *Bytes) *ResourceRequest {
	return &ResourceRequest{
		Token:  t,
		Method: method,
		Domain: domain,
		Url:    u,
		Form:   form,
		Body:   body,
	}
}

func (r *ResourceRequest) Do() (*Bytes, error) {
	c := &HTTPRequest{
		Method: r.Method,
		Url:    r.Url.Normalize(r.Domain),
		Form:   r.Form,
		Body:   r.Body,
	}
	return c.Do(r.Token)
}
