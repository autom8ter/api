//go:generate godocdown -o README.md

package common

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/objectify"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/google"
	ojwt "golang.org/x/oauth2/jwt"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

func init() {
	util = objectify.Default()
	if ENVContext == nil {
		ENVContext = context.WithValue(context.TODO(), "env", os.Environ())
	}
}

var (
	util       *objectify.Handler
	ENVContext context.Context
)

func (p *JSONWebKeys) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *JSONWebKeys) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *Jwks) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (c *Jwks) TokenCert(token *jwt.Token) (string, error) {
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

func (s *JWT) Validate(fn func(a *JWT) error) error {
	return fn(s)
}

func (s *OAuth2) Validate(fn func(a *OAuth2) error) error {
	return fn(s)
}

func (s *ClientCredentials) Validate(fn func(a *ClientCredentials) error) error {
	return fn(s)
}

func (c *OAuth2) Exchange() (*oauth2.Token, error) {
	a := c.Config()
	tok, err := a.Exchange(context.TODO(), c.Code)
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func (c *OAuth2) Token() (*oauth2.Token, error) {
	return c.Exchange()
}

func (c *OAuth2) Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.AuthUrl,
			TokenURL: c.TokenUrl,
		},
		RedirectURL: c.Redirect,
		Scopes:      c.Scopes,
	}
}

func (c *JWT) Config() *ojwt.Config {
	d, _ := ptypes.Timestamp(c.Expires)
	return &ojwt.Config{
		Email:        c.Email,
		PrivateKey:   c.PrivateKey,
		PrivateKeyID: c.PriveKeyId,
		Subject:      c.Subject,
		Scopes:       c.Scopes,
		TokenURL:     c.TokenUrl,
		Expires:      d.Sub(time.Now()),
	}
}

func (c *JWT) Token() (*oauth2.Token, error) {
	a := c.Config()

	src := a.TokenSource(context.TODO())
	tok, err := src.Token()
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func (c *ClientCredentials) Config() *clientcredentials.Config {
	rl := url.Values{}
	for k, v := range c.EndpointParams {
		rl.Set(k, v)
	}
	return &clientcredentials.Config{
		ClientID:       c.ClientId,
		ClientSecret:   c.ClientSecret,
		TokenURL:       c.TokenUrl,
		Scopes:         c.Scopes,
		EndpointParams: rl,
	}
}

func (c *ClientCredentials) Token() (*oauth2.Token, error) {
	a := c.Config()
	return a.Token(context.TODO())
}

func (c *ClientCredentials) PerRPCCredentials() (credentials.PerRPCCredentials, error) {
	tok, err := c.Token()
	if err != nil {
		return nil, err
	}
	return oauth.NewOauthAccess(tok), nil
}

func (c *JWT) PerRPCCredentials() (credentials.PerRPCCredentials, error) {
	tok, err := c.Token()
	if err != nil {
		return nil, err
	}
	return oauth.NewOauthAccess(tok), nil
}

func (c *OAuth2) PerRPCCredentials() (credentials.PerRPCCredentials, error) {
	tok, err := c.Token()
	if err != nil {
		return nil, err
	}
	return oauth.NewOauthAccess(tok), nil
}

func (c *ClientCredentials) Client(ctx context.Context) *http.Client {
	return oauth2.NewClient(ctx, c)
}

func (c *OAuth2) Client(ctx context.Context) *http.Client {
	return oauth2.NewClient(ctx, c)
}

func (c *JWT) Client(ctx context.Context) *http.Client {
	return oauth2.NewClient(ctx, c)
}

func (c *OAuth2) AuthCodeURL(state string, audience string) string {

	aud := oauth2.SetAuthURLParam("audience", audience)

	return c.Config().AuthCodeURL(state, aud)
}

func (c *OAuth2) SetCode(r *http.Request) {
	c.Code = r.URL.Query().Get("code")
}

func NewClientCredentials(tokenURL, clientID, clientSecret string, params map[string]string, scopes []string) *ClientCredentials {
	return &ClientCredentials{
		ClientId:       clientID,
		ClientSecret:   clientSecret,
		TokenUrl:       tokenURL,
		Scopes:         scopes,
		EndpointParams: params,
	}
}

func NewJWT(tokenURL, email, privateKey, privKeyID, subject string, audience string, timer time.Time, scopes []string) *JWT {
	d, _ := ptypes.TimestampProto(timer)

	return &JWT{
		Email:      email,
		PrivateKey: []byte(privateKey),
		PriveKeyId: privKeyID,
		Subject:    subject,
		Scopes:     scopes,
		TokenUrl:   tokenURL,
		Expires:    d,
		Audience:   audience,
	}
}

func NewOAuth2(clientID string, clientSecret string, tokenURL string, authURL string, redirect string, scopes []string) *OAuth2 {
	return &OAuth2{
		ClientId:     clientID,
		ClientSecret: clientSecret,
		TokenUrl:     tokenURL,
		AuthUrl:      authURL,
		Scopes:       scopes,
		Redirect:     redirect,
	}
}

func NewGoogleDefaultCredentials(scopes []string) *DefaultGCPCredentials {
	return &DefaultGCPCredentials{
		Scopes: scopes,
	}
}

func (d *DefaultGCPCredentials) FindCredentials() (*google.Credentials, error) {
	return google.FindDefaultCredentials(context.TODO(), d.Scopes...)
}

func (d *DefaultGCPCredentials) Token() (*oauth2.Token, error) {
	creds, err := d.FindCredentials()
	if err != nil {
		return nil, err
	}
	t, err := creds.TokenSource.Token()
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *DefaultGCPCredentials) Client(ctx context.Context) *http.Client {
	return oauth2.NewClient(ctx, c)
}

func (c *DefaultGCPCredentials) PerRPCCredentials() (credentials.PerRPCCredentials, error) {
	tok, err := c.Token()
	if err != nil {
		return nil, err
	}
	return oauth.NewOauthAccess(tok), nil
}

func (s *DefaultGCPCredentials) Debugf(format string) {
	util.Debugf("%s", util.MarshalJSONPB(s))
}

func (s *ClientCredentials) Debugf(format string) {
	util.Debugf("%s", util.MarshalJSONPB(s))
}

func (s *JWT) Debugf(format string) {
	util.Debugf("%s", util.MarshalJSONPB(s))
}

func (s *JSONWebKeys) Debugf(format string) {
	util.Debugf("%s", util.MarshalJSONPB(s))
}

func (s *OAuth2) Debugf(format string) {
	util.Debugf("%s", util.MarshalJSONPB(s))
}

func (s *DefaultGCPCredentials) Validate(fn func(a *DefaultGCPCredentials) error) error {
	return fn(s)
}

func (s *HTTPTask) Debugf(format string) {
	util.Debugf("%s", util.MarshalJSONPB(s))
}

func (s *HTTPTask) Validate(fn func(a *HTTPTask) error) error {
	return fn(s)
}

func (s *HTTPTask) request() (*http.Request, error) {
	u, err := url.Parse(s.Url)
	if err != nil {
		return nil, err
	}
	r := &http.Request{
		Method: s.Method,
		URL:    u,
	}
	for k, v := range s.Headers {
		r.Header.Set(k, v)
	}
	for k, v := range s.Form {
		r.Form.Set(k, v)
	}
	if s.Username != "" && s.Password != "" {
		r.SetBasicAuth(s.Username, s.Password)
	}
	r.WithContext(ENVContext)
	return r, nil
}

func (s *HTTPTask) Do() error {
	r, err := s.request()
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_, err = http.Post(s.CallbackUrl, "application/json", bytes.NewBuffer(bits))
	if err != nil {
		return err
	}
	return nil
}

func Serve(addr string, plugins ...driver.Plugin) error {
	return engine.Serve(addr, true, plugins...)
}

func (c *Common) GetCategory() string {
	return c.Object.TypeUrl
}

func (c *Common) MetaKey(key string) string {
	return c.GetMeta()[key]
}

func (c *Common) Unmarshal(msg proto.Message) error {
	return util.UnarshalAnyPB(c.Object, msg)
}

func (c *Common) AddMeta(key string, val string) {
	c.Meta[key] = val
}

func ToCommon(id string, meta map[string]string, msg proto.Message) (*Common, error) {
	any, err := util.MarshalAnyPB(msg)
	if err != nil {
		return nil, err
	}
	return &Common{
		Identifier: id,
		Object:     any,
		Meta:       meta,
	}, nil
}
