# api
--
    import "github.com/autom8ter/api"

Package api is a reverse proxy.

It translates gRPC into RESTful JSON APIs.

## Usage

#### func  RegisterDBServiceHandler

```go
func RegisterDBServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
```
RegisterDBServiceHandler registers the http handlers for service DBService to
"mux". The handlers forward requests to the grpc endpoint over "conn".

#### func  RegisterDBServiceHandlerClient

```go
func RegisterDBServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client DBServiceClient) error
```
RegisterDBServiceHandlerClient registers the http handlers for service DBService
to "mux". The handlers forward requests to the grpc endpoint over the given
implementation of "DBServiceClient". Note: the gRPC framework executes
interceptors within the gRPC handler. If the passed in "DBServiceClient" doesn't
go through the normal gRPC flow (creating a gRPC client etc.) then it will be up
to the passed in "DBServiceClient" to call the correct interceptors.

#### func  RegisterDBServiceHandlerFromEndpoint

```go
func RegisterDBServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
```
RegisterDBServiceHandlerFromEndpoint is same as RegisterDBServiceHandler but
automatically dials to "endpoint" and closes the connection when "ctx" gets
done.

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

#### type Address

```go
type Address struct {
	City                 string   `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Line1                string   `protobuf:"bytes,3,opt,name=line1,proto3" json:"line1,omitempty"`
	Line2                string   `protobuf:"bytes,4,opt,name=line2,proto3" json:"line2,omitempty"`
	Zip                  string   `protobuf:"bytes,5,opt,name=zip,proto3" json:"zip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Address) Descriptor

```go
func (*Address) Descriptor() ([]byte, []int)
```

#### func (*Address) GetCity

```go
func (m *Address) GetCity() string
```

#### func (*Address) GetLine1

```go
func (m *Address) GetLine1() string
```

#### func (*Address) GetLine2

```go
func (m *Address) GetLine2() string
```

#### func (*Address) GetState

```go
func (m *Address) GetState() string
```

#### func (*Address) GetZip

```go
func (m *Address) GetZip() string
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
	Description          string            `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	AccountBalance       string            `protobuf:"bytes,2,opt,name=account_balance,json=accountBalance,proto3" json:"account_balance,omitempty"`
	Plan                 *Plan             `protobuf:"bytes,3,opt,name=plan,proto3" json:"plan,omitempty"`
	Tags                 map[string]string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Card                 *Card             `protobuf:"bytes,5,opt,name=card,proto3" json:"card,omitempty"`
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
func (m *AppMetadata) GetAccountBalance() string
```

#### func (*AppMetadata) GetCard

```go
func (m *AppMetadata) GetCard() *Card
```

#### func (*AppMetadata) GetDescription

```go
func (m *AppMetadata) GetDescription() string
```

#### func (*AppMetadata) GetPlan

```go
func (m *AppMetadata) GetPlan() *Plan
```

#### func (*AppMetadata) GetTags

```go
func (m *AppMetadata) GetTags() map[string]string
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

#### type Identity

```go
type Identity struct {
	Connection           string   `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Provider             string   `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
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

#### type Plan

```go
type Plan struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Active               bool     `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	Amount               int64    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Interval             string   `protobuf:"bytes,4,opt,name=interval,proto3" json:"interval,omitempty"`
	Nickname             string   `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Product              *Product `protobuf:"bytes,6,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Plan) Descriptor

```go
func (*Plan) Descriptor() ([]byte, []int)
```

#### func (*Plan) GetActive

```go
func (m *Plan) GetActive() bool
```

#### func (*Plan) GetAmount

```go
func (m *Plan) GetAmount() int64
```

#### func (*Plan) GetId

```go
func (m *Plan) GetId() string
```

#### func (*Plan) GetInterval

```go
func (m *Plan) GetInterval() string
```

#### func (*Plan) GetNickname

```go
func (m *Plan) GetNickname() string
```

#### func (*Plan) GetProduct

```go
func (m *Plan) GetProduct() *Product
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
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Url                  string            `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Tags                 map[string]string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func (*Product) Descriptor

```go
func (*Product) Descriptor() ([]byte, []int)
```

#### func (*Product) GetDescription

```go
func (m *Product) GetDescription() string
```

#### func (*Product) GetId

```go
func (m *Product) GetId() string
```

#### func (*Product) GetTags

```go
func (m *Product) GetTags() map[string]string
```

#### func (*Product) GetUrl

```go
func (m *Product) GetUrl() string
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

#### type Role

```go
type Role struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Role) Descriptor

```go
func (*Role) Descriptor() ([]byte, []int)
```

#### func (*Role) GetDescription

```go
func (m *Role) GetDescription() string
```

#### func (*Role) GetId

```go
func (m *Role) GetId() string
```

#### func (*Role) GetName

```go
func (m *Role) GetName() string
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

#### type User

```go
type User struct {
	UserId               string            `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GivenName            string            `protobuf:"bytes,3,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	FamilyName           string            `protobuf:"bytes,4,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	Gender               string            `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthdate            string            `protobuf:"bytes,6,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	Email                string            `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber          string            `protobuf:"bytes,8,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Picture              string            `protobuf:"bytes,9,opt,name=picture,proto3" json:"picture,omitempty"`
	Nickname             string            `protobuf:"bytes,14,opt,name=nickname,proto3" json:"nickname,omitempty"`
	UserMetadata         map[string]string `protobuf:"bytes,10,rep,name=user_metadata,json=userMetadata,proto3" json:"user_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AppMetadata          map[string]string `protobuf:"bytes,11,rep,name=app_metadata,json=appMetadata,proto3" json:"app_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LastIp               string            `protobuf:"bytes,12,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
	Blocked              bool              `protobuf:"varint,13,opt,name=blocked,proto3" json:"blocked,omitempty"`
	Multifactor          []string          `protobuf:"bytes,15,rep,name=multifactor,proto3" json:"multifactor,omitempty"`
	CreatedAt            string            `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string            `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	PhoneVerified        bool              `protobuf:"varint,19,opt,name=phone_verified,json=phoneVerified,proto3" json:"phone_verified,omitempty"`
	EmailVerified        bool              `protobuf:"varint,20,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	Password             string            `protobuf:"bytes,21,opt,name=password,proto3" json:"password,omitempty"`
	Identities           []*Identity       `protobuf:"bytes,22,rep,name=identities,proto3" json:"identities,omitempty"`
	Roles                []*Role           `protobuf:"bytes,23,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func (*User) Descriptor

```go
func (*User) Descriptor() ([]byte, []int)
```

#### func (*User) GetAppMetadata

```go
func (m *User) GetAppMetadata() map[string]string
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

#### func (*User) GetEmailVerified

```go
func (m *User) GetEmailVerified() bool
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

#### func (*User) GetPassword

```go
func (m *User) GetPassword() string
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

#### func (*User) GetRoles

```go
func (m *User) GetRoles() []*Role
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
func (m *User) GetUserMetadata() map[string]string
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

#### type UserMetadata

```go
type UserMetadata struct {
	Status               string            `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Bio                  map[string]string `protobuf:"bytes,2,rep,name=bio,proto3" json:"bio,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Address              *Address          `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Tags                 map[string]string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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
func (m *UserMetadata) GetBio() map[string]string
```

#### func (*UserMetadata) GetStatus

```go
func (m *UserMetadata) GetStatus() string
```

#### func (*UserMetadata) GetTags

```go
func (m *UserMetadata) GetTags() map[string]string
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
