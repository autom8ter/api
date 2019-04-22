# api
--
    import "github.com/autom8ter/api/go/api"

Package api is a reverse proxy.

It translates gRPC into RESTful JSON APIs.

## Usage

```go
var Util = objectify.Default()
```

#### func  RegisterProfileServiceHandler

```go
func RegisterProfileServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
```
RegisterProfileServiceHandler registers the http handlers for service
ProfileService to "mux". The handlers forward requests to the grpc endpoint over
"conn".

#### func  RegisterProfileServiceHandlerClient

```go
func RegisterProfileServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ProfileServiceClient) error
```
RegisterProfileServiceHandlerClient registers the http handlers for service
ProfileService to "mux". The handlers forward requests to the grpc endpoint over
the given implementation of "ProfileServiceClient". Note: the gRPC framework
executes interceptors within the gRPC handler. If the passed in
"ProfileServiceClient" doesn't go through the normal gRPC flow (creating a gRPC
client etc.) then it will be up to the passed in "ProfileServiceClient" to call
the correct interceptors.

#### func  RegisterProfileServiceHandlerFromEndpoint

```go
func RegisterProfileServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
```
RegisterProfileServiceHandlerFromEndpoint is same as
RegisterProfileServiceHandler but automatically dials to "endpoint" and closes
the connection when "ctx" gets done.

#### func  RegisterProfileServiceServer

```go
func RegisterProfileServiceServer(s *grpc.Server, srv ProfileServiceServer)
```

#### type GetProfileByEmail

```go
type GetProfileByEmail struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*GetProfileByEmail) Descriptor

```go
func (*GetProfileByEmail) Descriptor() ([]byte, []int)
```

#### func (*GetProfileByEmail) GetEmail

```go
func (m *GetProfileByEmail) GetEmail() string
```

#### func (*GetProfileByEmail) ProtoMessage

```go
func (*GetProfileByEmail) ProtoMessage()
```

#### func (*GetProfileByEmail) Reset

```go
func (m *GetProfileByEmail) Reset()
```

#### func (*GetProfileByEmail) String

```go
func (m *GetProfileByEmail) String() string
```

#### func (*GetProfileByEmail) XXX_DiscardUnknown

```go
func (m *GetProfileByEmail) XXX_DiscardUnknown()
```

#### func (*GetProfileByEmail) XXX_Marshal

```go
func (m *GetProfileByEmail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*GetProfileByEmail) XXX_Merge

```go
func (m *GetProfileByEmail) XXX_Merge(src proto.Message)
```

#### func (*GetProfileByEmail) XXX_Size

```go
func (m *GetProfileByEmail) XXX_Size() int
```

#### func (*GetProfileByEmail) XXX_Unmarshal

```go
func (m *GetProfileByEmail) XXX_Unmarshal(b []byte) error
```

#### type Identity

```go
type Identity struct {
	Provider             string   `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Connection           string   `protobuf:"bytes,3,opt,name=connection,proto3" json:"connection,omitempty"`
	IsSocial             bool     `protobuf:"varint,4,opt,name=isSocial,proto3" json:"isSocial,omitempty"`
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
func (m *Identity) GetIsSocial() bool
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

#### type Profile

```go
type Profile struct {
	Email                string        `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	EmailVerified        bool          `protobuf:"varint,2,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	Name                 string        `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	GivenName            string        `protobuf:"bytes,4,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	FamilyName           string        `protobuf:"bytes,5,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	Picture              string        `protobuf:"bytes,6,opt,name=picture,proto3" json:"picture,omitempty"`
	Locale               string        `protobuf:"bytes,7,opt,name=locale,proto3" json:"locale,omitempty"`
	UserId               string        `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Nickname             string        `protobuf:"bytes,9,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Connection           string        `protobuf:"bytes,10,opt,name=connection,proto3" json:"connection,omitempty"`
	Identities           []*Identity   `protobuf:"bytes,11,rep,name=identities,proto3" json:"identities,omitempty"`
	LastIp               string        `protobuf:"bytes,12,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
	LoginCount           int64         `protobuf:"varint,13,opt,name=login_count,json=loginCount,proto3" json:"login_count,omitempty"`
	UpdatedAt            string        `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt            string        `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UserMetadata         *UserMetadata `protobuf:"bytes,20,opt,name=user_metadata,json=userMetadata,proto3" json:"user_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
```


#### func (*Profile) Descriptor

```go
func (*Profile) Descriptor() ([]byte, []int)
```

#### func (*Profile) GetConnection

```go
func (m *Profile) GetConnection() string
```

#### func (*Profile) GetCreatedAt

```go
func (m *Profile) GetCreatedAt() string
```

#### func (*Profile) GetEmail

```go
func (m *Profile) GetEmail() string
```

#### func (*Profile) GetEmailVerified

```go
func (m *Profile) GetEmailVerified() bool
```

#### func (*Profile) GetFamilyName

```go
func (m *Profile) GetFamilyName() string
```

#### func (*Profile) GetGivenName

```go
func (m *Profile) GetGivenName() string
```

#### func (*Profile) GetIdentities

```go
func (m *Profile) GetIdentities() []*Identity
```

#### func (*Profile) GetLastIp

```go
func (m *Profile) GetLastIp() string
```

#### func (*Profile) GetLocale

```go
func (m *Profile) GetLocale() string
```

#### func (*Profile) GetLoginCount

```go
func (m *Profile) GetLoginCount() int64
```

#### func (*Profile) GetName

```go
func (m *Profile) GetName() string
```

#### func (*Profile) GetNickname

```go
func (m *Profile) GetNickname() string
```

#### func (*Profile) GetPicture

```go
func (m *Profile) GetPicture() string
```

#### func (*Profile) GetUpdatedAt

```go
func (m *Profile) GetUpdatedAt() string
```

#### func (*Profile) GetUserId

```go
func (m *Profile) GetUserId() string
```

#### func (*Profile) GetUserMetadata

```go
func (m *Profile) GetUserMetadata() *UserMetadata
```

#### func (*Profile) ProtoMessage

```go
func (*Profile) ProtoMessage()
```

#### func (*Profile) Reset

```go
func (m *Profile) Reset()
```

#### func (*Profile) String

```go
func (m *Profile) String() string
```

#### func (*Profile) XXX_DiscardUnknown

```go
func (m *Profile) XXX_DiscardUnknown()
```

#### func (*Profile) XXX_Marshal

```go
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Profile) XXX_Merge

```go
func (m *Profile) XXX_Merge(src proto.Message)
```

#### func (*Profile) XXX_Size

```go
func (m *Profile) XXX_Size() int
```

#### func (*Profile) XXX_Unmarshal

```go
func (m *Profile) XXX_Unmarshal(b []byte) error
```

#### type ProfileServiceClient

```go
type ProfileServiceClient interface {
	GetProfile(ctx context.Context, in *GetProfileByEmail, opts ...grpc.CallOption) (*Profile, error)
}
```

ProfileServiceClient is the client API for ProfileService service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to
https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

#### func  NewProfileServiceClient

```go
func NewProfileServiceClient(cc *grpc.ClientConn) ProfileServiceClient
```

#### type ProfileServiceServer

```go
type ProfileServiceServer interface {
	GetProfile(context.Context, *GetProfileByEmail) (*Profile, error)
}
```

ProfileServiceServer is the server API for ProfileService service.

#### type UserMetadata

```go
type UserMetadata struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Plan                 string   `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	PayToken             string   `protobuf:"bytes,3,opt,name=pay_token,json=payToken,proto3" json:"pay_token,omitempty"`
	LastContact          string   `protobuf:"bytes,4,opt,name=last_contact,json=lastContact,proto3" json:"last_contact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*UserMetadata) Descriptor

```go
func (*UserMetadata) Descriptor() ([]byte, []int)
```

#### func (*UserMetadata) GetLastContact

```go
func (m *UserMetadata) GetLastContact() string
```

#### func (*UserMetadata) GetPayToken

```go
func (m *UserMetadata) GetPayToken() string
```

#### func (*UserMetadata) GetPhone

```go
func (m *UserMetadata) GetPhone() string
```

#### func (*UserMetadata) GetPlan

```go
func (m *UserMetadata) GetPlan() string
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
