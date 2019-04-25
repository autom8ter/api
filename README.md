# api
--
    import "github.com/autom8ter/api"

Package api is a reverse proxy.

It translates gRPC into RESTful JSON APIs.

## Usage

```go
var AUTH_SESSION = "auth-session"
```

```go
var Context context.Context
```

```go
var DEFAULT_OAUTH_REDIRECT = "http://localhost:8080/callback"
```

```go
var DEFAULT_OAUTH_SCOPES = []Scope{Scope_OPENID, Scope_PROFILE, Scope_EMAIL}
```

```go
var HTTPMethod_name = map[int32]string{
	0: "GET",
	1: "POST",
}
```

```go
var HTTPMethod_value = map[string]int32{
	"GET":  0,
	"POST": 1,
}
```

```go
var Scope_name = map[int32]string{
	0:  "OPENID",
	1:  "PROFILE",
	2:  "EMAIL",
	3:  "READ_USERS",
	4:  "READ_USER_IDP_TOKENS",
	5:  "CREATE_USERS",
	6:  "READ_STATS",
	7:  "READ_EMAIL_TEMPLATES",
	8:  "UPDATE_EMAIL_TEMPLATES",
	9:  "CREATE_EMAIL_TEMPLATES",
	10: "READ_RULES",
	11: "UPDATE_RULES",
	12: "CREATE_RULES",
	13: "DELETE_RULES",
	14: "READ_ROLES",
	15: "UPDATE_ROLES",
	16: "CREATE_ROLES",
	17: "DELETE_ROLES",
	18: "READ_LOGS",
}
```

```go
var Scope_value = map[string]int32{
	"OPENID":                 0,
	"PROFILE":                1,
	"EMAIL":                  2,
	"READ_USERS":             3,
	"READ_USER_IDP_TOKENS":   4,
	"CREATE_USERS":           5,
	"READ_STATS":             6,
	"READ_EMAIL_TEMPLATES":   7,
	"UPDATE_EMAIL_TEMPLATES": 8,
	"CREATE_EMAIL_TEMPLATES": 9,
	"READ_RULES":             10,
	"UPDATE_RULES":           11,
	"CREATE_RULES":           12,
	"DELETE_RULES":           13,
	"READ_ROLES":             14,
	"UPDATE_ROLES":           15,
	"CREATE_ROLES":           16,
	"DELETE_ROLES":           17,
	"READ_LOGS":              18,
}
```

```go
var Util = objectify.Default()
```

#### func  ENVFromContext

```go
func ENVFromContext() []string
```

#### func  FuncMap

```go
func FuncMap() template.FuncMap
```

#### func  InitSessions

```go
func InitSessions(secret string) error
```

#### func  NormalizeScopes

```go
func NormalizeScopes(scopes ...Scope) []string
```

#### func  RegisterContactServiceHandler

```go
func RegisterContactServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
```
RegisterContactServiceHandler registers the http handlers for service
ContactService to "mux". The handlers forward requests to the grpc endpoint over
"conn".

#### func  RegisterContactServiceHandlerClient

```go
func RegisterContactServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ContactServiceClient) error
```
RegisterContactServiceHandlerClient registers the http handlers for service
ContactService to "mux". The handlers forward requests to the grpc endpoint over
the given implementation of "ContactServiceClient". Note: the gRPC framework
executes interceptors within the gRPC handler. If the passed in
"ContactServiceClient" doesn't go through the normal gRPC flow (creating a gRPC
client etc.) then it will be up to the passed in "ContactServiceClient" to call
the correct interceptors.

#### func  RegisterContactServiceHandlerFromEndpoint

```go
func RegisterContactServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
```
RegisterContactServiceHandlerFromEndpoint is same as
RegisterContactServiceHandler but automatically dials to "endpoint" and closes
the connection when "ctx" gets done.

#### func  RegisterContactServiceServer

```go
func RegisterContactServiceServer(s *grpc.Server, srv ContactServiceServer)
```

#### func  RegisterPaymentServiceHandler

```go
func RegisterPaymentServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
```
RegisterPaymentServiceHandler registers the http handlers for service
PaymentService to "mux". The handlers forward requests to the grpc endpoint over
"conn".

#### func  RegisterPaymentServiceHandlerClient

```go
func RegisterPaymentServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client PaymentServiceClient) error
```
RegisterPaymentServiceHandlerClient registers the http handlers for service
PaymentService to "mux". The handlers forward requests to the grpc endpoint over
the given implementation of "PaymentServiceClient". Note: the gRPC framework
executes interceptors within the gRPC handler. If the passed in
"PaymentServiceClient" doesn't go through the normal gRPC flow (creating a gRPC
client etc.) then it will be up to the passed in "PaymentServiceClient" to call
the correct interceptors.

#### func  RegisterPaymentServiceHandlerFromEndpoint

```go
func RegisterPaymentServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
```
RegisterPaymentServiceHandlerFromEndpoint is same as
RegisterPaymentServiceHandler but automatically dials to "endpoint" and closes
the connection when "ctx" gets done.

#### func  RegisterPaymentServiceServer

```go
func RegisterPaymentServiceServer(s *grpc.Server, srv PaymentServiceServer)
```

#### func  RegisterUserServiceHandler

```go
func RegisterUserServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
```
RegisterUserServiceHandler registers the http handlers for service UserService
to "mux". The handlers forward requests to the grpc endpoint over "conn".

#### func  RegisterUserServiceHandlerClient

```go
func RegisterUserServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client UserServiceClient) error
```
RegisterUserServiceHandlerClient registers the http handlers for service
UserService to "mux". The handlers forward requests to the grpc endpoint over
the given implementation of "UserServiceClient". Note: the gRPC framework
executes interceptors within the gRPC handler. If the passed in
"UserServiceClient" doesn't go through the normal gRPC flow (creating a gRPC
client etc.) then it will be up to the passed in "UserServiceClient" to call the
correct interceptors.

#### func  RegisterUserServiceHandlerFromEndpoint

```go
func RegisterUserServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
```
RegisterUserServiceHandlerFromEndpoint is same as RegisterUserServiceHandler but
automatically dials to "endpoint" and closes the connection when "ctx" gets
done.

#### func  RegisterUserServiceServer

```go
func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer)
```

#### func  RegisterUtilityServiceHandler

```go
func RegisterUtilityServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
```
RegisterUtilityServiceHandler registers the http handlers for service
UtilityService to "mux". The handlers forward requests to the grpc endpoint over
"conn".

#### func  RegisterUtilityServiceHandlerClient

```go
func RegisterUtilityServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client UtilityServiceClient) error
```
RegisterUtilityServiceHandlerClient registers the http handlers for service
UtilityService to "mux". The handlers forward requests to the grpc endpoint over
the given implementation of "UtilityServiceClient". Note: the gRPC framework
executes interceptors within the gRPC handler. If the passed in
"UtilityServiceClient" doesn't go through the normal gRPC flow (creating a gRPC
client etc.) then it will be up to the passed in "UtilityServiceClient" to call
the correct interceptors.

#### func  RegisterUtilityServiceHandlerFromEndpoint

```go
func RegisterUtilityServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
```
RegisterUtilityServiceHandlerFromEndpoint is same as
RegisterUtilityServiceHandler but automatically dials to "endpoint" and closes
the connection when "ctx" gets done.

#### func  RegisterUtilityServiceServer

```go
func RegisterUtilityServiceServer(s *grpc.Server, srv UtilityServiceServer)
```

#### func  RenderFile

```go
func RenderFile(name string, data []byte) http.HandlerFunc
```

#### func  WriteFile

```go
func WriteFile(name string) http.HandlerFunc
```

#### type AppMetadata

```go
type AppMetadata struct {
	Metadata             map[string]string `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func (*AppMetadata) Descriptor

```go
func (*AppMetadata) Descriptor() ([]byte, []int)
```

#### func (*AppMetadata) GetMetadata

```go
func (m *AppMetadata) GetMetadata() map[string]string
```

#### func (*AppMetadata) ProtoMessage

```go
func (*AppMetadata) ProtoMessage()
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

#### type Auth

```go
type Auth struct {
	Domain               string   `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         string   `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	Redirect             string   `protobuf:"bytes,4,opt,name=redirect,proto3" json:"redirect,omitempty"`
	Audience             string   `protobuf:"bytes,5,opt,name=audience,proto3" json:"audience,omitempty"`
	Scopes               []Scope  `protobuf:"varint,6,rep,packed,name=scopes,proto3,enum=api.Scope" json:"scopes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Auth) AuthURL

```go
func (c *Auth) AuthURL() string
```

#### func (*Auth) ClientGrantsURL

```go
func (c *Auth) ClientGrantsURL() string
```

#### func (*Auth) ClientsURL

```go
func (c *Auth) ClientsURL() string
```

#### func (*Auth) ConnectionsURL

```go
func (c *Auth) ConnectionsURL() string
```

#### func (*Auth) CustomDomainsURL

```go
func (c *Auth) CustomDomainsURL() string
```

#### func (*Auth) Descriptor

```go
func (*Auth) Descriptor() ([]byte, []int)
```

#### func (*Auth) DeviceCredentials

```go
func (c *Auth) DeviceCredentials(name string) string
```

#### func (*Auth) EmailTemplateURL

```go
func (c *Auth) EmailTemplateURL(name string) string
```

#### func (*Auth) EmailURL

```go
func (c *Auth) EmailURL() string
```

#### func (*Auth) GetAudience

```go
func (m *Auth) GetAudience() string
```

#### func (*Auth) GetClientId

```go
func (m *Auth) GetClientId() string
```

#### func (*Auth) GetClientSecret

```go
func (m *Auth) GetClientSecret() string
```

#### func (*Auth) GetDomain

```go
func (m *Auth) GetDomain() string
```

#### func (*Auth) GetJWKS

```go
func (c *Auth) GetJWKS() (*Jwks, error)
```

#### func (*Auth) GetRedirect

```go
func (m *Auth) GetRedirect() string
```

#### func (*Auth) GetScopes

```go
func (m *Auth) GetScopes() []Scope
```

#### func (*Auth) GrantsURL

```go
func (c *Auth) GrantsURL() string
```

#### func (*Auth) JWKSURL

```go
func (c *Auth) JWKSURL() string
```

#### func (*Auth) ListenAndServe

```go
func (a *Auth) ListenAndServe(addr string, c *RouterPaths, home, loggedIn http.HandlerFunc) error
```

#### func (*Auth) LogoutURL

```go
func (c *Auth) LogoutURL(r *RouterPaths) (string, error)
```

#### func (*Auth) LogsURL

```go
func (c *Auth) LogsURL() string
```

#### func (*Auth) ProtoMessage

```go
func (*Auth) ProtoMessage()
```

#### func (*Auth) RequireLogin

```go
func (a *Auth) RequireLogin(paths *RouterPaths, next http.HandlerFunc) http.HandlerFunc
```

#### func (*Auth) Reset

```go
func (m *Auth) Reset()
```

#### func (*Auth) RolesURL

```go
func (c *Auth) RolesURL() string
```

#### func (*Auth) Router

```go
func (a *Auth) Router(c *RouterPaths, home, loggedIn http.HandlerFunc) *mux.Router
```

#### func (*Auth) RulesURL

```go
func (c *Auth) RulesURL() string
```

#### func (*Auth) SearchUsersURL

```go
func (c *Auth) SearchUsersURL() string
```

#### func (*Auth) StatsURL

```go
func (c *Auth) StatsURL() string
```

#### func (*Auth) String

```go
func (m *Auth) String() string
```

#### func (*Auth) TenantURL

```go
func (c *Auth) TenantURL() string
```

#### func (*Auth) TokenURL

```go
func (c *Auth) TokenURL() string
```

#### func (*Auth) UserInfoURL

```go
func (c *Auth) UserInfoURL() string
```

#### func (*Auth) UsersURL

```go
func (c *Auth) UsersURL() string
```

#### func (*Auth) XXX_DiscardUnknown

```go
func (m *Auth) XXX_DiscardUnknown()
```

#### func (*Auth) XXX_Marshal

```go
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Auth) XXX_Merge

```go
func (m *Auth) XXX_Merge(src proto.Message)
```

#### func (*Auth) XXX_Size

```go
func (m *Auth) XXX_Size() int
```

#### func (*Auth) XXX_Unmarshal

```go
func (m *Auth) XXX_Unmarshal(b []byte) error
```

#### type Bytes

```go
type Bytes struct {
	Bits                 []byte   `protobuf:"bytes,1,opt,name=bits,proto3" json:"bits,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  AsBytes

```go
func AsBytes(obj interface{}) *Bytes
```

#### func  BytesFromBytes

```go
func BytesFromBytes(bits []byte) *Bytes
```

#### func  BytesFromFile

```go
func BytesFromFile(fileName string) (*Bytes, error)
```

#### func  BytesFromReader

```go
func BytesFromReader(r io.Reader) (*Bytes, error)
```

#### func  BytesFromString

```go
func BytesFromString(str string) *Bytes
```

#### func  NewBytes

```go
func NewBytes() *Bytes
```

#### func (*Bytes) BitString

```go
func (m *Bytes) BitString() string
```

#### func (*Bytes) Clear

```go
func (m *Bytes) Clear()
```

#### func (*Bytes) Compile

```go
func (m *Bytes) Compile(w io.Writer, t *Template) error
```

#### func (*Bytes) CompileHTTP

```go
func (m *Bytes) CompileHTTP(t *Template) http.HandlerFunc
```

#### func (*Bytes) Contains

```go
func (m *Bytes) Contains(str string) bool
```

#### func (*Bytes) Descriptor

```go
func (*Bytes) Descriptor() ([]byte, []int)
```

#### func (*Bytes) GetBits

```go
func (m *Bytes) GetBits() []byte
```

#### func (*Bytes) JSON

```go
func (m *Bytes) JSON() []byte
```

#### func (*Bytes) Length

```go
func (m *Bytes) Length() int
```

#### func (*Bytes) ProtoMessage

```go
func (*Bytes) ProtoMessage()
```

#### func (*Bytes) Read

```go
func (m *Bytes) Read(p []byte) (n int, err error)
```

#### func (*Bytes) Reset

```go
func (m *Bytes) Reset()
```

#### func (*Bytes) String

```go
func (m *Bytes) String() string
```

#### func (*Bytes) UnMarshalJSON

```go
func (m *Bytes) UnMarshalJSON(obj interface{}) error
```

#### func (*Bytes) UnMarshalProto

```go
func (m *Bytes) UnMarshalProto(obj interface{}) error
```

#### func (*Bytes) Write

```go
func (m *Bytes) Write(p []byte) (n int, err error)
```

#### func (*Bytes) WriteHTTP

```go
func (m *Bytes) WriteHTTP() http.HandlerFunc
```

#### func (*Bytes) XML

```go
func (m *Bytes) XML() []byte
```

#### func (*Bytes) XXX_DiscardUnknown

```go
func (m *Bytes) XXX_DiscardUnknown()
```

#### func (*Bytes) XXX_Marshal

```go
func (m *Bytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Bytes) XXX_Merge

```go
func (m *Bytes) XXX_Merge(src proto.Message)
```

#### func (*Bytes) XXX_Size

```go
func (m *Bytes) XXX_Size() int
```

#### func (*Bytes) XXX_Unmarshal

```go
func (m *Bytes) XXX_Unmarshal(b []byte) error
```

#### func (*Bytes) YAML

```go
func (m *Bytes) YAML() []byte
```

#### type Call

```go
type Call struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	App                  string   `protobuf:"bytes,3,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Call) Descriptor

```go
func (*Call) Descriptor() ([]byte, []int)
```

#### func (*Call) GetApp

```go
func (m *Call) GetApp() string
```

#### func (*Call) GetFrom

```go
func (m *Call) GetFrom() string
```

#### func (*Call) GetTo

```go
func (m *Call) GetTo() string
```

#### func (*Call) ProtoMessage

```go
func (*Call) ProtoMessage()
```

#### func (*Call) Reset

```go
func (m *Call) Reset()
```

#### func (*Call) String

```go
func (m *Call) String() string
```

#### func (*Call) XXX_DiscardUnknown

```go
func (m *Call) XXX_DiscardUnknown()
```

#### func (*Call) XXX_Marshal

```go
func (m *Call) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Call) XXX_Merge

```go
func (m *Call) XXX_Merge(src proto.Message)
```

#### func (*Call) XXX_Size

```go
func (m *Call) XXX_Size() int
```

#### func (*Call) XXX_Unmarshal

```go
func (m *Call) XXX_Unmarshal(b []byte) error
```

#### type CallBlast

```go
type CallBlast struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   []string `protobuf:"bytes,2,rep,name=to,proto3" json:"to,omitempty"`
	App                  string   `protobuf:"bytes,3,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*CallBlast) Descriptor

```go
func (*CallBlast) Descriptor() ([]byte, []int)
```

#### func (*CallBlast) GetApp

```go
func (m *CallBlast) GetApp() string
```

#### func (*CallBlast) GetFrom

```go
func (m *CallBlast) GetFrom() string
```

#### func (*CallBlast) GetTo

```go
func (m *CallBlast) GetTo() []string
```

#### func (*CallBlast) ProtoMessage

```go
func (*CallBlast) ProtoMessage()
```

#### func (*CallBlast) Reset

```go
func (m *CallBlast) Reset()
```

#### func (*CallBlast) String

```go
func (m *CallBlast) String() string
```

#### func (*CallBlast) XXX_DiscardUnknown

```go
func (m *CallBlast) XXX_DiscardUnknown()
```

#### func (*CallBlast) XXX_Marshal

```go
func (m *CallBlast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*CallBlast) XXX_Merge

```go
func (m *CallBlast) XXX_Merge(src proto.Message)
```

#### func (*CallBlast) XXX_Size

```go
func (m *CallBlast) XXX_Size() int
```

#### func (*CallBlast) XXX_Unmarshal

```go
func (m *CallBlast) XXX_Unmarshal(b []byte) error
```

#### type Card

```go
type Card struct {
	Number               string   `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	ExpMonth             string   `protobuf:"bytes,2,opt,name=exp_month,json=expMonth,proto3" json:"exp_month,omitempty"`
	ExpYear              string   `protobuf:"bytes,3,opt,name=exp_year,json=expYear,proto3" json:"exp_year,omitempty"`
	Cvc                  string   `protobuf:"bytes,4,opt,name=cvc,proto3" json:"cvc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Card) Descriptor

```go
func (*Card) Descriptor() ([]byte, []int)
```

#### func (*Card) GetCvc

```go
func (m *Card) GetCvc() string
```

#### func (*Card) GetExpMonth

```go
func (m *Card) GetExpMonth() string
```

#### func (*Card) GetExpYear

```go
func (m *Card) GetExpYear() string
```

#### func (*Card) GetNumber

```go
func (m *Card) GetNumber() string
```

#### func (*Card) ProtoMessage

```go
func (*Card) ProtoMessage()
```

#### func (*Card) Reset

```go
func (m *Card) Reset()
```

#### func (*Card) String

```go
func (m *Card) String() string
```

#### func (*Card) XXX_DiscardUnknown

```go
func (m *Card) XXX_DiscardUnknown()
```

#### func (*Card) XXX_Marshal

```go
func (m *Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Card) XXX_Merge

```go
func (m *Card) XXX_Merge(src proto.Message)
```

#### func (*Card) XXX_Size

```go
func (m *Card) XXX_Size() int
```

#### func (*Card) XXX_Unmarshal

```go
func (m *Card) XXX_Unmarshal(b []byte) error
```

#### type ClientSet

```go
type ClientSet struct {
	Utility UtilityServiceClient
	Contact ContactServiceClient
	User    UserServiceClient
	Payment PaymentServiceClient
}
```


#### func  NewClientSet

```go
func NewClientSet(conn *grpc.ClientConn) *ClientSet
```

#### type ContactServiceClient

```go
type ContactServiceClient interface {
	SendSMS(ctx context.Context, in *SMS, opts ...grpc.CallOption) (*Bytes, error)
	SendSMSBlast(ctx context.Context, in *SMSBlast, opts ...grpc.CallOption) (ContactService_SendSMSBlastClient, error)
	GetSMS(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Bytes, error)
	SendEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*Message, error)
	SendEmailBlast(ctx context.Context, in *EmailBlastRequest, opts ...grpc.CallOption) (ContactService_SendEmailBlastClient, error)
	SendCall(ctx context.Context, in *Call, opts ...grpc.CallOption) (*Bytes, error)
	SendCallBlast(ctx context.Context, in *CallBlast, opts ...grpc.CallOption) (ContactService_SendCallBlastClient, error)
}
```

ContactServiceClient is the client API for ContactService service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to
https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

#### func  NewContactServiceClient

```go
func NewContactServiceClient(cc *grpc.ClientConn) ContactServiceClient
```

#### type ContactServiceServer

```go
type ContactServiceServer interface {
	SendSMS(context.Context, *SMS) (*Bytes, error)
	SendSMSBlast(*SMSBlast, ContactService_SendSMSBlastServer) error
	GetSMS(context.Context, *Identifier) (*Bytes, error)
	SendEmail(context.Context, *EmailRequest) (*Message, error)
	SendEmailBlast(*EmailBlastRequest, ContactService_SendEmailBlastServer) error
	SendCall(context.Context, *Call) (*Bytes, error)
	SendCallBlast(*CallBlast, ContactService_SendCallBlastServer) error
}
```

ContactServiceServer is the server API for ContactService service.

#### type ContactService_SendCallBlastClient

```go
type ContactService_SendCallBlastClient interface {
	Recv() (*Bytes, error)
	grpc.ClientStream
}
```


#### type ContactService_SendCallBlastServer

```go
type ContactService_SendCallBlastServer interface {
	Send(*Bytes) error
	grpc.ServerStream
}
```


#### type ContactService_SendEmailBlastClient

```go
type ContactService_SendEmailBlastClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}
```


#### type ContactService_SendEmailBlastServer

```go
type ContactService_SendEmailBlastServer interface {
	Send(*Message) error
	grpc.ServerStream
}
```


#### type ContactService_SendSMSBlastClient

```go
type ContactService_SendSMSBlastClient interface {
	Recv() (*Bytes, error)
	grpc.ClientStream
}
```


#### type ContactService_SendSMSBlastServer

```go
type ContactService_SendSMSBlastServer interface {
	Send(*Bytes) error
	grpc.ServerStream
}
```


#### type Email

```go
type Email struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Subject              string   `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Plain                string   `protobuf:"bytes,4,opt,name=plain,proto3" json:"plain,omitempty"`
	Html                 string   `protobuf:"bytes,5,opt,name=html,proto3" json:"html,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Email) Descriptor

```go
func (*Email) Descriptor() ([]byte, []int)
```

#### func (*Email) GetAddress

```go
func (m *Email) GetAddress() string
```

#### func (*Email) GetHtml

```go
func (m *Email) GetHtml() string
```

#### func (*Email) GetName

```go
func (m *Email) GetName() string
```

#### func (*Email) GetPlain

```go
func (m *Email) GetPlain() string
```

#### func (*Email) GetSubject

```go
func (m *Email) GetSubject() string
```

#### func (*Email) ProtoMessage

```go
func (*Email) ProtoMessage()
```

#### func (*Email) Reset

```go
func (m *Email) Reset()
```

#### func (*Email) String

```go
func (m *Email) String() string
```

#### func (*Email) XXX_DiscardUnknown

```go
func (m *Email) XXX_DiscardUnknown()
```

#### func (*Email) XXX_Marshal

```go
func (m *Email) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Email) XXX_Merge

```go
func (m *Email) XXX_Merge(src proto.Message)
```

#### func (*Email) XXX_Size

```go
func (m *Email) XXX_Size() int
```

#### func (*Email) XXX_Unmarshal

```go
func (m *Email) XXX_Unmarshal(b []byte) error
```

#### type EmailBlast

```go
type EmailBlast struct {
	NameAddress          map[string]string `protobuf:"bytes,1,rep,name=name_address,json=nameAddress,proto3" json:"name_address,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Subject              string            `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Plain                string            `protobuf:"bytes,3,opt,name=plain,proto3" json:"plain,omitempty"`
	Html                 string            `protobuf:"bytes,4,opt,name=html,proto3" json:"html,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func (*EmailBlast) Descriptor

```go
func (*EmailBlast) Descriptor() ([]byte, []int)
```

#### func (*EmailBlast) GetHtml

```go
func (m *EmailBlast) GetHtml() string
```

#### func (*EmailBlast) GetNameAddress

```go
func (m *EmailBlast) GetNameAddress() map[string]string
```

#### func (*EmailBlast) GetPlain

```go
func (m *EmailBlast) GetPlain() string
```

#### func (*EmailBlast) GetSubject

```go
func (m *EmailBlast) GetSubject() string
```

#### func (*EmailBlast) ProtoMessage

```go
func (*EmailBlast) ProtoMessage()
```

#### func (*EmailBlast) Reset

```go
func (m *EmailBlast) Reset()
```

#### func (*EmailBlast) String

```go
func (m *EmailBlast) String() string
```

#### func (*EmailBlast) XXX_DiscardUnknown

```go
func (m *EmailBlast) XXX_DiscardUnknown()
```

#### func (*EmailBlast) XXX_Marshal

```go
func (m *EmailBlast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*EmailBlast) XXX_Merge

```go
func (m *EmailBlast) XXX_Merge(src proto.Message)
```

#### func (*EmailBlast) XXX_Size

```go
func (m *EmailBlast) XXX_Size() int
```

#### func (*EmailBlast) XXX_Unmarshal

```go
func (m *EmailBlast) XXX_Unmarshal(b []byte) error
```

#### type EmailBlastRequest

```go
type EmailBlastRequest struct {
	FromName             string      `protobuf:"bytes,1,opt,name=from_name,json=fromName,proto3" json:"from_name,omitempty"`
	FromEmail            string      `protobuf:"bytes,2,opt,name=from_email,json=fromEmail,proto3" json:"from_email,omitempty"`
	Blast                *EmailBlast `protobuf:"bytes,3,opt,name=blast,proto3" json:"blast,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}
```


#### func (*EmailBlastRequest) Descriptor

```go
func (*EmailBlastRequest) Descriptor() ([]byte, []int)
```

#### func (*EmailBlastRequest) GetBlast

```go
func (m *EmailBlastRequest) GetBlast() *EmailBlast
```

#### func (*EmailBlastRequest) GetFromEmail

```go
func (m *EmailBlastRequest) GetFromEmail() string
```

#### func (*EmailBlastRequest) GetFromName

```go
func (m *EmailBlastRequest) GetFromName() string
```

#### func (*EmailBlastRequest) ProtoMessage

```go
func (*EmailBlastRequest) ProtoMessage()
```

#### func (*EmailBlastRequest) Reset

```go
func (m *EmailBlastRequest) Reset()
```

#### func (*EmailBlastRequest) String

```go
func (m *EmailBlastRequest) String() string
```

#### func (*EmailBlastRequest) XXX_DiscardUnknown

```go
func (m *EmailBlastRequest) XXX_DiscardUnknown()
```

#### func (*EmailBlastRequest) XXX_Marshal

```go
func (m *EmailBlastRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*EmailBlastRequest) XXX_Merge

```go
func (m *EmailBlastRequest) XXX_Merge(src proto.Message)
```

#### func (*EmailBlastRequest) XXX_Size

```go
func (m *EmailBlastRequest) XXX_Size() int
```

#### func (*EmailBlastRequest) XXX_Unmarshal

```go
func (m *EmailBlastRequest) XXX_Unmarshal(b []byte) error
```

#### type EmailRequest

```go
type EmailRequest struct {
	FromName             string   `protobuf:"bytes,1,opt,name=from_name,json=fromName,proto3" json:"from_name,omitempty"`
	FromEmail            string   `protobuf:"bytes,2,opt,name=from_email,json=fromEmail,proto3" json:"from_email,omitempty"`
	Email                *Email   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*EmailRequest) Descriptor

```go
func (*EmailRequest) Descriptor() ([]byte, []int)
```

#### func (*EmailRequest) GetEmail

```go
func (m *EmailRequest) GetEmail() *Email
```

#### func (*EmailRequest) GetFromEmail

```go
func (m *EmailRequest) GetFromEmail() string
```

#### func (*EmailRequest) GetFromName

```go
func (m *EmailRequest) GetFromName() string
```

#### func (*EmailRequest) ProtoMessage

```go
func (*EmailRequest) ProtoMessage()
```

#### func (*EmailRequest) Reset

```go
func (m *EmailRequest) Reset()
```

#### func (*EmailRequest) String

```go
func (m *EmailRequest) String() string
```

#### func (*EmailRequest) XXX_DiscardUnknown

```go
func (m *EmailRequest) XXX_DiscardUnknown()
```

#### func (*EmailRequest) XXX_Marshal

```go
func (m *EmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*EmailRequest) XXX_Merge

```go
func (m *EmailRequest) XXX_Merge(src proto.Message)
```

#### func (*EmailRequest) XXX_Size

```go
func (m *EmailRequest) XXX_Size() int
```

#### func (*EmailRequest) XXX_Unmarshal

```go
func (m *EmailRequest) XXX_Unmarshal(b []byte) error
```

#### type Empty

```go
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Empty) Descriptor

```go
func (*Empty) Descriptor() ([]byte, []int)
```

#### func (*Empty) ProtoMessage

```go
func (*Empty) ProtoMessage()
```

#### func (*Empty) Reset

```go
func (m *Empty) Reset()
```

#### func (*Empty) String

```go
func (m *Empty) String() string
```

#### func (*Empty) XXX_DiscardUnknown

```go
func (m *Empty) XXX_DiscardUnknown()
```

#### func (*Empty) XXX_Marshal

```go
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Empty) XXX_Merge

```go
func (m *Empty) XXX_Merge(src proto.Message)
```

#### func (*Empty) XXX_Size

```go
func (m *Empty) XXX_Size() int
```

#### func (*Empty) XXX_Unmarshal

```go
func (m *Empty) XXX_Unmarshal(b []byte) error
```

#### type HTTPMethod

```go
type HTTPMethod int32
```


```go
const (
	HTTPMethod_GET  HTTPMethod = 0
	HTTPMethod_POST HTTPMethod = 1
)
```

#### func (HTTPMethod) EnumDescriptor

```go
func (HTTPMethod) EnumDescriptor() ([]byte, []int)
```

#### func (HTTPMethod) Normalize

```go
func (h HTTPMethod) Normalize() string
```

#### func (HTTPMethod) String

```go
func (x HTTPMethod) String() string
```

#### type HTTPRequest

```go
type HTTPRequest struct {
	Method               HTTPMethod        `protobuf:"varint,1,opt,name=method,proto3,enum=api.HTTPMethod" json:"method,omitempty"`
	Url                  string            `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Token                string            `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Account              string            `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	ContentType          string            `protobuf:"bytes,5,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Headers              map[string]string `protobuf:"bytes,6,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Form                 map[string]string `protobuf:"bytes,7,rep,name=form,proto3" json:"form,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Cookies              map[string]string `protobuf:"bytes,8,rep,name=cookies,proto3" json:"cookies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 *Bytes            `protobuf:"bytes,9,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func (*HTTPRequest) Descriptor

```go
func (*HTTPRequest) Descriptor() ([]byte, []int)
```

#### func (*HTTPRequest) Do

```go
func (h *HTTPRequest) Do() (*Bytes, error)
```

#### func (*HTTPRequest) GetAccount

```go
func (m *HTTPRequest) GetAccount() string
```

#### func (*HTTPRequest) GetBody

```go
func (m *HTTPRequest) GetBody() *Bytes
```

#### func (*HTTPRequest) GetContentType

```go
func (m *HTTPRequest) GetContentType() string
```

#### func (*HTTPRequest) GetCookies

```go
func (m *HTTPRequest) GetCookies() map[string]string
```

#### func (*HTTPRequest) GetForm

```go
func (m *HTTPRequest) GetForm() map[string]string
```

#### func (*HTTPRequest) GetHeaders

```go
func (m *HTTPRequest) GetHeaders() map[string]string
```

#### func (*HTTPRequest) GetMethod

```go
func (m *HTTPRequest) GetMethod() HTTPMethod
```

#### func (*HTTPRequest) GetToken

```go
func (m *HTTPRequest) GetToken() string
```

#### func (*HTTPRequest) GetUrl

```go
func (m *HTTPRequest) GetUrl() string
```

#### func (*HTTPRequest) ProtoMessage

```go
func (*HTTPRequest) ProtoMessage()
```

#### func (*HTTPRequest) Reset

```go
func (m *HTTPRequest) Reset()
```

#### func (*HTTPRequest) String

```go
func (m *HTTPRequest) String() string
```

#### func (*HTTPRequest) XXX_DiscardUnknown

```go
func (m *HTTPRequest) XXX_DiscardUnknown()
```

#### func (*HTTPRequest) XXX_Marshal

```go
func (m *HTTPRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*HTTPRequest) XXX_Merge

```go
func (m *HTTPRequest) XXX_Merge(src proto.Message)
```

#### func (*HTTPRequest) XXX_Size

```go
func (m *HTTPRequest) XXX_Size() int
```

#### func (*HTTPRequest) XXX_Unmarshal

```go
func (m *HTTPRequest) XXX_Unmarshal(b []byte) error
```

#### type Identifier

```go
type Identifier struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Identifier) Descriptor

```go
func (*Identifier) Descriptor() ([]byte, []int)
```

#### func (*Identifier) GetId

```go
func (m *Identifier) GetId() string
```

#### func (*Identifier) ProtoMessage

```go
func (*Identifier) ProtoMessage()
```

#### func (*Identifier) Reset

```go
func (m *Identifier) Reset()
```

#### func (*Identifier) String

```go
func (m *Identifier) String() string
```

#### func (*Identifier) XXX_DiscardUnknown

```go
func (m *Identifier) XXX_DiscardUnknown()
```

#### func (*Identifier) XXX_Marshal

```go
func (m *Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Identifier) XXX_Merge

```go
func (m *Identifier) XXX_Merge(src proto.Message)
```

#### func (*Identifier) XXX_Size

```go
func (m *Identifier) XXX_Size() int
```

#### func (*Identifier) XXX_Unmarshal

```go
func (m *Identifier) XXX_Unmarshal(b []byte) error
```

#### type Identity

```go
type Identity struct {
	Connection           string   `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Provider             string   `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	IsSocial             string   `protobuf:"bytes,4,opt,name=isSocial,proto3" json:"isSocial,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Identity) Descriptor

```go
func (*Identity) Descriptor() ([]byte, []int)
```

#### func (*Identity) GetConnection

```go
func (m *Identity) GetConnection() string
```

#### func (*Identity) GetIsSocial

```go
func (m *Identity) GetIsSocial() string
```

#### func (*Identity) GetProvider

```go
func (m *Identity) GetProvider() string
```

#### func (*Identity) GetUserId

```go
func (m *Identity) GetUserId() string
```

#### func (*Identity) ProtoMessage

```go
func (*Identity) ProtoMessage()
```

#### func (*Identity) Reset

```go
func (m *Identity) Reset()
```

#### func (*Identity) String

```go
func (m *Identity) String() string
```

#### func (*Identity) XXX_DiscardUnknown

```go
func (m *Identity) XXX_DiscardUnknown()
```

#### func (*Identity) XXX_Marshal

```go
func (m *Identity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Identity) XXX_Merge

```go
func (m *Identity) XXX_Merge(src proto.Message)
```

#### func (*Identity) XXX_Size

```go
func (m *Identity) XXX_Size() int
```

#### func (*Identity) XXX_Unmarshal

```go
func (m *Identity) XXX_Unmarshal(b []byte) error
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

#### func (*Jwks) GetCert

```go
func (c *Jwks) GetCert(token *jwt.Token) (string, error)
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

#### type ManagementToken

```go
type ManagementToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*ManagementToken) Descriptor

```go
func (*ManagementToken) Descriptor() ([]byte, []int)
```

#### func (*ManagementToken) GetToken

```go
func (m *ManagementToken) GetToken() string
```

#### func (*ManagementToken) ProtoMessage

```go
func (*ManagementToken) ProtoMessage()
```

#### func (*ManagementToken) Reset

```go
func (m *ManagementToken) Reset()
```

#### func (*ManagementToken) String

```go
func (m *ManagementToken) String() string
```

#### func (*ManagementToken) XXX_DiscardUnknown

```go
func (m *ManagementToken) XXX_DiscardUnknown()
```

#### func (*ManagementToken) XXX_Marshal

```go
func (m *ManagementToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*ManagementToken) XXX_Merge

```go
func (m *ManagementToken) XXX_Merge(src proto.Message)
```

#### func (*ManagementToken) XXX_Size

```go
func (m *ManagementToken) XXX_Size() int
```

#### func (*ManagementToken) XXX_Unmarshal

```go
func (m *ManagementToken) XXX_Unmarshal(b []byte) error
```

#### type Message

```go
type Message struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Message) Descriptor

```go
func (*Message) Descriptor() ([]byte, []int)
```

#### func (*Message) GetValue

```go
func (m *Message) GetValue() string
```

#### func (*Message) ProtoMessage

```go
func (*Message) ProtoMessage()
```

#### func (*Message) Reset

```go
func (m *Message) Reset()
```

#### func (*Message) String

```go
func (m *Message) String() string
```

#### func (*Message) XXX_DiscardUnknown

```go
func (m *Message) XXX_DiscardUnknown()
```

#### func (*Message) XXX_Marshal

```go
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Message) XXX_Merge

```go
func (m *Message) XXX_Merge(src proto.Message)
```

#### func (*Message) XXX_Size

```go
func (m *Message) XXX_Size() int
```

#### func (*Message) XXX_Unmarshal

```go
func (m *Message) XXX_Unmarshal(b []byte) error
```

#### type PaymentServiceClient

```go
type PaymentServiceClient interface {
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*Bytes, error)
	Unsubscribe(ctx context.Context, in *UnSubscribeRequest, opts ...grpc.CallOption) (*Bytes, error)
}
```

PaymentServiceClient is the client API for PaymentService service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to
https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

#### func  NewPaymentServiceClient

```go
func NewPaymentServiceClient(cc *grpc.ClientConn) PaymentServiceClient
```

#### type PaymentServiceServer

```go
type PaymentServiceServer interface {
	Subscribe(context.Context, *SubscribeRequest) (*Bytes, error)
	Unsubscribe(context.Context, *UnSubscribeRequest) (*Bytes, error)
}
```

PaymentServiceServer is the server API for PaymentService service.

#### type RenderRequest

```go
type RenderRequest struct {
	Template             *Template `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	Data                 *Bytes    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}
```


#### func (*RenderRequest) Descriptor

```go
func (*RenderRequest) Descriptor() ([]byte, []int)
```

#### func (*RenderRequest) GetData

```go
func (m *RenderRequest) GetData() *Bytes
```

#### func (*RenderRequest) GetTemplate

```go
func (m *RenderRequest) GetTemplate() *Template
```

#### func (*RenderRequest) ProtoMessage

```go
func (*RenderRequest) ProtoMessage()
```

#### func (*RenderRequest) Reset

```go
func (m *RenderRequest) Reset()
```

#### func (*RenderRequest) String

```go
func (m *RenderRequest) String() string
```

#### func (*RenderRequest) XXX_DiscardUnknown

```go
func (m *RenderRequest) XXX_DiscardUnknown()
```

#### func (*RenderRequest) XXX_Marshal

```go
func (m *RenderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*RenderRequest) XXX_Merge

```go
func (m *RenderRequest) XXX_Merge(src proto.Message)
```

#### func (*RenderRequest) XXX_Size

```go
func (m *RenderRequest) XXX_Size() int
```

#### func (*RenderRequest) XXX_Unmarshal

```go
func (m *RenderRequest) XXX_Unmarshal(b []byte) error
```

#### type RouterPaths

```go
type RouterPaths struct {
	HomePath     string
	LoggedInPath string
	LoginPath    string
	LogoutPath   string
	CallbackPath string
	HomeURL      string
}
```


#### type SMS

```go
type SMS struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Message              *Message `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	MediaURL             string   `protobuf:"bytes,4,opt,name=mediaURL,proto3" json:"mediaURL,omitempty"`
	Callback             string   `protobuf:"bytes,5,opt,name=callback,proto3" json:"callback,omitempty"`
	App                  string   `protobuf:"bytes,6,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*SMS) Descriptor

```go
func (*SMS) Descriptor() ([]byte, []int)
```

#### func (*SMS) GetApp

```go
func (m *SMS) GetApp() string
```

#### func (*SMS) GetCallback

```go
func (m *SMS) GetCallback() string
```

#### func (*SMS) GetMediaURL

```go
func (m *SMS) GetMediaURL() string
```

#### func (*SMS) GetMessage

```go
func (m *SMS) GetMessage() *Message
```

#### func (*SMS) GetService

```go
func (m *SMS) GetService() string
```

#### func (*SMS) GetTo

```go
func (m *SMS) GetTo() string
```

#### func (*SMS) ProtoMessage

```go
func (*SMS) ProtoMessage()
```

#### func (*SMS) Reset

```go
func (m *SMS) Reset()
```

#### func (*SMS) String

```go
func (m *SMS) String() string
```

#### func (*SMS) XXX_DiscardUnknown

```go
func (m *SMS) XXX_DiscardUnknown()
```

#### func (*SMS) XXX_Marshal

```go
func (m *SMS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*SMS) XXX_Merge

```go
func (m *SMS) XXX_Merge(src proto.Message)
```

#### func (*SMS) XXX_Size

```go
func (m *SMS) XXX_Size() int
```

#### func (*SMS) XXX_Unmarshal

```go
func (m *SMS) XXX_Unmarshal(b []byte) error
```

#### type SMSBlast

```go
type SMSBlast struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	To                   []string `protobuf:"bytes,2,rep,name=to,proto3" json:"to,omitempty"`
	Message              *Message `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	MediaURL             string   `protobuf:"bytes,4,opt,name=mediaURL,proto3" json:"mediaURL,omitempty"`
	Callback             string   `protobuf:"bytes,5,opt,name=callback,proto3" json:"callback,omitempty"`
	App                  string   `protobuf:"bytes,6,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*SMSBlast) Descriptor

```go
func (*SMSBlast) Descriptor() ([]byte, []int)
```

#### func (*SMSBlast) GetApp

```go
func (m *SMSBlast) GetApp() string
```

#### func (*SMSBlast) GetCallback

```go
func (m *SMSBlast) GetCallback() string
```

#### func (*SMSBlast) GetMediaURL

```go
func (m *SMSBlast) GetMediaURL() string
```

#### func (*SMSBlast) GetMessage

```go
func (m *SMSBlast) GetMessage() *Message
```

#### func (*SMSBlast) GetService

```go
func (m *SMSBlast) GetService() string
```

#### func (*SMSBlast) GetTo

```go
func (m *SMSBlast) GetTo() []string
```

#### func (*SMSBlast) ProtoMessage

```go
func (*SMSBlast) ProtoMessage()
```

#### func (*SMSBlast) Reset

```go
func (m *SMSBlast) Reset()
```

#### func (*SMSBlast) String

```go
func (m *SMSBlast) String() string
```

#### func (*SMSBlast) XXX_DiscardUnknown

```go
func (m *SMSBlast) XXX_DiscardUnknown()
```

#### func (*SMSBlast) XXX_Marshal

```go
func (m *SMSBlast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*SMSBlast) XXX_Merge

```go
func (m *SMSBlast) XXX_Merge(src proto.Message)
```

#### func (*SMSBlast) XXX_Size

```go
func (m *SMSBlast) XXX_Size() int
```

#### func (*SMSBlast) XXX_Unmarshal

```go
func (m *SMSBlast) XXX_Unmarshal(b []byte) error
```

#### type Scope

```go
type Scope int32
```


```go
const (
	Scope_OPENID                 Scope = 0
	Scope_PROFILE                Scope = 1
	Scope_EMAIL                  Scope = 2
	Scope_READ_USERS             Scope = 3
	Scope_READ_USER_IDP_TOKENS   Scope = 4
	Scope_CREATE_USERS           Scope = 5
	Scope_READ_STATS             Scope = 6
	Scope_READ_EMAIL_TEMPLATES   Scope = 7
	Scope_UPDATE_EMAIL_TEMPLATES Scope = 8
	Scope_CREATE_EMAIL_TEMPLATES Scope = 9
	Scope_READ_RULES             Scope = 10
	Scope_UPDATE_RULES           Scope = 11
	Scope_CREATE_RULES           Scope = 12
	Scope_DELETE_RULES           Scope = 13
	Scope_READ_ROLES             Scope = 14
	Scope_UPDATE_ROLES           Scope = 15
	Scope_CREATE_ROLES           Scope = 16
	Scope_DELETE_ROLES           Scope = 17
	Scope_READ_LOGS              Scope = 18
)
```

#### func (Scope) EnumDescriptor

```go
func (Scope) EnumDescriptor() ([]byte, []int)
```

#### func (Scope) Normalize

```go
func (s Scope) Normalize() string
```

#### func (Scope) String

```go
func (x Scope) String() string
```

#### type SubscribeRequest

```go
type SubscribeRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Plan                 string   `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	Card                 *Card    `protobuf:"bytes,3,opt,name=card,proto3" json:"card,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*SubscribeRequest) Descriptor

```go
func (*SubscribeRequest) Descriptor() ([]byte, []int)
```

#### func (*SubscribeRequest) GetCard

```go
func (m *SubscribeRequest) GetCard() *Card
```

#### func (*SubscribeRequest) GetEmail

```go
func (m *SubscribeRequest) GetEmail() string
```

#### func (*SubscribeRequest) GetPlan

```go
func (m *SubscribeRequest) GetPlan() string
```

#### func (*SubscribeRequest) ProtoMessage

```go
func (*SubscribeRequest) ProtoMessage()
```

#### func (*SubscribeRequest) Reset

```go
func (m *SubscribeRequest) Reset()
```

#### func (*SubscribeRequest) String

```go
func (m *SubscribeRequest) String() string
```

#### func (*SubscribeRequest) XXX_DiscardUnknown

```go
func (m *SubscribeRequest) XXX_DiscardUnknown()
```

#### func (*SubscribeRequest) XXX_Marshal

```go
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*SubscribeRequest) XXX_Merge

```go
func (m *SubscribeRequest) XXX_Merge(src proto.Message)
```

#### func (*SubscribeRequest) XXX_Size

```go
func (m *SubscribeRequest) XXX_Size() int
```

#### func (*SubscribeRequest) XXX_Unmarshal

```go
func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error
```

#### type Template

```go
type Template struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  NewTemplateFromFile

```go
func NewTemplateFromFile(filename string) (*Template, error)
```

#### func (*Template) AsTemplate

```go
func (m *Template) AsTemplate() (*template.Template, error)
```

#### func (*Template) Descriptor

```go
func (*Template) Descriptor() ([]byte, []int)
```

#### func (*Template) GetName

```go
func (m *Template) GetName() string
```

#### func (*Template) GetText

```go
func (m *Template) GetText() string
```

#### func (*Template) IsTemplate

```go
func (m *Template) IsTemplate() bool
```

#### func (*Template) ProtoMessage

```go
func (*Template) ProtoMessage()
```

#### func (*Template) Render

```go
func (m *Template) Render(w io.Writer, data interface{}) error
```

#### func (*Template) RenderBytes

```go
func (m *Template) RenderBytes(w io.Writer, bits *Bytes) error
```

#### func (*Template) RenderUser

```go
func (t *Template) RenderUser() http.HandlerFunc
```

#### func (*Template) Reset

```go
func (m *Template) Reset()
```

#### func (*Template) String

```go
func (m *Template) String() string
```

#### func (*Template) XXX_DiscardUnknown

```go
func (m *Template) XXX_DiscardUnknown()
```

#### func (*Template) XXX_Marshal

```go
func (m *Template) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Template) XXX_Merge

```go
func (m *Template) XXX_Merge(src proto.Message)
```

#### func (*Template) XXX_Size

```go
func (m *Template) XXX_Size() int
```

#### func (*Template) XXX_Unmarshal

```go
func (m *Template) XXX_Unmarshal(b []byte) error
```

#### type UnSubscribeRequest

```go
type UnSubscribeRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Plan                 string   `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*UnSubscribeRequest) Descriptor

```go
func (*UnSubscribeRequest) Descriptor() ([]byte, []int)
```

#### func (*UnSubscribeRequest) GetEmail

```go
func (m *UnSubscribeRequest) GetEmail() string
```

#### func (*UnSubscribeRequest) GetPlan

```go
func (m *UnSubscribeRequest) GetPlan() string
```

#### func (*UnSubscribeRequest) ProtoMessage

```go
func (*UnSubscribeRequest) ProtoMessage()
```

#### func (*UnSubscribeRequest) Reset

```go
func (m *UnSubscribeRequest) Reset()
```

#### func (*UnSubscribeRequest) String

```go
func (m *UnSubscribeRequest) String() string
```

#### func (*UnSubscribeRequest) XXX_DiscardUnknown

```go
func (m *UnSubscribeRequest) XXX_DiscardUnknown()
```

#### func (*UnSubscribeRequest) XXX_Marshal

```go
func (m *UnSubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*UnSubscribeRequest) XXX_Merge

```go
func (m *UnSubscribeRequest) XXX_Merge(src proto.Message)
```

#### func (*UnSubscribeRequest) XXX_Size

```go
func (m *UnSubscribeRequest) XXX_Size() int
```

#### func (*UnSubscribeRequest) XXX_Unmarshal

```go
func (m *UnSubscribeRequest) XXX_Unmarshal(b []byte) error
```

#### type User

```go
type User struct {
	UserId               string        `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GivenName            string        `protobuf:"bytes,3,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	FamilyName           string        `protobuf:"bytes,4,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	Gender               string        `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthdate            string        `protobuf:"bytes,6,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	Email                string        `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber          string        `protobuf:"bytes,8,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Picture              string        `protobuf:"bytes,9,opt,name=picture,proto3" json:"picture,omitempty"`
	UserMetadata         *UserMetadata `protobuf:"bytes,10,opt,name=user_metadata,json=userMetadata,proto3" json:"user_metadata,omitempty"`
	AppMetadata          *AppMetadata  `protobuf:"bytes,11,opt,name=app_metadata,json=appMetadata,proto3" json:"app_metadata,omitempty"`
	LastIp               string        `protobuf:"bytes,12,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
	Blocked              bool          `protobuf:"varint,13,opt,name=blocked,proto3" json:"blocked,omitempty"`
	Nickname             string        `protobuf:"bytes,14,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Multifactor          []string      `protobuf:"bytes,15,rep,name=multifactor,proto3" json:"multifactor,omitempty"`
	CreatedAt            string        `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string        `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	PhoneVerified        bool          `protobuf:"varint,19,opt,name=phone_verified,json=phoneVerified,proto3" json:"phone_verified,omitempty"`
	Identities           []*Identity   `protobuf:"bytes,20,rep,name=identities,proto3" json:"identities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
```


#### func (*User) Descriptor

```go
func (*User) Descriptor() ([]byte, []int)
```

#### func (*User) GetAppMetadata

```go
func (m *User) GetAppMetadata() *AppMetadata
```

#### func (*User) GetBirthdate

```go
func (m *User) GetBirthdate() string
```

#### func (*User) GetBlocked

```go
func (m *User) GetBlocked() bool
```

#### func (*User) GetCreatedAt

```go
func (m *User) GetCreatedAt() string
```

#### func (*User) GetEmail

```go
func (m *User) GetEmail() string
```

#### func (*User) GetFamilyName

```go
func (m *User) GetFamilyName() string
```

#### func (*User) GetGender

```go
func (m *User) GetGender() string
```

#### func (*User) GetGivenName

```go
func (m *User) GetGivenName() string
```

#### func (*User) GetIdentities

```go
func (m *User) GetIdentities() []*Identity
```

#### func (*User) GetLastIp

```go
func (m *User) GetLastIp() string
```

#### func (*User) GetMultifactor

```go
func (m *User) GetMultifactor() []string
```

#### func (*User) GetName

```go
func (m *User) GetName() string
```

#### func (*User) GetNickname

```go
func (m *User) GetNickname() string
```

#### func (*User) GetPhoneNumber

```go
func (m *User) GetPhoneNumber() string
```

#### func (*User) GetPhoneVerified

```go
func (m *User) GetPhoneVerified() bool
```

#### func (*User) GetPicture

```go
func (m *User) GetPicture() string
```

#### func (*User) GetUpdatedAt

```go
func (m *User) GetUpdatedAt() string
```

#### func (*User) GetUserId

```go
func (m *User) GetUserId() string
```

#### func (*User) GetUserMetadata

```go
func (m *User) GetUserMetadata() *UserMetadata
```

#### func (*User) ProtoMessage

```go
func (*User) ProtoMessage()
```

#### func (*User) Reset

```go
func (m *User) Reset()
```

#### func (*User) String

```go
func (m *User) String() string
```

#### func (*User) XXX_DiscardUnknown

```go
func (m *User) XXX_DiscardUnknown()
```

#### func (*User) XXX_Marshal

```go
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*User) XXX_Merge

```go
func (m *User) XXX_Merge(src proto.Message)
```

#### func (*User) XXX_Size

```go
func (m *User) XXX_Size() int
```

#### func (*User) XXX_Unmarshal

```go
func (m *User) XXX_Unmarshal(b []byte) error
```

#### type UserByEmailRequest

```go
type UserByEmailRequest struct {
	Token                *ManagementToken `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Email                string           `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}
```


#### func (*UserByEmailRequest) Descriptor

```go
func (*UserByEmailRequest) Descriptor() ([]byte, []int)
```

#### func (*UserByEmailRequest) GetEmail

```go
func (m *UserByEmailRequest) GetEmail() string
```

#### func (*UserByEmailRequest) GetToken

```go
func (m *UserByEmailRequest) GetToken() *ManagementToken
```

#### func (*UserByEmailRequest) ProtoMessage

```go
func (*UserByEmailRequest) ProtoMessage()
```

#### func (*UserByEmailRequest) Reset

```go
func (m *UserByEmailRequest) Reset()
```

#### func (*UserByEmailRequest) String

```go
func (m *UserByEmailRequest) String() string
```

#### func (*UserByEmailRequest) XXX_DiscardUnknown

```go
func (m *UserByEmailRequest) XXX_DiscardUnknown()
```

#### func (*UserByEmailRequest) XXX_Marshal

```go
func (m *UserByEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*UserByEmailRequest) XXX_Merge

```go
func (m *UserByEmailRequest) XXX_Merge(src proto.Message)
```

#### func (*UserByEmailRequest) XXX_Size

```go
func (m *UserByEmailRequest) XXX_Size() int
```

#### func (*UserByEmailRequest) XXX_Unmarshal

```go
func (m *UserByEmailRequest) XXX_Unmarshal(b []byte) error
```

#### type UserMetadata

```go
type UserMetadata struct {
	Metadata             map[string]string `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func (*UserMetadata) Descriptor

```go
func (*UserMetadata) Descriptor() ([]byte, []int)
```

#### func (*UserMetadata) GetMetadata

```go
func (m *UserMetadata) GetMetadata() map[string]string
```

#### func (*UserMetadata) ProtoMessage

```go
func (*UserMetadata) ProtoMessage()
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

#### type UserRequest

```go
type UserRequest struct {
	String_              *ManagementToken `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`
	User                 *User            `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}
```


#### func (*UserRequest) Descriptor

```go
func (*UserRequest) Descriptor() ([]byte, []int)
```

#### func (*UserRequest) GetString_

```go
func (m *UserRequest) GetString_() *ManagementToken
```

#### func (*UserRequest) GetUser

```go
func (m *UserRequest) GetUser() *User
```

#### func (*UserRequest) ProtoMessage

```go
func (*UserRequest) ProtoMessage()
```

#### func (*UserRequest) Reset

```go
func (m *UserRequest) Reset()
```

#### func (*UserRequest) String

```go
func (m *UserRequest) String() string
```

#### func (*UserRequest) XXX_DiscardUnknown

```go
func (m *UserRequest) XXX_DiscardUnknown()
```

#### func (*UserRequest) XXX_Marshal

```go
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*UserRequest) XXX_Merge

```go
func (m *UserRequest) XXX_Merge(src proto.Message)
```

#### func (*UserRequest) XXX_Size

```go
func (m *UserRequest) XXX_Size() int
```

#### func (*UserRequest) XXX_Unmarshal

```go
func (m *UserRequest) XXX_Unmarshal(b []byte) error
```

#### type UserServiceClient

```go
type UserServiceClient interface {
	GetUser(ctx context.Context, in *UserByEmailRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Bytes, error)
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Bytes, error)
	DeleteUser(ctx context.Context, in *UserByEmailRequest, opts ...grpc.CallOption) (*Bytes, error)
	ListUsers(ctx context.Context, in *ManagementToken, opts ...grpc.CallOption) (UserService_ListUsersClient, error)
}
```

UserServiceClient is the client API for UserService service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to
https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

#### func  NewUserServiceClient

```go
func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient
```

#### type UserServiceServer

```go
type UserServiceServer interface {
	GetUser(context.Context, *UserByEmailRequest) (*User, error)
	UpdateUser(context.Context, *UserRequest) (*Bytes, error)
	CreateUser(context.Context, *UserRequest) (*Bytes, error)
	DeleteUser(context.Context, *UserByEmailRequest) (*Bytes, error)
	ListUsers(*ManagementToken, UserService_ListUsersServer) error
}
```

UserServiceServer is the server API for UserService service.

#### type UserService_ListUsersClient

```go
type UserService_ListUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}
```


#### type UserService_ListUsersServer

```go
type UserService_ListUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}
```


#### type UtilityServiceClient

```go
type UtilityServiceClient interface {
	Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	EchoSpanish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	EchoChinese(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	EchoEnglish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	EchoHindi(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	EchoArabic(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	MarshalJSON(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*Bytes, error)
	MarshalYAML(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*Bytes, error)
	MarshalXML(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*Bytes, error)
	Render(ctx context.Context, in *RenderRequest, opts ...grpc.CallOption) (*Bytes, error)
}
```

UtilityServiceClient is the client API for UtilityService service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to
https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

#### func  NewUtilityServiceClient

```go
func NewUtilityServiceClient(cc *grpc.ClientConn) UtilityServiceClient
```

#### type UtilityServiceServer

```go
type UtilityServiceServer interface {
	Echo(context.Context, *Message) (*Message, error)
	EchoSpanish(context.Context, *Message) (*Message, error)
	EchoChinese(context.Context, *Message) (*Message, error)
	EchoEnglish(context.Context, *Message) (*Message, error)
	EchoHindi(context.Context, *Message) (*Message, error)
	EchoArabic(context.Context, *Message) (*Message, error)
	MarshalJSON(context.Context, *Bytes) (*Bytes, error)
	MarshalYAML(context.Context, *Bytes) (*Bytes, error)
	MarshalXML(context.Context, *Bytes) (*Bytes, error)
	Render(context.Context, *RenderRequest) (*Bytes, error)
}
```

UtilityServiceServer is the server API for UtilityService service.
