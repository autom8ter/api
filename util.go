//go:generate godocdown -o README.md

package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/autom8ter/api/common"
	"github.com/autom8ter/objectify"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/proto"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/google"
	ojwt "golang.org/x/oauth2/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"io"
	"net/http"
	"net/url"
)

func init() {
	util = objectify.Default()
}

var (
	util *objectify.Handler
)

type ClientSet struct {
	Debug         DebugServiceClient
	Subscriptions SubscriptionServiceClient
	Users         UserServiceClient
	Auth          AuthenticationServiceClient
	Events        EventServiceClient
}

func NewClientSet(conn *grpc.ClientConn) *ClientSet {
	return &ClientSet{
		Debug:         NewDebugServiceClient(conn),
		Subscriptions: NewSubscriptionServiceClient(conn),
		Users:         NewUserServiceClient(conn),
		Auth:          NewAuthenticationServiceClient(conn),
		Events:        NewEventServiceClient(conn),
	}
}

func (p *Plan) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Product) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *AppMetadata) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *UserMetadata) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Role) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *User) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *JSONWebKeys) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Identity) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Card) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Jwks) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Document) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func NewDocument(category, name string, data map[string]interface{}) *Document {
	return &Document{
		Category: common.ToString(category),
		Name:     common.ToString(name),
		Data:     common.StringMapFromData(data),
	}
}

func (p *Jwks) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *Card) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *Identity) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *User) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *Plan) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *Product) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *UserMetadata) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *AppMetadata) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *Document) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *Role) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

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
			cert = "-----BEGIN CERTIFICATE-----\n" + c.Keys[k].X5C.Strings[0].Text + "\n-----END CERTIFICATE-----"
		}
	}
	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}
	return cert, nil
}

func (s *OAuth2) ToContext(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, key, s)
}

func (s *DefaultGCPCredentials) ToContext(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, key, s)
}
func (s *ClientCredentials) ToContext(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, key, s)
}

func (s *JWT) ToContext(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, key, s)
}

func (s *Event) JSONString() *common.String {
	return common.MessageToJSONString(s)
}

func (s *Event) UnmarshalProtofrom(bits []byte) error {
	return util.UnmarshalProto(bits, s)
}

func (s *Event) UnmarshalJSONfrom(bits []byte) error {
	return util.UnmarshalJSON(bits, s)
}

func (s *JWT) JSONString() *common.String {
	return common.MessageToJSONString(s)
}

func (s *OAuth2) JSONString() *common.String {
	return common.MessageToJSONString(s)
}
func (s *ClientCredentials) JSONString() *common.String {
	return common.MessageToJSONString(s)
}

func (s *JWT) UnmarshalProtoFrom(b []byte) error {
	return util.UnmarshalProto(b, s)
}

func (s *JWT) UnmarshalJSONFrom(b []byte) error {
	return util.UnmarshalJSON(b, s)
}

func (s *OAuth2) UnmarshalProtoFrom(b []byte) error {
	return util.UnmarshalProto(b, s)
}

func (s *OAuth2) UnmarshalJSONFrom(b []byte) error {
	return util.UnmarshalJSON(b, s)
}

func (s *ClientCredentials) UnmarshalProtoFrom(b []byte) error {
	return util.UnmarshalProto(b, s)
}

func (s *ClientCredentials) UnmarshalJSONFrom(b []byte) error {
	return util.UnmarshalJSON(b, s)
}

func (s *JWT) DeepEqual(y interface{}) bool {
	return util.DeepEqual(s, y)
}

func (s *OAuth2) DeepEqual(y interface{}) bool {
	return util.DeepEqual(s, y)
}

func (s *ClientCredentials) DeepEqual(y interface{}) bool {
	return util.DeepEqual(s, y)
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

func (s *JWT) Debugf(format string) {
	str := common.MessageToJSONString(s)
	str.Debugf(format)
}

func (s *ClientCredentials) Debugf(format string) {
	str := common.MessageToJSONString(s)
	str.Debugf(format)
}

func (s *OAuth2) Debugf(format string) {
	str := common.MessageToJSONString(s)
	str.Debugf(format)
}

func (c *OAuth2) Token() (*oauth2.Token, error) {
	a := c.Config()
	tok, err := a.Exchange(context.TODO(), c.Code.Text)
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func (c *OAuth2) Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.ClientId.Text,
		ClientSecret: c.ClientSecret.Text,
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.AuthUrl.Text,
			TokenURL: c.TokenUrl.Text,
		},
		RedirectURL: c.Redirect.Text,
		Scopes:      c.Scopes.Array(),
	}
}

func (c *JWT) Config() *ojwt.Config {
	return &ojwt.Config{
		Email:        "",
		PrivateKey:   nil,
		PrivateKeyID: "",
		Subject:      "",
		Scopes:       c.Scopes.Array(),
		TokenURL:     "",
		Expires:      0,
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
	for k, v := range c.EndpointParams.StringMap {
		rl.Set(k, v.Text)
	}
	return &clientcredentials.Config{
		ClientID:       c.ClientId.Text,
		ClientSecret:   c.ClientSecret.Text,
		TokenURL:       c.TokenUrl.Text,
		Scopes:         c.Scopes.Array(),
		EndpointParams: rl,
	}
}

func (c *ClientCredentials) Token() (*oauth2.Token, error) {
	a := c.Config()
	src := a.TokenSource(context.TODO())
	tok, err := src.Token()
	if err != nil {
		return nil, err
	}
	return tok, nil
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
	c.Code = common.ToString(r.URL.Query().Get("code"))
}

func (c *JWT) NewAPIClientSet(ctx context.Context, addr string) (*ClientSet, error) {
	creds, err := c.PerRPCCredentials()
	if err != nil {
		return nil, err
	}
	conn, err := grpc.DialContext(ctx, addr, grpc.WithPerRPCCredentials(creds))
	if err != nil {
		return nil, err
	}
	return NewClientSet(conn), nil
}

func (c *OAuth2) NewAPIClientSet(ctx context.Context, addr string) (*ClientSet, error) {
	creds, err := c.PerRPCCredentials()
	if err != nil {
		return nil, err
	}
	conn, err := grpc.DialContext(ctx, addr, grpc.WithPerRPCCredentials(creds))
	if err != nil {
		return nil, err
	}
	return NewClientSet(conn), nil

}

func (c *ClientCredentials) NewAPIClientSet(ctx context.Context, addr string) (*ClientSet, error) {
	creds, err := c.PerRPCCredentials()
	if err != nil {
		return nil, err
	}
	conn, err := grpc.DialContext(ctx, addr, grpc.WithPerRPCCredentials(creds))
	return NewClientSet(conn), nil
}

func NewClientCredentials(tokenURL, clientID, clientSecret string, params map[string]string, scopes []string) *ClientCredentials {
	return &ClientCredentials{
		ClientId:       common.ToString(clientID),
		ClientSecret:   common.ToString(clientSecret),
		TokenUrl:       common.ToString(tokenURL),
		Scopes:         common.ToStringArray(scopes),
		EndpointParams: common.ToStringMap(params),
	}
}

func NewJWT(tokenURL, email, privateKey, privKeyID, subject string, expiry, audience string, scopes []string) *JWT {
	return &JWT{
		Email:      common.ToString(email),
		PrivateKey: []byte(privateKey),
		PriveKeyId: common.ToString(privKeyID),
		Subject:    common.ToString(subject),
		Scopes:     common.ToStringArray(scopes),
		TokenUrl:   common.ToString(tokenURL),
		Expires:    common.ToString(expiry),
		Audience:   common.ToString(audience),
	}
}

func NewOAuth2(clientID string, clientSecret string, tokenURL string, authURL string, redirect string, scopes []string) *OAuth2 {
	return &OAuth2{
		ClientId:     common.ToString(clientID),
		ClientSecret: common.ToString(clientSecret),
		TokenUrl:     common.ToString(tokenURL),
		AuthUrl:      common.ToString(authURL),
		Scopes:       common.ToStringArray(scopes),
		Redirect:     common.ToString(redirect),
	}
}

func NewGoogleDefaultCredentials(scopes []string) *DefaultGCPCredentials {
	return &DefaultGCPCredentials{
		Scopes: common.ToStringArray(scopes),
	}
}

func (d *DefaultGCPCredentials) FindCredentials() (*google.Credentials, error) {
	return google.FindDefaultCredentials(context.TODO(), d.Scopes.Array()...)
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

func (c *DefaultGCPCredentials) NewAPIClientSet(ctx context.Context, addr string) (*ClientSet, error) {
	creds, err := c.PerRPCCredentials()
	if err != nil {
		return nil, err
	}
	conn, err := grpc.DialContext(ctx, addr, grpc.WithPerRPCCredentials(creds))
	if err != nil {
		return nil, err
	}
	return NewClientSet(conn), nil

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
	str := common.MessageToJSONString(s)
	str.Debugf(format)
}

func (s *DefaultGCPCredentials) Validate(fn func(a *DefaultGCPCredentials) error) error {
	return fn(s)
}

func (d *Document) Render(s *common.String, w io.Writer) error {
	return s.Render(d.Name.Text, w, d.Data)
}
