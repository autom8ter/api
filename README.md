# api
--
    import "github.com/autom8ter/api"

Package api is a reverse proxy.

It translates gRPC into RESTful JSON APIs.

## Usage

```go
var EnvContext context.Context
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

#### func  RegisterAuthenticationServiceHandler

```go
func RegisterAuthenticationServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
```
RegisterAuthenticationServiceHandler registers the http handlers for service
AuthenticationService to "mux". The handlers forward requests to the grpc
endpoint over "conn".

#### func  RegisterAuthenticationServiceHandlerClient

```go
func RegisterAuthenticationServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client AuthenticationServiceClient) error
```
RegisterAuthenticationServiceHandlerClient registers the http handlers for
service AuthenticationService to "mux". The handlers forward requests to the
grpc endpoint over the given implementation of "AuthenticationServiceClient".
Note: the gRPC framework executes interceptors within the gRPC handler. If the
passed in "AuthenticationServiceClient" doesn't go through the normal gRPC flow
(creating a gRPC client etc.) then it will be up to the passed in
"AuthenticationServiceClient" to call the correct interceptors.

#### func  RegisterAuthenticationServiceHandlerFromEndpoint

```go
func RegisterAuthenticationServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
```
RegisterAuthenticationServiceHandlerFromEndpoint is same as
RegisterAuthenticationServiceHandler but automatically dials to "endpoint" and
closes the connection when "ctx" gets done.

#### func  RegisterAuthenticationServiceServer

```go
func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer)
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

#### func  RegisterManagementServiceHandler

```go
func RegisterManagementServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
```
RegisterManagementServiceHandler registers the http handlers for service
ManagementService to "mux". The handlers forward requests to the grpc endpoint
over "conn".

#### func  RegisterManagementServiceHandlerClient

```go
func RegisterManagementServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ManagementServiceClient) error
```
RegisterManagementServiceHandlerClient registers the http handlers for service
ManagementService to "mux". The handlers forward requests to the grpc endpoint
over the given implementation of "ManagementServiceClient". Note: the gRPC
framework executes interceptors within the gRPC handler. If the passed in
"ManagementServiceClient" doesn't go through the normal gRPC flow (creating a
gRPC client etc.) then it will be up to the passed in "ManagementServiceClient"
to call the correct interceptors.

#### func  RegisterManagementServiceHandlerFromEndpoint

```go
func RegisterManagementServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
```
RegisterManagementServiceHandlerFromEndpoint is same as
RegisterManagementServiceHandler but automatically dials to "endpoint" and
closes the connection when "ctx" gets done.

#### func  RegisterManagementServiceServer

```go
func RegisterManagementServiceServer(s *grpc.Server, srv ManagementServiceServer)
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

#### type AddUserRolesRequest

```go
type AddUserRolesRequest struct {
	Email                *common.Identifier `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Roles                []*Role            `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func (*AddUserRolesRequest) Descriptor

```go
func (*AddUserRolesRequest) Descriptor() ([]byte, []int)
```

#### func (*AddUserRolesRequest) GetEmail

```go
func (m *AddUserRolesRequest) GetEmail() *common.Identifier
```

#### func (*AddUserRolesRequest) GetRoles

```go
func (m *AddUserRolesRequest) GetRoles() []*Role
```

#### func (*AddUserRolesRequest) ProtoMessage

```go
func (*AddUserRolesRequest) ProtoMessage()
```

#### func (*AddUserRolesRequest) Reset

```go
func (m *AddUserRolesRequest) Reset()
```

#### func (*AddUserRolesRequest) String

```go
func (m *AddUserRolesRequest) String() string
```

#### func (*AddUserRolesRequest) XXX_DiscardUnknown

```go
func (m *AddUserRolesRequest) XXX_DiscardUnknown()
```

#### func (*AddUserRolesRequest) XXX_Marshal

```go
func (m *AddUserRolesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*AddUserRolesRequest) XXX_Merge

```go
func (m *AddUserRolesRequest) XXX_Merge(src proto.Message)
```

#### func (*AddUserRolesRequest) XXX_Size

```go
func (m *AddUserRolesRequest) XXX_Size() int
```

#### func (*AddUserRolesRequest) XXX_Unmarshal

```go
func (m *AddUserRolesRequest) XXX_Unmarshal(b []byte) error
```

#### type AuthenticationServiceClient

```go
type AuthenticationServiceClient interface {
	GetUser(ctx context.Context, in *common.AuthToken, opts ...grpc.CallOption) (*User, error)
}
```

AuthenticationServiceClient is the client API for AuthenticationService service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to
https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

#### func  NewAuthenticationServiceClient

```go
func NewAuthenticationServiceClient(cc *grpc.ClientConn) AuthenticationServiceClient
```

#### type AuthenticationServiceServer

```go
type AuthenticationServiceServer interface {
	GetUser(context.Context, *common.AuthToken) (*User, error)
}
```

AuthenticationServiceServer is the server API for AuthenticationService service.

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

#### func (*Call) JSONString

```go
func (p *Call) JSONString() *common.String
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

#### func (*Call) UnmarshalProtoFrom

```go
func (p *Call) UnmarshalProtoFrom(bits []byte) error
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

#### func (*CallResponse) GetStatus

```go
func (m *CallResponse) GetStatus() *common.String
```

#### func (*CallResponse) GetTo

```go
func (m *CallResponse) GetTo() *common.String
```

#### func (*CallResponse) JSONString

```go
func (p *CallResponse) JSONString() *common.String
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

#### func (*CallResponse) UnmarshalJSONFrom

```go
func (p *CallResponse) UnmarshalJSONFrom(bits []byte) error
```

#### func (*CallResponse) UnmarshalProtoFrom

```go
func (p *CallResponse) UnmarshalProtoFrom(bits []byte) error
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

#### func (*Card) JSONString

```go
func (p *Card) JSONString() *common.String
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

#### func (*Card) UnmarshalJSONFrom

```go
func (p *Card) UnmarshalJSONFrom(bits []byte) error
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

#### type ClientCredentials

```go
type ClientCredentials struct {
	ClientId             *common.String      `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         *common.String      `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	TokenUrl             *common.String      `protobuf:"bytes,3,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	Scopes               *common.StringArray `protobuf:"bytes,4,opt,name=scopes,proto3" json:"scopes,omitempty"`
	EndpointParams       *common.StringMap   `protobuf:"bytes,5,opt,name=endpoint_params,json=endpointParams,proto3" json:"endpoint_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}
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

#### func (*ClientCredentials) DeepEqual

```go
func (s *ClientCredentials) DeepEqual(y interface{}) bool
```

#### func (*ClientCredentials) Descriptor

```go
func (*ClientCredentials) Descriptor() ([]byte, []int)
```

#### func (*ClientCredentials) GetClientId

```go
func (m *ClientCredentials) GetClientId() *common.String
```

#### func (*ClientCredentials) GetClientSecret

```go
func (m *ClientCredentials) GetClientSecret() *common.String
```

#### func (*ClientCredentials) GetEndpointParams

```go
func (m *ClientCredentials) GetEndpointParams() *common.StringMap
```

#### func (*ClientCredentials) GetRequestMetadata

```go
func (c *ClientCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
```

#### func (*ClientCredentials) GetScopes

```go
func (m *ClientCredentials) GetScopes() *common.StringArray
```

#### func (*ClientCredentials) GetTokenUrl

```go
func (m *ClientCredentials) GetTokenUrl() *common.String
```

#### func (*ClientCredentials) JSONString

```go
func (s *ClientCredentials) JSONString() *common.String
```

#### func (*ClientCredentials) NewAPIClientSet

```go
func (c *ClientCredentials) NewAPIClientSet(addr string) (*ClientSet, error)
```

#### func (*ClientCredentials) PerRPCCredentials

```go
func (c *ClientCredentials) PerRPCCredentials() (credentials.PerRPCCredentials, error)
```

#### func (*ClientCredentials) ProtoMessage

```go
func (*ClientCredentials) ProtoMessage()
```

#### func (*ClientCredentials) RequireTransportSecurity

```go
func (c *ClientCredentials) RequireTransportSecurity() bool
```

#### func (*ClientCredentials) Reset

```go
func (m *ClientCredentials) Reset()
```

#### func (*ClientCredentials) String

```go
func (m *ClientCredentials) String() string
```

#### func (*ClientCredentials) ToContext

```go
func (s *ClientCredentials) ToContext(ctx context.Context, key string) context.Context
```

#### func (*ClientCredentials) Token

```go
func (c *ClientCredentials) Token() (*oauth2.Token, error)
```

#### func (*ClientCredentials) UnmarshalJSONFrom

```go
func (s *ClientCredentials) UnmarshalJSONFrom(b []byte) error
```

#### func (*ClientCredentials) UnmarshalProtoFrom

```go
func (s *ClientCredentials) UnmarshalProtoFrom(b []byte) error
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

#### type ClientSet

```go
type ClientSet struct {
	Utility    UtilityServiceClient
	Contact    ContactServiceClient
	Payment    PaymentServiceClient
	Management ManagementServiceClient
	Auth       AuthenticationServiceClient
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
	SendFax(ctx context.Context, in *Fax, opts ...grpc.CallOption) (*FaxResponse, error)
	SendFaxBlast(ctx context.Context, in *FaxBlast, opts ...grpc.CallOption) (ContactService_SendFaxBlastClient, error)
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
	SendFax(context.Context, *Fax) (*FaxResponse, error)
	SendFaxBlast(*FaxBlast, ContactService_SendFaxBlastServer) error
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


#### type ContactService_SendFaxBlastClient

```go
type ContactService_SendFaxBlastClient interface {
	Recv() (*FaxResponse, error)
	grpc.ClientStream
}
```


#### type ContactService_SendFaxBlastServer

```go
type ContactService_SendFaxBlastServer interface {
	Send(*FaxResponse) error
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

#### func (*Email) JSONString

```go
func (p *Email) JSONString() *common.String
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

#### func (*Email) UnmarshalProtoFrom

```go
func (p *Email) UnmarshalProtoFrom(bits []byte) error
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

#### func (*EmailBlastRequest) JSONString

```go
func (p *EmailBlastRequest) JSONString() *common.String
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

#### func (*EmailBlastRequest) UnmarshalProtoFrom

```go
func (p *EmailBlastRequest) UnmarshalProtoFrom(bits []byte) error
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

#### func (*EmailRequest) JSONString

```go
func (p *EmailRequest) JSONString() *common.String
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

#### func (*EmailRequest) UnmarshalProtoFrom

```go
func (p *EmailRequest) UnmarshalProtoFrom(bits []byte) error
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

#### type Event

```go
type Event struct {
	Date                 *common.String `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Type                 *common.String `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ClientId             *common.String `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientName           *common.String `protobuf:"bytes,4,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	Ip                   *common.String `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	LocationInfo         *common.String `protobuf:"bytes,6,opt,name=location_info,json=locationInfo,proto3" json:"location_info,omitempty"`
	Details              *common.String `protobuf:"bytes,7,opt,name=details,proto3" json:"details,omitempty"`
	UserId               *common.String `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*Event) Descriptor

```go
func (*Event) Descriptor() ([]byte, []int)
```

#### func (*Event) GetClientId

```go
func (m *Event) GetClientId() *common.String
```

#### func (*Event) GetClientName

```go
func (m *Event) GetClientName() *common.String
```

#### func (*Event) GetDate

```go
func (m *Event) GetDate() *common.String
```

#### func (*Event) GetDetails

```go
func (m *Event) GetDetails() *common.String
```

#### func (*Event) GetIp

```go
func (m *Event) GetIp() *common.String
```

#### func (*Event) GetLocationInfo

```go
func (m *Event) GetLocationInfo() *common.String
```

#### func (*Event) GetType

```go
func (m *Event) GetType() *common.String
```

#### func (*Event) GetUserId

```go
func (m *Event) GetUserId() *common.String
```

#### func (*Event) JSONString

```go
func (s *Event) JSONString() *common.String
```

#### func (*Event) ProtoMessage

```go
func (*Event) ProtoMessage()
```

#### func (*Event) Reset

```go
func (m *Event) Reset()
```

#### func (*Event) String

```go
func (m *Event) String() string
```

#### func (*Event) UnmarshalJSONfrom

```go
func (s *Event) UnmarshalJSONfrom(bits []byte) error
```

#### func (*Event) UnmarshalProtofrom

```go
func (s *Event) UnmarshalProtofrom(bits []byte) error
```

#### func (*Event) XXX_DiscardUnknown

```go
func (m *Event) XXX_DiscardUnknown()
```

#### func (*Event) XXX_Marshal

```go
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Event) XXX_Merge

```go
func (m *Event) XXX_Merge(src proto.Message)
```

#### func (*Event) XXX_Size

```go
func (m *Event) XXX_Size() int
```

#### func (*Event) XXX_Unmarshal

```go
func (m *Event) XXX_Unmarshal(b []byte) error
```

#### type Fax

```go
type Fax struct {
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


#### func (*Fax) Descriptor

```go
func (*Fax) Descriptor() ([]byte, []int)
```

#### func (*Fax) GetCallback

```go
func (m *Fax) GetCallback() *common.String
```

#### func (*Fax) GetFrom

```go
func (m *Fax) GetFrom() *common.String
```

#### func (*Fax) GetMediaUrl

```go
func (m *Fax) GetMediaUrl() *common.String
```

#### func (*Fax) GetQuality

```go
func (m *Fax) GetQuality() *common.String
```

#### func (*Fax) GetStoreMedia

```go
func (m *Fax) GetStoreMedia() bool
```

#### func (*Fax) GetTo

```go
func (m *Fax) GetTo() *common.String
```

#### func (*Fax) JSONString

```go
func (p *Fax) JSONString() *common.String
```

#### func (*Fax) ProtoMessage

```go
func (*Fax) ProtoMessage()
```

#### func (*Fax) Reset

```go
func (m *Fax) Reset()
```

#### func (*Fax) String

```go
func (m *Fax) String() string
```

#### func (*Fax) UnmarshalProtoFrom

```go
func (p *Fax) UnmarshalProtoFrom(bits []byte) error
```

#### func (*Fax) XXX_DiscardUnknown

```go
func (m *Fax) XXX_DiscardUnknown()
```

#### func (*Fax) XXX_Marshal

```go
func (m *Fax) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Fax) XXX_Merge

```go
func (m *Fax) XXX_Merge(src proto.Message)
```

#### func (*Fax) XXX_Size

```go
func (m *Fax) XXX_Size() int
```

#### func (*Fax) XXX_Unmarshal

```go
func (m *Fax) XXX_Unmarshal(b []byte) error
```

#### type FaxBlast

```go
type FaxBlast struct {
	To                   *common.StringArray `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	From                 *common.String      `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	MediaUrl             *common.String      `protobuf:"bytes,3,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	Quality              *common.String      `protobuf:"bytes,4,opt,name=quality,proto3" json:"quality,omitempty"`
	Callback             *common.String      `protobuf:"bytes,5,opt,name=callback,proto3" json:"callback,omitempty"`
	StoreMedia           bool                `protobuf:"varint,6,opt,name=store_media,json=storeMedia,proto3" json:"store_media,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}
```


#### func (*FaxBlast) Descriptor

```go
func (*FaxBlast) Descriptor() ([]byte, []int)
```

#### func (*FaxBlast) GetCallback

```go
func (m *FaxBlast) GetCallback() *common.String
```

#### func (*FaxBlast) GetFrom

```go
func (m *FaxBlast) GetFrom() *common.String
```

#### func (*FaxBlast) GetMediaUrl

```go
func (m *FaxBlast) GetMediaUrl() *common.String
```

#### func (*FaxBlast) GetQuality

```go
func (m *FaxBlast) GetQuality() *common.String
```

#### func (*FaxBlast) GetStoreMedia

```go
func (m *FaxBlast) GetStoreMedia() bool
```

#### func (*FaxBlast) GetTo

```go
func (m *FaxBlast) GetTo() *common.StringArray
```

#### func (*FaxBlast) JSONString

```go
func (p *FaxBlast) JSONString() *common.String
```

#### func (*FaxBlast) ProtoMessage

```go
func (*FaxBlast) ProtoMessage()
```

#### func (*FaxBlast) Reset

```go
func (m *FaxBlast) Reset()
```

#### func (*FaxBlast) String

```go
func (m *FaxBlast) String() string
```

#### func (*FaxBlast) UnmarshalProtoFrom

```go
func (p *FaxBlast) UnmarshalProtoFrom(bits []byte) error
```

#### func (*FaxBlast) XXX_DiscardUnknown

```go
func (m *FaxBlast) XXX_DiscardUnknown()
```

#### func (*FaxBlast) XXX_Marshal

```go
func (m *FaxBlast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*FaxBlast) XXX_Merge

```go
func (m *FaxBlast) XXX_Merge(src proto.Message)
```

#### func (*FaxBlast) XXX_Size

```go
func (m *FaxBlast) XXX_Size() int
```

#### func (*FaxBlast) XXX_Unmarshal

```go
func (m *FaxBlast) XXX_Unmarshal(b []byte) error
```

#### type FaxResponse

```go
type FaxResponse struct {
	Id                   *common.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MediaUrl             *common.String     `protobuf:"bytes,3,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	To                   *common.String     `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	From                 *common.String     `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	Status               *common.String     `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
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

#### func (*FaxResponse) GetStatus

```go
func (m *FaxResponse) GetStatus() *common.String
```

#### func (*FaxResponse) GetTo

```go
func (m *FaxResponse) GetTo() *common.String
```

#### func (*FaxResponse) JSONString

```go
func (p *FaxResponse) JSONString() *common.String
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

#### func (*FaxResponse) UnmarshalJSONFrom

```go
func (p *FaxResponse) UnmarshalJSONFrom(bits []byte) error
```

#### func (*FaxResponse) UnmarshalProtoFrom

```go
func (p *FaxResponse) UnmarshalProtoFrom(bits []byte) error
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

#### func (*Identity) JSONString

```go
func (p *Identity) JSONString() *common.String
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

#### func (*Identity) UnmarshalJSONFrom

```go
func (p *Identity) UnmarshalJSONFrom(bits []byte) error
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

#### func (*JSONWebKeys) JSONString

```go
func (p *JSONWebKeys) JSONString() *common.String
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
	Email                *common.String      `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	PrivateKey           []byte              `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	PriveKeyId           *common.String      `protobuf:"bytes,3,opt,name=prive_key_id,json=priveKeyId,proto3" json:"prive_key_id,omitempty"`
	Subject              *common.String      `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Scopes               *common.StringArray `protobuf:"bytes,5,opt,name=scopes,proto3" json:"scopes,omitempty"`
	TokenUrl             *common.String      `protobuf:"bytes,6,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	Expires              *common.String      `protobuf:"bytes,7,opt,name=expires,proto3" json:"expires,omitempty"`
	Audience             *common.String      `protobuf:"bytes,8,opt,name=audience,proto3" json:"audience,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}
```


#### func (*JWT) AuthCodeURL

```go
func (c *JWT) AuthCodeURL(state string, audience string) string
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

#### func (*JWT) DeepEqual

```go
func (s *JWT) DeepEqual(y interface{}) bool
```

#### func (*JWT) Descriptor

```go
func (*JWT) Descriptor() ([]byte, []int)
```

#### func (*JWT) GetAudience

```go
func (m *JWT) GetAudience() *common.String
```

#### func (*JWT) GetEmail

```go
func (m *JWT) GetEmail() *common.String
```

#### func (*JWT) GetExpires

```go
func (m *JWT) GetExpires() *common.String
```

#### func (*JWT) GetPrivateKey

```go
func (m *JWT) GetPrivateKey() []byte
```

#### func (*JWT) GetPriveKeyId

```go
func (m *JWT) GetPriveKeyId() *common.String
```

#### func (*JWT) GetRequestMetadata

```go
func (c *JWT) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
```

#### func (*JWT) GetScopes

```go
func (m *JWT) GetScopes() *common.StringArray
```

#### func (*JWT) GetSubject

```go
func (m *JWT) GetSubject() *common.String
```

#### func (*JWT) GetTokenUrl

```go
func (m *JWT) GetTokenUrl() *common.String
```

#### func (*JWT) JSONString

```go
func (s *JWT) JSONString() *common.String
```

#### func (*JWT) NewAPIClientSet

```go
func (c *JWT) NewAPIClientSet(addr string) (*ClientSet, error)
```

#### func (*JWT) PerRPCCredentials

```go
func (c *JWT) PerRPCCredentials() (credentials.PerRPCCredentials, error)
```

#### func (*JWT) ProtoMessage

```go
func (*JWT) ProtoMessage()
```

#### func (*JWT) RequireTransportSecurity

```go
func (c *JWT) RequireTransportSecurity() bool
```

#### func (*JWT) Reset

```go
func (m *JWT) Reset()
```

#### func (*JWT) String

```go
func (m *JWT) String() string
```

#### func (*JWT) ToContext

```go
func (s *JWT) ToContext(ctx context.Context, key string) context.Context
```

#### func (*JWT) Token

```go
func (c *JWT) Token() (*oauth2.Token, error)
```

#### func (*JWT) UnmarshalJSONFrom

```go
func (s *JWT) UnmarshalJSONFrom(b []byte) error
```

#### func (*JWT) UnmarshalProtoFrom

```go
func (s *JWT) UnmarshalProtoFrom(b []byte) error
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

#### func (*Jwks) JSONString

```go
func (p *Jwks) JSONString() *common.String
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

#### func (*Jwks) UnmarshalJSONFrom

```go
func (p *Jwks) UnmarshalJSONFrom(bits []byte) error
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

#### type ManagementServiceClient

```go
type ManagementServiceClient interface {
	QueryUsers(ctx context.Context, in *Query, opts ...grpc.CallOption) (ManagementService_QueryUsersClient, error)
	GetUser(ctx context.Context, in *common.Identifier, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UserRoles(ctx context.Context, in *common.Identifier, opts ...grpc.CallOption) (ManagementService_UserRolesClient, error)
	AddUserRoles(ctx context.Context, in *AddUserRolesRequest, opts ...grpc.CallOption) (*common.String, error)
	QueryUserEvents(ctx context.Context, in *Query, opts ...grpc.CallOption) (ManagementService_QueryUserEventsClient, error)
}
```

ManagementServiceClient is the client API for ManagementService service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to
https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

#### func  NewManagementServiceClient

```go
func NewManagementServiceClient(cc *grpc.ClientConn) ManagementServiceClient
```

#### type ManagementServiceServer

```go
type ManagementServiceServer interface {
	QueryUsers(*Query, ManagementService_QueryUsersServer) error
	GetUser(context.Context, *common.Identifier) (*User, error)
	UpdateUser(context.Context, *User) (*User, error)
	UserRoles(*common.Identifier, ManagementService_UserRolesServer) error
	AddUserRoles(context.Context, *AddUserRolesRequest) (*common.String, error)
	QueryUserEvents(*Query, ManagementService_QueryUserEventsServer) error
}
```

ManagementServiceServer is the server API for ManagementService service.

#### type ManagementService_QueryUserEventsClient

```go
type ManagementService_QueryUserEventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}
```


#### type ManagementService_QueryUserEventsServer

```go
type ManagementService_QueryUserEventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}
```


#### type ManagementService_QueryUsersClient

```go
type ManagementService_QueryUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}
```


#### type ManagementService_QueryUsersServer

```go
type ManagementService_QueryUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}
```


#### type ManagementService_UserRolesClient

```go
type ManagementService_UserRolesClient interface {
	Recv() (*Role, error)
	grpc.ClientStream
}
```


#### type ManagementService_UserRolesServer

```go
type ManagementService_UserRolesServer interface {
	Send(*Role) error
	grpc.ServerStream
}
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

#### type OAuth2

```go
type OAuth2 struct {
	ClientId             *common.String      `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         *common.String      `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	TokenUrl             *common.String      `protobuf:"bytes,3,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	AuthUrl              *common.String      `protobuf:"bytes,4,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
	Scopes               *common.StringArray `protobuf:"bytes,5,opt,name=scopes,proto3" json:"scopes,omitempty"`
	Redirect             *common.String      `protobuf:"bytes,6,opt,name=redirect,proto3" json:"redirect,omitempty"`
	Code                 *common.String      `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}
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

#### func (*OAuth2) DeepEqual

```go
func (s *OAuth2) DeepEqual(y interface{}) bool
```

#### func (*OAuth2) Descriptor

```go
func (*OAuth2) Descriptor() ([]byte, []int)
```

#### func (*OAuth2) GetAuthUrl

```go
func (m *OAuth2) GetAuthUrl() *common.String
```

#### func (*OAuth2) GetClientId

```go
func (m *OAuth2) GetClientId() *common.String
```

#### func (*OAuth2) GetClientSecret

```go
func (m *OAuth2) GetClientSecret() *common.String
```

#### func (*OAuth2) GetCode

```go
func (m *OAuth2) GetCode() *common.String
```

#### func (*OAuth2) GetRedirect

```go
func (m *OAuth2) GetRedirect() *common.String
```

#### func (*OAuth2) GetRequestMetadata

```go
func (c *OAuth2) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
```

#### func (*OAuth2) GetScopes

```go
func (m *OAuth2) GetScopes() *common.StringArray
```

#### func (*OAuth2) GetTokenUrl

```go
func (m *OAuth2) GetTokenUrl() *common.String
```

#### func (*OAuth2) JSONString

```go
func (s *OAuth2) JSONString() *common.String
```

#### func (*OAuth2) NewAPIClientSet

```go
func (c *OAuth2) NewAPIClientSet(addr string) (*ClientSet, error)
```

#### func (*OAuth2) PerRPCCredentials

```go
func (c *OAuth2) PerRPCCredentials() (credentials.PerRPCCredentials, error)
```

#### func (*OAuth2) ProtoMessage

```go
func (*OAuth2) ProtoMessage()
```

#### func (*OAuth2) RequireTransportSecurity

```go
func (c *OAuth2) RequireTransportSecurity() bool
```

#### func (*OAuth2) Reset

```go
func (m *OAuth2) Reset()
```

#### func (*OAuth2) String

```go
func (m *OAuth2) String() string
```

#### func (*OAuth2) ToContext

```go
func (s *OAuth2) ToContext(ctx context.Context, key string) context.Context
```

#### func (*OAuth2) Token

```go
func (c *OAuth2) Token() (*oauth2.Token, error)
```

#### func (*OAuth2) UnmarshalJSONFrom

```go
func (s *OAuth2) UnmarshalJSONFrom(b []byte) error
```

#### func (*OAuth2) UnmarshalProtoFrom

```go
func (s *OAuth2) UnmarshalProtoFrom(b []byte) error
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

#### func (*PhoneNumber) JSONString

```go
func (p *PhoneNumber) JSONString() *common.String
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

#### func (*PhoneNumber) UnmarshalJSONFrom

```go
func (p *PhoneNumber) UnmarshalJSONFrom(bits []byte) error
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

#### func (*PhoneNumberResource) JSONString

```go
func (p *PhoneNumberResource) JSONString() *common.String
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

#### func (*PhoneNumberResource) UnmarshalJSONFrom

```go
func (p *PhoneNumberResource) UnmarshalJSONFrom(bits []byte) error
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

#### func (Plan) String

```go
func (x Plan) String() string
```

#### type Query

```go
type Query struct {
	Query                *common.String `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	Fields               *common.String `protobuf:"bytes,5,opt,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*Query) Descriptor

```go
func (*Query) Descriptor() ([]byte, []int)
```

#### func (*Query) GetFields

```go
func (m *Query) GetFields() *common.String
```

#### func (*Query) GetQuery

```go
func (m *Query) GetQuery() *common.String
```

#### func (*Query) ProtoMessage

```go
func (*Query) ProtoMessage()
```

#### func (*Query) Reset

```go
func (m *Query) Reset()
```

#### func (*Query) String

```go
func (m *Query) String() string
```

#### func (*Query) XXX_DiscardUnknown

```go
func (m *Query) XXX_DiscardUnknown()
```

#### func (*Query) XXX_Marshal

```go
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Query) XXX_Merge

```go
func (m *Query) XXX_Merge(src proto.Message)
```

#### func (*Query) XXX_Size

```go
func (m *Query) XXX_Size() int
```

#### func (*Query) XXX_Unmarshal

```go
func (m *Query) XXX_Unmarshal(b []byte) error
```

#### type RenderRequest

```go
type RenderRequest struct {
	Name                 *common.String `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Text                 *common.String `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Data                 *common.String `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
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
func (m *RenderRequest) GetData() *common.String
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

#### func (*SMS) JSONString

```go
func (p *SMS) JSONString() *common.String
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

#### func (*SMS) UnmarshalProtoFrom

```go
func (p *SMS) UnmarshalProtoFrom(bits []byte) error
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

#### func (*SMSResponse) JSONString

```go
func (p *SMSResponse) JSONString() *common.String
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

#### func (*SMSResponse) UnmarshalJSONFrom

```go
func (p *SMSResponse) UnmarshalJSONFrom(bits []byte) error
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
	Amount               *common.Int64      `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	DaysUntilDue         *common.Int64      `protobuf:"bytes,3,opt,name=days_until_due,json=daysUntilDue,proto3" json:"days_until_due,omitempty"`
	Annotations          *common.StringMap  `protobuf:"bytes,10,opt,name=annotations,proto3" json:"annotations,omitempty"`
	Plan                 Plan               `protobuf:"varint,4,opt,name=plan,proto3,enum=api.Plan" json:"plan,omitempty"`
	User                 *User              `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Status               *common.String     `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func (*SubscriptionResponse) Descriptor

```go
func (*SubscriptionResponse) Descriptor() ([]byte, []int)
```

#### func (*SubscriptionResponse) GetAmount

```go
func (m *SubscriptionResponse) GetAmount() *common.Int64
```

#### func (*SubscriptionResponse) GetAnnotations

```go
func (m *SubscriptionResponse) GetAnnotations() *common.StringMap
```

#### func (*SubscriptionResponse) GetDaysUntilDue

```go
func (m *SubscriptionResponse) GetDaysUntilDue() *common.Int64
```

#### func (*SubscriptionResponse) GetId

```go
func (m *SubscriptionResponse) GetId() *common.Identifier
```

#### func (*SubscriptionResponse) GetPlan

```go
func (m *SubscriptionResponse) GetPlan() Plan
```

#### func (*SubscriptionResponse) GetStatus

```go
func (m *SubscriptionResponse) GetStatus() *common.String
```

#### func (*SubscriptionResponse) GetUser

```go
func (m *SubscriptionResponse) GetUser() *User
```

#### func (*SubscriptionResponse) JSONString

```go
func (p *SubscriptionResponse) JSONString() *common.String
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

#### func (*SubscriptionResponse) UnmarshalJSONFrom

```go
func (p *SubscriptionResponse) UnmarshalJSONFrom(bits []byte) error
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

#### func (*User) JSONString

```go
func (p *User) JSONString() *common.String
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

#### func (*User) UnmarshalJSONFrom

```go
func (p *User) UnmarshalJSONFrom(bits []byte) error
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
