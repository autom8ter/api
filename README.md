# api
--
    import "github.com/autom8ter/api"

Package api is a reverse proxy.

It translates gRPC into RESTful JSON APIs.

## Usage

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

#### func  RegisterDebugServiceHandler

```go
func RegisterDebugServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
```
RegisterDebugServiceHandler registers the http handlers for service DebugService
to "mux". The handlers forward requests to the grpc endpoint over "conn".

#### func  RegisterDebugServiceHandlerClient

```go
func RegisterDebugServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client DebugServiceClient) error
```
RegisterDebugServiceHandlerClient registers the http handlers for service
DebugService to "mux". The handlers forward requests to the grpc endpoint over
the given implementation of "DebugServiceClient". Note: the gRPC framework
executes interceptors within the gRPC handler. If the passed in
"DebugServiceClient" doesn't go through the normal gRPC flow (creating a gRPC
client etc.) then it will be up to the passed in "DebugServiceClient" to call
the correct interceptors.

#### func  RegisterDebugServiceHandlerFromEndpoint

```go
func RegisterDebugServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
```
RegisterDebugServiceHandlerFromEndpoint is same as RegisterDebugServiceHandler
but automatically dials to "endpoint" and closes the connection when "ctx" gets
done.

#### func  RegisterDebugServiceServer

```go
func RegisterDebugServiceServer(s *grpc.Server, srv DebugServiceServer)
```

#### func  RegisterDocumentServiceHandler

```go
func RegisterDocumentServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
```
RegisterDocumentServiceHandler registers the http handlers for service
DocumentService to "mux". The handlers forward requests to the grpc endpoint
over "conn".

#### func  RegisterDocumentServiceHandlerClient

```go
func RegisterDocumentServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client DocumentServiceClient) error
```
RegisterDocumentServiceHandlerClient registers the http handlers for service
DocumentService to "mux". The handlers forward requests to the grpc endpoint
over the given implementation of "DocumentServiceClient". Note: the gRPC
framework executes interceptors within the gRPC handler. If the passed in
"DocumentServiceClient" doesn't go through the normal gRPC flow (creating a gRPC
client etc.) then it will be up to the passed in "DocumentServiceClient" to call
the correct interceptors.

#### func  RegisterDocumentServiceHandlerFromEndpoint

```go
func RegisterDocumentServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
```
RegisterDocumentServiceHandlerFromEndpoint is same as
RegisterDocumentServiceHandler but automatically dials to "endpoint" and closes
the connection when "ctx" gets done.

#### func  RegisterDocumentServiceServer

```go
func RegisterDocumentServiceServer(s *grpc.Server, srv DocumentServiceServer)
```

#### func  RegisterEventServiceHandler

```go
func RegisterEventServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
```
RegisterEventServiceHandler registers the http handlers for service EventService
to "mux". The handlers forward requests to the grpc endpoint over "conn".

#### func  RegisterEventServiceHandlerClient

```go
func RegisterEventServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client EventServiceClient) error
```
RegisterEventServiceHandlerClient registers the http handlers for service
EventService to "mux". The handlers forward requests to the grpc endpoint over
the given implementation of "EventServiceClient". Note: the gRPC framework
executes interceptors within the gRPC handler. If the passed in
"EventServiceClient" doesn't go through the normal gRPC flow (creating a gRPC
client etc.) then it will be up to the passed in "EventServiceClient" to call
the correct interceptors.

#### func  RegisterEventServiceHandlerFromEndpoint

```go
func RegisterEventServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
```
RegisterEventServiceHandlerFromEndpoint is same as RegisterEventServiceHandler
but automatically dials to "endpoint" and closes the connection when "ctx" gets
done.

#### func  RegisterEventServiceServer

```go
func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer)
```

#### func  RegisterSubscriptionServiceHandler

```go
func RegisterSubscriptionServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
```
RegisterSubscriptionServiceHandler registers the http handlers for service
SubscriptionService to "mux". The handlers forward requests to the grpc endpoint
over "conn".

#### func  RegisterSubscriptionServiceHandlerClient

```go
func RegisterSubscriptionServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client SubscriptionServiceClient) error
```
RegisterSubscriptionServiceHandlerClient registers the http handlers for service
SubscriptionService to "mux". The handlers forward requests to the grpc endpoint
over the given implementation of "SubscriptionServiceClient". Note: the gRPC
framework executes interceptors within the gRPC handler. If the passed in
"SubscriptionServiceClient" doesn't go through the normal gRPC flow (creating a
gRPC client etc.) then it will be up to the passed in
"SubscriptionServiceClient" to call the correct interceptors.

#### func  RegisterSubscriptionServiceHandlerFromEndpoint

```go
func RegisterSubscriptionServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
```
RegisterSubscriptionServiceHandlerFromEndpoint is same as
RegisterSubscriptionServiceHandler but automatically dials to "endpoint" and
closes the connection when "ctx" gets done.

#### func  RegisterSubscriptionServiceServer

```go
func RegisterSubscriptionServiceServer(s *grpc.Server, srv SubscriptionServiceServer)
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

#### type Address

```go
type Address struct {
	City                 *common.String `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	State                *common.String `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Line1                *common.String `protobuf:"bytes,3,opt,name=line1,proto3" json:"line1,omitempty"`
	Line2                *common.String `protobuf:"bytes,4,opt,name=line2,proto3" json:"line2,omitempty"`
	Zip                  *common.String `protobuf:"bytes,5,opt,name=zip,proto3" json:"zip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*Address) Descriptor

```go
func (*Address) Descriptor() ([]byte, []int)
```

#### func (*Address) GetCity

```go
func (m *Address) GetCity() *common.String
```

#### func (*Address) GetLine1

```go
func (m *Address) GetLine1() *common.String
```

#### func (*Address) GetLine2

```go
func (m *Address) GetLine2() *common.String
```

#### func (*Address) GetState

```go
func (m *Address) GetState() *common.String
```

#### func (*Address) GetZip

```go
func (m *Address) GetZip() *common.String
```

#### func (*Address) ProtoMessage

```go
func (*Address) ProtoMessage()
```

#### func (*Address) Reset

```go
func (m *Address) Reset()
```

#### func (*Address) String

```go
func (m *Address) String() string
```

#### func (*Address) XXX_DiscardUnknown

```go
func (m *Address) XXX_DiscardUnknown()
```

#### func (*Address) XXX_Marshal

```go
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Address) XXX_Merge

```go
func (m *Address) XXX_Merge(src proto.Message)
```

#### func (*Address) XXX_Size

```go
func (m *Address) XXX_Size() int
```

#### func (*Address) XXX_Unmarshal

```go
func (m *Address) XXX_Unmarshal(b []byte) error
```

#### type AppMetadata

```go
type AppMetadata struct {
	Description          *common.String    `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	AccountBalance       *common.String    `protobuf:"bytes,2,opt,name=account_balance,json=accountBalance,proto3" json:"account_balance,omitempty"`
	Plan                 *Plan             `protobuf:"bytes,3,opt,name=plan,proto3" json:"plan,omitempty"`
	Tags                 *common.StringMap `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func (*AppMetadata) Descriptor

```go
func (*AppMetadata) Descriptor() ([]byte, []int)
```

#### func (*AppMetadata) GetAccountBalance

```go
func (m *AppMetadata) GetAccountBalance() *common.String
```

#### func (*AppMetadata) GetDescription

```go
func (m *AppMetadata) GetDescription() *common.String
```

#### func (*AppMetadata) GetPlan

```go
func (m *AppMetadata) GetPlan() *Plan
```

#### func (*AppMetadata) GetTags

```go
func (m *AppMetadata) GetTags() *common.StringMap
```

#### func (*AppMetadata) JSONString

```go
func (p *AppMetadata) JSONString() *common.String
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

#### func (*AppMetadata) UnmarshalJSONFrom

```go
func (p *AppMetadata) UnmarshalJSONFrom(bits []byte) error
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

#### type Card

```go
type Card struct {
	Number               *common.Identifier `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	ExpMonth             *common.String     `protobuf:"bytes,2,opt,name=exp_month,json=expMonth,proto3" json:"exp_month,omitempty"`
	ExpYear              *common.String     `protobuf:"bytes,3,opt,name=exp_year,json=expYear,proto3" json:"exp_year,omitempty"`
	Cvc                  *common.String     `protobuf:"bytes,4,opt,name=cvc,proto3" json:"cvc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
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
func (m *Card) GetNumber() *common.Identifier
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

#### type CategoryQuery

```go
type CategoryQuery struct {
	Category             *common.String `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*CategoryQuery) Descriptor

```go
func (*CategoryQuery) Descriptor() ([]byte, []int)
```

#### func (*CategoryQuery) GetCategory

```go
func (m *CategoryQuery) GetCategory() *common.String
```

#### func (*CategoryQuery) ProtoMessage

```go
func (*CategoryQuery) ProtoMessage()
```

#### func (*CategoryQuery) Reset

```go
func (m *CategoryQuery) Reset()
```

#### func (*CategoryQuery) String

```go
func (m *CategoryQuery) String() string
```

#### func (*CategoryQuery) XXX_DiscardUnknown

```go
func (m *CategoryQuery) XXX_DiscardUnknown()
```

#### func (*CategoryQuery) XXX_Marshal

```go
func (m *CategoryQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*CategoryQuery) XXX_Merge

```go
func (m *CategoryQuery) XXX_Merge(src proto.Message)
```

#### func (*CategoryQuery) XXX_Size

```go
func (m *CategoryQuery) XXX_Size() int
```

#### func (*CategoryQuery) XXX_Unmarshal

```go
func (m *CategoryQuery) XXX_Unmarshal(b []byte) error
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
func (c *ClientCredentials) NewAPIClientSet(ctx context.Context, addr string) (*ClientSet, error)
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
	Debug         DebugServiceClient
	Subscriptions SubscriptionServiceClient
	Users         UserServiceClient
	Auth          AuthenticationServiceClient
	Events        EventServiceClient
}
```


#### func  NewClientSet

```go
func NewClientSet(conn *grpc.ClientConn) *ClientSet
```

#### type DebugServiceClient

```go
type DebugServiceClient interface {
	Echo(ctx context.Context, in *common.String, opts ...grpc.CallOption) (*common.String, error)
}
```

DebugServiceClient is the client API for DebugService service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to
https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

#### func  NewDebugServiceClient

```go
func NewDebugServiceClient(cc *grpc.ClientConn) DebugServiceClient
```

#### type DebugServiceServer

```go
type DebugServiceServer interface {
	Echo(context.Context, *common.String) (*common.String, error)
}
```

DebugServiceServer is the server API for DebugService service.

#### type DefaultGCPCredentials

```go
type DefaultGCPCredentials struct {
	Scopes               *common.StringArray `protobuf:"bytes,1,opt,name=scopes,proto3" json:"scopes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
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
func (m *DefaultGCPCredentials) GetScopes() *common.StringArray
```

#### func (*DefaultGCPCredentials) NewAPIClientSet

```go
func (c *DefaultGCPCredentials) NewAPIClientSet(ctx context.Context, addr string) (*ClientSet, error)
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

#### func (*DefaultGCPCredentials) ToContext

```go
func (s *DefaultGCPCredentials) ToContext(ctx context.Context, key string) context.Context
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

#### type Document

```go
type Document struct {
	Category             *common.String    `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Name                 *common.String    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data                 *common.StringMap `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func  NewDocument

```go
func NewDocument(category, name string, data map[string]interface{}) *Document
```

#### func (*Document) Descriptor

```go
func (*Document) Descriptor() ([]byte, []int)
```

#### func (*Document) GetCategory

```go
func (m *Document) GetCategory() *common.String
```

#### func (*Document) GetData

```go
func (m *Document) GetData() *common.StringMap
```

#### func (*Document) GetName

```go
func (m *Document) GetName() *common.String
```

#### func (*Document) JSONString

```go
func (p *Document) JSONString() *common.String
```

#### func (*Document) ProtoMessage

```go
func (*Document) ProtoMessage()
```

#### func (*Document) Render

```go
func (d *Document) Render(s *common.String, w io.Writer) error
```

#### func (*Document) Reset

```go
func (m *Document) Reset()
```

#### func (*Document) String

```go
func (m *Document) String() string
```

#### func (*Document) UnmarshalJSONFrom

```go
func (p *Document) UnmarshalJSONFrom(bits []byte) error
```

#### func (*Document) XXX_DiscardUnknown

```go
func (m *Document) XXX_DiscardUnknown()
```

#### func (*Document) XXX_Marshal

```go
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Document) XXX_Merge

```go
func (m *Document) XXX_Merge(src proto.Message)
```

#### func (*Document) XXX_Size

```go
func (m *Document) XXX_Size() int
```

#### func (*Document) XXX_Unmarshal

```go
func (m *Document) XXX_Unmarshal(b []byte) error
```

#### type DocumentQuery

```go
type DocumentQuery struct {
	Category             *common.String     `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Name                 *common.Identifier `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func (*DocumentQuery) Descriptor

```go
func (*DocumentQuery) Descriptor() ([]byte, []int)
```

#### func (*DocumentQuery) GetCategory

```go
func (m *DocumentQuery) GetCategory() *common.String
```

#### func (*DocumentQuery) GetName

```go
func (m *DocumentQuery) GetName() *common.Identifier
```

#### func (*DocumentQuery) ProtoMessage

```go
func (*DocumentQuery) ProtoMessage()
```

#### func (*DocumentQuery) Reset

```go
func (m *DocumentQuery) Reset()
```

#### func (*DocumentQuery) String

```go
func (m *DocumentQuery) String() string
```

#### func (*DocumentQuery) XXX_DiscardUnknown

```go
func (m *DocumentQuery) XXX_DiscardUnknown()
```

#### func (*DocumentQuery) XXX_Marshal

```go
func (m *DocumentQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*DocumentQuery) XXX_Merge

```go
func (m *DocumentQuery) XXX_Merge(src proto.Message)
```

#### func (*DocumentQuery) XXX_Size

```go
func (m *DocumentQuery) XXX_Size() int
```

#### func (*DocumentQuery) XXX_Unmarshal

```go
func (m *DocumentQuery) XXX_Unmarshal(b []byte) error
```

#### type DocumentServiceClient

```go
type DocumentServiceClient interface {
	GetDocument(ctx context.Context, in *DocumentQuery, opts ...grpc.CallOption) (*Document, error)
	Update(ctx context.Context, in *Document, opts ...grpc.CallOption) (*common.Empty, error)
	Create(ctx context.Context, in *Document, opts ...grpc.CallOption) (*common.Empty, error)
	Delete(ctx context.Context, in *DocumentQuery, opts ...grpc.CallOption) (*common.Empty, error)
	ListDocuments(ctx context.Context, in *CategoryQuery, opts ...grpc.CallOption) (DocumentService_ListDocumentsClient, error)
}
```

DocumentServiceClient is the client API for DocumentService service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to
https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

#### func  NewDocumentServiceClient

```go
func NewDocumentServiceClient(cc *grpc.ClientConn) DocumentServiceClient
```

#### type DocumentServiceServer

```go
type DocumentServiceServer interface {
	GetDocument(context.Context, *DocumentQuery) (*Document, error)
	Update(context.Context, *Document) (*common.Empty, error)
	Create(context.Context, *Document) (*common.Empty, error)
	Delete(context.Context, *DocumentQuery) (*common.Empty, error)
	ListDocuments(*CategoryQuery, DocumentService_ListDocumentsServer) error
}
```

DocumentServiceServer is the server API for DocumentService service.

#### type DocumentService_ListDocumentsClient

```go
type DocumentService_ListDocumentsClient interface {
	Recv() (*Document, error)
	grpc.ClientStream
}
```


#### type DocumentService_ListDocumentsServer

```go
type DocumentService_ListDocumentsServer interface {
	Send(*Document) error
	grpc.ServerStream
}
```


#### type Event

```go
type Event struct {
	Date                 *common.String    `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Type                 *common.String    `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ClientId             *common.String    `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientName           *common.String    `protobuf:"bytes,4,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	Ip                   *common.String    `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	LocationInfo         *common.String    `protobuf:"bytes,6,opt,name=location_info,json=locationInfo,proto3" json:"location_info,omitempty"`
	Details              *common.String    `protobuf:"bytes,7,opt,name=details,proto3" json:"details,omitempty"`
	UserId               *common.String    `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Annotations          *common.StringMap `protobuf:"bytes,9,opt,name=annotations,proto3" json:"annotations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func (*Event) Descriptor

```go
func (*Event) Descriptor() ([]byte, []int)
```

#### func (*Event) GetAnnotations

```go
func (m *Event) GetAnnotations() *common.StringMap
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

#### type EventQuery

```go
type EventQuery struct {
	Date                 *common.String `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Type                 *common.String `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ClientId             *common.String `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	UserId               *common.String `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}
```


#### func (*EventQuery) Descriptor

```go
func (*EventQuery) Descriptor() ([]byte, []int)
```

#### func (*EventQuery) GetClientId

```go
func (m *EventQuery) GetClientId() *common.String
```

#### func (*EventQuery) GetDate

```go
func (m *EventQuery) GetDate() *common.String
```

#### func (*EventQuery) GetType

```go
func (m *EventQuery) GetType() *common.String
```

#### func (*EventQuery) GetUserId

```go
func (m *EventQuery) GetUserId() *common.String
```

#### func (*EventQuery) ProtoMessage

```go
func (*EventQuery) ProtoMessage()
```

#### func (*EventQuery) Reset

```go
func (m *EventQuery) Reset()
```

#### func (*EventQuery) String

```go
func (m *EventQuery) String() string
```

#### func (*EventQuery) XXX_DiscardUnknown

```go
func (m *EventQuery) XXX_DiscardUnknown()
```

#### func (*EventQuery) XXX_Marshal

```go
func (m *EventQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*EventQuery) XXX_Merge

```go
func (m *EventQuery) XXX_Merge(src proto.Message)
```

#### func (*EventQuery) XXX_Size

```go
func (m *EventQuery) XXX_Size() int
```

#### func (*EventQuery) XXX_Unmarshal

```go
func (m *EventQuery) XXX_Unmarshal(b []byte) error
```

#### type EventServiceClient

```go
type EventServiceClient interface {
	GetEvents(ctx context.Context, in *EventQuery, opts ...grpc.CallOption) (EventService_GetEventsClient, error)
}
```

EventServiceClient is the client API for EventService service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to
https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

#### func  NewEventServiceClient

```go
func NewEventServiceClient(cc *grpc.ClientConn) EventServiceClient
```

#### type EventServiceServer

```go
type EventServiceServer interface {
	GetEvents(*EventQuery, EventService_GetEventsServer) error
}
```

EventServiceServer is the server API for EventService service.

#### type EventService_GetEventsClient

```go
type EventService_GetEventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}
```


#### type EventService_GetEventsServer

```go
type EventService_GetEventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}
```


#### type Identity

```go
type Identity struct {
	Connection           *common.String     `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	UserId               *common.Identifier `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Provider             *common.String     `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	IsSocial             *common.Bool       `protobuf:"bytes,4,opt,name=isSocial,proto3" json:"isSocial,omitempty"`
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
func (m *Identity) GetIsSocial() *common.Bool
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


#### func  NewJWT

```go
func NewJWT(tokenURL, email, privateKey, privKeyID, subject string, expiry, audience string, scopes []string) *JWT
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
func (c *JWT) NewAPIClientSet(ctx context.Context, addr string) (*ClientSet, error)
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
func (c *OAuth2) NewAPIClientSet(ctx context.Context, addr string) (*ClientSet, error)
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

#### type Plan

```go
type Plan struct {
	Id                   *common.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Active               *common.Bool       `protobuf:"bytes,2,opt,name=active,proto3" json:"active,omitempty"`
	Amount               *common.Int64      `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Interval             *common.String     `protobuf:"bytes,4,opt,name=interval,proto3" json:"interval,omitempty"`
	Nickname             *common.String     `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Product              *Product           `protobuf:"bytes,6,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func (*Plan) Descriptor

```go
func (*Plan) Descriptor() ([]byte, []int)
```

#### func (*Plan) GetActive

```go
func (m *Plan) GetActive() *common.Bool
```

#### func (*Plan) GetAmount

```go
func (m *Plan) GetAmount() *common.Int64
```

#### func (*Plan) GetId

```go
func (m *Plan) GetId() *common.Identifier
```

#### func (*Plan) GetInterval

```go
func (m *Plan) GetInterval() *common.String
```

#### func (*Plan) GetNickname

```go
func (m *Plan) GetNickname() *common.String
```

#### func (*Plan) GetProduct

```go
func (m *Plan) GetProduct() *Product
```

#### func (*Plan) JSONString

```go
func (p *Plan) JSONString() *common.String
```

#### func (*Plan) ProtoMessage

```go
func (*Plan) ProtoMessage()
```

#### func (*Plan) Reset

```go
func (m *Plan) Reset()
```

#### func (*Plan) String

```go
func (m *Plan) String() string
```

#### func (*Plan) UnmarshalJSONFrom

```go
func (p *Plan) UnmarshalJSONFrom(bits []byte) error
```

#### func (*Plan) XXX_DiscardUnknown

```go
func (m *Plan) XXX_DiscardUnknown()
```

#### func (*Plan) XXX_Marshal

```go
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Plan) XXX_Merge

```go
func (m *Plan) XXX_Merge(src proto.Message)
```

#### func (*Plan) XXX_Size

```go
func (m *Plan) XXX_Size() int
```

#### func (*Plan) XXX_Unmarshal

```go
func (m *Plan) XXX_Unmarshal(b []byte) error
```

#### type Product

```go
type Product struct {
	Id                   *common.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          *common.String     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Url                  *common.String     `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func (*Product) Descriptor

```go
func (*Product) Descriptor() ([]byte, []int)
```

#### func (*Product) GetDescription

```go
func (m *Product) GetDescription() *common.String
```

#### func (*Product) GetId

```go
func (m *Product) GetId() *common.Identifier
```

#### func (*Product) GetUrl

```go
func (m *Product) GetUrl() *common.String
```

#### func (*Product) JSONString

```go
func (p *Product) JSONString() *common.String
```

#### func (*Product) ProtoMessage

```go
func (*Product) ProtoMessage()
```

#### func (*Product) Reset

```go
func (m *Product) Reset()
```

#### func (*Product) String

```go
func (m *Product) String() string
```

#### func (*Product) UnmarshalJSONFrom

```go
func (p *Product) UnmarshalJSONFrom(bits []byte) error
```

#### func (*Product) XXX_DiscardUnknown

```go
func (m *Product) XXX_DiscardUnknown()
```

#### func (*Product) XXX_Marshal

```go
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Product) XXX_Merge

```go
func (m *Product) XXX_Merge(src proto.Message)
```

#### func (*Product) XXX_Size

```go
func (m *Product) XXX_Size() int
```

#### func (*Product) XXX_Unmarshal

```go
func (m *Product) XXX_Unmarshal(b []byte) error
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

#### func (*Role) JSONString

```go
func (p *Role) JSONString() *common.String
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

#### func (*Role) UnmarshalJSONFrom

```go
func (p *Role) UnmarshalJSONFrom(bits []byte) error
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

#### type SubscribeRequest

```go
type SubscribeRequest struct {
	Email                *common.Identifier `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Plan                 *common.Identifier `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	Card                 *Card              `protobuf:"bytes,3,opt,name=card,proto3" json:"card,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
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
func (m *SubscribeRequest) GetEmail() *common.Identifier
```

#### func (*SubscribeRequest) GetPlan

```go
func (m *SubscribeRequest) GetPlan() *common.Identifier
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

#### type SubscriptionServiceClient

```go
type SubscriptionServiceClient interface {
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*User, error)
	Unsubscribe(ctx context.Context, in *UnSubscribeRequest, opts ...grpc.CallOption) (*User, error)
}
```

SubscriptionServiceClient is the client API for SubscriptionService service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to
https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

#### func  NewSubscriptionServiceClient

```go
func NewSubscriptionServiceClient(cc *grpc.ClientConn) SubscriptionServiceClient
```

#### type SubscriptionServiceServer

```go
type SubscriptionServiceServer interface {
	Subscribe(context.Context, *SubscribeRequest) (*User, error)
	Unsubscribe(context.Context, *UnSubscribeRequest) (*User, error)
}
```

SubscriptionServiceServer is the server API for SubscriptionService service.

#### type UnSubscribeRequest

```go
type UnSubscribeRequest struct {
	Email                *common.Identifier `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Plan                 *common.String     `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func (*UnSubscribeRequest) Descriptor

```go
func (*UnSubscribeRequest) Descriptor() ([]byte, []int)
```

#### func (*UnSubscribeRequest) GetEmail

```go
func (m *UnSubscribeRequest) GetEmail() *common.Identifier
```

#### func (*UnSubscribeRequest) GetPlan

```go
func (m *UnSubscribeRequest) GetPlan() *common.String
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
	Nickname             *common.String      `protobuf:"bytes,14,opt,name=nickname,proto3" json:"nickname,omitempty"`
	UserMetadata         *common.StringMap   `protobuf:"bytes,10,opt,name=user_metadata,json=userMetadata,proto3" json:"user_metadata,omitempty"`
	AppMetadata          *common.StringMap   `protobuf:"bytes,11,opt,name=app_metadata,json=appMetadata,proto3" json:"app_metadata,omitempty"`
	LastIp               *common.String      `protobuf:"bytes,12,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
	Blocked              *common.Bool        `protobuf:"bytes,13,opt,name=blocked,proto3" json:"blocked,omitempty"`
	Multifactor          *common.StringArray `protobuf:"bytes,15,opt,name=multifactor,proto3" json:"multifactor,omitempty"`
	CreatedAt            *common.String      `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *common.String      `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	PhoneVerified        *common.Bool        `protobuf:"bytes,19,opt,name=phone_verified,json=phoneVerified,proto3" json:"phone_verified,omitempty"`
	EmailVerified        *common.Bool        `protobuf:"bytes,20,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
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
func (m *User) GetBlocked() *common.Bool
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
func (m *User) GetEmailVerified() *common.Bool
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
func (m *User) GetPhoneVerified() *common.Bool
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

#### type UserMetadata

```go
type UserMetadata struct {
	Status               *common.String    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Bio                  *common.StringMap `protobuf:"bytes,2,opt,name=bio,proto3" json:"bio,omitempty"`
	Address              *Address          `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Tags                 *common.StringMap `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func (*UserMetadata) Descriptor

```go
func (*UserMetadata) Descriptor() ([]byte, []int)
```

#### func (*UserMetadata) GetAddress

```go
func (m *UserMetadata) GetAddress() *Address
```

#### func (*UserMetadata) GetBio

```go
func (m *UserMetadata) GetBio() *common.StringMap
```

#### func (*UserMetadata) GetStatus

```go
func (m *UserMetadata) GetStatus() *common.String
```

#### func (*UserMetadata) GetTags

```go
func (m *UserMetadata) GetTags() *common.StringMap
```

#### func (*UserMetadata) JSONString

```go
func (p *UserMetadata) JSONString() *common.String
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

#### func (*UserMetadata) UnmarshalJSONFrom

```go
func (p *UserMetadata) UnmarshalJSONFrom(bits []byte) error
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
	QueryUsers(ctx context.Context, in *Query, opts ...grpc.CallOption) (UserService_QueryUsersClient, error)
	GetUser(ctx context.Context, in *common.Identifier, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UserRoles(ctx context.Context, in *common.Identifier, opts ...grpc.CallOption) (UserService_UserRolesClient, error)
	AddUserRoles(ctx context.Context, in *AddUserRolesRequest, opts ...grpc.CallOption) (*common.String, error)
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
	QueryUsers(*Query, UserService_QueryUsersServer) error
	GetUser(context.Context, *common.Identifier) (*User, error)
	UpdateUser(context.Context, *User) (*User, error)
	UserRoles(*common.Identifier, UserService_UserRolesServer) error
	AddUserRoles(context.Context, *AddUserRolesRequest) (*common.String, error)
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
