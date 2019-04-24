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
var DEFAULT_REDIRECT = "http://localhost:8080/callback"
```

```go
var DEFAULT_SCOPES = []string{"openid", "profile", "email"}
```

```go
var SESSIONS_PATH = os.Getenv("PWD") + "/sessions"
```

#### func  FuncMap

```go
func FuncMap() template.FuncMap
```

#### func  InitSessions

```go
func InitSessions(secret string) error
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

#### type AppMetadata

```go
type AppMetadata struct {
	Plan                 string   `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	PayToken             string   `protobuf:"bytes,2,opt,name=pay_token,json=payToken,proto3" json:"pay_token,omitempty"`
	Delinquent           string   `protobuf:"bytes,3,opt,name=delinquent,proto3" json:"delinquent,omitempty"`
	Tags                 []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*AppMetadata) Descriptor

```go
func (*AppMetadata) Descriptor() ([]byte, []int)
```

#### func (*AppMetadata) GetDelinquent

```go
func (m *AppMetadata) GetDelinquent() string
```

#### func (*AppMetadata) GetPayToken

```go
func (m *AppMetadata) GetPayToken() string
```

#### func (*AppMetadata) GetPlan

```go
func (m *AppMetadata) GetPlan() string
```

#### func (*AppMetadata) GetTags

```go
func (m *AppMetadata) GetTags() []string
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


#### func (*Auth) Descriptor

```go
func (*Auth) Descriptor() ([]byte, []int)
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

#### func (*Auth) GetRedirect

```go
func (m *Auth) GetRedirect() string
```

#### func (*Auth) GetScopes

```go
func (m *Auth) GetScopes() []string
```

#### func (*Auth) ListenAndServe

```go
func (a *Auth) ListenAndServe(addr string, c *RouterPaths, home, loggedIn http.HandlerFunc) error
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

#### func (*Auth) Router

```go
func (a *Auth) Router(c *RouterPaths, home, loggedIn http.HandlerFunc) *mux.Router
```

#### func (*Auth) String

```go
func (m *Auth) String() string
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
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Callback             string   `protobuf:"bytes,2,opt,name=callback,proto3" json:"callback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Call) Descriptor

```go
func (*Call) Descriptor() ([]byte, []int)
```

#### func (*Call) GetCallback

```go
func (m *Call) GetCallback() string
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

#### type ClientSet

```go
type ClientSet struct {
	Utility UtilityServiceClient
	Contact ContactServiceClient
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
	SendEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Identifier, error)
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
	SendEmail(context.Context, *Email) (*Identifier, error)
	SendCall(context.Context, *Call) (*Identifier, error)
}
```

ContactServiceServer is the server API for ContactService service.

#### type Email

```go
type Email struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Subject              string   `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Message              *Message `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
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

#### func (*Email) GetMessage

```go
func (m *Email) GetMessage() *Message
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
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Message              *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	MediaURL             string   `protobuf:"bytes,3,opt,name=mediaURL,proto3" json:"mediaURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*SMS) Descriptor

```go
func (*SMS) Descriptor() ([]byte, []int)
```

#### func (*SMS) GetMediaURL

```go
func (m *SMS) GetMediaURL() string
```

#### func (*SMS) GetMessage

```go
func (m *SMS) GetMessage() *Message
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
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	PreferredContact     string   `protobuf:"bytes,2,opt,name=preferred_contact,json=preferredContact,proto3" json:"preferred_contact,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Tags                 []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
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

#### func (*UserMetadata) GetStatus

```go
func (m *UserMetadata) GetStatus() string
```

#### func (*UserMetadata) GetTags

```go
func (m *UserMetadata) GetTags() []string
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
