# api
--
    import "github.com/autom8ter/api"

Package api is a reverse proxy.

It translates gRPC into RESTful JSON APIs.

## Usage

```go
var (
	DEFAULT_OAUTH_REDIRECT = common.ToString("http://localhost:8080/callback")
	DEFAULT_OAUTH_SCOPES   = []Scope{Scope_OPENID, Scope_PROFILE, Scope_EMAIL}
)
```

```go
var Plan_name = map[int32]string{
	0: "FREE",
	1: "BASIC",
	2: "PREMIUM",
}
```

```go
var Plan_value = map[string]int32{
	"FREE":    0,
	"BASIC":   1,
	"PREMIUM": 2,
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
var URL_name = map[int32]string{
	0:  "USER_INFOURL",
	1:  "TOKENURL",
	2:  "AUTHORIZEURL",
	3:  "USERSURL",
	4:  "CLIENTSURL",
	5:  "GRANTSURL",
	6:  "RULESURL",
	7:  "ROLESURL",
	8:  "LOGSURL",
	9:  "STATSURL",
	10: "CONNECTIONSURL",
	11: "TENANTSURL",
	12: "EMAIL_TEMPLATEURL",
	13: "EMAILURL",
	14: "SEARCH_USERSURL",
	18: "DEVICEURL",
	19: "JWKSURL",
	20: "CLIENT_GRANTSURL",
}
```

```go
var URL_value = map[string]int32{
	"USER_INFOURL":      0,
	"TOKENURL":          1,
	"AUTHORIZEURL":      2,
	"USERSURL":          3,
	"CLIENTSURL":        4,
	"GRANTSURL":         5,
	"RULESURL":          6,
	"ROLESURL":          7,
	"LOGSURL":           8,
	"STATSURL":          9,
	"CONNECTIONSURL":    10,
	"TENANTSURL":        11,
	"EMAIL_TEMPLATEURL": 12,
	"EMAILURL":          13,
	"SEARCH_USERSURL":   14,
	"DEVICEURL":         18,
	"JWKSURL":           19,
	"CLIENT_GRANTSURL":  20,
}
```

#### func  ChatServiceURL

```go
func ChatServiceURL() *common.String
```

#### func  IncomingPhoneNumbersURL

```go
func IncomingPhoneNumbersURL(account *common.String) *common.String
```

#### func  NormalizeScopes

```go
func NormalizeScopes(scopes ...Scope) *common.StringArray
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

#### func  RenderUserFunc

```go
func RenderUserFunc(t *common.String) http.HandlerFunc
```

#### func  SearchUSPhoneNumbersURL

```go
func SearchUSPhoneNumbersURL(account *common.String) *common.String
```

#### type AdminServiceClient

```go
type AdminServiceClient interface {
	StartCache(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*common.Empty, error)
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
	StartCache(context.Context, *common.Empty) (*common.Empty, error)
}
```

AdminServiceServer is the server API for AdminService service.

#### type Auth

```go
type Auth struct {
	Domain               *common.String `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	ClientId             *common.String `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         *common.String `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	Redirect             *common.String `protobuf:"bytes,4,opt,name=redirect,proto3" json:"redirect,omitempty"`
	Scopes               []Scope        `protobuf:"varint,5,rep,packed,name=scopes,proto3,enum=api.Scope" json:"scopes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*Auth) AuthCodeURL

```go
func (a *Auth) AuthCodeURL(state string, u URL) string
```

#### func (*Auth) DefaultIfEmpty

```go
func (a *Auth) DefaultIfEmpty()
```

#### func (*Auth) Descriptor

```go
func (*Auth) Descriptor() ([]byte, []int)
```

#### func (*Auth) GetClientId

```go
func (m *Auth) GetClientId() *common.String
```

#### func (*Auth) GetClientSecret

```go
func (m *Auth) GetClientSecret() *common.String
```

#### func (*Auth) GetDomain

```go
func (m *Auth) GetDomain() *common.String
```

#### func (*Auth) GetRedirect

```go
func (m *Auth) GetRedirect() *common.String
```

#### func (*Auth) GetScopes

```go
func (m *Auth) GetScopes() []Scope
```

#### func (*Auth) ProtoMessage

```go
func (*Auth) ProtoMessage()
```

#### func (*Auth) Reset

```go
func (m *Auth) Reset()
```

#### func (*Auth) String

```go
func (m *Auth) String() string
```

#### func (*Auth) Token

```go
func (a *Auth) Token(ctx context.Context, code string) (*common.Token, error)
```

#### func (*Auth) Validate

```go
func (a *Auth) Validate() error
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

#### type Call

```go
type Call struct {
	From                 *common.String `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *common.String `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	App                  *common.String `protobuf:"bytes,3,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*Call) Descriptor

```go
func (*Call) Descriptor() ([]byte, []int)
```

#### func (*Call) GetApp

```go
func (m *Call) GetApp() *common.String
```

#### func (*Call) GetFrom

```go
func (m *Call) GetFrom() *common.String
```

#### func (*Call) GetTo

```go
func (m *Call) GetTo() *common.String
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
	From                 *common.String      `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *common.StringArray `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	App                  *common.String      `protobuf:"bytes,3,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}
```


#### func (*CallBlast) Descriptor

```go
func (*CallBlast) Descriptor() ([]byte, []int)
```

#### func (*CallBlast) GetApp

```go
func (m *CallBlast) GetApp() *common.String
```

#### func (*CallBlast) GetFrom

```go
func (m *CallBlast) GetFrom() *common.String
```

#### func (*CallBlast) GetTo

```go
func (m *CallBlast) GetTo() *common.StringArray
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

#### type CallResponse

```go
type CallResponse struct {
	Id                   *common.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	To                   *common.String     `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	From                 *common.String     `protobuf:"bytes,6,opt,name=from,proto3" json:"from,omitempty"`
	MediaUrl             *common.String     `protobuf:"bytes,7,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	Body                 *common.String     `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
	Status               *common.String     `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	AnsweredBy           *common.String     `protobuf:"bytes,10,opt,name=answered_by,json=answeredBy,proto3" json:"answered_by,omitempty"`
	ForwardedFrom        *common.String     `protobuf:"bytes,11,opt,name=forwarded_from,json=forwardedFrom,proto3" json:"forwarded_from,omitempty"`
	CallerName           *common.String     `protobuf:"bytes,12,opt,name=caller_name,json=callerName,proto3" json:"caller_name,omitempty"`
	Annotations          *common.StringMap  `protobuf:"bytes,13,opt,name=annotations,proto3" json:"annotations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func (*CallResponse) Descriptor

```go
func (*CallResponse) Descriptor() ([]byte, []int)
```

#### func (*CallResponse) GetAnnotations

```go
func (m *CallResponse) GetAnnotations() *common.StringMap
```

#### func (*CallResponse) GetAnsweredBy

```go
func (m *CallResponse) GetAnsweredBy() *common.String
```

#### func (*CallResponse) GetBody

```go
func (m *CallResponse) GetBody() *common.String
```

#### func (*CallResponse) GetCallerName

```go
func (m *CallResponse) GetCallerName() *common.String
```

#### func (*CallResponse) GetForwardedFrom

```go
func (m *CallResponse) GetForwardedFrom() *common.String
```

#### func (*CallResponse) GetFrom

```go
func (m *CallResponse) GetFrom() *common.String
```

#### func (*CallResponse) GetId

```go
func (m *CallResponse) GetId() *common.Identifier
```

#### func (*CallResponse) GetMediaUrl

```go
func (m *CallResponse) GetMediaUrl() *common.String
```

#### func (*CallResponse) GetStatus

```go
func (m *CallResponse) GetStatus() *common.String
```

#### func (*CallResponse) GetTo

```go
func (m *CallResponse) GetTo() *common.String
```

#### func (*CallResponse) ProtoMessage

```go
func (*CallResponse) ProtoMessage()
```

#### func (*CallResponse) Reset

```go
func (m *CallResponse) Reset()
```

#### func (*CallResponse) String

```go
func (m *CallResponse) String() string
```

#### func (*CallResponse) XXX_DiscardUnknown

```go
func (m *CallResponse) XXX_DiscardUnknown()
```

#### func (*CallResponse) XXX_Marshal

```go
func (m *CallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*CallResponse) XXX_Merge

```go
func (m *CallResponse) XXX_Merge(src proto.Message)
```

#### func (*CallResponse) XXX_Size

```go
func (m *CallResponse) XXX_Size() int
```

#### func (*CallResponse) XXX_Unmarshal

```go
func (m *CallResponse) XXX_Unmarshal(b []byte) error
```

#### type Card

```go
type Card struct {
	Number               *common.String `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	ExpMonth             *common.String `protobuf:"bytes,2,opt,name=exp_month,json=expMonth,proto3" json:"exp_month,omitempty"`
	ExpYear              *common.String `protobuf:"bytes,3,opt,name=exp_year,json=expYear,proto3" json:"exp_year,omitempty"`
	Cvc                  *common.String `protobuf:"bytes,4,opt,name=cvc,proto3" json:"cvc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*Card) Descriptor

```go
func (*Card) Descriptor() ([]byte, []int)
```

#### func (*Card) GetCvc

```go
func (m *Card) GetCvc() *common.String
```

#### func (*Card) GetExpMonth

```go
func (m *Card) GetExpMonth() *common.String
```

#### func (*Card) GetExpYear

```go
func (m *Card) GetExpYear() *common.String
```

#### func (*Card) GetNumber

```go
func (m *Card) GetNumber() *common.String
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
	Payment PaymentServiceClient
	User    UserServiceClient
	Admin   AdminServiceClient
}
```


#### func  NewClientSet

```go
func NewClientSet(conn *grpc.ClientConn) *ClientSet
```

#### type ContactServiceClient

```go
type ContactServiceClient interface {
	SendSMS(ctx context.Context, in *SMS, opts ...grpc.CallOption) (*SMSResponse, error)
	SendSMSBlast(ctx context.Context, in *SMSBlast, opts ...grpc.CallOption) (ContactService_SendSMSBlastClient, error)
	GetSMS(ctx context.Context, in *common.Identifier, opts ...grpc.CallOption) (*SMSResponse, error)
	SendEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*common.String, error)
	SendEmailBlast(ctx context.Context, in *EmailBlastRequest, opts ...grpc.CallOption) (ContactService_SendEmailBlastClient, error)
	SendCall(ctx context.Context, in *Call, opts ...grpc.CallOption) (*CallResponse, error)
	SendCallBlast(ctx context.Context, in *CallBlast, opts ...grpc.CallOption) (ContactService_SendCallBlastClient, error)
	SendFax(ctx context.Context, in *FaxRequest, opts ...grpc.CallOption) (*FaxResponse, error)
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
	SendSMS(context.Context, *SMS) (*SMSResponse, error)
	SendSMSBlast(*SMSBlast, ContactService_SendSMSBlastServer) error
	GetSMS(context.Context, *common.Identifier) (*SMSResponse, error)
	SendEmail(context.Context, *EmailRequest) (*common.String, error)
	SendEmailBlast(*EmailBlastRequest, ContactService_SendEmailBlastServer) error
	SendCall(context.Context, *Call) (*CallResponse, error)
	SendCallBlast(*CallBlast, ContactService_SendCallBlastServer) error
	SendFax(context.Context, *FaxRequest) (*FaxResponse, error)
}
```

ContactServiceServer is the server API for ContactService service.

#### type ContactService_SendCallBlastClient

```go
type ContactService_SendCallBlastClient interface {
	Recv() (*CallResponse, error)
	grpc.ClientStream
}
```


#### type ContactService_SendCallBlastServer

```go
type ContactService_SendCallBlastServer interface {
	Send(*CallResponse) error
	grpc.ServerStream
}
```


#### type ContactService_SendEmailBlastClient

```go
type ContactService_SendEmailBlastClient interface {
	Recv() (*common.String, error)
	grpc.ClientStream
}
```


#### type ContactService_SendEmailBlastServer

```go
type ContactService_SendEmailBlastServer interface {
	Send(*common.String) error
	grpc.ServerStream
}
```


#### type ContactService_SendSMSBlastClient

```go
type ContactService_SendSMSBlastClient interface {
	Recv() (*SMSResponse, error)
	grpc.ClientStream
}
```


#### type ContactService_SendSMSBlastServer

```go
type ContactService_SendSMSBlastServer interface {
	Send(*SMSResponse) error
	grpc.ServerStream
}
```


#### type Email

```go
type Email struct {
	Name                 *common.String `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              *common.String `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Subject              *common.String `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Plain                *common.String `protobuf:"bytes,4,opt,name=plain,proto3" json:"plain,omitempty"`
	Html                 *common.String `protobuf:"bytes,5,opt,name=html,proto3" json:"html,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*Email) Descriptor

```go
func (*Email) Descriptor() ([]byte, []int)
```

#### func (*Email) GetAddress

```go
func (m *Email) GetAddress() *common.String
```

#### func (*Email) GetHtml

```go
func (m *Email) GetHtml() *common.String
```

#### func (*Email) GetName

```go
func (m *Email) GetName() *common.String
```

#### func (*Email) GetPlain

```go
func (m *Email) GetPlain() *common.String
```

#### func (*Email) GetSubject

```go
func (m *Email) GetSubject() *common.String
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
	NameAddress          *common.StringMap `protobuf:"bytes,1,opt,name=name_address,json=nameAddress,proto3" json:"name_address,omitempty"`
	Subject              *common.String    `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Plain                *common.String    `protobuf:"bytes,3,opt,name=plain,proto3" json:"plain,omitempty"`
	Html                 *common.String    `protobuf:"bytes,4,opt,name=html,proto3" json:"html,omitempty"`
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
func (m *EmailBlast) GetHtml() *common.String
```

#### func (*EmailBlast) GetNameAddress

```go
func (m *EmailBlast) GetNameAddress() *common.StringMap
```

#### func (*EmailBlast) GetPlain

```go
func (m *EmailBlast) GetPlain() *common.String
```

#### func (*EmailBlast) GetSubject

```go
func (m *EmailBlast) GetSubject() *common.String
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
	FromName             *common.String `protobuf:"bytes,1,opt,name=from_name,json=fromName,proto3" json:"from_name,omitempty"`
	FromEmail            *common.String `protobuf:"bytes,2,opt,name=from_email,json=fromEmail,proto3" json:"from_email,omitempty"`
	Blast                *EmailBlast    `protobuf:"bytes,3,opt,name=blast,proto3" json:"blast,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
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
func (m *EmailBlastRequest) GetFromEmail() *common.String
```

#### func (*EmailBlastRequest) GetFromName

```go
func (m *EmailBlastRequest) GetFromName() *common.String
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
	FromName             *common.String `protobuf:"bytes,1,opt,name=from_name,json=fromName,proto3" json:"from_name,omitempty"`
	FromEmail            *common.String `protobuf:"bytes,2,opt,name=from_email,json=fromEmail,proto3" json:"from_email,omitempty"`
	Email                *Email         `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
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
func (m *EmailRequest) GetFromEmail() *common.String
```

#### func (*EmailRequest) GetFromName

```go
func (m *EmailRequest) GetFromName() *common.String
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

#### type FaxRequest

```go
type FaxRequest struct {
	To                   *common.String `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	From                 *common.String `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	MediaUrl             *common.String `protobuf:"bytes,3,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	Quality              *common.String `protobuf:"bytes,4,opt,name=quality,proto3" json:"quality,omitempty"`
	Callback             *common.String `protobuf:"bytes,5,opt,name=callback,proto3" json:"callback,omitempty"`
	StoreMedia           bool           `protobuf:"varint,6,opt,name=store_media,json=storeMedia,proto3" json:"store_media,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*FaxRequest) Descriptor

```go
func (*FaxRequest) Descriptor() ([]byte, []int)
```

#### func (*FaxRequest) GetCallback

```go
func (m *FaxRequest) GetCallback() *common.String
```

#### func (*FaxRequest) GetFrom

```go
func (m *FaxRequest) GetFrom() *common.String
```

#### func (*FaxRequest) GetMediaUrl

```go
func (m *FaxRequest) GetMediaUrl() *common.String
```

#### func (*FaxRequest) GetQuality

```go
func (m *FaxRequest) GetQuality() *common.String
```

#### func (*FaxRequest) GetStoreMedia

```go
func (m *FaxRequest) GetStoreMedia() bool
```

#### func (*FaxRequest) GetTo

```go
func (m *FaxRequest) GetTo() *common.String
```

#### func (*FaxRequest) ProtoMessage

```go
func (*FaxRequest) ProtoMessage()
```

#### func (*FaxRequest) Reset

```go
func (m *FaxRequest) Reset()
```

#### func (*FaxRequest) String

```go
func (m *FaxRequest) String() string
```

#### func (*FaxRequest) XXX_DiscardUnknown

```go
func (m *FaxRequest) XXX_DiscardUnknown()
```

#### func (*FaxRequest) XXX_Marshal

```go
func (m *FaxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*FaxRequest) XXX_Merge

```go
func (m *FaxRequest) XXX_Merge(src proto.Message)
```

#### func (*FaxRequest) XXX_Size

```go
func (m *FaxRequest) XXX_Size() int
```

#### func (*FaxRequest) XXX_Unmarshal

```go
func (m *FaxRequest) XXX_Unmarshal(b []byte) error
```

#### type FaxResponse

```go
type FaxResponse struct {
	Id                   *common.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ContentType          *common.String     `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	MediaUrl             *common.String     `protobuf:"bytes,3,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	To                   *common.String     `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	From                 *common.String     `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	Annotations          *common.StringMap  `protobuf:"bytes,10,opt,name=annotations,proto3" json:"annotations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func (*FaxResponse) Descriptor

```go
func (*FaxResponse) Descriptor() ([]byte, []int)
```

#### func (*FaxResponse) GetAnnotations

```go
func (m *FaxResponse) GetAnnotations() *common.StringMap
```

#### func (*FaxResponse) GetContentType

```go
func (m *FaxResponse) GetContentType() *common.String
```

#### func (*FaxResponse) GetFrom

```go
func (m *FaxResponse) GetFrom() *common.String
```

#### func (*FaxResponse) GetId

```go
func (m *FaxResponse) GetId() *common.Identifier
```

#### func (*FaxResponse) GetMediaUrl

```go
func (m *FaxResponse) GetMediaUrl() *common.String
```

#### func (*FaxResponse) GetTo

```go
func (m *FaxResponse) GetTo() *common.String
```

#### func (*FaxResponse) ProtoMessage

```go
func (*FaxResponse) ProtoMessage()
```

#### func (*FaxResponse) Reset

```go
func (m *FaxResponse) Reset()
```

#### func (*FaxResponse) String

```go
func (m *FaxResponse) String() string
```

#### func (*FaxResponse) XXX_DiscardUnknown

```go
func (m *FaxResponse) XXX_DiscardUnknown()
```

#### func (*FaxResponse) XXX_Marshal

```go
func (m *FaxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*FaxResponse) XXX_Merge

```go
func (m *FaxResponse) XXX_Merge(src proto.Message)
```

#### func (*FaxResponse) XXX_Size

```go
func (m *FaxResponse) XXX_Size() int
```

#### func (*FaxResponse) XXX_Unmarshal

```go
func (m *FaxResponse) XXX_Unmarshal(b []byte) error
```

#### type IDBody

```go
type IDBody struct {
	Id                   *common.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body                 *common.Bytes      `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func (*IDBody) Descriptor

```go
func (*IDBody) Descriptor() ([]byte, []int)
```

#### func (*IDBody) GetBody

```go
func (m *IDBody) GetBody() *common.Bytes
```

#### func (*IDBody) GetId

```go
func (m *IDBody) GetId() *common.Identifier
```

#### func (*IDBody) ProtoMessage

```go
func (*IDBody) ProtoMessage()
```

#### func (*IDBody) Reset

```go
func (m *IDBody) Reset()
```

#### func (*IDBody) String

```go
func (m *IDBody) String() string
```

#### func (*IDBody) XXX_DiscardUnknown

```go
func (m *IDBody) XXX_DiscardUnknown()
```

#### func (*IDBody) XXX_Marshal

```go
func (m *IDBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*IDBody) XXX_Merge

```go
func (m *IDBody) XXX_Merge(src proto.Message)
```

#### func (*IDBody) XXX_Size

```go
func (m *IDBody) XXX_Size() int
```

#### func (*IDBody) XXX_Unmarshal

```go
func (m *IDBody) XXX_Unmarshal(b []byte) error
```

#### type IDStrings

```go
type IDStrings struct {
	Id                   *common.Identifier  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Strings              *common.StringArray `protobuf:"bytes,2,opt,name=strings,proto3" json:"strings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}
```


#### func (*IDStrings) Descriptor

```go
func (*IDStrings) Descriptor() ([]byte, []int)
```

#### func (*IDStrings) GetId

```go
func (m *IDStrings) GetId() *common.Identifier
```

#### func (*IDStrings) GetStrings

```go
func (m *IDStrings) GetStrings() *common.StringArray
```

#### func (*IDStrings) ProtoMessage

```go
func (*IDStrings) ProtoMessage()
```

#### func (*IDStrings) Reset

```go
func (m *IDStrings) Reset()
```

#### func (*IDStrings) String

```go
func (m *IDStrings) String() string
```

#### func (*IDStrings) XXX_DiscardUnknown

```go
func (m *IDStrings) XXX_DiscardUnknown()
```

#### func (*IDStrings) XXX_Marshal

```go
func (m *IDStrings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*IDStrings) XXX_Merge

```go
func (m *IDStrings) XXX_Merge(src proto.Message)
```

#### func (*IDStrings) XXX_Size

```go
func (m *IDStrings) XXX_Size() int
```

#### func (*IDStrings) XXX_Unmarshal

```go
func (m *IDStrings) XXX_Unmarshal(b []byte) error
```

#### type Identity

```go
type Identity struct {
	Connection           *common.String     `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	UserId               *common.Identifier `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Provider             *common.String     `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	IsSocial             bool               `protobuf:"varint,4,opt,name=isSocial,proto3" json:"isSocial,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func (*Identity) Descriptor

```go
func (*Identity) Descriptor() ([]byte, []int)
```

#### func (*Identity) GetConnection

```go
func (m *Identity) GetConnection() *common.String
```

#### func (*Identity) GetIsSocial

```go
func (m *Identity) GetIsSocial() bool
```

#### func (*Identity) GetProvider

```go
func (m *Identity) GetProvider() *common.String
```

#### func (*Identity) GetUserId

```go
func (m *Identity) GetUserId() *common.Identifier
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
	Kty                  *common.String      `protobuf:"bytes,1,opt,name=kty,proto3" json:"kty,omitempty"`
	Kid                  *common.Identifier  `protobuf:"bytes,2,opt,name=kid,proto3" json:"kid,omitempty"`
	Use                  *common.String      `protobuf:"bytes,3,opt,name=use,proto3" json:"use,omitempty"`
	N                    *common.String      `protobuf:"bytes,4,opt,name=n,proto3" json:"n,omitempty"`
	E                    *common.String      `protobuf:"bytes,5,opt,name=e,proto3" json:"e,omitempty"`
	X5C                  *common.StringArray `protobuf:"bytes,6,opt,name=x5c,proto3" json:"x5c,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}
```


#### func (*JSONWebKeys) Descriptor

```go
func (*JSONWebKeys) Descriptor() ([]byte, []int)
```

#### func (*JSONWebKeys) GetE

```go
func (m *JSONWebKeys) GetE() *common.String
```

#### func (*JSONWebKeys) GetKid

```go
func (m *JSONWebKeys) GetKid() *common.Identifier
```

#### func (*JSONWebKeys) GetKty

```go
func (m *JSONWebKeys) GetKty() *common.String
```

#### func (*JSONWebKeys) GetN

```go
func (m *JSONWebKeys) GetN() *common.String
```

#### func (*JSONWebKeys) GetUse

```go
func (m *JSONWebKeys) GetUse() *common.String
```

#### func (*JSONWebKeys) GetX5C

```go
func (m *JSONWebKeys) GetX5C() *common.StringArray
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

#### type NumberCapabilities

```go
type NumberCapabilities struct {
	Voice                bool     `protobuf:"varint,1,opt,name=voice,proto3" json:"voice,omitempty"`
	Sms                  bool     `protobuf:"varint,2,opt,name=sms,proto3" json:"sms,omitempty"`
	Mms                  bool     `protobuf:"varint,3,opt,name=mms,proto3" json:"mms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*NumberCapabilities) Descriptor

```go
func (*NumberCapabilities) Descriptor() ([]byte, []int)
```

#### func (*NumberCapabilities) GetMms

```go
func (m *NumberCapabilities) GetMms() bool
```

#### func (*NumberCapabilities) GetSms

```go
func (m *NumberCapabilities) GetSms() bool
```

#### func (*NumberCapabilities) GetVoice

```go
func (m *NumberCapabilities) GetVoice() bool
```

#### func (*NumberCapabilities) ProtoMessage

```go
func (*NumberCapabilities) ProtoMessage()
```

#### func (*NumberCapabilities) Reset

```go
func (m *NumberCapabilities) Reset()
```

#### func (*NumberCapabilities) String

```go
func (m *NumberCapabilities) String() string
```

#### func (*NumberCapabilities) XXX_DiscardUnknown

```go
func (m *NumberCapabilities) XXX_DiscardUnknown()
```

#### func (*NumberCapabilities) XXX_Marshal

```go
func (m *NumberCapabilities) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*NumberCapabilities) XXX_Merge

```go
func (m *NumberCapabilities) XXX_Merge(src proto.Message)
```

#### func (*NumberCapabilities) XXX_Size

```go
func (m *NumberCapabilities) XXX_Size() int
```

#### func (*NumberCapabilities) XXX_Unmarshal

```go
func (m *NumberCapabilities) XXX_Unmarshal(b []byte) error
```

#### type PaymentServiceClient

```go
type PaymentServiceClient interface {
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscriptionResponse, error)
	Unsubscribe(ctx context.Context, in *UnSubscribeRequest, opts ...grpc.CallOption) (*SubscriptionResponse, error)
	PurchasePhoneNumber(ctx context.Context, in *PhoneNumber, opts ...grpc.CallOption) (*PhoneNumberResource, error)
	SearchPhoneNumber(ctx context.Context, in *SearchPhoneNumberRequest, opts ...grpc.CallOption) (PaymentService_SearchPhoneNumberClient, error)
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
	Subscribe(context.Context, *SubscribeRequest) (*SubscriptionResponse, error)
	Unsubscribe(context.Context, *UnSubscribeRequest) (*SubscriptionResponse, error)
	PurchasePhoneNumber(context.Context, *PhoneNumber) (*PhoneNumberResource, error)
	SearchPhoneNumber(*SearchPhoneNumberRequest, PaymentService_SearchPhoneNumberServer) error
}
```

PaymentServiceServer is the server API for PaymentService service.

#### type PaymentService_SearchPhoneNumberClient

```go
type PaymentService_SearchPhoneNumberClient interface {
	Recv() (*PhoneNumber, error)
	grpc.ClientStream
}
```


#### type PaymentService_SearchPhoneNumberServer

```go
type PaymentService_SearchPhoneNumberServer interface {
	Send(*PhoneNumber) error
	grpc.ServerStream
}
```


#### type PhoneNumber

```go
type PhoneNumber struct {
	FriendlyName         *common.String      `protobuf:"bytes,1,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	PhoneNumber          *common.String      `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Region               *common.String      `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Capabilities         *NumberCapabilities `protobuf:"bytes,4,opt,name=capabilities,proto3" json:"capabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}
```


#### func (*PhoneNumber) Descriptor

```go
func (*PhoneNumber) Descriptor() ([]byte, []int)
```

#### func (*PhoneNumber) GetCapabilities

```go
func (m *PhoneNumber) GetCapabilities() *NumberCapabilities
```

#### func (*PhoneNumber) GetFriendlyName

```go
func (m *PhoneNumber) GetFriendlyName() *common.String
```

#### func (*PhoneNumber) GetPhoneNumber

```go
func (m *PhoneNumber) GetPhoneNumber() *common.String
```

#### func (*PhoneNumber) GetRegion

```go
func (m *PhoneNumber) GetRegion() *common.String
```

#### func (*PhoneNumber) ProtoMessage

```go
func (*PhoneNumber) ProtoMessage()
```

#### func (*PhoneNumber) Reset

```go
func (m *PhoneNumber) Reset()
```

#### func (*PhoneNumber) String

```go
func (m *PhoneNumber) String() string
```

#### func (*PhoneNumber) XXX_DiscardUnknown

```go
func (m *PhoneNumber) XXX_DiscardUnknown()
```

#### func (*PhoneNumber) XXX_Marshal

```go
func (m *PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*PhoneNumber) XXX_Merge

```go
func (m *PhoneNumber) XXX_Merge(src proto.Message)
```

#### func (*PhoneNumber) XXX_Size

```go
func (m *PhoneNumber) XXX_Size() int
```

#### func (*PhoneNumber) XXX_Unmarshal

```go
func (m *PhoneNumber) XXX_Unmarshal(b []byte) error
```

#### type PhoneNumberResource

```go
type PhoneNumberResource struct {
	Number               *PhoneNumber       `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Id                   *common.Identifier `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Uri                  *common.String     `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func (*PhoneNumberResource) Descriptor

```go
func (*PhoneNumberResource) Descriptor() ([]byte, []int)
```

#### func (*PhoneNumberResource) GetId

```go
func (m *PhoneNumberResource) GetId() *common.Identifier
```

#### func (*PhoneNumberResource) GetNumber

```go
func (m *PhoneNumberResource) GetNumber() *PhoneNumber
```

#### func (*PhoneNumberResource) GetUri

```go
func (m *PhoneNumberResource) GetUri() *common.String
```

#### func (*PhoneNumberResource) ProtoMessage

```go
func (*PhoneNumberResource) ProtoMessage()
```

#### func (*PhoneNumberResource) Reset

```go
func (m *PhoneNumberResource) Reset()
```

#### func (*PhoneNumberResource) String

```go
func (m *PhoneNumberResource) String() string
```

#### func (*PhoneNumberResource) XXX_DiscardUnknown

```go
func (m *PhoneNumberResource) XXX_DiscardUnknown()
```

#### func (*PhoneNumberResource) XXX_Marshal

```go
func (m *PhoneNumberResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*PhoneNumberResource) XXX_Merge

```go
func (m *PhoneNumberResource) XXX_Merge(src proto.Message)
```

#### func (*PhoneNumberResource) XXX_Size

```go
func (m *PhoneNumberResource) XXX_Size() int
```

#### func (*PhoneNumberResource) XXX_Unmarshal

```go
func (m *PhoneNumberResource) XXX_Unmarshal(b []byte) error
```

#### type Plan

```go
type Plan int32
```


```go
const (
	Plan_FREE    Plan = 0
	Plan_BASIC   Plan = 1
	Plan_PREMIUM Plan = 2
)
```

#### func (Plan) EnumDescriptor

```go
func (Plan) EnumDescriptor() ([]byte, []int)
```

#### func (Plan) Normalize

```go
func (p Plan) Normalize() *common.String
```

#### func (Plan) String

```go
func (x Plan) String() string
```

#### type RenderRequest

```go
type RenderRequest struct {
	Name                 *common.String `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Text                 *common.String `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Data                 *common.Bytes  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*RenderRequest) Descriptor

```go
func (*RenderRequest) Descriptor() ([]byte, []int)
```

#### func (*RenderRequest) GetData

```go
func (m *RenderRequest) GetData() *common.Bytes
```

#### func (*RenderRequest) GetName

```go
func (m *RenderRequest) GetName() *common.String
```

#### func (*RenderRequest) GetText

```go
func (m *RenderRequest) GetText() *common.String
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

#### type Role

```go
type Role struct {
	Id                   *common.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *common.String     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          *common.String     `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func (*Role) Descriptor

```go
func (*Role) Descriptor() ([]byte, []int)
```

#### func (*Role) GetDescription

```go
func (m *Role) GetDescription() *common.String
```

#### func (*Role) GetId

```go
func (m *Role) GetId() *common.Identifier
```

#### func (*Role) GetName

```go
func (m *Role) GetName() *common.String
```

#### func (*Role) ProtoMessage

```go
func (*Role) ProtoMessage()
```

#### func (*Role) Reset

```go
func (m *Role) Reset()
```

#### func (*Role) String

```go
func (m *Role) String() string
```

#### func (*Role) XXX_DiscardUnknown

```go
func (m *Role) XXX_DiscardUnknown()
```

#### func (*Role) XXX_Marshal

```go
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Role) XXX_Merge

```go
func (m *Role) XXX_Merge(src proto.Message)
```

#### func (*Role) XXX_Size

```go
func (m *Role) XXX_Size() int
```

#### func (*Role) XXX_Unmarshal

```go
func (m *Role) XXX_Unmarshal(b []byte) error
```

#### type SMS

```go
type SMS struct {
	Service              *common.String `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	To                   *common.String `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Message              *common.String `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	MediaURL             *common.String `protobuf:"bytes,4,opt,name=mediaURL,proto3" json:"mediaURL,omitempty"`
	Callback             *common.String `protobuf:"bytes,5,opt,name=callback,proto3" json:"callback,omitempty"`
	App                  *common.String `protobuf:"bytes,6,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*SMS) Descriptor

```go
func (*SMS) Descriptor() ([]byte, []int)
```

#### func (*SMS) GetApp

```go
func (m *SMS) GetApp() *common.String
```

#### func (*SMS) GetCallback

```go
func (m *SMS) GetCallback() *common.String
```

#### func (*SMS) GetMediaURL

```go
func (m *SMS) GetMediaURL() *common.String
```

#### func (*SMS) GetMessage

```go
func (m *SMS) GetMessage() *common.String
```

#### func (*SMS) GetService

```go
func (m *SMS) GetService() *common.String
```

#### func (*SMS) GetTo

```go
func (m *SMS) GetTo() *common.String
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
	Service              *common.String      `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	To                   *common.StringArray `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Message              *common.String      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	MediaURL             *common.String      `protobuf:"bytes,4,opt,name=mediaURL,proto3" json:"mediaURL,omitempty"`
	Callback             *common.String      `protobuf:"bytes,5,opt,name=callback,proto3" json:"callback,omitempty"`
	App                  *common.String      `protobuf:"bytes,6,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}
```


#### func (*SMSBlast) Descriptor

```go
func (*SMSBlast) Descriptor() ([]byte, []int)
```

#### func (*SMSBlast) GetApp

```go
func (m *SMSBlast) GetApp() *common.String
```

#### func (*SMSBlast) GetCallback

```go
func (m *SMSBlast) GetCallback() *common.String
```

#### func (*SMSBlast) GetMediaURL

```go
func (m *SMSBlast) GetMediaURL() *common.String
```

#### func (*SMSBlast) GetMessage

```go
func (m *SMSBlast) GetMessage() *common.String
```

#### func (*SMSBlast) GetService

```go
func (m *SMSBlast) GetService() *common.String
```

#### func (*SMSBlast) GetTo

```go
func (m *SMSBlast) GetTo() *common.StringArray
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

#### type SMSResponse

```go
type SMSResponse struct {
	Id                   *common.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	To                   *common.String     `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	From                 *common.String     `protobuf:"bytes,6,opt,name=from,proto3" json:"from,omitempty"`
	MediaUrl             *common.String     `protobuf:"bytes,7,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	Body                 *common.String     `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
	Status               *common.String     `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	Annotations          *common.StringMap  `protobuf:"bytes,10,opt,name=annotations,proto3" json:"annotations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func (*SMSResponse) Descriptor

```go
func (*SMSResponse) Descriptor() ([]byte, []int)
```

#### func (*SMSResponse) GetAnnotations

```go
func (m *SMSResponse) GetAnnotations() *common.StringMap
```

#### func (*SMSResponse) GetBody

```go
func (m *SMSResponse) GetBody() *common.String
```

#### func (*SMSResponse) GetFrom

```go
func (m *SMSResponse) GetFrom() *common.String
```

#### func (*SMSResponse) GetId

```go
func (m *SMSResponse) GetId() *common.Identifier
```

#### func (*SMSResponse) GetMediaUrl

```go
func (m *SMSResponse) GetMediaUrl() *common.String
```

#### func (*SMSResponse) GetStatus

```go
func (m *SMSResponse) GetStatus() *common.String
```

#### func (*SMSResponse) GetTo

```go
func (m *SMSResponse) GetTo() *common.String
```

#### func (*SMSResponse) ProtoMessage

```go
func (*SMSResponse) ProtoMessage()
```

#### func (*SMSResponse) Reset

```go
func (m *SMSResponse) Reset()
```

#### func (*SMSResponse) String

```go
func (m *SMSResponse) String() string
```

#### func (*SMSResponse) XXX_DiscardUnknown

```go
func (m *SMSResponse) XXX_DiscardUnknown()
```

#### func (*SMSResponse) XXX_Marshal

```go
func (m *SMSResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*SMSResponse) XXX_Merge

```go
func (m *SMSResponse) XXX_Merge(src proto.Message)
```

#### func (*SMSResponse) XXX_Size

```go
func (m *SMSResponse) XXX_Size() int
```

#### func (*SMSResponse) XXX_Unmarshal

```go
func (m *SMSResponse) XXX_Unmarshal(b []byte) error
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
func (s Scope) Normalize() *common.String
```

#### func (Scope) String

```go
func (x Scope) String() string
```

#### type SearchPhoneNumberRequest

```go
type SearchPhoneNumberRequest struct {
	State                *common.String      `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Capabilities         *NumberCapabilities `protobuf:"bytes,2,opt,name=capabilities,proto3" json:"capabilities,omitempty"`
	TotalResults         *common.Int64       `protobuf:"bytes,3,opt,name=total_results,json=totalResults,proto3" json:"total_results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}
```


#### func (*SearchPhoneNumberRequest) Descriptor

```go
func (*SearchPhoneNumberRequest) Descriptor() ([]byte, []int)
```

#### func (*SearchPhoneNumberRequest) GetCapabilities

```go
func (m *SearchPhoneNumberRequest) GetCapabilities() *NumberCapabilities
```

#### func (*SearchPhoneNumberRequest) GetState

```go
func (m *SearchPhoneNumberRequest) GetState() *common.String
```

#### func (*SearchPhoneNumberRequest) GetTotalResults

```go
func (m *SearchPhoneNumberRequest) GetTotalResults() *common.Int64
```

#### func (*SearchPhoneNumberRequest) ProtoMessage

```go
func (*SearchPhoneNumberRequest) ProtoMessage()
```

#### func (*SearchPhoneNumberRequest) Reset

```go
func (m *SearchPhoneNumberRequest) Reset()
```

#### func (*SearchPhoneNumberRequest) String

```go
func (m *SearchPhoneNumberRequest) String() string
```

#### func (*SearchPhoneNumberRequest) XXX_DiscardUnknown

```go
func (m *SearchPhoneNumberRequest) XXX_DiscardUnknown()
```

#### func (*SearchPhoneNumberRequest) XXX_Marshal

```go
func (m *SearchPhoneNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*SearchPhoneNumberRequest) XXX_Merge

```go
func (m *SearchPhoneNumberRequest) XXX_Merge(src proto.Message)
```

#### func (*SearchPhoneNumberRequest) XXX_Size

```go
func (m *SearchPhoneNumberRequest) XXX_Size() int
```

#### func (*SearchPhoneNumberRequest) XXX_Unmarshal

```go
func (m *SearchPhoneNumberRequest) XXX_Unmarshal(b []byte) error
```

#### type SubscribeRequest

```go
type SubscribeRequest struct {
	Email                *common.String `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Plan                 Plan           `protobuf:"varint,2,opt,name=plan,proto3,enum=api.Plan" json:"plan,omitempty"`
	Card                 *Card          `protobuf:"bytes,3,opt,name=card,proto3" json:"card,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
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
func (m *SubscribeRequest) GetEmail() *common.String
```

#### func (*SubscribeRequest) GetPlan

```go
func (m *SubscribeRequest) GetPlan() Plan
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

#### type SubscriptionResponse

```go
type SubscriptionResponse struct {
	Id                   *common.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MonthlyCharge        *common.Int64      `protobuf:"bytes,2,opt,name=monthly_charge,json=monthlyCharge,proto3" json:"monthly_charge,omitempty"`
	NextCharge           *common.String     `protobuf:"bytes,3,opt,name=next_charge,json=nextCharge,proto3" json:"next_charge,omitempty"`
	Annotations          *common.StringMap  `protobuf:"bytes,10,opt,name=annotations,proto3" json:"annotations,omitempty"`
	Plan                 Plan               `protobuf:"varint,4,opt,name=plan,proto3,enum=api.Plan" json:"plan,omitempty"`
	User                 *User              `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func (*SubscriptionResponse) Descriptor

```go
func (*SubscriptionResponse) Descriptor() ([]byte, []int)
```

#### func (*SubscriptionResponse) GetAnnotations

```go
func (m *SubscriptionResponse) GetAnnotations() *common.StringMap
```

#### func (*SubscriptionResponse) GetId

```go
func (m *SubscriptionResponse) GetId() *common.Identifier
```

#### func (*SubscriptionResponse) GetMonthlyCharge

```go
func (m *SubscriptionResponse) GetMonthlyCharge() *common.Int64
```

#### func (*SubscriptionResponse) GetNextCharge

```go
func (m *SubscriptionResponse) GetNextCharge() *common.String
```

#### func (*SubscriptionResponse) GetPlan

```go
func (m *SubscriptionResponse) GetPlan() Plan
```

#### func (*SubscriptionResponse) GetUser

```go
func (m *SubscriptionResponse) GetUser() *User
```

#### func (*SubscriptionResponse) ProtoMessage

```go
func (*SubscriptionResponse) ProtoMessage()
```

#### func (*SubscriptionResponse) Reset

```go
func (m *SubscriptionResponse) Reset()
```

#### func (*SubscriptionResponse) String

```go
func (m *SubscriptionResponse) String() string
```

#### func (*SubscriptionResponse) XXX_DiscardUnknown

```go
func (m *SubscriptionResponse) XXX_DiscardUnknown()
```

#### func (*SubscriptionResponse) XXX_Marshal

```go
func (m *SubscriptionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*SubscriptionResponse) XXX_Merge

```go
func (m *SubscriptionResponse) XXX_Merge(src proto.Message)
```

#### func (*SubscriptionResponse) XXX_Size

```go
func (m *SubscriptionResponse) XXX_Size() int
```

#### func (*SubscriptionResponse) XXX_Unmarshal

```go
func (m *SubscriptionResponse) XXX_Unmarshal(b []byte) error
```

#### type TokenQuery

```go
type TokenQuery struct {
	Token                *common.Token  `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Query                *common.String `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*TokenQuery) Descriptor

```go
func (*TokenQuery) Descriptor() ([]byte, []int)
```

#### func (*TokenQuery) GetQuery

```go
func (m *TokenQuery) GetQuery() *common.String
```

#### func (*TokenQuery) GetToken

```go
func (m *TokenQuery) GetToken() *common.Token
```

#### func (*TokenQuery) ProtoMessage

```go
func (*TokenQuery) ProtoMessage()
```

#### func (*TokenQuery) Reset

```go
func (m *TokenQuery) Reset()
```

#### func (*TokenQuery) String

```go
func (m *TokenQuery) String() string
```

#### func (*TokenQuery) XXX_DiscardUnknown

```go
func (m *TokenQuery) XXX_DiscardUnknown()
```

#### func (*TokenQuery) XXX_Marshal

```go
func (m *TokenQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*TokenQuery) XXX_Merge

```go
func (m *TokenQuery) XXX_Merge(src proto.Message)
```

#### func (*TokenQuery) XXX_Size

```go
func (m *TokenQuery) XXX_Size() int
```

#### func (*TokenQuery) XXX_Unmarshal

```go
func (m *TokenQuery) XXX_Unmarshal(b []byte) error
```

#### type URL

```go
type URL int32
```


```go
const (
	URL_USER_INFOURL      URL = 0
	URL_TOKENURL          URL = 1
	URL_AUTHORIZEURL      URL = 2
	URL_USERSURL          URL = 3
	URL_CLIENTSURL        URL = 4
	URL_GRANTSURL         URL = 5
	URL_RULESURL          URL = 6
	URL_ROLESURL          URL = 7
	URL_LOGSURL           URL = 8
	URL_STATSURL          URL = 9
	URL_CONNECTIONSURL    URL = 10
	URL_TENANTSURL        URL = 11
	URL_EMAIL_TEMPLATEURL URL = 12
	URL_EMAILURL          URL = 13
	URL_SEARCH_USERSURL   URL = 14
	URL_DEVICEURL         URL = 18
	URL_JWKSURL           URL = 19
	URL_CLIENT_GRANTSURL  URL = 20
)
```

#### func (URL) EnumDescriptor

```go
func (URL) EnumDescriptor() ([]byte, []int)
```

#### func (URL) Normalize

```go
func (u URL) Normalize(domain *common.String) *common.String
```

#### func (URL) String

```go
func (x URL) String() string
```

#### type UnSubscribeRequest

```go
type UnSubscribeRequest struct {
	Email                *common.String `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Plan                 Plan           `protobuf:"varint,2,opt,name=plan,proto3,enum=api.Plan" json:"plan,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*UnSubscribeRequest) Descriptor

```go
func (*UnSubscribeRequest) Descriptor() ([]byte, []int)
```

#### func (*UnSubscribeRequest) GetEmail

```go
func (m *UnSubscribeRequest) GetEmail() *common.String
```

#### func (*UnSubscribeRequest) GetPlan

```go
func (m *UnSubscribeRequest) GetPlan() Plan
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
	UserId               *common.Identifier  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 *common.String      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GivenName            *common.String      `protobuf:"bytes,3,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	FamilyName           *common.String      `protobuf:"bytes,4,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	Gender               *common.String      `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthdate            *common.String      `protobuf:"bytes,6,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	Email                *common.Identifier  `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber          *common.Identifier  `protobuf:"bytes,8,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Picture              *common.String      `protobuf:"bytes,9,opt,name=picture,proto3" json:"picture,omitempty"`
	UserMetadata         *common.StringMap   `protobuf:"bytes,10,opt,name=user_metadata,json=userMetadata,proto3" json:"user_metadata,omitempty"`
	AppMetadata          *common.StringMap   `protobuf:"bytes,11,opt,name=app_metadata,json=appMetadata,proto3" json:"app_metadata,omitempty"`
	LastIp               *common.String      `protobuf:"bytes,12,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
	Blocked              bool                `protobuf:"varint,13,opt,name=blocked,proto3" json:"blocked,omitempty"`
	Nickname             *common.String      `protobuf:"bytes,14,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Multifactor          *common.StringArray `protobuf:"bytes,15,opt,name=multifactor,proto3" json:"multifactor,omitempty"`
	CreatedAt            *common.String      `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *common.String      `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	PhoneVerified        bool                `protobuf:"varint,19,opt,name=phone_verified,json=phoneVerified,proto3" json:"phone_verified,omitempty"`
	EmailVerified        bool                `protobuf:"varint,20,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	Password             *common.String      `protobuf:"bytes,21,opt,name=password,proto3" json:"password,omitempty"`
	Identities           []*Identity         `protobuf:"bytes,22,rep,name=identities,proto3" json:"identities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}
```


#### func  UserFromSession

```go
func UserFromSession(session *sessions.Session) (*User, error)
```

#### func (*User) Descriptor

```go
func (*User) Descriptor() ([]byte, []int)
```

#### func (*User) GetAppMetadata

```go
func (m *User) GetAppMetadata() *common.StringMap
```

#### func (*User) GetBirthdate

```go
func (m *User) GetBirthdate() *common.String
```

#### func (*User) GetBlocked

```go
func (m *User) GetBlocked() bool
```

#### func (*User) GetCreatedAt

```go
func (m *User) GetCreatedAt() *common.String
```

#### func (*User) GetEmail

```go
func (m *User) GetEmail() *common.Identifier
```

#### func (*User) GetEmailVerified

```go
func (m *User) GetEmailVerified() bool
```

#### func (*User) GetFamilyName

```go
func (m *User) GetFamilyName() *common.String
```

#### func (*User) GetGender

```go
func (m *User) GetGender() *common.String
```

#### func (*User) GetGivenName

```go
func (m *User) GetGivenName() *common.String
```

#### func (*User) GetIdentities

```go
func (m *User) GetIdentities() []*Identity
```

#### func (*User) GetLastIp

```go
func (m *User) GetLastIp() *common.String
```

#### func (*User) GetMultifactor

```go
func (m *User) GetMultifactor() *common.StringArray
```

#### func (*User) GetName

```go
func (m *User) GetName() *common.String
```

#### func (*User) GetNickname

```go
func (m *User) GetNickname() *common.String
```

#### func (*User) GetPassword

```go
func (m *User) GetPassword() *common.String
```

#### func (*User) GetPhoneNumber

```go
func (m *User) GetPhoneNumber() *common.Identifier
```

#### func (*User) GetPhoneVerified

```go
func (m *User) GetPhoneVerified() bool
```

#### func (*User) GetPicture

```go
func (m *User) GetPicture() *common.String
```

#### func (*User) GetUpdatedAt

```go
func (m *User) GetUpdatedAt() *common.String
```

#### func (*User) GetUserId

```go
func (m *User) GetUserId() *common.Identifier
```

#### func (*User) GetUserMetadata

```go
func (m *User) GetUserMetadata() *common.StringMap
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

#### type UserServiceClient

```go
type UserServiceClient interface {
	QueryUsers(ctx context.Context, in *TokenQuery, opts ...grpc.CallOption) (UserService_QueryUsersClient, error)
	CreateUser(ctx context.Context, in *common.Bytes, opts ...grpc.CallOption) (*User, error)
	GetUser(ctx context.Context, in *common.Identifier, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *common.Identifier, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *IDBody, opts ...grpc.CallOption) (*User, error)
	UserRoles(ctx context.Context, in *common.Identifier, opts ...grpc.CallOption) (UserService_UserRolesClient, error)
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
	QueryUsers(*TokenQuery, UserService_QueryUsersServer) error
	CreateUser(context.Context, *common.Bytes) (*User, error)
	GetUser(context.Context, *common.Identifier) (*User, error)
	DeleteUser(context.Context, *common.Identifier) (*User, error)
	UpdateUser(context.Context, *IDBody) (*User, error)
	UserRoles(*common.Identifier, UserService_UserRolesServer) error
}
```

UserServiceServer is the server API for UserService service.

#### type UserService_QueryUsersClient

```go
type UserService_QueryUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}
```


#### type UserService_QueryUsersServer

```go
type UserService_QueryUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}
```


#### type UserService_UserRolesClient

```go
type UserService_UserRolesClient interface {
	Recv() (*Role, error)
	grpc.ClientStream
}
```


#### type UserService_UserRolesServer

```go
type UserService_UserRolesServer interface {
	Send(*Role) error
	grpc.ServerStream
}
```


#### type UtilityServiceClient

```go
type UtilityServiceClient interface {
	Echo(ctx context.Context, in *common.String, opts ...grpc.CallOption) (*common.String, error)
	EchoSpanish(ctx context.Context, in *common.String, opts ...grpc.CallOption) (*common.String, error)
	EchoChinese(ctx context.Context, in *common.String, opts ...grpc.CallOption) (*common.String, error)
	EchoEnglish(ctx context.Context, in *common.String, opts ...grpc.CallOption) (*common.String, error)
	EchoHindi(ctx context.Context, in *common.String, opts ...grpc.CallOption) (*common.String, error)
	EchoArabic(ctx context.Context, in *common.String, opts ...grpc.CallOption) (*common.String, error)
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
	Echo(context.Context, *common.String) (*common.String, error)
	EchoSpanish(context.Context, *common.String) (*common.String, error)
	EchoChinese(context.Context, *common.String) (*common.String, error)
	EchoEnglish(context.Context, *common.String) (*common.String, error)
	EchoHindi(context.Context, *common.String) (*common.String, error)
	EchoArabic(context.Context, *common.String) (*common.String, error)
}
```

UtilityServiceServer is the server API for UtilityService service.
