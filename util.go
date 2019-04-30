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
	"google.golang.org/grpc"
	"net/http"
	"net/url"
	"reflect"
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

func (s *ConfigSet) ToContext(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, key, s)
}

func (s *Config) ToContext(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, key, s)
}

func (s *ConfigSet) UnmarshalJSONFrom(b []byte) error {
	return util.UnmarshalJSON(b, s)
}
func (s *ConfigSet) UnmarshalProtoFrom(b []byte) error {
	return util.UnmarshalProto(b, s)
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

func (s *ConfigSet) JSONString() *common.String {
	return common.MessageToJSONString(s)
}

func (s *Config) JSONString() *common.String {
	return common.MessageToJSONString(s)
}

func (s *Config) UnmarshalProtoFrom(b []byte) error {
	return util.UnmarshalProto(b, s)
}

func (s *Config) UnmarshalJSONFrom(b []byte) error {
	return util.UnmarshalJSON(b, s)
}

func (s *ConfigSet) Get(key string) *Config {
	return s.Configs[key]
}

func (s *ConfigSet) Put(key string, c *Config) {
	s.Configs[key] = c
}

func (s *ConfigSet) Exists(key string) bool {
	t := s.Configs[key]
	if t == nil {
		return false
	}
	return true
}

func (s *ConfigSet) Length() int {
	return len(s.Configs)
}

func (s *Config) DeepEqual(y interface{}) bool {
	return reflect.DeepEqual(s, y)
}

func (s *ConfigSet) DeepEqual(y interface{}) bool {
	return util.DeepEqual(s, y)
}

func (s *ConfigSet) Validate(fn func(set *ConfigSet) error) error {
	return fn(s)
}

func (s *ConfigSet) Debugf(format string) {
	str := common.MessageToJSONString(s)
	str.Debugf(format)
}

func (s *Config) Debugf(format string) {
	str := common.MessageToJSONString(s)
	str.Debugf(format)
}

func (c *Config) Oauth2Client(ctx context.Context, code *common.String) (*http.Client, error) {
	a := &oauth2.Config{
		ClientID:     c.ClientId.Text,
		ClientSecret: c.ClientSecret.Text,
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.AuthUrl.Text,
			TokenURL: c.TokenUrl.Text,
		},
		RedirectURL: c.Redirect.Text,
		Scopes:      c.Scopes.Array(),
	}
	tok, err := a.Exchange(ctx, code.Text)
	if err != nil {
		return nil, err
	}
	return a.Client(ctx, tok), nil
}

func (c *Config) ClientCredentials() *clientcredentials.Config {
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

func (c *Config) ClientCredentialsClient(ctx context.Context) *http.Client {
	return c.ClientCredentials().Client(ctx)
}
