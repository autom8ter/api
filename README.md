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
var DEFAULT_REDIRECT = "http://localhost:8080/callback"
```

```go
var DEFAULT_SCOPES = []string{"openid", "profile", "email"}
```

```go
var HTTPClient *http.Client
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

#### func  RegisterAdminServiceHandler

```go
func RegisterAdminServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
```
RegisterAdminServiceHandler registers the http handlers for service AdminService
to "mux". The handlers forward requests to the grpc endpoint over "conn".

#### func  RegisterAdminServiceHandlerClient

```go
func RegisterAdminServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client AdminServiceClient) error
```
RegisterAdminServiceHandlerClient registers the http handlers for service
AdminService to "mux". The handlers forward requests to the grpc endpoint over
the given implementation of "AdminServiceClient". Note: the gRPC framework
executes interceptors within the gRPC handler. If the passed in
"AdminServiceClient" doesn't go through the normal gRPC flow (creating a gRPC
client etc.) then it will be up to the passed in "AdminServiceClient" to call
the correct interceptors.

#### func  RegisterAdminServiceHandlerFromEndpoint

```go
func RegisterAdminServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
```
RegisterAdminServiceHandlerFromEndpoint is same as RegisterAdminServiceHandler
but automatically dials to "endpoint" and closes the connection when "ctx" gets
done.

#### func  RegisterAdminServiceServer

```go
func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer)
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

#### func  RenderFileWithData

```go
func RenderFileWithData(name string, data interface{}) http.HandlerFunc
```

#### func  RenderFileWithUserInfo

```go
func RenderFileWithUserInfo(filename string) http.HandlerFunc
```

#### type AdminServiceClient

```go
type AdminServiceClient interface {
	GetDashboard(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Dashboard, error)
	EmailUser(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Message, error)
}
```

AdminServiceClient is the client API for AdminService service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to
https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

#### func  NewAdminServiceClient

```go
func NewAdminServiceClient(cc *grpc.ClientConn) AdminServiceClient
```

#### type AdminServiceServer

```go
type AdminServiceServer interface {
	GetDashboard(context.Context, *Secret) (*Dashboard, error)
	EmailUser(context.Context, *Email) (*Message, error)
}
```

AdminServiceServer is the server API for AdminService service.

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
	Scopes               []string `protobuf:"bytes,6,rep,name=scopes,proto3" json:"scopes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Auth) AuthURL

```go
func (c *Auth) AuthURL() string
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
func (c *Auth) GetJWKS(token *jwt.Token) (*Jwks, error)
```

#### func (*Auth) GetRedirect

```go
func (m *Auth) GetRedirect() string
```

#### func (*Auth) GetScopes

```go
func (m *Auth) GetScopes() []string
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
func (a *Auth) RequireLogin(c *RouterPaths, next http.HandlerFunc) http.HandlerFunc
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


#### func (*Bytes) Descriptor

```go
func (*Bytes) Descriptor() ([]byte, []int)
```

#### func (*Bytes) GetBits

```go
func (m *Bytes) GetBits() []byte
```

#### func (*Bytes) ProtoMessage

```go
func (*Bytes) ProtoMessage()
```

#### func (*Bytes) Reset

```go
func (m *Bytes) Reset()
```

#### func (*Bytes) String

```go
func (m *Bytes) String() string
```

#### func (*Bytes) Write

```go
func (m *Bytes) Write(p []byte) (n int, err error)
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

#### type ChargesWidget

```go
type ChargesWidget struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Total                float64  `protobuf:"fixed64,2,opt,name=total,proto3" json:"total,omitempty"`
	DollarsPerCharge     float64  `protobuf:"fixed64,3,opt,name=dollars_per_charge,json=dollarsPerCharge,proto3" json:"dollars_per_charge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*ChargesWidget) Descriptor

```go
func (*ChargesWidget) Descriptor() ([]byte, []int)
```

#### func (*ChargesWidget) GetCount

```go
func (m *ChargesWidget) GetCount() int64
```

#### func (*ChargesWidget) GetDollarsPerCharge

```go
func (m *ChargesWidget) GetDollarsPerCharge() float64
```

#### func (*ChargesWidget) GetTotal

```go
func (m *ChargesWidget) GetTotal() float64
```

#### func (*ChargesWidget) ProtoMessage

```go
func (*ChargesWidget) ProtoMessage()
```

#### func (*ChargesWidget) Reset

```go
func (m *ChargesWidget) Reset()
```

#### func (*ChargesWidget) String

```go
func (m *ChargesWidget) String() string
```

#### func (*ChargesWidget) XXX_DiscardUnknown

```go
func (m *ChargesWidget) XXX_DiscardUnknown()
```

#### func (*ChargesWidget) XXX_Marshal

```go
func (m *ChargesWidget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*ChargesWidget) XXX_Merge

```go
func (m *ChargesWidget) XXX_Merge(src proto.Message)
```

#### func (*ChargesWidget) XXX_Size

```go
func (m *ChargesWidget) XXX_Size() int
```

#### func (*ChargesWidget) XXX_Unmarshal

```go
func (m *ChargesWidget) XXX_Unmarshal(b []byte) error
```

#### type ClientSet

```go
type ClientSet struct {
	Utility UtilityServiceClient
	Contact ContactServiceClient
	User    UserServiceClient
	Admin   AdminServiceClient
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
	SendSMS(ctx context.Context, in *SMS, opts ...grpc.CallOption) (*Identifier, error)
	GetSMS(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*SMSStatus, error)
	SendEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*Message, error)
	SendCall(ctx context.Context, in *Call, opts ...grpc.CallOption) (*Identifier, error)
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
	SendSMS(context.Context, *SMS) (*Identifier, error)
	GetSMS(context.Context, *Identifier) (*SMSStatus, error)
	SendEmail(context.Context, *EmailRequest) (*Message, error)
	SendCall(context.Context, *Call) (*Identifier, error)
}
```

ContactServiceServer is the server API for ContactService service.

#### type CustomClaims

```go
type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}
```


#### func (*CustomClaims) CheckScope

```go
func (c *CustomClaims) CheckScope(scope string, tokenString string) bool
```

#### type CustomersWidget

```go
type CustomersWidget struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*CustomersWidget) Descriptor

```go
func (*CustomersWidget) Descriptor() ([]byte, []int)
```

#### func (*CustomersWidget) GetCount

```go
func (m *CustomersWidget) GetCount() int64
```

#### func (*CustomersWidget) ProtoMessage

```go
func (*CustomersWidget) ProtoMessage()
```

#### func (*CustomersWidget) Reset

```go
func (m *CustomersWidget) Reset()
```

#### func (*CustomersWidget) String

```go
func (m *CustomersWidget) String() string
```

#### func (*CustomersWidget) XXX_DiscardUnknown

```go
func (m *CustomersWidget) XXX_DiscardUnknown()
```

#### func (*CustomersWidget) XXX_Marshal

```go
func (m *CustomersWidget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*CustomersWidget) XXX_Merge

```go
func (m *CustomersWidget) XXX_Merge(src proto.Message)
```

#### func (*CustomersWidget) XXX_Size

```go
func (m *CustomersWidget) XXX_Size() int
```

#### func (*CustomersWidget) XXX_Unmarshal

```go
func (m *CustomersWidget) XXX_Unmarshal(b []byte) error
```

#### type Dashboard

```go
type Dashboard struct {
	Users                *UsersWidget         `protobuf:"bytes,1,opt,name=users,proto3" json:"users,omitempty"`
	Customers            *CustomersWidget     `protobuf:"bytes,2,opt,name=customers,proto3" json:"customers,omitempty"`
	Plans                *PlansWidget         `protobuf:"bytes,3,opt,name=plans,proto3" json:"plans,omitempty"`
	Subscriptions        *SubscriptionsWidget `protobuf:"bytes,4,opt,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	Charges              *ChargesWidget       `protobuf:"bytes,5,opt,name=charges,proto3" json:"charges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}
```


#### func (*Dashboard) Descriptor

```go
func (*Dashboard) Descriptor() ([]byte, []int)
```

#### func (*Dashboard) GetCharges

```go
func (m *Dashboard) GetCharges() *ChargesWidget
```

#### func (*Dashboard) GetCustomers

```go
func (m *Dashboard) GetCustomers() *CustomersWidget
```

#### func (*Dashboard) GetPlans

```go
func (m *Dashboard) GetPlans() *PlansWidget
```

#### func (*Dashboard) GetSubscriptions

```go
func (m *Dashboard) GetSubscriptions() *SubscriptionsWidget
```

#### func (*Dashboard) GetUsers

```go
func (m *Dashboard) GetUsers() *UsersWidget
```

#### func (*Dashboard) ProtoMessage

```go
func (*Dashboard) ProtoMessage()
```

#### func (*Dashboard) Reset

```go
func (m *Dashboard) Reset()
```

#### func (*Dashboard) String

```go
func (m *Dashboard) String() string
```

#### func (*Dashboard) XXX_DiscardUnknown

```go
func (m *Dashboard) XXX_DiscardUnknown()
```

#### func (*Dashboard) XXX_Marshal

```go
func (m *Dashboard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Dashboard) XXX_Merge

```go
func (m *Dashboard) XXX_Merge(src proto.Message)
```

#### func (*Dashboard) XXX_Size

```go
func (m *Dashboard) XXX_Size() int
```

#### func (*Dashboard) XXX_Unmarshal

```go
func (m *Dashboard) XXX_Unmarshal(b []byte) error
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
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*Identifier, error)
	Unsubscribe(ctx context.Context, in *UnSubscribeRequest, opts ...grpc.CallOption) (*Identifier, error)
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
	Subscribe(context.Context, *SubscribeRequest) (*Identifier, error)
	Unsubscribe(context.Context, *UnSubscribeRequest) (*Identifier, error)
}
```

PaymentServiceServer is the server API for PaymentService service.

#### type PlansWidget

```go
type PlansWidget struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*PlansWidget) Descriptor

```go
func (*PlansWidget) Descriptor() ([]byte, []int)
```

#### func (*PlansWidget) GetCount

```go
func (m *PlansWidget) GetCount() int64
```

#### func (*PlansWidget) ProtoMessage

```go
func (*PlansWidget) ProtoMessage()
```

#### func (*PlansWidget) Reset

```go
func (m *PlansWidget) Reset()
```

#### func (*PlansWidget) String

```go
func (m *PlansWidget) String() string
```

#### func (*PlansWidget) XXX_DiscardUnknown

```go
func (m *PlansWidget) XXX_DiscardUnknown()
```

#### func (*PlansWidget) XXX_Marshal

```go
func (m *PlansWidget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*PlansWidget) XXX_Merge

```go
func (m *PlansWidget) XXX_Merge(src proto.Message)
```

#### func (*PlansWidget) XXX_Size

```go
func (m *PlansWidget) XXX_Size() int
```

#### func (*PlansWidget) XXX_Unmarshal

```go
func (m *PlansWidget) XXX_Unmarshal(b []byte) error
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

#### type SMSStatus

```go
type SMSStatus struct {
	Id                   *Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sms                  *SMS        `protobuf:"bytes,2,opt,name=sms,proto3" json:"sms,omitempty"`
	Status               string      `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Uri                  string      `protobuf:"bytes,4,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}
```


#### func (*SMSStatus) Descriptor

```go
func (*SMSStatus) Descriptor() ([]byte, []int)
```

#### func (*SMSStatus) GetId

```go
func (m *SMSStatus) GetId() *Identifier
```

#### func (*SMSStatus) GetSms

```go
func (m *SMSStatus) GetSms() *SMS
```

#### func (*SMSStatus) GetStatus

```go
func (m *SMSStatus) GetStatus() string
```

#### func (*SMSStatus) GetUri

```go
func (m *SMSStatus) GetUri() string
```

#### func (*SMSStatus) ProtoMessage

```go
func (*SMSStatus) ProtoMessage()
```

#### func (*SMSStatus) Reset

```go
func (m *SMSStatus) Reset()
```

#### func (*SMSStatus) String

```go
func (m *SMSStatus) String() string
```

#### func (*SMSStatus) XXX_DiscardUnknown

```go
func (m *SMSStatus) XXX_DiscardUnknown()
```

#### func (*SMSStatus) XXX_Marshal

```go
func (m *SMSStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*SMSStatus) XXX_Merge

```go
func (m *SMSStatus) XXX_Merge(src proto.Message)
```

#### func (*SMSStatus) XXX_Size

```go
func (m *SMSStatus) XXX_Size() int
```

#### func (*SMSStatus) XXX_Unmarshal

```go
func (m *SMSStatus) XXX_Unmarshal(b []byte) error
```

#### type Scope

```go
type Scope int
```


```go
const (
	OPENID Scope = iota
	PROFILE
	EMAIL
)
```

#### func (Scope) String

```go
func (s Scope) String() string
```

#### type Secret

```go
type Secret struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Secret) Descriptor

```go
func (*Secret) Descriptor() ([]byte, []int)
```

#### func (*Secret) GetText

```go
func (m *Secret) GetText() string
```

#### func (*Secret) ProtoMessage

```go
func (*Secret) ProtoMessage()
```

#### func (*Secret) Reset

```go
func (m *Secret) Reset()
```

#### func (*Secret) String

```go
func (m *Secret) String() string
```

#### func (*Secret) XXX_DiscardUnknown

```go
func (m *Secret) XXX_DiscardUnknown()
```

#### func (*Secret) XXX_Marshal

```go
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Secret) XXX_Merge

```go
func (m *Secret) XXX_Merge(src proto.Message)
```

#### func (*Secret) XXX_Size

```go
func (m *Secret) XXX_Size() int
```

#### func (*Secret) XXX_Unmarshal

```go
func (m *Secret) XXX_Unmarshal(b []byte) error
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

#### type SubscriptionsWidget

```go
type SubscriptionsWidget struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*SubscriptionsWidget) Descriptor

```go
func (*SubscriptionsWidget) Descriptor() ([]byte, []int)
```

#### func (*SubscriptionsWidget) GetCount

```go
func (m *SubscriptionsWidget) GetCount() int64
```

#### func (*SubscriptionsWidget) ProtoMessage

```go
func (*SubscriptionsWidget) ProtoMessage()
```

#### func (*SubscriptionsWidget) Reset

```go
func (m *SubscriptionsWidget) Reset()
```

#### func (*SubscriptionsWidget) String

```go
func (m *SubscriptionsWidget) String() string
```

#### func (*SubscriptionsWidget) XXX_DiscardUnknown

```go
func (m *SubscriptionsWidget) XXX_DiscardUnknown()
```

#### func (*SubscriptionsWidget) XXX_Marshal

```go
func (m *SubscriptionsWidget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*SubscriptionsWidget) XXX_Merge

```go
func (m *SubscriptionsWidget) XXX_Merge(src proto.Message)
```

#### func (*SubscriptionsWidget) XXX_Size

```go
func (m *SubscriptionsWidget) XXX_Size() int
```

#### func (*SubscriptionsWidget) XXX_Unmarshal

```go
func (m *SubscriptionsWidget) XXX_Unmarshal(b []byte) error
```

#### type Template

```go
type Template struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Template) Descriptor

```go
func (*Template) Descriptor() ([]byte, []int)
```

#### func (*Template) GetData

```go
func (m *Template) GetData() []byte
```

#### func (*Template) GetText

```go
func (m *Template) GetText() string
```

#### func (*Template) ProtoMessage

```go
func (*Template) ProtoMessage()
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

#### type UserInfo

```go
type UserInfo struct {
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

#### func (*UserInfo) GetBlocked

```go
func (m *UserInfo) GetBlocked() bool
```

#### func (*UserInfo) GetCreatedAt

```go
func (m *UserInfo) GetCreatedAt() string
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

#### func (*UserInfo) GetIdentities

```go
func (m *UserInfo) GetIdentities() []*Identity
```

#### func (*UserInfo) GetLastIp

```go
func (m *UserInfo) GetLastIp() string
```

#### func (*UserInfo) GetMultifactor

```go
func (m *UserInfo) GetMultifactor() []string
```

#### func (*UserInfo) GetName

```go
func (m *UserInfo) GetName() string
```

#### func (*UserInfo) GetNickname

```go
func (m *UserInfo) GetNickname() string
```

#### func (*UserInfo) GetPhoneNumber

```go
func (m *UserInfo) GetPhoneNumber() string
```

#### func (*UserInfo) GetPhoneVerified

```go
func (m *UserInfo) GetPhoneVerified() bool
```

#### func (*UserInfo) GetPicture

```go
func (m *UserInfo) GetPicture() string
```

#### func (*UserInfo) GetUpdatedAt

```go
func (m *UserInfo) GetUpdatedAt() string
```

#### func (*UserInfo) GetUserId

```go
func (m *UserInfo) GetUserId() string
```

#### func (*UserInfo) GetUserMetadata

```go
func (m *UserInfo) GetUserMetadata() *UserMetadata
```

#### func (*UserInfo) ProtoMessage

```go
func (*UserInfo) ProtoMessage()
```

#### func (*UserInfo) Reset

```go
func (m *UserInfo) Reset()
```

#### func (*UserInfo) String

```go
func (m *UserInfo) String() string
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

#### type UserServiceClient

```go
type UserServiceClient interface {
	GetUser(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*UserInfo, error)
	UpdateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Identifier, error)
	CreateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Identifier, error)
	DeleteUser(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Identifier, error)
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
	GetUser(context.Context, *Identifier) (*UserInfo, error)
	UpdateUser(context.Context, *UserInfo) (*Identifier, error)
	CreateUser(context.Context, *UserInfo) (*Identifier, error)
	DeleteUser(context.Context, *Identifier) (*Identifier, error)
}
```

UserServiceServer is the server API for UserService service.

#### type UsersWidget

```go
type UsersWidget struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*UsersWidget) Descriptor

```go
func (*UsersWidget) Descriptor() ([]byte, []int)
```

#### func (*UsersWidget) GetCount

```go
func (m *UsersWidget) GetCount() int64
```

#### func (*UsersWidget) ProtoMessage

```go
func (*UsersWidget) ProtoMessage()
```

#### func (*UsersWidget) Reset

```go
func (m *UsersWidget) Reset()
```

#### func (*UsersWidget) String

```go
func (m *UsersWidget) String() string
```

#### func (*UsersWidget) XXX_DiscardUnknown

```go
func (m *UsersWidget) XXX_DiscardUnknown()
```

#### func (*UsersWidget) XXX_Marshal

```go
func (m *UsersWidget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*UsersWidget) XXX_Merge

```go
func (m *UsersWidget) XXX_Merge(src proto.Message)
```

#### func (*UsersWidget) XXX_Size

```go
func (m *UsersWidget) XXX_Size() int
```

#### func (*UsersWidget) XXX_Unmarshal

```go
func (m *UsersWidget) XXX_Unmarshal(b []byte) error
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
	Render(ctx context.Context, in *Template, opts ...grpc.CallOption) (*Bytes, error)
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
	Render(context.Context, *Template) (*Bytes, error)
}
```

UtilityServiceServer is the server API for UtilityService service.
