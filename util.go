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
	ojwt "golang.org/x/oauth2/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"net/http"
	"net/url"
)

func init() {
	EnvContext = common.StringArrayFromEnv().ToContext(context.TODO(), "env")
	util = objectify.Default()
}

var (
	util *objectify.Handler
)

var EnvContext context.Context

type ClientSet struct {
	Utility    UtilityServiceClient
	Contact    ContactServiceClient
	Payment    PaymentServiceClient
	Management ManagementServiceClient
	Auth       AuthenticationServiceClient
}

func NewClientSet(conn *grpc.ClientConn) *ClientSet {
	return &ClientSet{
		Utility:    NewUtilityServiceClient(conn),
		Contact:    NewContactServiceClient(conn),
		Payment:    NewPaymentServiceClient(conn),
		Auth:       NewAuthenticationServiceClient(conn),
		Management: NewManagementServiceClient(conn),
	}
}

func (p *SubscriptionResponse) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *SMSResponse) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *CallResponse) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *User) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *PhoneNumberResource) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *PhoneNumber) JSONString() *common.String {
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

func (p *SMSResponse) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *CallResponse) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *FaxResponse) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *SubscriptionResponse) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *PhoneNumberResource) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *PhoneNumber) UnmarshalJSONFrom(bits []byte) error {
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

func (p *FaxResponse) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *CallResponse) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *SMS) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *Call) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *Fax) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *FaxBlast) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *SMS) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Call) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Fax) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *FaxBlast) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *FaxResponse) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Email) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *EmailBlastRequest) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *EmailRequest) JSONString() *common.String {
	return common.MessageToJSONString(p)
}
func (p *Email) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *EmailBlastRequest) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *EmailRequest) UnmarshalProtoFrom(bits []byte) error {
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

func (c *OAuth2) GetRequestMetadata(ctx context.Context, uri ...string) (*common.StringMap, error) {
	src := &oauth.TokenSource{
		TokenSource: c,
	}
	m, err := src.GetRequestMetadata(ctx, uri...)
	if err != nil {
		return nil, err
	}
	return common.ToStringMap(m), nil
}

func (c *ClientCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (*common.StringMap, error) {
	src := &oauth.TokenSource{
		TokenSource: c,
	}
	m, err := src.GetRequestMetadata(ctx, uri...)
	if err != nil {
		return nil, err
	}
	return common.ToStringMap(m), nil
}

func (c *JWT) GetRequestMetadata(ctx context.Context, uri ...string) (*common.StringMap, error) {
	src := &oauth.TokenSource{
		TokenSource: c,
	}
	m, err := src.GetRequestMetadata(ctx, uri...)
	if err != nil {
		return nil, err
	}
	return common.ToStringMap(m), nil
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

func (c *OAuth2) AuthCodeURL(ctx context.Context) *http.Client {
	return oauth2.NewClient(ctx, c)
}
