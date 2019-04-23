# api
--
    import "github.com/autom8ter/api/go/api"

Package api is a reverse proxy.

It translates gRPC into RESTful JSON APIs.

## Usage

```go
var AUTH_SESSION_NAME = "autom8ter"
```

```go
var Debug bool
```

```go
var Util = objectify.Default()
```

#### func  Debugf

```go
func Debugf(format string, args ...interface{})
```

#### func  Fatalf

```go
func Fatalf(format string, args ...interface{})
```

#### func  RegisterCustomerServiceHandler

```go
func RegisterCustomerServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
```
RegisterCustomerServiceHandler registers the http handlers for service
CustomerService to "mux". The handlers forward requests to the grpc endpoint
over "conn".

#### func  RegisterCustomerServiceHandlerClient

```go
func RegisterCustomerServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CustomerServiceClient) error
```
RegisterCustomerServiceHandlerClient registers the http handlers for service
CustomerService to "mux". The handlers forward requests to the grpc endpoint
over the given implementation of "CustomerServiceClient". Note: the gRPC
framework executes interceptors within the gRPC handler. If the passed in
"CustomerServiceClient" doesn't go through the normal gRPC flow (creating a gRPC
client etc.) then it will be up to the passed in "CustomerServiceClient" to call
the correct interceptors.

#### func  RegisterCustomerServiceHandlerFromEndpoint

```go
func RegisterCustomerServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
```
RegisterCustomerServiceHandlerFromEndpoint is same as
RegisterCustomerServiceHandler but automatically dials to "endpoint" and closes
the connection when "ctx" gets done.

#### func  RegisterCustomerServiceServer

```go
func RegisterCustomerServiceServer(s *grpc.Server, srv CustomerServiceServer)
```

#### func  Warnf

```go
func Warnf(format string, args ...interface{})
```

#### type AccessToken

```go
type AccessToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*AccessToken) Descriptor

```go
func (*AccessToken) Descriptor() ([]byte, []int)
```

#### func (*AccessToken) GetToken

```go
func (m *AccessToken) GetToken() string
```

#### func (*AccessToken) ProtoMessage

```go
func (*AccessToken) ProtoMessage()
```

#### func (*AccessToken) Reset

```go
func (m *AccessToken) Reset()
```

#### func (*AccessToken) String

```go
func (m *AccessToken) String() string
```

#### func (*AccessToken) XXX_DiscardUnknown

```go
func (m *AccessToken) XXX_DiscardUnknown()
```

#### func (*AccessToken) XXX_Marshal

```go
func (m *AccessToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*AccessToken) XXX_Merge

```go
func (m *AccessToken) XXX_Merge(src proto.Message)
```

#### func (*AccessToken) XXX_Size

```go
func (m *AccessToken) XXX_Size() int
```

#### func (*AccessToken) XXX_Unmarshal

```go
func (m *AccessToken) XXX_Unmarshal(b []byte) error
```

#### type AppMetadata

```go
type AppMetadata struct {
	Plan                 string   `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	PayToken             string   `protobuf:"bytes,2,opt,name=pay_token,json=payToken,proto3" json:"pay_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*AppMetadata) Descriptor

```go
func (*AppMetadata) Descriptor() ([]byte, []int)
```

#### func (*AppMetadata) GetPayToken

```go
func (m *AppMetadata) GetPayToken() string
```

#### func (*AppMetadata) GetPlan

```go
func (m *AppMetadata) GetPlan() string
```

#### func (*AppMetadata) JSONString

```go
func (p *AppMetadata) JSONString() string
```

#### func (*AppMetadata) ProtoMessage

```go
func (*AppMetadata) ProtoMessage()
```

#### func (*AppMetadata) Render

```go
func (p *AppMetadata) Render(tmpl *template.Template, w io.Writer) error
```

#### func (*AppMetadata) Reset

```go
func (m *AppMetadata) Reset()
```

#### func (*AppMetadata) String

```go
func (m *AppMetadata) String() string
```

#### func (*AppMetadata) XXX_DiscardUnknown

```go
func (m *AppMetadata) XXX_DiscardUnknown()
```

#### func (*AppMetadata) XXX_Marshal

```go
func (m *AppMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*AppMetadata) XXX_Merge

```go
func (m *AppMetadata) XXX_Merge(src proto.Message)
```

#### func (*AppMetadata) XXX_Size

```go
func (m *AppMetadata) XXX_Size() int
```

#### func (*AppMetadata) XXX_Unmarshal

```go
func (m *AppMetadata) XXX_Unmarshal(b []byte) error
```

#### type Auth0

```go
type Auth0 struct {
	Domain               string   `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         string   `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	Scopes               []string `protobuf:"bytes,4,rep,name=scopes,proto3" json:"scopes,omitempty"`
	Redirect             string   `protobuf:"bytes,5,opt,name=redirect,proto3" json:"redirect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Auth0) Descriptor

```go
func (*Auth0) Descriptor() ([]byte, []int)
```

#### func (*Auth0) GetClientId

```go
func (m *Auth0) GetClientId() string
```

#### func (*Auth0) GetClientSecret

```go
func (m *Auth0) GetClientSecret() string
```

#### func (*Auth0) GetDomain

```go
func (m *Auth0) GetDomain() string
```

#### func (*Auth0) GetRedirect

```go
func (m *Auth0) GetRedirect() string
```

#### func (*Auth0) GetScopes

```go
func (m *Auth0) GetScopes() []string
```

#### func (*Auth0) ProtoMessage

```go
func (*Auth0) ProtoMessage()
```

#### func (*Auth0) Reset

```go
func (m *Auth0) Reset()
```

#### func (*Auth0) String

```go
func (m *Auth0) String() string
```

#### func (*Auth0) XXX_DiscardUnknown

```go
func (m *Auth0) XXX_DiscardUnknown()
```

#### func (*Auth0) XXX_Marshal

```go
func (m *Auth0) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Auth0) XXX_Merge

```go
func (m *Auth0) XXX_Merge(src proto.Message)
```

#### func (*Auth0) XXX_Size

```go
func (m *Auth0) XXX_Size() int
```

#### func (*Auth0) XXX_Unmarshal

```go
func (m *Auth0) XXX_Unmarshal(b []byte) error
```

#### type ClientSet

```go
type ClientSet struct {
	Customers CustomerServiceClient
}
```


#### func  NewClientSet

```go
func NewClientSet(d *Dialer) *ClientSet
```

#### type CreateCustomerRequest

```go
type CreateCustomerRequest struct {
	UserInfo             *UserInfo `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}
```


#### func (*CreateCustomerRequest) Descriptor

```go
func (*CreateCustomerRequest) Descriptor() ([]byte, []int)
```

#### func (*CreateCustomerRequest) GetUserInfo

```go
func (m *CreateCustomerRequest) GetUserInfo() *UserInfo
```

#### func (*CreateCustomerRequest) ProtoMessage

```go
func (*CreateCustomerRequest) ProtoMessage()
```

#### func (*CreateCustomerRequest) Reset

```go
func (m *CreateCustomerRequest) Reset()
```

#### func (*CreateCustomerRequest) String

```go
func (m *CreateCustomerRequest) String() string
```

#### func (*CreateCustomerRequest) XXX_DiscardUnknown

```go
func (m *CreateCustomerRequest) XXX_DiscardUnknown()
```

#### func (*CreateCustomerRequest) XXX_Marshal

```go
func (m *CreateCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*CreateCustomerRequest) XXX_Merge

```go
func (m *CreateCustomerRequest) XXX_Merge(src proto.Message)
```

#### func (*CreateCustomerRequest) XXX_Size

```go
func (m *CreateCustomerRequest) XXX_Size() int
```

#### func (*CreateCustomerRequest) XXX_Unmarshal

```go
func (m *CreateCustomerRequest) XXX_Unmarshal(b []byte) error
```

#### type CreateCustomerResponse

```go
type CreateCustomerResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*CreateCustomerResponse) Descriptor

```go
func (*CreateCustomerResponse) Descriptor() ([]byte, []int)
```

#### func (*CreateCustomerResponse) GetId

```go
func (m *CreateCustomerResponse) GetId() string
```

#### func (*CreateCustomerResponse) ProtoMessage

```go
func (*CreateCustomerResponse) ProtoMessage()
```

#### func (*CreateCustomerResponse) Reset

```go
func (m *CreateCustomerResponse) Reset()
```

#### func (*CreateCustomerResponse) String

```go
func (m *CreateCustomerResponse) String() string
```

#### func (*CreateCustomerResponse) XXX_DiscardUnknown

```go
func (m *CreateCustomerResponse) XXX_DiscardUnknown()
```

#### func (*CreateCustomerResponse) XXX_Marshal

```go
func (m *CreateCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*CreateCustomerResponse) XXX_Merge

```go
func (m *CreateCustomerResponse) XXX_Merge(src proto.Message)
```

#### func (*CreateCustomerResponse) XXX_Size

```go
func (m *CreateCustomerResponse) XXX_Size() int
```

#### func (*CreateCustomerResponse) XXX_Unmarshal

```go
func (m *CreateCustomerResponse) XXX_Unmarshal(b []byte) error
```

#### type CustomerServiceClient

```go
type CustomerServiceClient interface {
	CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerResponse, error)
}
```

CustomerServiceClient is the client API for CustomerService service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to
https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

#### func  NewCustomerServiceClient

```go
func NewCustomerServiceClient(cc *grpc.ClientConn) CustomerServiceClient
```

#### type CustomerServiceServer

```go
type CustomerServiceServer interface {
	CreateCustomer(context.Context, *CreateCustomerRequest) (*CreateCustomerResponse, error)
}
```

CustomerServiceServer is the server API for CustomerService service.

#### type Dialer

```go
type Dialer struct {
}
```


#### func  NewDialer

```go
func NewDialer(ctx context.Context, addr string, r *http.Request) (*Dialer, error)
```

#### func (*Dialer) Conn

```go
func (d *Dialer) Conn() *grpc.ClientConn
```

#### type IDToken

```go
type IDToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*IDToken) Descriptor

```go
func (*IDToken) Descriptor() ([]byte, []int)
```

#### func (*IDToken) GetToken

```go
func (m *IDToken) GetToken() string
```

#### func (*IDToken) ProtoMessage

```go
func (*IDToken) ProtoMessage()
```

#### func (*IDToken) Reset

```go
func (m *IDToken) Reset()
```

#### func (*IDToken) String

```go
func (m *IDToken) String() string
```

#### func (*IDToken) XXX_DiscardUnknown

```go
func (m *IDToken) XXX_DiscardUnknown()
```

#### func (*IDToken) XXX_Marshal

```go
func (m *IDToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*IDToken) XXX_Merge

```go
func (m *IDToken) XXX_Merge(src proto.Message)
```

#### func (*IDToken) XXX_Size

```go
func (m *IDToken) XXX_Size() int
```

#### func (*IDToken) XXX_Unmarshal

```go
func (m *IDToken) XXX_Unmarshal(b []byte) error
```

#### type Paths

```go
type Paths struct {
	Home                 string   `protobuf:"bytes,1,opt,name=home,proto3" json:"home,omitempty"`
	Dashboard            string   `protobuf:"bytes,2,opt,name=dashboard,proto3" json:"dashboard,omitempty"`
	Settings             string   `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	Logout               string   `protobuf:"bytes,4,opt,name=logout,proto3" json:"logout,omitempty"`
	Callback             string   `protobuf:"bytes,5,opt,name=callback,proto3" json:"callback,omitempty"`
	Login                string   `protobuf:"bytes,6,opt,name=login,proto3" json:"login,omitempty"`
	Subscribe            string   `protobuf:"bytes,7,opt,name=subscribe,proto3" json:"subscribe,omitempty"`
	Unsubscribe          string   `protobuf:"bytes,8,opt,name=unsubscribe,proto3" json:"unsubscribe,omitempty"`
	Faq                  string   `protobuf:"bytes,9,opt,name=faq,proto3" json:"faq,omitempty"`
	Support              string   `protobuf:"bytes,10,opt,name=support,proto3" json:"support,omitempty"`
	Terms                string   `protobuf:"bytes,11,opt,name=terms,proto3" json:"terms,omitempty"`
	Privacy              string   `protobuf:"bytes,12,opt,name=privacy,proto3" json:"privacy,omitempty"`
	Debug                string   `protobuf:"bytes,13,opt,name=debug,proto3" json:"debug,omitempty"`
	Blog                 string   `protobuf:"bytes,14,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  DefaultPaths

```go
func DefaultPaths() *Paths
```

#### func (*Paths) Descriptor

```go
func (*Paths) Descriptor() ([]byte, []int)
```

#### func (*Paths) GetBlog

```go
func (m *Paths) GetBlog() string
```

#### func (*Paths) GetCallback

```go
func (m *Paths) GetCallback() string
```

#### func (*Paths) GetDashboard

```go
func (m *Paths) GetDashboard() string
```

#### func (*Paths) GetDebug

```go
func (m *Paths) GetDebug() string
```

#### func (*Paths) GetFaq

```go
func (m *Paths) GetFaq() string
```

#### func (*Paths) GetHome

```go
func (m *Paths) GetHome() string
```

#### func (*Paths) GetLogin

```go
func (m *Paths) GetLogin() string
```

#### func (*Paths) GetLogout

```go
func (m *Paths) GetLogout() string
```

#### func (*Paths) GetPrivacy

```go
func (m *Paths) GetPrivacy() string
```

#### func (*Paths) GetSettings

```go
func (m *Paths) GetSettings() string
```

#### func (*Paths) GetSubscribe

```go
func (m *Paths) GetSubscribe() string
```

#### func (*Paths) GetSupport

```go
func (m *Paths) GetSupport() string
```

#### func (*Paths) GetTerms

```go
func (m *Paths) GetTerms() string
```

#### func (*Paths) GetUnsubscribe

```go
func (m *Paths) GetUnsubscribe() string
```

#### func (*Paths) JSONString

```go
func (p *Paths) JSONString() string
```

#### func (*Paths) ProtoMessage

```go
func (*Paths) ProtoMessage()
```

#### func (*Paths) Render

```go
func (p *Paths) Render(tmpl *template.Template, w io.Writer) error
```

#### func (*Paths) Reset

```go
func (m *Paths) Reset()
```

#### func (*Paths) String

```go
func (m *Paths) String() string
```

#### func (*Paths) XXX_DiscardUnknown

```go
func (m *Paths) XXX_DiscardUnknown()
```

#### func (*Paths) XXX_Marshal

```go
func (m *Paths) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Paths) XXX_Merge

```go
func (m *Paths) XXX_Merge(src proto.Message)
```

#### func (*Paths) XXX_Size

```go
func (m *Paths) XXX_Size() int
```

#### func (*Paths) XXX_Unmarshal

```go
func (m *Paths) XXX_Unmarshal(b []byte) error
```

#### type RefreshToken

```go
type RefreshToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*RefreshToken) Descriptor

```go
func (*RefreshToken) Descriptor() ([]byte, []int)
```

#### func (*RefreshToken) GetToken

```go
func (m *RefreshToken) GetToken() string
```

#### func (*RefreshToken) ProtoMessage

```go
func (*RefreshToken) ProtoMessage()
```

#### func (*RefreshToken) Reset

```go
func (m *RefreshToken) Reset()
```

#### func (*RefreshToken) String

```go
func (m *RefreshToken) String() string
```

#### func (*RefreshToken) XXX_DiscardUnknown

```go
func (m *RefreshToken) XXX_DiscardUnknown()
```

#### func (*RefreshToken) XXX_Marshal

```go
func (m *RefreshToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*RefreshToken) XXX_Merge

```go
func (m *RefreshToken) XXX_Merge(src proto.Message)
```

#### func (*RefreshToken) XXX_Size

```go
func (m *RefreshToken) XXX_Size() int
```

#### func (*RefreshToken) XXX_Unmarshal

```go
func (m *RefreshToken) XXX_Unmarshal(b []byte) error
```

#### type Tokens

```go
type Tokens struct {
	Id                   *IDToken      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Access               *AccessToken  `protobuf:"bytes,2,opt,name=access,proto3" json:"access,omitempty"`
	Refresh              *RefreshToken `protobuf:"bytes,3,opt,name=refresh,proto3" json:"refresh,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
```


#### func (*Tokens) Descriptor

```go
func (*Tokens) Descriptor() ([]byte, []int)
```

#### func (*Tokens) GetAccess

```go
func (m *Tokens) GetAccess() *AccessToken
```

#### func (*Tokens) GetId

```go
func (m *Tokens) GetId() *IDToken
```

#### func (*Tokens) GetRefresh

```go
func (m *Tokens) GetRefresh() *RefreshToken
```

#### func (*Tokens) JSONString

```go
func (p *Tokens) JSONString() string
```

#### func (*Tokens) ProtoMessage

```go
func (*Tokens) ProtoMessage()
```

#### func (*Tokens) Render

```go
func (p *Tokens) Render(tmpl *template.Template, w io.Writer) error
```

#### func (*Tokens) Reset

```go
func (m *Tokens) Reset()
```

#### func (*Tokens) String

```go
func (m *Tokens) String() string
```

#### func (*Tokens) XXX_DiscardUnknown

```go
func (m *Tokens) XXX_DiscardUnknown()
```

#### func (*Tokens) XXX_Marshal

```go
func (m *Tokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Tokens) XXX_Merge

```go
func (m *Tokens) XXX_Merge(src proto.Message)
```

#### func (*Tokens) XXX_Size

```go
func (m *Tokens) XXX_Size() int
```

#### func (*Tokens) XXX_Unmarshal

```go
func (m *Tokens) XXX_Unmarshal(b []byte) error
```

#### type UserInfo

```go
type UserInfo struct {
	Name                 string        `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	GivenName            string        `protobuf:"bytes,7,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	FamilyName           string        `protobuf:"bytes,8,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	Gender               string        `protobuf:"bytes,9,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthdate            string        `protobuf:"bytes,10,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	Email                string        `protobuf:"bytes,11,opt,name=email,proto3" json:"email,omitempty"`
	Picture              string        `protobuf:"bytes,12,opt,name=picture,proto3" json:"picture,omitempty"`
	UserMetadata         *UserMetadata `protobuf:"bytes,13,opt,name=user_metadata,json=userMetadata,proto3" json:"user_metadata,omitempty"`
	AppMetadata          *AppMetadata  `protobuf:"bytes,14,opt,name=app_metadata,json=appMetadata,proto3" json:"app_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
```


#### func (*UserInfo) Descriptor

```go
func (*UserInfo) Descriptor() ([]byte, []int)
```

#### func (*UserInfo) GetAppMetadata

```go
func (m *UserInfo) GetAppMetadata() *AppMetadata
```

#### func (*UserInfo) GetBirthdate

```go
func (m *UserInfo) GetBirthdate() string
```

#### func (*UserInfo) GetEmail

```go
func (m *UserInfo) GetEmail() string
```

#### func (*UserInfo) GetFamilyName

```go
func (m *UserInfo) GetFamilyName() string
```

#### func (*UserInfo) GetGender

```go
func (m *UserInfo) GetGender() string
```

#### func (*UserInfo) GetGivenName

```go
func (m *UserInfo) GetGivenName() string
```

#### func (*UserInfo) GetName

```go
func (m *UserInfo) GetName() string
```

#### func (*UserInfo) GetPicture

```go
func (m *UserInfo) GetPicture() string
```

#### func (*UserInfo) GetUserMetadata

```go
func (m *UserInfo) GetUserMetadata() *UserMetadata
```

#### func (*UserInfo) JSONString

```go
func (p *UserInfo) JSONString() string
```

#### func (*UserInfo) ProtoMessage

```go
func (*UserInfo) ProtoMessage()
```

#### func (*UserInfo) Render

```go
func (p *UserInfo) Render(tmpl *template.Template, w io.Writer) error
```

#### func (*UserInfo) Reset

```go
func (m *UserInfo) Reset()
```

#### func (*UserInfo) String

```go
func (m *UserInfo) String() string
```

#### func (*UserInfo) Validate

```go
func (p *UserInfo) Validate() error
```

#### func (*UserInfo) XXX_DiscardUnknown

```go
func (m *UserInfo) XXX_DiscardUnknown()
```

#### func (*UserInfo) XXX_Marshal

```go
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*UserInfo) XXX_Merge

```go
func (m *UserInfo) XXX_Merge(src proto.Message)
```

#### func (*UserInfo) XXX_Size

```go
func (m *UserInfo) XXX_Size() int
```

#### func (*UserInfo) XXX_Unmarshal

```go
func (m *UserInfo) XXX_Unmarshal(b []byte) error
```

#### type UserMetadata

```go
type UserMetadata struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	PreferredContact     string   `protobuf:"bytes,2,opt,name=preferred_contact,json=preferredContact,proto3" json:"preferred_contact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*UserMetadata) Descriptor

```go
func (*UserMetadata) Descriptor() ([]byte, []int)
```

#### func (*UserMetadata) GetPhone

```go
func (m *UserMetadata) GetPhone() string
```

#### func (*UserMetadata) GetPreferredContact

```go
func (m *UserMetadata) GetPreferredContact() string
```

#### func (*UserMetadata) JSONString

```go
func (p *UserMetadata) JSONString() string
```

#### func (*UserMetadata) ProtoMessage

```go
func (*UserMetadata) ProtoMessage()
```

#### func (*UserMetadata) Render

```go
func (p *UserMetadata) Render(tmpl *template.Template, w io.Writer) error
```

#### func (*UserMetadata) Reset

```go
func (m *UserMetadata) Reset()
```

#### func (*UserMetadata) String

```go
func (m *UserMetadata) String() string
```

#### func (*UserMetadata) XXX_DiscardUnknown

```go
func (m *UserMetadata) XXX_DiscardUnknown()
```

#### func (*UserMetadata) XXX_Marshal

```go
func (m *UserMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*UserMetadata) XXX_Merge

```go
func (m *UserMetadata) XXX_Merge(src proto.Message)
```

#### func (*UserMetadata) XXX_Size

```go
func (m *UserMetadata) XXX_Size() int
```

#### func (*UserMetadata) XXX_Unmarshal

```go
func (m *UserMetadata) XXX_Unmarshal(b []byte) error
```
