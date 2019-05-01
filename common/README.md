# common
--
    import "github.com/autom8ter/api/common"


## Usage

```go
var (
	ENVContext context.Context
)
```

#### func  Serve

```go
func Serve(addr string, plugins ...driver.Plugin) error
```

#### type ClientCredentials

```go
type ClientCredentials struct {
	ClientId             string            `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         string            `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	TokenUrl             string            `protobuf:"bytes,3,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	Scopes               []string          `protobuf:"bytes,4,rep,name=scopes,proto3" json:"scopes,omitempty"`
	EndpointParams       map[string]string `protobuf:"bytes,5,rep,name=endpoint_params,json=endpointParams,proto3" json:"endpoint_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func  NewClientCredentials

```go
func NewClientCredentials(tokenURL, clientID, clientSecret string, params map[string]string, scopes []string) *ClientCredentials
```

#### func (*ClientCredentials) Client

```go
func (c *ClientCredentials) Client(ctx context.Context) *http.Client
```

#### func (*ClientCredentials) Config

```go
func (c *ClientCredentials) Config() *clientcredentials.Config
```

#### func (*ClientCredentials) Debugf

```go
func (s *ClientCredentials) Debugf(format string)
```

#### func (*ClientCredentials) Descriptor

```go
func (*ClientCredentials) Descriptor() ([]byte, []int)
```

#### func (*ClientCredentials) GetClientId

```go
func (m *ClientCredentials) GetClientId() string
```

#### func (*ClientCredentials) GetClientSecret

```go
func (m *ClientCredentials) GetClientSecret() string
```

#### func (*ClientCredentials) GetEndpointParams

```go
func (m *ClientCredentials) GetEndpointParams() map[string]string
```

#### func (*ClientCredentials) GetScopes

```go
func (m *ClientCredentials) GetScopes() []string
```

#### func (*ClientCredentials) GetTokenUrl

```go
func (m *ClientCredentials) GetTokenUrl() string
```

#### func (*ClientCredentials) PerRPCCredentials

```go
func (c *ClientCredentials) PerRPCCredentials() (credentials.PerRPCCredentials, error)
```

#### func (*ClientCredentials) ProtoMessage

```go
func (*ClientCredentials) ProtoMessage()
```

#### func (*ClientCredentials) Reset

```go
func (m *ClientCredentials) Reset()
```

#### func (*ClientCredentials) String

```go
func (m *ClientCredentials) String() string
```

#### func (*ClientCredentials) Token

```go
func (c *ClientCredentials) Token() (*oauth2.Token, error)
```

#### func (*ClientCredentials) Validate

```go
func (s *ClientCredentials) Validate(fn func(a *ClientCredentials) error) error
```

#### func (*ClientCredentials) XXX_DiscardUnknown

```go
func (m *ClientCredentials) XXX_DiscardUnknown()
```

#### func (*ClientCredentials) XXX_Marshal

```go
func (m *ClientCredentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*ClientCredentials) XXX_Merge

```go
func (m *ClientCredentials) XXX_Merge(src proto.Message)
```

#### func (*ClientCredentials) XXX_Size

```go
func (m *ClientCredentials) XXX_Size() int
```

#### func (*ClientCredentials) XXX_Unmarshal

```go
func (m *ClientCredentials) XXX_Unmarshal(b []byte) error
```

#### type Common

```go
type Common struct {
	Identifer            string            `protobuf:"bytes,1,opt,name=identifer,proto3" json:"identifer,omitempty"`
	Object               *any.Any          `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	Meta                 map[string]string `protobuf:"bytes,3,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func (*Common) Descriptor

```go
func (*Common) Descriptor() ([]byte, []int)
```

#### func (*Common) GetIdentifer

```go
func (m *Common) GetIdentifer() string
```

#### func (*Common) GetMeta

```go
func (m *Common) GetMeta() map[string]string
```

#### func (*Common) GetObject

```go
func (m *Common) GetObject() *any.Any
```

#### func (*Common) ProtoMessage

```go
func (*Common) ProtoMessage()
```

#### func (*Common) Reset

```go
func (m *Common) Reset()
```

#### func (*Common) String

```go
func (m *Common) String() string
```

#### func (*Common) XXX_DiscardUnknown

```go
func (m *Common) XXX_DiscardUnknown()
```

#### func (*Common) XXX_Marshal

```go
func (m *Common) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Common) XXX_Merge

```go
func (m *Common) XXX_Merge(src proto.Message)
```

#### func (*Common) XXX_Size

```go
func (m *Common) XXX_Size() int
```

#### func (*Common) XXX_Unmarshal

```go
func (m *Common) XXX_Unmarshal(b []byte) error
```

#### type DefaultGCPCredentials

```go
type DefaultGCPCredentials struct {
	Scopes               []string `protobuf:"bytes,1,rep,name=scopes,proto3" json:"scopes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  NewGoogleDefaultCredentials

```go
func NewGoogleDefaultCredentials(scopes []string) *DefaultGCPCredentials
```

#### func (*DefaultGCPCredentials) Client

```go
func (c *DefaultGCPCredentials) Client(ctx context.Context) *http.Client
```

#### func (*DefaultGCPCredentials) Debugf

```go
func (s *DefaultGCPCredentials) Debugf(format string)
```

#### func (*DefaultGCPCredentials) Descriptor

```go
func (*DefaultGCPCredentials) Descriptor() ([]byte, []int)
```

#### func (*DefaultGCPCredentials) FindCredentials

```go
func (d *DefaultGCPCredentials) FindCredentials() (*google.Credentials, error)
```

#### func (*DefaultGCPCredentials) GetScopes

```go
func (m *DefaultGCPCredentials) GetScopes() []string
```

#### func (*DefaultGCPCredentials) PerRPCCredentials

```go
func (c *DefaultGCPCredentials) PerRPCCredentials() (credentials.PerRPCCredentials, error)
```

#### func (*DefaultGCPCredentials) ProtoMessage

```go
func (*DefaultGCPCredentials) ProtoMessage()
```

#### func (*DefaultGCPCredentials) Reset

```go
func (m *DefaultGCPCredentials) Reset()
```

#### func (*DefaultGCPCredentials) String

```go
func (m *DefaultGCPCredentials) String() string
```

#### func (*DefaultGCPCredentials) Token

```go
func (d *DefaultGCPCredentials) Token() (*oauth2.Token, error)
```

#### func (*DefaultGCPCredentials) Validate

```go
func (s *DefaultGCPCredentials) Validate(fn func(a *DefaultGCPCredentials) error) error
```

#### func (*DefaultGCPCredentials) XXX_DiscardUnknown

```go
func (m *DefaultGCPCredentials) XXX_DiscardUnknown()
```

#### func (*DefaultGCPCredentials) XXX_Marshal

```go
func (m *DefaultGCPCredentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*DefaultGCPCredentials) XXX_Merge

```go
func (m *DefaultGCPCredentials) XXX_Merge(src proto.Message)
```

#### func (*DefaultGCPCredentials) XXX_Size

```go
func (m *DefaultGCPCredentials) XXX_Size() int
```

#### func (*DefaultGCPCredentials) XXX_Unmarshal

```go
func (m *DefaultGCPCredentials) XXX_Unmarshal(b []byte) error
```

#### type HTTPTask

```go
type HTTPTask struct {
	Url                  string               `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Method               string               `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Headers              map[string]string    `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Form                 map[string]string    `protobuf:"bytes,4,rep,name=form,proto3" json:"form,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Username             string               `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password             string               `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Body                 *any.Any             `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	Schedule             *timestamp.Timestamp `protobuf:"bytes,8,opt,name=schedule,proto3" json:"schedule,omitempty"`
	CallbackUrl          string               `protobuf:"bytes,9,opt,name=callback_url,json=callbackUrl,proto3" json:"callback_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}
```


#### func (*HTTPTask) Debugf

```go
func (s *HTTPTask) Debugf(format string)
```

#### func (*HTTPTask) Descriptor

```go
func (*HTTPTask) Descriptor() ([]byte, []int)
```

#### func (*HTTPTask) GetBody

```go
func (m *HTTPTask) GetBody() *any.Any
```

#### func (*HTTPTask) GetCallbackUrl

```go
func (m *HTTPTask) GetCallbackUrl() string
```

#### func (*HTTPTask) GetForm

```go
func (m *HTTPTask) GetForm() map[string]string
```

#### func (*HTTPTask) GetHeaders

```go
func (m *HTTPTask) GetHeaders() map[string]string
```

#### func (*HTTPTask) GetMethod

```go
func (m *HTTPTask) GetMethod() string
```

#### func (*HTTPTask) GetPassword

```go
func (m *HTTPTask) GetPassword() string
```

#### func (*HTTPTask) GetSchedule

```go
func (m *HTTPTask) GetSchedule() *timestamp.Timestamp
```

#### func (*HTTPTask) GetUrl

```go
func (m *HTTPTask) GetUrl() string
```

#### func (*HTTPTask) GetUsername

```go
func (m *HTTPTask) GetUsername() string
```

#### func (*HTTPTask) POST

```go
func (s *HTTPTask) POST() (*http.Request, error)
```

#### func (*HTTPTask) ProtoMessage

```go
func (*HTTPTask) ProtoMessage()
```

#### func (*HTTPTask) Reset

```go
func (m *HTTPTask) Reset()
```

#### func (*HTTPTask) String

```go
func (m *HTTPTask) String() string
```

#### func (*HTTPTask) Validate

```go
func (s *HTTPTask) Validate(fn func(a *HTTPTask) error) error
```

#### func (*HTTPTask) XXX_DiscardUnknown

```go
func (m *HTTPTask) XXX_DiscardUnknown()
```

#### func (*HTTPTask) XXX_Marshal

```go
func (m *HTTPTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*HTTPTask) XXX_Merge

```go
func (m *HTTPTask) XXX_Merge(src proto.Message)
```

#### func (*HTTPTask) XXX_Size

```go
func (m *HTTPTask) XXX_Size() int
```

#### func (*HTTPTask) XXX_Unmarshal

```go
func (m *HTTPTask) XXX_Unmarshal(b []byte) error
```

#### type JSONWebKeys

```go
type JSONWebKeys struct {
	Kty                  string   `protobuf:"bytes,1,opt,name=kty,proto3" json:"kty,omitempty"`
	Kid                  string   `protobuf:"bytes,2,opt,name=kid,proto3" json:"kid,omitempty"`
	Use                  string   `protobuf:"bytes,3,opt,name=use,proto3" json:"use,omitempty"`
	N                    string   `protobuf:"bytes,4,opt,name=n,proto3" json:"n,omitempty"`
	E                    string   `protobuf:"bytes,5,opt,name=e,proto3" json:"e,omitempty"`
	X5C                  []string `protobuf:"bytes,6,rep,name=x5c,proto3" json:"x5c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*JSONWebKeys) Debugf

```go
func (s *JSONWebKeys) Debugf(format string)
```

#### func (*JSONWebKeys) Descriptor

```go
func (*JSONWebKeys) Descriptor() ([]byte, []int)
```

#### func (*JSONWebKeys) GetE

```go
func (m *JSONWebKeys) GetE() string
```

#### func (*JSONWebKeys) GetKid

```go
func (m *JSONWebKeys) GetKid() string
```

#### func (*JSONWebKeys) GetKty

```go
func (m *JSONWebKeys) GetKty() string
```

#### func (*JSONWebKeys) GetN

```go
func (m *JSONWebKeys) GetN() string
```

#### func (*JSONWebKeys) GetUse

```go
func (m *JSONWebKeys) GetUse() string
```

#### func (*JSONWebKeys) GetX5C

```go
func (m *JSONWebKeys) GetX5C() []string
```

#### func (*JSONWebKeys) ProtoMessage

```go
func (*JSONWebKeys) ProtoMessage()
```

#### func (*JSONWebKeys) Reset

```go
func (m *JSONWebKeys) Reset()
```

#### func (*JSONWebKeys) String

```go
func (m *JSONWebKeys) String() string
```

#### func (*JSONWebKeys) UnmarshalJSONFrom

```go
func (p *JSONWebKeys) UnmarshalJSONFrom(bits []byte) error
```

#### func (*JSONWebKeys) UnmarshalProtoFrom

```go
func (p *JSONWebKeys) UnmarshalProtoFrom(bits []byte) error
```

#### func (*JSONWebKeys) XXX_DiscardUnknown

```go
func (m *JSONWebKeys) XXX_DiscardUnknown()
```

#### func (*JSONWebKeys) XXX_Marshal

```go
func (m *JSONWebKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*JSONWebKeys) XXX_Merge

```go
func (m *JSONWebKeys) XXX_Merge(src proto.Message)
```

#### func (*JSONWebKeys) XXX_Size

```go
func (m *JSONWebKeys) XXX_Size() int
```

#### func (*JSONWebKeys) XXX_Unmarshal

```go
func (m *JSONWebKeys) XXX_Unmarshal(b []byte) error
```

#### type JWT

```go
type JWT struct {
	Email                string               `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	PrivateKey           []byte               `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	PriveKeyId           string               `protobuf:"bytes,3,opt,name=prive_key_id,json=priveKeyId,proto3" json:"prive_key_id,omitempty"`
	Subject              string               `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Scopes               []string             `protobuf:"bytes,5,rep,name=scopes,proto3" json:"scopes,omitempty"`
	TokenUrl             string               `protobuf:"bytes,6,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	Expires              *timestamp.Timestamp `protobuf:"bytes,7,opt,name=expires,proto3" json:"expires,omitempty"`
	Audience             string               `protobuf:"bytes,8,opt,name=audience,proto3" json:"audience,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}
```


#### func  NewJWT

```go
func NewJWT(tokenURL, email, privateKey, privKeyID, subject string, audience string, timer time.Time, scopes []string) *JWT
```

#### func (*JWT) Client

```go
func (c *JWT) Client(ctx context.Context) *http.Client
```

#### func (*JWT) Config

```go
func (c *JWT) Config() *ojwt.Config
```

#### func (*JWT) Debugf

```go
func (s *JWT) Debugf(format string)
```

#### func (*JWT) Descriptor

```go
func (*JWT) Descriptor() ([]byte, []int)
```

#### func (*JWT) GetAudience

```go
func (m *JWT) GetAudience() string
```

#### func (*JWT) GetEmail

```go
func (m *JWT) GetEmail() string
```

#### func (*JWT) GetExpires

```go
func (m *JWT) GetExpires() *timestamp.Timestamp
```

#### func (*JWT) GetPrivateKey

```go
func (m *JWT) GetPrivateKey() []byte
```

#### func (*JWT) GetPriveKeyId

```go
func (m *JWT) GetPriveKeyId() string
```

#### func (*JWT) GetScopes

```go
func (m *JWT) GetScopes() []string
```

#### func (*JWT) GetSubject

```go
func (m *JWT) GetSubject() string
```

#### func (*JWT) GetTokenUrl

```go
func (m *JWT) GetTokenUrl() string
```

#### func (*JWT) PerRPCCredentials

```go
func (c *JWT) PerRPCCredentials() (credentials.PerRPCCredentials, error)
```

#### func (*JWT) ProtoMessage

```go
func (*JWT) ProtoMessage()
```

#### func (*JWT) Reset

```go
func (m *JWT) Reset()
```

#### func (*JWT) String

```go
func (m *JWT) String() string
```

#### func (*JWT) Token

```go
func (c *JWT) Token() (*oauth2.Token, error)
```

#### func (*JWT) Validate

```go
func (s *JWT) Validate(fn func(a *JWT) error) error
```

#### func (*JWT) XXX_DiscardUnknown

```go
func (m *JWT) XXX_DiscardUnknown()
```

#### func (*JWT) XXX_Marshal

```go
func (m *JWT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*JWT) XXX_Merge

```go
func (m *JWT) XXX_Merge(src proto.Message)
```

#### func (*JWT) XXX_Size

```go
func (m *JWT) XXX_Size() int
```

#### func (*JWT) XXX_Unmarshal

```go
func (m *JWT) XXX_Unmarshal(b []byte) error
```

#### type Jwks

```go
type Jwks struct {
	Keys                 []*JSONWebKeys `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*Jwks) Descriptor

```go
func (*Jwks) Descriptor() ([]byte, []int)
```

#### func (*Jwks) GetKeys

```go
func (m *Jwks) GetKeys() []*JSONWebKeys
```

#### func (*Jwks) ProtoMessage

```go
func (*Jwks) ProtoMessage()
```

#### func (*Jwks) Reset

```go
func (m *Jwks) Reset()
```

#### func (*Jwks) String

```go
func (m *Jwks) String() string
```

#### func (*Jwks) TokenCert

```go
func (c *Jwks) TokenCert(token *jwt.Token) (string, error)
```

#### func (*Jwks) UnmarshalProtoFrom

```go
func (p *Jwks) UnmarshalProtoFrom(bits []byte) error
```

#### func (*Jwks) XXX_DiscardUnknown

```go
func (m *Jwks) XXX_DiscardUnknown()
```

#### func (*Jwks) XXX_Marshal

```go
func (m *Jwks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Jwks) XXX_Merge

```go
func (m *Jwks) XXX_Merge(src proto.Message)
```

#### func (*Jwks) XXX_Size

```go
func (m *Jwks) XXX_Size() int
```

#### func (*Jwks) XXX_Unmarshal

```go
func (m *Jwks) XXX_Unmarshal(b []byte) error
```

#### type OAuth2

```go
type OAuth2 struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         string   `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	TokenUrl             string   `protobuf:"bytes,3,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	AuthUrl              string   `protobuf:"bytes,4,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
	Scopes               []string `protobuf:"bytes,5,rep,name=scopes,proto3" json:"scopes,omitempty"`
	Redirect             string   `protobuf:"bytes,6,opt,name=redirect,proto3" json:"redirect,omitempty"`
	Code                 string   `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  NewOAuth2

```go
func NewOAuth2(clientID string, clientSecret string, tokenURL string, authURL string, redirect string, scopes []string) *OAuth2
```

#### func (*OAuth2) AuthCodeURL

```go
func (c *OAuth2) AuthCodeURL(state string, audience string) string
```

#### func (*OAuth2) Client

```go
func (c *OAuth2) Client(ctx context.Context) *http.Client
```

#### func (*OAuth2) Config

```go
func (c *OAuth2) Config() *oauth2.Config
```

#### func (*OAuth2) Debugf

```go
func (s *OAuth2) Debugf(format string)
```

#### func (*OAuth2) Descriptor

```go
func (*OAuth2) Descriptor() ([]byte, []int)
```

#### func (*OAuth2) Exchange

```go
func (c *OAuth2) Exchange() (*oauth2.Token, error)
```

#### func (*OAuth2) GetAuthUrl

```go
func (m *OAuth2) GetAuthUrl() string
```

#### func (*OAuth2) GetClientId

```go
func (m *OAuth2) GetClientId() string
```

#### func (*OAuth2) GetClientSecret

```go
func (m *OAuth2) GetClientSecret() string
```

#### func (*OAuth2) GetCode

```go
func (m *OAuth2) GetCode() string
```

#### func (*OAuth2) GetRedirect

```go
func (m *OAuth2) GetRedirect() string
```

#### func (*OAuth2) GetScopes

```go
func (m *OAuth2) GetScopes() []string
```

#### func (*OAuth2) GetTokenUrl

```go
func (m *OAuth2) GetTokenUrl() string
```

#### func (*OAuth2) PerRPCCredentials

```go
func (c *OAuth2) PerRPCCredentials() (credentials.PerRPCCredentials, error)
```

#### func (*OAuth2) ProtoMessage

```go
func (*OAuth2) ProtoMessage()
```

#### func (*OAuth2) Reset

```go
func (m *OAuth2) Reset()
```

#### func (*OAuth2) SetCode

```go
func (c *OAuth2) SetCode(r *http.Request)
```

#### func (*OAuth2) String

```go
func (m *OAuth2) String() string
```

#### func (*OAuth2) Token

```go
func (c *OAuth2) Token() (*oauth2.Token, error)
```

#### func (*OAuth2) Validate

```go
func (s *OAuth2) Validate(fn func(a *OAuth2) error) error
```

#### func (*OAuth2) XXX_DiscardUnknown

```go
func (m *OAuth2) XXX_DiscardUnknown()
```

#### func (*OAuth2) XXX_Marshal

```go
func (m *OAuth2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*OAuth2) XXX_Merge

```go
func (m *OAuth2) XXX_Merge(src proto.Message)
```

#### func (*OAuth2) XXX_Size

```go
func (m *OAuth2) XXX_Size() int
```

#### func (*OAuth2) XXX_Unmarshal

```go
func (m *OAuth2) XXX_Unmarshal(b []byte) error
```

#### type PlainText

```go
type PlainText struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*PlainText) Descriptor

```go
func (*PlainText) Descriptor() ([]byte, []int)
```

#### func (*PlainText) GetText

```go
func (m *PlainText) GetText() string
```

#### func (*PlainText) ProtoMessage

```go
func (*PlainText) ProtoMessage()
```

#### func (*PlainText) Reset

```go
func (m *PlainText) Reset()
```

#### func (*PlainText) String

```go
func (m *PlainText) String() string
```

#### func (*PlainText) XXX_DiscardUnknown

```go
func (m *PlainText) XXX_DiscardUnknown()
```

#### func (*PlainText) XXX_Marshal

```go
func (m *PlainText) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*PlainText) XXX_Merge

```go
func (m *PlainText) XXX_Merge(src proto.Message)
```

#### func (*PlainText) XXX_Size

```go
func (m *PlainText) XXX_Size() int
```

#### func (*PlainText) XXX_Unmarshal

```go
func (m *PlainText) XXX_Unmarshal(b []byte) error
```
