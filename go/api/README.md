# api
--
    import "github.com/autom8ter/api/go/api"

Package api is a reverse proxy.

It translates gRPC into RESTful JSON APIs.

## Usage

```go
var CardType_name = map[int32]string{
	0: "VISA",
	1: "MASTERCARD",
	2: "DISCOVER",
	3: "AMEX",
}
```

```go
var CardType_value = map[string]int32{
	"VISA":       0,
	"MASTERCARD": 1,
	"DISCOVER":   2,
	"AMEX":       3,
}
```

```go
var CustomerIndex_name = map[int32]string{
	0: "ID",
	1: "EMAIL",
	2: "PHONE",
}
```

```go
var CustomerIndex_value = map[string]int32{
	"ID":    0,
	"EMAIL": 1,
	"PHONE": 2,
}
```

```go
var Grant_name = map[int32]string{
	0: "TWILIO",
	1: "SENDGRID",
	2: "STRIPE",
	3: "SLACK",
	4: "GCP",
	5: "AUTOM8TER",
}
```

```go
var Grant_value = map[string]int32{
	"TWILIO":    0,
	"SENDGRID":  1,
	"STRIPE":    2,
	"SLACK":     3,
	"GCP":       4,
	"AUTOM8TER": 5,
}
```

```go
var SigningMethod_name = map[int32]string{
	0: "HMAC",
	1: "ECDSA",
	2: "RSA",
	3: "RSAPPS",
}
```

```go
var SigningMethod_value = map[string]int32{
	"HMAC":   0,
	"ECDSA":  1,
	"RSA":    2,
	"RSAPPS": 3,
}
```

```go
var Util = objectify.Default()
```

#### func  Cmd

```go
func Cmd(name, description string, fn func() error) *cobra.Command
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

#### type Access

```go
type Access struct {
	Autom8TerAccount     string   `protobuf:"bytes,1,opt,name=autom8ter_account,json=autom8terAccount,proto3" json:"autom8ter_account,omitempty"`
	Autom8TerKey         string   `protobuf:"bytes,2,opt,name=autom8ter_key,json=autom8terKey,proto3" json:"autom8ter_key,omitempty"`
	TwilioAccount        string   `protobuf:"bytes,3,opt,name=twilio_account,json=twilioAccount,proto3" json:"twilio_account,omitempty"`
	TwilioKey            string   `protobuf:"bytes,4,opt,name=twilio_key,json=twilioKey,proto3" json:"twilio_key,omitempty"`
	SendgridAccount      string   `protobuf:"bytes,5,opt,name=sendgrid_account,json=sendgridAccount,proto3" json:"sendgrid_account,omitempty"`
	SendgridKey          string   `protobuf:"bytes,6,opt,name=sendgrid_key,json=sendgridKey,proto3" json:"sendgrid_key,omitempty"`
	StripeAccount        string   `protobuf:"bytes,7,opt,name=stripe_account,json=stripeAccount,proto3" json:"stripe_account,omitempty"`
	StripeKey            string   `protobuf:"bytes,8,opt,name=stripe_key,json=stripeKey,proto3" json:"stripe_key,omitempty"`
	SlackAccount         string   `protobuf:"bytes,9,opt,name=slack_account,json=slackAccount,proto3" json:"slack_account,omitempty"`
	SlackKey             string   `protobuf:"bytes,10,opt,name=slack_key,json=slackKey,proto3" json:"slack_key,omitempty"`
	GcpProject           string   `protobuf:"bytes,11,opt,name=gcp_project,json=gcpProject,proto3" json:"gcp_project,omitempty"`
	GcpKey               string   `protobuf:"bytes,12,opt,name=gcp_key,json=gcpKey,proto3" json:"gcp_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  AccessFromEnv

```go
func AccessFromEnv() *Access
```
TWILIO_ACCOUNT TWILIO_KEY SENDGRID_ACCOUNT SENDGRID_KEY STRIPE_ACCOUNT
STRIPE_KEY SLACK_ACCOUNT SLACK_KEY GCP_PROJECT GCP_KEY

#### func  AccessFromJSON

```go
func AccessFromJSON(j *JSON) *Access
```

#### func (*Access) Descriptor

```go
func (*Access) Descriptor() ([]byte, []int)
```

#### func (*Access) GetAutom8TerAccount

```go
func (m *Access) GetAutom8TerAccount() string
```

#### func (*Access) GetAutom8TerKey

```go
func (m *Access) GetAutom8TerKey() string
```

#### func (*Access) GetGcpKey

```go
func (m *Access) GetGcpKey() string
```

#### func (*Access) GetGcpProject

```go
func (m *Access) GetGcpProject() string
```

#### func (*Access) GetSendgridAccount

```go
func (m *Access) GetSendgridAccount() string
```

#### func (*Access) GetSendgridKey

```go
func (m *Access) GetSendgridKey() string
```

#### func (*Access) GetSlackAccount

```go
func (m *Access) GetSlackAccount() string
```

#### func (*Access) GetSlackKey

```go
func (m *Access) GetSlackKey() string
```

#### func (*Access) GetStripeAccount

```go
func (m *Access) GetStripeAccount() string
```

#### func (*Access) GetStripeKey

```go
func (m *Access) GetStripeKey() string
```

#### func (*Access) GetTwilioAccount

```go
func (m *Access) GetTwilioAccount() string
```

#### func (*Access) GetTwilioKey

```go
func (m *Access) GetTwilioKey() string
```

#### func (*Access) ProtoMessage

```go
func (*Access) ProtoMessage()
```

#### func (*Access) Reset

```go
func (m *Access) Reset()
```

#### func (*Access) String

```go
func (m *Access) String() string
```

#### func (*Access) XXX_DiscardUnknown

```go
func (m *Access) XXX_DiscardUnknown()
```

#### func (*Access) XXX_Marshal

```go
func (m *Access) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Access) XXX_Merge

```go
func (m *Access) XXX_Merge(src proto.Message)
```

#### func (*Access) XXX_Size

```go
func (m *Access) XXX_Size() int
```

#### func (*Access) XXX_Unmarshal

```go
func (m *Access) XXX_Unmarshal(b []byte) error
```

#### type AddCustomerRequest

```go
type AddCustomerRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Plan                 string   `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Address              *Address `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*AddCustomerRequest) Descriptor

```go
func (*AddCustomerRequest) Descriptor() ([]byte, []int)
```

#### func (*AddCustomerRequest) GetAddress

```go
func (m *AddCustomerRequest) GetAddress() *Address
```

#### func (*AddCustomerRequest) GetDescription

```go
func (m *AddCustomerRequest) GetDescription() string
```

#### func (*AddCustomerRequest) GetEmail

```go
func (m *AddCustomerRequest) GetEmail() string
```

#### func (*AddCustomerRequest) GetName

```go
func (m *AddCustomerRequest) GetName() string
```

#### func (*AddCustomerRequest) GetPhone

```go
func (m *AddCustomerRequest) GetPhone() string
```

#### func (*AddCustomerRequest) GetPlan

```go
func (m *AddCustomerRequest) GetPlan() string
```

#### func (*AddCustomerRequest) ProtoMessage

```go
func (*AddCustomerRequest) ProtoMessage()
```

#### func (*AddCustomerRequest) Reset

```go
func (m *AddCustomerRequest) Reset()
```

#### func (*AddCustomerRequest) String

```go
func (m *AddCustomerRequest) String() string
```

#### func (*AddCustomerRequest) XXX_DiscardUnknown

```go
func (m *AddCustomerRequest) XXX_DiscardUnknown()
```

#### func (*AddCustomerRequest) XXX_Marshal

```go
func (m *AddCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*AddCustomerRequest) XXX_Merge

```go
func (m *AddCustomerRequest) XXX_Merge(src proto.Message)
```

#### func (*AddCustomerRequest) XXX_Size

```go
func (m *AddCustomerRequest) XXX_Size() int
```

#### func (*AddCustomerRequest) XXX_Unmarshal

```go
func (m *AddCustomerRequest) XXX_Unmarshal(b []byte) error
```

#### type Address

```go
type Address struct {
	City                 string   `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	Country              string   `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Line1                string   `protobuf:"bytes,3,opt,name=line1,proto3" json:"line1,omitempty"`
	Line2                string   `protobuf:"bytes,4,opt,name=line2,proto3" json:"line2,omitempty"`
	PostalCode           string   `protobuf:"bytes,5,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	State                string   `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
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

#### func (*Address) GetCountry

```go
func (m *Address) GetCountry() string
```

#### func (*Address) GetLine1

```go
func (m *Address) GetLine1() string
```

#### func (*Address) GetLine2

```go
func (m *Address) GetLine2() string
```

#### func (*Address) GetPostalCode

```go
func (m *Address) GetPostalCode() string
```

#### func (*Address) GetState

```go
func (m *Address) GetState() string
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

#### type Attachment

```go
type Attachment struct {
	Color                string             `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	Fallback             string             `protobuf:"bytes,2,opt,name=fallback,proto3" json:"fallback,omitempty"`
	CallbackId           string             `protobuf:"bytes,3,opt,name=callback_id,json=callbackId,proto3" json:"callback_id,omitempty"`
	Id                   int64              `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId             string             `protobuf:"bytes,5,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	AuthorName           string             `protobuf:"bytes,6,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
	AuthorLink           string             `protobuf:"bytes,7,opt,name=author_link,json=authorLink,proto3" json:"author_link,omitempty"`
	AuthorIcon           string             `protobuf:"bytes,8,opt,name=author_icon,json=authorIcon,proto3" json:"author_icon,omitempty"`
	Title                string             `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`
	TitlePrefix          string             `protobuf:"bytes,10,opt,name=title_prefix,json=titlePrefix,proto3" json:"title_prefix,omitempty"`
	Pretext              string             `protobuf:"bytes,11,opt,name=pretext,proto3" json:"pretext,omitempty"`
	Text                 string             `protobuf:"bytes,12,opt,name=text,proto3" json:"text,omitempty"`
	ImageUrl             string             `protobuf:"bytes,13,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	ThumbUrl             string             `protobuf:"bytes,14,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url,omitempty"`
	Fields               []*AttachmentField `protobuf:"bytes,15,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func  NewAttachment

```go
func NewAttachment(opts ...AttachmentFunc) *Attachment
```

#### func (*Attachment) AsMap

```go
func (j *Attachment) AsMap() map[string]interface{}
```

#### func (*Attachment) CompileHTML

```go
func (j *Attachment) CompileHTML(text string, w io.Writer) error
```

#### func (*Attachment) CompileTXT

```go
func (j *Attachment) CompileTXT(text string, w io.Writer) error
```

#### func (*Attachment) Descriptor

```go
func (*Attachment) Descriptor() ([]byte, []int)
```

#### func (*Attachment) GetAuthorIcon

```go
func (m *Attachment) GetAuthorIcon() string
```

#### func (*Attachment) GetAuthorId

```go
func (m *Attachment) GetAuthorId() string
```

#### func (*Attachment) GetAuthorLink

```go
func (m *Attachment) GetAuthorLink() string
```

#### func (*Attachment) GetAuthorName

```go
func (m *Attachment) GetAuthorName() string
```

#### func (*Attachment) GetCallbackId

```go
func (m *Attachment) GetCallbackId() string
```

#### func (*Attachment) GetColor

```go
func (m *Attachment) GetColor() string
```

#### func (*Attachment) GetFallback

```go
func (m *Attachment) GetFallback() string
```

#### func (*Attachment) GetFields

```go
func (m *Attachment) GetFields() []*AttachmentField
```

#### func (*Attachment) GetId

```go
func (m *Attachment) GetId() int64
```

#### func (*Attachment) GetImageUrl

```go
func (m *Attachment) GetImageUrl() string
```

#### func (*Attachment) GetPretext

```go
func (m *Attachment) GetPretext() string
```

#### func (*Attachment) GetText

```go
func (m *Attachment) GetText() string
```

#### func (*Attachment) GetThumbUrl

```go
func (m *Attachment) GetThumbUrl() string
```

#### func (*Attachment) GetTitle

```go
func (m *Attachment) GetTitle() string
```

#### func (*Attachment) GetTitlePrefix

```go
func (m *Attachment) GetTitlePrefix() string
```

#### func (*Attachment) MarshalJSON

```go
func (j *Attachment) MarshalJSON() ([]byte, error)
```

#### func (*Attachment) ProtoMessage

```go
func (*Attachment) ProtoMessage()
```

#### func (*Attachment) Reset

```go
func (m *Attachment) Reset()
```

#### func (*Attachment) String

```go
func (m *Attachment) String() string
```

#### func (*Attachment) UnMarshalJSON

```go
func (j *Attachment) UnMarshalJSON(obj interface{}) error
```

#### func (*Attachment) XXX_DiscardUnknown

```go
func (m *Attachment) XXX_DiscardUnknown()
```

#### func (*Attachment) XXX_Marshal

```go
func (m *Attachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Attachment) XXX_Merge

```go
func (m *Attachment) XXX_Merge(src proto.Message)
```

#### func (*Attachment) XXX_Size

```go
func (m *Attachment) XXX_Size() int
```

#### func (*Attachment) XXX_Unmarshal

```go
func (m *Attachment) XXX_Unmarshal(b []byte) error
```

#### type AttachmentAction

```go
type AttachmentAction struct {
	Name                 string                         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Text                 string                         `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Style                string                         `protobuf:"bytes,3,opt,name=style,proto3" json:"style,omitempty"`
	Type                 string                         `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Value                string                         `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	DataSource           string                         `protobuf:"bytes,6,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	MinQueryLength       int64                          `protobuf:"varint,7,opt,name=min_query_length,json=minQueryLength,proto3" json:"min_query_length,omitempty"`
	Options              []*AttachmentActionOption      `protobuf:"bytes,8,rep,name=options,proto3" json:"options,omitempty"`
	SelectedOptions      []*AttachmentActionOption      `protobuf:"bytes,9,rep,name=selected_options,json=selectedOptions,proto3" json:"selected_options,omitempty"`
	OptionGroups         []*AttachmentActionOptionGroup `protobuf:"bytes,10,rep,name=option_groups,json=optionGroups,proto3" json:"option_groups,omitempty"`
	Confirm              *AttachmentConfirmationField   `protobuf:"bytes,11,opt,name=confirm,proto3" json:"confirm,omitempty"`
	Url                  string                         `protobuf:"bytes,12,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}
```


#### func  NewAttachmentAction

```go
func NewAttachmentAction(opts ...AttachmentActionFunc) *AttachmentAction
```

#### func (*AttachmentAction) Descriptor

```go
func (*AttachmentAction) Descriptor() ([]byte, []int)
```

#### func (*AttachmentAction) GetConfirm

```go
func (m *AttachmentAction) GetConfirm() *AttachmentConfirmationField
```

#### func (*AttachmentAction) GetDataSource

```go
func (m *AttachmentAction) GetDataSource() string
```

#### func (*AttachmentAction) GetMinQueryLength

```go
func (m *AttachmentAction) GetMinQueryLength() int64
```

#### func (*AttachmentAction) GetName

```go
func (m *AttachmentAction) GetName() string
```

#### func (*AttachmentAction) GetOptionGroups

```go
func (m *AttachmentAction) GetOptionGroups() []*AttachmentActionOptionGroup
```

#### func (*AttachmentAction) GetOptions

```go
func (m *AttachmentAction) GetOptions() []*AttachmentActionOption
```

#### func (*AttachmentAction) GetSelectedOptions

```go
func (m *AttachmentAction) GetSelectedOptions() []*AttachmentActionOption
```

#### func (*AttachmentAction) GetStyle

```go
func (m *AttachmentAction) GetStyle() string
```

#### func (*AttachmentAction) GetText

```go
func (m *AttachmentAction) GetText() string
```

#### func (*AttachmentAction) GetType

```go
func (m *AttachmentAction) GetType() string
```

#### func (*AttachmentAction) GetUrl

```go
func (m *AttachmentAction) GetUrl() string
```

#### func (*AttachmentAction) GetValue

```go
func (m *AttachmentAction) GetValue() string
```

#### func (*AttachmentAction) ProtoMessage

```go
func (*AttachmentAction) ProtoMessage()
```

#### func (*AttachmentAction) Reset

```go
func (m *AttachmentAction) Reset()
```

#### func (*AttachmentAction) String

```go
func (m *AttachmentAction) String() string
```

#### func (*AttachmentAction) XXX_DiscardUnknown

```go
func (m *AttachmentAction) XXX_DiscardUnknown()
```

#### func (*AttachmentAction) XXX_Marshal

```go
func (m *AttachmentAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*AttachmentAction) XXX_Merge

```go
func (m *AttachmentAction) XXX_Merge(src proto.Message)
```

#### func (*AttachmentAction) XXX_Size

```go
func (m *AttachmentAction) XXX_Size() int
```

#### func (*AttachmentAction) XXX_Unmarshal

```go
func (m *AttachmentAction) XXX_Unmarshal(b []byte) error
```

#### type AttachmentActionFunc

```go
type AttachmentActionFunc func(a *AttachmentAction)
```


#### type AttachmentActionOption

```go
type AttachmentActionOption struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*AttachmentActionOption) Descriptor

```go
func (*AttachmentActionOption) Descriptor() ([]byte, []int)
```

#### func (*AttachmentActionOption) GetDescription

```go
func (m *AttachmentActionOption) GetDescription() string
```

#### func (*AttachmentActionOption) GetTitle

```go
func (m *AttachmentActionOption) GetTitle() string
```

#### func (*AttachmentActionOption) GetValue

```go
func (m *AttachmentActionOption) GetValue() string
```

#### func (*AttachmentActionOption) ProtoMessage

```go
func (*AttachmentActionOption) ProtoMessage()
```

#### func (*AttachmentActionOption) Reset

```go
func (m *AttachmentActionOption) Reset()
```

#### func (*AttachmentActionOption) String

```go
func (m *AttachmentActionOption) String() string
```

#### func (*AttachmentActionOption) XXX_DiscardUnknown

```go
func (m *AttachmentActionOption) XXX_DiscardUnknown()
```

#### func (*AttachmentActionOption) XXX_Marshal

```go
func (m *AttachmentActionOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*AttachmentActionOption) XXX_Merge

```go
func (m *AttachmentActionOption) XXX_Merge(src proto.Message)
```

#### func (*AttachmentActionOption) XXX_Size

```go
func (m *AttachmentActionOption) XXX_Size() int
```

#### func (*AttachmentActionOption) XXX_Unmarshal

```go
func (m *AttachmentActionOption) XXX_Unmarshal(b []byte) error
```

#### type AttachmentActionOptionGroup

```go
type AttachmentActionOptionGroup struct {
	Text                 string                    `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Options              []*AttachmentActionOption `protobuf:"bytes,2,rep,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}
```


#### func (*AttachmentActionOptionGroup) Descriptor

```go
func (*AttachmentActionOptionGroup) Descriptor() ([]byte, []int)
```

#### func (*AttachmentActionOptionGroup) GetOptions

```go
func (m *AttachmentActionOptionGroup) GetOptions() []*AttachmentActionOption
```

#### func (*AttachmentActionOptionGroup) GetText

```go
func (m *AttachmentActionOptionGroup) GetText() string
```

#### func (*AttachmentActionOptionGroup) ProtoMessage

```go
func (*AttachmentActionOptionGroup) ProtoMessage()
```

#### func (*AttachmentActionOptionGroup) Reset

```go
func (m *AttachmentActionOptionGroup) Reset()
```

#### func (*AttachmentActionOptionGroup) String

```go
func (m *AttachmentActionOptionGroup) String() string
```

#### func (*AttachmentActionOptionGroup) XXX_DiscardUnknown

```go
func (m *AttachmentActionOptionGroup) XXX_DiscardUnknown()
```

#### func (*AttachmentActionOptionGroup) XXX_Marshal

```go
func (m *AttachmentActionOptionGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*AttachmentActionOptionGroup) XXX_Merge

```go
func (m *AttachmentActionOptionGroup) XXX_Merge(src proto.Message)
```

#### func (*AttachmentActionOptionGroup) XXX_Size

```go
func (m *AttachmentActionOptionGroup) XXX_Size() int
```

#### func (*AttachmentActionOptionGroup) XXX_Unmarshal

```go
func (m *AttachmentActionOptionGroup) XXX_Unmarshal(b []byte) error
```

#### type AttachmentConfirmationField

```go
type AttachmentConfirmationField struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	OkText               string   `protobuf:"bytes,3,opt,name=ok_text,json=okText,proto3" json:"ok_text,omitempty"`
	DismissText          string   `protobuf:"bytes,4,opt,name=dismiss_text,json=dismissText,proto3" json:"dismiss_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*AttachmentConfirmationField) Descriptor

```go
func (*AttachmentConfirmationField) Descriptor() ([]byte, []int)
```

#### func (*AttachmentConfirmationField) GetDismissText

```go
func (m *AttachmentConfirmationField) GetDismissText() string
```

#### func (*AttachmentConfirmationField) GetOkText

```go
func (m *AttachmentConfirmationField) GetOkText() string
```

#### func (*AttachmentConfirmationField) GetText

```go
func (m *AttachmentConfirmationField) GetText() string
```

#### func (*AttachmentConfirmationField) GetTitle

```go
func (m *AttachmentConfirmationField) GetTitle() string
```

#### func (*AttachmentConfirmationField) ProtoMessage

```go
func (*AttachmentConfirmationField) ProtoMessage()
```

#### func (*AttachmentConfirmationField) Reset

```go
func (m *AttachmentConfirmationField) Reset()
```

#### func (*AttachmentConfirmationField) String

```go
func (m *AttachmentConfirmationField) String() string
```

#### func (*AttachmentConfirmationField) XXX_DiscardUnknown

```go
func (m *AttachmentConfirmationField) XXX_DiscardUnknown()
```

#### func (*AttachmentConfirmationField) XXX_Marshal

```go
func (m *AttachmentConfirmationField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*AttachmentConfirmationField) XXX_Merge

```go
func (m *AttachmentConfirmationField) XXX_Merge(src proto.Message)
```

#### func (*AttachmentConfirmationField) XXX_Size

```go
func (m *AttachmentConfirmationField) XXX_Size() int
```

#### func (*AttachmentConfirmationField) XXX_Unmarshal

```go
func (m *AttachmentConfirmationField) XXX_Unmarshal(b []byte) error
```

#### type AttachmentField

```go
type AttachmentField struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Short                bool     `protobuf:"varint,3,opt,name=short,proto3" json:"short,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*AttachmentField) Descriptor

```go
func (*AttachmentField) Descriptor() ([]byte, []int)
```

#### func (*AttachmentField) GetShort

```go
func (m *AttachmentField) GetShort() bool
```

#### func (*AttachmentField) GetTitle

```go
func (m *AttachmentField) GetTitle() string
```

#### func (*AttachmentField) GetValue

```go
func (m *AttachmentField) GetValue() string
```

#### func (*AttachmentField) ProtoMessage

```go
func (*AttachmentField) ProtoMessage()
```

#### func (*AttachmentField) Reset

```go
func (m *AttachmentField) Reset()
```

#### func (*AttachmentField) String

```go
func (m *AttachmentField) String() string
```

#### func (*AttachmentField) XXX_DiscardUnknown

```go
func (m *AttachmentField) XXX_DiscardUnknown()
```

#### func (*AttachmentField) XXX_Marshal

```go
func (m *AttachmentField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*AttachmentField) XXX_Merge

```go
func (m *AttachmentField) XXX_Merge(src proto.Message)
```

#### func (*AttachmentField) XXX_Size

```go
func (m *AttachmentField) XXX_Size() int
```

#### func (*AttachmentField) XXX_Unmarshal

```go
func (m *AttachmentField) XXX_Unmarshal(b []byte) error
```

#### type AttachmentFunc

```go
type AttachmentFunc func(a *Attachment)
```


#### type BankAccount

```go
type BankAccount struct {
	AccountNumber        string   `protobuf:"bytes,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	RoutingNumber        string   `protobuf:"bytes,2,opt,name=routing_number,json=routingNumber,proto3" json:"routing_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*BankAccount) Descriptor

```go
func (*BankAccount) Descriptor() ([]byte, []int)
```

#### func (*BankAccount) GetAccountNumber

```go
func (m *BankAccount) GetAccountNumber() string
```

#### func (*BankAccount) GetRoutingNumber

```go
func (m *BankAccount) GetRoutingNumber() string
```

#### func (*BankAccount) ProtoMessage

```go
func (*BankAccount) ProtoMessage()
```

#### func (*BankAccount) Reset

```go
func (m *BankAccount) Reset()
```

#### func (*BankAccount) String

```go
func (m *BankAccount) String() string
```

#### func (*BankAccount) XXX_DiscardUnknown

```go
func (m *BankAccount) XXX_DiscardUnknown()
```

#### func (*BankAccount) XXX_Marshal

```go
func (m *BankAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*BankAccount) XXX_Merge

```go
func (m *BankAccount) XXX_Merge(src proto.Message)
```

#### func (*BankAccount) XXX_Size

```go
func (m *BankAccount) XXX_Size() int
```

#### func (*BankAccount) XXX_Unmarshal

```go
func (m *BankAccount) XXX_Unmarshal(b []byte) error
```

#### type Call

```go
type Call struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Callback             string   `protobuf:"bytes,5,opt,name=callback,proto3" json:"callback,omitempty"`
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

#### type CallRequest

```go
type CallRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CallbackUrl          string   `protobuf:"bytes,2,opt,name=callback_url,json=callbackUrl,proto3" json:"callback_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*CallRequest) Descriptor

```go
func (*CallRequest) Descriptor() ([]byte, []int)
```

#### func (*CallRequest) GetCallbackUrl

```go
func (m *CallRequest) GetCallbackUrl() string
```

#### func (*CallRequest) GetUserId

```go
func (m *CallRequest) GetUserId() string
```

#### func (*CallRequest) ProtoMessage

```go
func (*CallRequest) ProtoMessage()
```

#### func (*CallRequest) Reset

```go
func (m *CallRequest) Reset()
```

#### func (*CallRequest) String

```go
func (m *CallRequest) String() string
```

#### func (*CallRequest) XXX_DiscardUnknown

```go
func (m *CallRequest) XXX_DiscardUnknown()
```

#### func (*CallRequest) XXX_Marshal

```go
func (m *CallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*CallRequest) XXX_Merge

```go
func (m *CallRequest) XXX_Merge(src proto.Message)
```

#### func (*CallRequest) XXX_Size

```go
func (m *CallRequest) XXX_Size() int
```

#### func (*CallRequest) XXX_Unmarshal

```go
func (m *CallRequest) XXX_Unmarshal(b []byte) error
```

#### type CancelSubscriptionRequest

```go
type CancelSubscriptionRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*CancelSubscriptionRequest) Descriptor

```go
func (*CancelSubscriptionRequest) Descriptor() ([]byte, []int)
```

#### func (*CancelSubscriptionRequest) GetEmail

```go
func (m *CancelSubscriptionRequest) GetEmail() string
```

#### func (*CancelSubscriptionRequest) ProtoMessage

```go
func (*CancelSubscriptionRequest) ProtoMessage()
```

#### func (*CancelSubscriptionRequest) Reset

```go
func (m *CancelSubscriptionRequest) Reset()
```

#### func (*CancelSubscriptionRequest) String

```go
func (m *CancelSubscriptionRequest) String() string
```

#### func (*CancelSubscriptionRequest) XXX_DiscardUnknown

```go
func (m *CancelSubscriptionRequest) XXX_DiscardUnknown()
```

#### func (*CancelSubscriptionRequest) XXX_Marshal

```go
func (m *CancelSubscriptionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*CancelSubscriptionRequest) XXX_Merge

```go
func (m *CancelSubscriptionRequest) XXX_Merge(src proto.Message)
```

#### func (*CancelSubscriptionRequest) XXX_Size

```go
func (m *CancelSubscriptionRequest) XXX_Size() int
```

#### func (*CancelSubscriptionRequest) XXX_Unmarshal

```go
func (m *CancelSubscriptionRequest) XXX_Unmarshal(b []byte) error
```

#### type Card

```go
type Card struct {
	CardType             CardType `protobuf:"varint,1,opt,name=card_type,json=cardType,proto3,enum=api.CardType" json:"card_type,omitempty"`
	CardNumber           string   `protobuf:"bytes,3,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	ExpMonth             string   `protobuf:"bytes,4,opt,name=exp_month,json=expMonth,proto3" json:"exp_month,omitempty"`
	ExpYear              string   `protobuf:"bytes,5,opt,name=exp_year,json=expYear,proto3" json:"exp_year,omitempty"`
	Cvc                  string   `protobuf:"bytes,6,opt,name=cvc,proto3" json:"cvc,omitempty"`
	Debit                bool     `protobuf:"varint,7,opt,name=debit,proto3" json:"debit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Card) Descriptor

```go
func (*Card) Descriptor() ([]byte, []int)
```

#### func (*Card) GetCardNumber

```go
func (m *Card) GetCardNumber() string
```

#### func (*Card) GetCardType

```go
func (m *Card) GetCardType() CardType
```

#### func (*Card) GetCvc

```go
func (m *Card) GetCvc() string
```

#### func (*Card) GetDebit

```go
func (m *Card) GetDebit() bool
```

#### func (*Card) GetExpMonth

```go
func (m *Card) GetExpMonth() string
```

#### func (*Card) GetExpYear

```go
func (m *Card) GetExpYear() string
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

#### type CardType

```go
type CardType int32
```


```go
const (
	CardType_VISA       CardType = 0
	CardType_MASTERCARD CardType = 1
	CardType_DISCOVER   CardType = 2
	CardType_AMEX       CardType = 3
)
```

#### func (CardType) EnumDescriptor

```go
func (CardType) EnumDescriptor() ([]byte, []int)
```

#### func (CardType) String

```go
func (x CardType) String() string
```

#### type ChannelReminder

```go
type ChannelReminder struct {
	ChannelId            string   `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Time                 string   `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*ChannelReminder) Descriptor

```go
func (*ChannelReminder) Descriptor() ([]byte, []int)
```

#### func (*ChannelReminder) GetChannelId

```go
func (m *ChannelReminder) GetChannelId() string
```

#### func (*ChannelReminder) GetText

```go
func (m *ChannelReminder) GetText() string
```

#### func (*ChannelReminder) GetTime

```go
func (m *ChannelReminder) GetTime() string
```

#### func (*ChannelReminder) ProtoMessage

```go
func (*ChannelReminder) ProtoMessage()
```

#### func (*ChannelReminder) Reset

```go
func (m *ChannelReminder) Reset()
```

#### func (*ChannelReminder) String

```go
func (m *ChannelReminder) String() string
```

#### func (*ChannelReminder) XXX_DiscardUnknown

```go
func (m *ChannelReminder) XXX_DiscardUnknown()
```

#### func (*ChannelReminder) XXX_Marshal

```go
func (m *ChannelReminder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*ChannelReminder) XXX_Merge

```go
func (m *ChannelReminder) XXX_Merge(src proto.Message)
```

#### func (*ChannelReminder) XXX_Size

```go
func (m *ChannelReminder) XXX_Size() int
```

#### func (*ChannelReminder) XXX_Unmarshal

```go
func (m *ChannelReminder) XXX_Unmarshal(b []byte) error
```

#### type Claims

```go
type Claims struct {
}
```

/////////////////////////////////////////////////////////////

#### func (*Claims) SignedGCPToken

```go
func (c *Claims) SignedGCPToken(secret string) (*SignedKey, error)
```

#### func (*Claims) SignedSendGridToken

```go
func (c *Claims) SignedSendGridToken(secret string) (*SignedKey, error)
```

#### func (*Claims) SignedSlackToken

```go
func (c *Claims) SignedSlackToken(secret string) (*SignedKey, error)
```

#### func (*Claims) SignedStripeToken

```go
func (c *Claims) SignedStripeToken(secret string) (*SignedKey, error)
```

#### func (*Claims) SignedTwilioToken

```go
func (c *Claims) SignedTwilioToken(secret string) (*SignedKey, error)
```

#### type ClientSet

```go
type ClientSet struct {
	UserSet UserServiceClient
}
```


#### func  NewClientSet

```go
func NewClientSet(ctx context.Context, addr string, opts ...grpc.DialOption) (*ClientSet, error)
```

#### type CreatePlanRequest

```go
type CreatePlanRequest struct {
	PlanId               string   `protobuf:"bytes,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	ServiceId            string   `protobuf:"bytes,3,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	ServiceName          string   `protobuf:"bytes,4,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	FriendlyName         string   `protobuf:"bytes,5,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*CreatePlanRequest) Descriptor

```go
func (*CreatePlanRequest) Descriptor() ([]byte, []int)
```

#### func (*CreatePlanRequest) GetAmount

```go
func (m *CreatePlanRequest) GetAmount() int64
```

#### func (*CreatePlanRequest) GetFriendlyName

```go
func (m *CreatePlanRequest) GetFriendlyName() string
```

#### func (*CreatePlanRequest) GetPlanId

```go
func (m *CreatePlanRequest) GetPlanId() string
```

#### func (*CreatePlanRequest) GetServiceId

```go
func (m *CreatePlanRequest) GetServiceId() string
```

#### func (*CreatePlanRequest) GetServiceName

```go
func (m *CreatePlanRequest) GetServiceName() string
```

#### func (*CreatePlanRequest) ProtoMessage

```go
func (*CreatePlanRequest) ProtoMessage()
```

#### func (*CreatePlanRequest) Reset

```go
func (m *CreatePlanRequest) Reset()
```

#### func (*CreatePlanRequest) String

```go
func (m *CreatePlanRequest) String() string
```

#### func (*CreatePlanRequest) XXX_DiscardUnknown

```go
func (m *CreatePlanRequest) XXX_DiscardUnknown()
```

#### func (*CreatePlanRequest) XXX_Marshal

```go
func (m *CreatePlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*CreatePlanRequest) XXX_Merge

```go
func (m *CreatePlanRequest) XXX_Merge(src proto.Message)
```

#### func (*CreatePlanRequest) XXX_Size

```go
func (m *CreatePlanRequest) XXX_Size() int
```

#### func (*CreatePlanRequest) XXX_Unmarshal

```go
func (m *CreatePlanRequest) XXX_Unmarshal(b []byte) error
```

#### type CreatePlanResponse

```go
type CreatePlanResponse struct {
	PlanId               string   `protobuf:"bytes,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*CreatePlanResponse) Descriptor

```go
func (*CreatePlanResponse) Descriptor() ([]byte, []int)
```

#### func (*CreatePlanResponse) GetPlanId

```go
func (m *CreatePlanResponse) GetPlanId() string
```

#### func (*CreatePlanResponse) ProtoMessage

```go
func (*CreatePlanResponse) ProtoMessage()
```

#### func (*CreatePlanResponse) Reset

```go
func (m *CreatePlanResponse) Reset()
```

#### func (*CreatePlanResponse) String

```go
func (m *CreatePlanResponse) String() string
```

#### func (*CreatePlanResponse) XXX_DiscardUnknown

```go
func (m *CreatePlanResponse) XXX_DiscardUnknown()
```

#### func (*CreatePlanResponse) XXX_Marshal

```go
func (m *CreatePlanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*CreatePlanResponse) XXX_Merge

```go
func (m *CreatePlanResponse) XXX_Merge(src proto.Message)
```

#### func (*CreatePlanResponse) XXX_Size

```go
func (m *CreatePlanResponse) XXX_Size() int
```

#### func (*CreatePlanResponse) XXX_Unmarshal

```go
func (m *CreatePlanResponse) XXX_Unmarshal(b []byte) error
```

#### type Customer

```go
type Customer struct {
	UserId               string            `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Plan                 string            `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	Name                 string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email                string            `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Description          string            `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Phone                string            `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Address              *Address          `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,9,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Deleted              bool              `protobuf:"varint,10,opt,name=deleted,proto3" json:"deleted,omitempty"`
	CreateDate           int64             `protobuf:"varint,20,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```

User is a user of the application

#### func (*Customer) AsMap

```go
func (j *Customer) AsMap() map[string]interface{}
```

#### func (*Customer) CompileHTML

```go
func (j *Customer) CompileHTML(text string, w io.Writer) error
```

#### func (*Customer) CompileTXT

```go
func (j *Customer) CompileTXT(text string, w io.Writer) error
```

#### func (*Customer) Descriptor

```go
func (*Customer) Descriptor() ([]byte, []int)
```

#### func (*Customer) GetAddress

```go
func (m *Customer) GetAddress() *Address
```

#### func (*Customer) GetCreateDate

```go
func (m *Customer) GetCreateDate() int64
```

#### func (*Customer) GetDeleted

```go
func (m *Customer) GetDeleted() bool
```

#### func (*Customer) GetDescription

```go
func (m *Customer) GetDescription() string
```

#### func (*Customer) GetEmail

```go
func (m *Customer) GetEmail() string
```

#### func (*Customer) GetMetadata

```go
func (m *Customer) GetMetadata() map[string]string
```

#### func (*Customer) GetName

```go
func (m *Customer) GetName() string
```

#### func (*Customer) GetPhone

```go
func (m *Customer) GetPhone() string
```

#### func (*Customer) GetPlan

```go
func (m *Customer) GetPlan() string
```

#### func (*Customer) GetUserId

```go
func (m *Customer) GetUserId() string
```

#### func (*Customer) MarshalJSON

```go
func (j *Customer) MarshalJSON() ([]byte, error)
```

#### func (*Customer) ProtoMessage

```go
func (*Customer) ProtoMessage()
```

#### func (*Customer) Reset

```go
func (m *Customer) Reset()
```

#### func (*Customer) String

```go
func (m *Customer) String() string
```

#### func (*Customer) UnMarshalJSON

```go
func (j *Customer) UnMarshalJSON(obj interface{}) error
```

#### func (*Customer) XXX_DiscardUnknown

```go
func (m *Customer) XXX_DiscardUnknown()
```

#### func (*Customer) XXX_Marshal

```go
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Customer) XXX_Merge

```go
func (m *Customer) XXX_Merge(src proto.Message)
```

#### func (*Customer) XXX_Size

```go
func (m *Customer) XXX_Size() int
```

#### func (*Customer) XXX_Unmarshal

```go
func (m *Customer) XXX_Unmarshal(b []byte) error
```

#### type CustomerIndex

```go
type CustomerIndex int32
```


```go
const (
	CustomerIndex_ID    CustomerIndex = 0
	CustomerIndex_EMAIL CustomerIndex = 1
	CustomerIndex_PHONE CustomerIndex = 2
)
```

#### func (CustomerIndex) EnumDescriptor

```go
func (CustomerIndex) EnumDescriptor() ([]byte, []int)
```

#### func (CustomerIndex) String

```go
func (x CustomerIndex) String() string
```

#### type Email

```go
type Email struct {
	From                 *EmailAddress   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Recipient            *RecipientEmail `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}
```


#### func (*Email) Descriptor

```go
func (*Email) Descriptor() ([]byte, []int)
```

#### func (*Email) GetFrom

```go
func (m *Email) GetFrom() *EmailAddress
```

#### func (*Email) GetRecipient

```go
func (m *Email) GetRecipient() *RecipientEmail
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

#### type EmailAddress

```go
type EmailAddress struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*EmailAddress) Descriptor

```go
func (*EmailAddress) Descriptor() ([]byte, []int)
```

#### func (*EmailAddress) GetAddress

```go
func (m *EmailAddress) GetAddress() string
```

#### func (*EmailAddress) GetName

```go
func (m *EmailAddress) GetName() string
```

#### func (*EmailAddress) ProtoMessage

```go
func (*EmailAddress) ProtoMessage()
```

#### func (*EmailAddress) Reset

```go
func (m *EmailAddress) Reset()
```

#### func (*EmailAddress) String

```go
func (m *EmailAddress) String() string
```

#### func (*EmailAddress) XXX_DiscardUnknown

```go
func (m *EmailAddress) XXX_DiscardUnknown()
```

#### func (*EmailAddress) XXX_Marshal

```go
func (m *EmailAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*EmailAddress) XXX_Merge

```go
func (m *EmailAddress) XXX_Merge(src proto.Message)
```

#### func (*EmailAddress) XXX_Size

```go
func (m *EmailAddress) XXX_Size() int
```

#### func (*EmailAddress) XXX_Unmarshal

```go
func (m *EmailAddress) XXX_Unmarshal(b []byte) error
```

#### type EmailRequest

```go
type EmailRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Subject              string   `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	PlainText            string   `protobuf:"bytes,3,opt,name=plain_text,json=plainText,proto3" json:"plain_text,omitempty"`
	HtmlAlt              string   `protobuf:"bytes,4,opt,name=html_alt,json=htmlAlt,proto3" json:"html_alt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*EmailRequest) Descriptor

```go
func (*EmailRequest) Descriptor() ([]byte, []int)
```

#### func (*EmailRequest) GetHtmlAlt

```go
func (m *EmailRequest) GetHtmlAlt() string
```

#### func (*EmailRequest) GetPlainText

```go
func (m *EmailRequest) GetPlainText() string
```

#### func (*EmailRequest) GetSubject

```go
func (m *EmailRequest) GetSubject() string
```

#### func (*EmailRequest) GetUserId

```go
func (m *EmailRequest) GetUserId() string
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

#### type Fax

```go
type Fax struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	MediaUrl             string   `protobuf:"bytes,3,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	Quality              string   `protobuf:"bytes,4,opt,name=quality,proto3" json:"quality,omitempty"`
	Callback             string   `protobuf:"bytes,5,opt,name=callback,proto3" json:"callback,omitempty"`
	StoreMedia           bool     `protobuf:"varint,6,opt,name=store_media,json=storeMedia,proto3" json:"store_media,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Fax) Descriptor

```go
func (*Fax) Descriptor() ([]byte, []int)
```

#### func (*Fax) GetCallback

```go
func (m *Fax) GetCallback() string
```

#### func (*Fax) GetFrom

```go
func (m *Fax) GetFrom() string
```

#### func (*Fax) GetMediaUrl

```go
func (m *Fax) GetMediaUrl() string
```

#### func (*Fax) GetQuality

```go
func (m *Fax) GetQuality() string
```

#### func (*Fax) GetStoreMedia

```go
func (m *Fax) GetStoreMedia() bool
```

#### func (*Fax) GetTo

```go
func (m *Fax) GetTo() string
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

#### type Grant

```go
type Grant int32
```


```go
const (
	Grant_TWILIO    Grant = 0
	Grant_SENDGRID  Grant = 1
	Grant_STRIPE    Grant = 2
	Grant_SLACK     Grant = 3
	Grant_GCP       Grant = 4
	Grant_AUTOM8TER Grant = 5
)
```

#### func (Grant) EnumDescriptor

```go
func (Grant) EnumDescriptor() ([]byte, []int)
```

#### func (Grant) String

```go
func (x Grant) String() string
```

#### type ItemRef

```go
type ItemRef struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	File                 string   `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*ItemRef) Descriptor

```go
func (*ItemRef) Descriptor() ([]byte, []int)
```

#### func (*ItemRef) GetChannel

```go
func (m *ItemRef) GetChannel() string
```

#### func (*ItemRef) GetComment

```go
func (m *ItemRef) GetComment() string
```

#### func (*ItemRef) GetFile

```go
func (m *ItemRef) GetFile() string
```

#### func (*ItemRef) ProtoMessage

```go
func (*ItemRef) ProtoMessage()
```

#### func (*ItemRef) Reset

```go
func (m *ItemRef) Reset()
```

#### func (*ItemRef) String

```go
func (m *ItemRef) String() string
```

#### func (*ItemRef) XXX_DiscardUnknown

```go
func (m *ItemRef) XXX_DiscardUnknown()
```

#### func (*ItemRef) XXX_Marshal

```go
func (m *ItemRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*ItemRef) XXX_Merge

```go
func (m *ItemRef) XXX_Merge(src proto.Message)
```

#### func (*ItemRef) XXX_Size

```go
func (m *ItemRef) XXX_Size() int
```

#### func (*ItemRef) XXX_Unmarshal

```go
func (m *ItemRef) XXX_Unmarshal(b []byte) error
```

#### type JSON

```go
type JSON struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*JSON) AsMap

```go
func (j *JSON) AsMap() map[string]interface{}
```

#### func (*JSON) CompileHTML

```go
func (j *JSON) CompileHTML(text string, w io.Writer) error
```

#### func (*JSON) CompileTXT

```go
func (j *JSON) CompileTXT(text string, w io.Writer) error
```

#### func (*JSON) Descriptor

```go
func (*JSON) Descriptor() ([]byte, []int)
```

#### func (*JSON) GetData

```go
func (m *JSON) GetData() []byte
```

#### func (*JSON) GetSize

```go
func (m *JSON) GetSize() int64
```

#### func (*JSON) MarshalJSON

```go
func (j *JSON) MarshalJSON() ([]byte, error)
```
////////////////////////////////////////////////////

#### func (*JSON) ProtoMessage

```go
func (*JSON) ProtoMessage()
```

#### func (*JSON) Reset

```go
func (m *JSON) Reset()
```

#### func (*JSON) String

```go
func (m *JSON) String() string
```

#### func (*JSON) UnMarshalJSON

```go
func (j *JSON) UnMarshalJSON(obj interface{}) error
```

#### func (*JSON) XXX_DiscardUnknown

```go
func (m *JSON) XXX_DiscardUnknown()
```

#### func (*JSON) XXX_Marshal

```go
func (m *JSON) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*JSON) XXX_Merge

```go
func (m *JSON) XXX_Merge(src proto.Message)
```

#### func (*JSON) XXX_Size

```go
func (m *JSON) XXX_Size() int
```

#### func (*JSON) XXX_Unmarshal

```go
func (m *JSON) XXX_Unmarshal(b []byte) error
```

#### type JSONMap

```go
type JSONMap struct {
	JsonMap              map[string]*JSON `protobuf:"bytes,1,rep,name=json_map,json=jsonMap,proto3" json:"json_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}
```


#### func (*JSONMap) Descriptor

```go
func (*JSONMap) Descriptor() ([]byte, []int)
```

#### func (*JSONMap) GetJsonMap

```go
func (m *JSONMap) GetJsonMap() map[string]*JSON
```

#### func (*JSONMap) ProtoMessage

```go
func (*JSONMap) ProtoMessage()
```

#### func (*JSONMap) Reset

```go
func (m *JSONMap) Reset()
```

#### func (*JSONMap) String

```go
func (m *JSONMap) String() string
```

#### func (*JSONMap) XXX_DiscardUnknown

```go
func (m *JSONMap) XXX_DiscardUnknown()
```

#### func (*JSONMap) XXX_Marshal

```go
func (m *JSONMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*JSONMap) XXX_Merge

```go
func (m *JSONMap) XXX_Merge(src proto.Message)
```

#### func (*JSONMap) XXX_Size

```go
func (m *JSONMap) XXX_Size() int
```

#### func (*JSONMap) XXX_Unmarshal

```go
func (m *JSONMap) XXX_Unmarshal(b []byte) error
```

#### type JWTMiddleware

```go
type JWTMiddleware struct {
}
```


#### func  NewJWTMiddleware

```go
func NewJWTMiddleware(key SignedKey) *JWTMiddleware
```

#### func (*JWTMiddleware) Check

```go
func (j *JWTMiddleware) Check(w http.ResponseWriter, r *http.Request) error
```

#### func (*JWTMiddleware) Wrap

```go
func (j *JWTMiddleware) Wrap(next http.HandlerFunc) http.HandlerFunc
```

#### type JWTToken

```go
type JWTToken struct {
	Raw                  string            `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	Method               SigningMethod     `protobuf:"varint,2,opt,name=method,proto3,enum=api.SigningMethod" json:"method,omitempty"`
	Header               map[string]string `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Claims               string            `protobuf:"bytes,4,opt,name=claims,proto3" json:"claims,omitempty"`
	Signature            string            `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Value                bool              `protobuf:"varint,6,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func (*JWTToken) Descriptor

```go
func (*JWTToken) Descriptor() ([]byte, []int)
```

#### func (*JWTToken) GetClaims

```go
func (m *JWTToken) GetClaims() string
```

#### func (*JWTToken) GetHeader

```go
func (m *JWTToken) GetHeader() map[string]string
```

#### func (*JWTToken) GetMethod

```go
func (m *JWTToken) GetMethod() SigningMethod
```

#### func (*JWTToken) GetRaw

```go
func (m *JWTToken) GetRaw() string
```

#### func (*JWTToken) GetSignature

```go
func (m *JWTToken) GetSignature() string
```

#### func (*JWTToken) GetValue

```go
func (m *JWTToken) GetValue() bool
```

#### func (*JWTToken) ProtoMessage

```go
func (*JWTToken) ProtoMessage()
```

#### func (*JWTToken) Reset

```go
func (m *JWTToken) Reset()
```

#### func (*JWTToken) String

```go
func (m *JWTToken) String() string
```

#### func (*JWTToken) XXX_DiscardUnknown

```go
func (m *JWTToken) XXX_DiscardUnknown()
```

#### func (*JWTToken) XXX_Marshal

```go
func (m *JWTToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*JWTToken) XXX_Merge

```go
func (m *JWTToken) XXX_Merge(src proto.Message)
```

#### func (*JWTToken) XXX_Size

```go
func (m *JWTToken) XXX_Size() int
```

#### func (*JWTToken) XXX_Unmarshal

```go
func (m *JWTToken) XXX_Unmarshal(b []byte) error
```

#### type LogConfig

```go
type LogConfig struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Channel              string   `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*LogConfig) Descriptor

```go
func (*LogConfig) Descriptor() ([]byte, []int)
```

#### func (*LogConfig) GetChannel

```go
func (m *LogConfig) GetChannel() string
```

#### func (*LogConfig) GetUsername

```go
func (m *LogConfig) GetUsername() string
```

#### func (*LogConfig) ProtoMessage

```go
func (*LogConfig) ProtoMessage()
```

#### func (*LogConfig) Reset

```go
func (m *LogConfig) Reset()
```

#### func (*LogConfig) String

```go
func (m *LogConfig) String() string
```

#### func (*LogConfig) XXX_DiscardUnknown

```go
func (m *LogConfig) XXX_DiscardUnknown()
```

#### func (*LogConfig) XXX_Marshal

```go
func (m *LogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*LogConfig) XXX_Merge

```go
func (m *LogConfig) XXX_Merge(src proto.Message)
```

#### func (*LogConfig) XXX_Size

```go
func (m *LogConfig) XXX_Size() int
```

#### func (*LogConfig) XXX_Unmarshal

```go
func (m *LogConfig) XXX_Unmarshal(b []byte) error
```

#### type LogHook

```go
type LogHook struct {
	Author               string   `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Icon                 string   `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*LogHook) Descriptor

```go
func (*LogHook) Descriptor() ([]byte, []int)
```

#### func (*LogHook) GetAuthor

```go
func (m *LogHook) GetAuthor() string
```

#### func (*LogHook) GetIcon

```go
func (m *LogHook) GetIcon() string
```

#### func (*LogHook) GetTitle

```go
func (m *LogHook) GetTitle() string
```

#### func (*LogHook) ProtoMessage

```go
func (*LogHook) ProtoMessage()
```

#### func (*LogHook) Reset

```go
func (m *LogHook) Reset()
```

#### func (*LogHook) String

```go
func (m *LogHook) String() string
```

#### func (*LogHook) XXX_DiscardUnknown

```go
func (m *LogHook) XXX_DiscardUnknown()
```

#### func (*LogHook) XXX_Marshal

```go
func (m *LogHook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*LogHook) XXX_Merge

```go
func (m *LogHook) XXX_Merge(src proto.Message)
```

#### func (*LogHook) XXX_Size

```go
func (m *LogHook) XXX_Size() int
```

#### func (*LogHook) XXX_Unmarshal

```go
func (m *LogHook) XXX_Unmarshal(b []byte) error
```

#### type MMSRequest

```go
type MMSRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	MediaUrl             string   `protobuf:"bytes,3,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*MMSRequest) Descriptor

```go
func (*MMSRequest) Descriptor() ([]byte, []int)
```

#### func (*MMSRequest) GetBody

```go
func (m *MMSRequest) GetBody() string
```

#### func (*MMSRequest) GetMediaUrl

```go
func (m *MMSRequest) GetMediaUrl() string
```

#### func (*MMSRequest) GetUserId

```go
func (m *MMSRequest) GetUserId() string
```

#### func (*MMSRequest) ProtoMessage

```go
func (*MMSRequest) ProtoMessage()
```

#### func (*MMSRequest) Reset

```go
func (m *MMSRequest) Reset()
```

#### func (*MMSRequest) String

```go
func (m *MMSRequest) String() string
```

#### func (*MMSRequest) XXX_DiscardUnknown

```go
func (m *MMSRequest) XXX_DiscardUnknown()
```

#### func (*MMSRequest) XXX_Marshal

```go
func (m *MMSRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*MMSRequest) XXX_Merge

```go
func (m *MMSRequest) XXX_Merge(src proto.Message)
```

#### func (*MMSRequest) XXX_Size

```go
func (m *MMSRequest) XXX_Size() int
```

#### func (*MMSRequest) XXX_Unmarshal

```go
func (m *MMSRequest) XXX_Unmarshal(b []byte) error
```

#### type Pin

```go
type Pin struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Item                 *ItemRef `protobuf:"bytes,4,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Pin) Descriptor

```go
func (*Pin) Descriptor() ([]byte, []int)
```

#### func (*Pin) GetItem

```go
func (m *Pin) GetItem() *ItemRef
```

#### func (*Pin) GetText

```go
func (m *Pin) GetText() string
```

#### func (*Pin) ProtoMessage

```go
func (*Pin) ProtoMessage()
```

#### func (*Pin) Reset

```go
func (m *Pin) Reset()
```

#### func (*Pin) String

```go
func (m *Pin) String() string
```

#### func (*Pin) XXX_DiscardUnknown

```go
func (m *Pin) XXX_DiscardUnknown()
```

#### func (*Pin) XXX_Marshal

```go
func (m *Pin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Pin) XXX_Merge

```go
func (m *Pin) XXX_Merge(src proto.Message)
```

#### func (*Pin) XXX_Size

```go
func (m *Pin) XXX_Size() int
```

#### func (*Pin) XXX_Unmarshal

```go
func (m *Pin) XXX_Unmarshal(b []byte) error
```

#### type PubSubFunc

```go
type PubSubFunc func(a *PubSubMessage)
```


#### type PubSubMessage

```go
type PubSubMessage struct {
	Message              *v1.PubsubMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func  NewPubSub

```go
func NewPubSub(opts ...PubSubFunc) *PubSubMessage
```

#### func (*PubSubMessage) AsMap

```go
func (j *PubSubMessage) AsMap() map[string]interface{}
```

#### func (*PubSubMessage) CompileHTML

```go
func (j *PubSubMessage) CompileHTML(text string, w io.Writer) error
```

#### func (*PubSubMessage) CompileTXT

```go
func (j *PubSubMessage) CompileTXT(text string, w io.Writer) error
```

#### func (*PubSubMessage) Descriptor

```go
func (*PubSubMessage) Descriptor() ([]byte, []int)
```

#### func (*PubSubMessage) GetMessage

```go
func (m *PubSubMessage) GetMessage() *v1.PubsubMessage
```

#### func (*PubSubMessage) MarshalJSON

```go
func (j *PubSubMessage) MarshalJSON() ([]byte, error)
```

#### func (*PubSubMessage) ProtoMessage

```go
func (*PubSubMessage) ProtoMessage()
```

#### func (*PubSubMessage) PubSubMessage

```go
func (p *PubSubMessage) PubSubMessage() *pubsub.PubsubMessage
```

#### func (*PubSubMessage) Reset

```go
func (m *PubSubMessage) Reset()
```

#### func (*PubSubMessage) String

```go
func (m *PubSubMessage) String() string
```

#### func (*PubSubMessage) UnMarshalJSON

```go
func (j *PubSubMessage) UnMarshalJSON(obj interface{}) error
```

#### func (*PubSubMessage) XXX_DiscardUnknown

```go
func (m *PubSubMessage) XXX_DiscardUnknown()
```

#### func (*PubSubMessage) XXX_Marshal

```go
func (m *PubSubMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*PubSubMessage) XXX_Merge

```go
func (m *PubSubMessage) XXX_Merge(src proto.Message)
```

#### func (*PubSubMessage) XXX_Size

```go
func (m *PubSubMessage) XXX_Size() int
```

#### func (*PubSubMessage) XXX_Unmarshal

```go
func (m *PubSubMessage) XXX_Unmarshal(b []byte) error
```

#### type PubSubTopic

```go
type PubSubTopic struct {
	Topic                *v1.Topic `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}
```


#### func  NewPubSubTopic

```go
func NewPubSubTopic(opts ...PubSubTopicFunc) *PubSubTopic
```

#### func (*PubSubTopic) Descriptor

```go
func (*PubSubTopic) Descriptor() ([]byte, []int)
```

#### func (*PubSubTopic) GetTopic

```go
func (m *PubSubTopic) GetTopic() *v1.Topic
```

#### func (*PubSubTopic) ProtoMessage

```go
func (*PubSubTopic) ProtoMessage()
```

#### func (*PubSubTopic) PubSubTopic

```go
func (p *PubSubTopic) PubSubTopic() *pubsub.Topic
```

#### func (*PubSubTopic) Reset

```go
func (m *PubSubTopic) Reset()
```

#### func (*PubSubTopic) String

```go
func (m *PubSubTopic) String() string
```

#### func (*PubSubTopic) XXX_DiscardUnknown

```go
func (m *PubSubTopic) XXX_DiscardUnknown()
```

#### func (*PubSubTopic) XXX_Marshal

```go
func (m *PubSubTopic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*PubSubTopic) XXX_Merge

```go
func (m *PubSubTopic) XXX_Merge(src proto.Message)
```

#### func (*PubSubTopic) XXX_Size

```go
func (m *PubSubTopic) XXX_Size() int
```

#### func (*PubSubTopic) XXX_Unmarshal

```go
func (m *PubSubTopic) XXX_Unmarshal(b []byte) error
```

#### type PubSubTopicFunc

```go
type PubSubTopicFunc func(a *PubSubTopic)
```


#### type RecipientEmail

```go
type RecipientEmail struct {
	To                   *EmailAddress `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Subject              string        `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	PlainText            string        `protobuf:"bytes,4,opt,name=plain_text,json=plainText,proto3" json:"plain_text,omitempty"`
	Html                 string        `protobuf:"bytes,5,opt,name=html,proto3" json:"html,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
```


#### func (*RecipientEmail) Descriptor

```go
func (*RecipientEmail) Descriptor() ([]byte, []int)
```

#### func (*RecipientEmail) GetHtml

```go
func (m *RecipientEmail) GetHtml() string
```

#### func (*RecipientEmail) GetPlainText

```go
func (m *RecipientEmail) GetPlainText() string
```

#### func (*RecipientEmail) GetSubject

```go
func (m *RecipientEmail) GetSubject() string
```

#### func (*RecipientEmail) GetTo

```go
func (m *RecipientEmail) GetTo() *EmailAddress
```

#### func (*RecipientEmail) ProtoMessage

```go
func (*RecipientEmail) ProtoMessage()
```

#### func (*RecipientEmail) Reset

```go
func (m *RecipientEmail) Reset()
```

#### func (*RecipientEmail) String

```go
func (m *RecipientEmail) String() string
```

#### func (*RecipientEmail) XXX_DiscardUnknown

```go
func (m *RecipientEmail) XXX_DiscardUnknown()
```

#### func (*RecipientEmail) XXX_Marshal

```go
func (m *RecipientEmail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*RecipientEmail) XXX_Merge

```go
func (m *RecipientEmail) XXX_Merge(src proto.Message)
```

#### func (*RecipientEmail) XXX_Size

```go
func (m *RecipientEmail) XXX_Size() int
```

#### func (*RecipientEmail) XXX_Unmarshal

```go
func (m *RecipientEmail) XXX_Unmarshal(b []byte) error
```

#### type SMS

```go
type SMS struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	MediaUrl             string   `protobuf:"bytes,4,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
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

#### func (*SMS) GetBody

```go
func (m *SMS) GetBody() string
```

#### func (*SMS) GetCallback

```go
func (m *SMS) GetCallback() string
```

#### func (*SMS) GetFrom

```go
func (m *SMS) GetFrom() string
```

#### func (*SMS) GetMediaUrl

```go
func (m *SMS) GetMediaUrl() string
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

#### type SMSRequest

```go
type SMSRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*SMSRequest) Descriptor

```go
func (*SMSRequest) Descriptor() ([]byte, []int)
```

#### func (*SMSRequest) GetBody

```go
func (m *SMSRequest) GetBody() string
```

#### func (*SMSRequest) GetUserId

```go
func (m *SMSRequest) GetUserId() string
```

#### func (*SMSRequest) ProtoMessage

```go
func (*SMSRequest) ProtoMessage()
```

#### func (*SMSRequest) Reset

```go
func (m *SMSRequest) Reset()
```

#### func (*SMSRequest) String

```go
func (m *SMSRequest) String() string
```

#### func (*SMSRequest) XXX_DiscardUnknown

```go
func (m *SMSRequest) XXX_DiscardUnknown()
```

#### func (*SMSRequest) XXX_Marshal

```go
func (m *SMSRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*SMSRequest) XXX_Merge

```go
func (m *SMSRequest) XXX_Merge(src proto.Message)
```

#### func (*SMSRequest) XXX_Size

```go
func (m *SMSRequest) XXX_Size() int
```

#### func (*SMSRequest) XXX_Unmarshal

```go
func (m *SMSRequest) XXX_Unmarshal(b []byte) error
```

#### type Server

```go
type Server struct {
	Addr string
	UserServiceServer
	driver.PluginFunc
}
```


#### func  NewServer

```go
func NewServer(addr string, server UserServiceServer) *Server
```

#### func (*Server) Serve

```go
func (s *Server) Serve(debug bool) error
```

#### type SignedKey

```go
type SignedKey struct {
	SignedKey            string   `protobuf:"bytes,1,opt,name=signed_key,json=signedKey,proto3" json:"signed_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*SignedKey) Descriptor

```go
func (*SignedKey) Descriptor() ([]byte, []int)
```

#### func (*SignedKey) GetSignedKey

```go
func (m *SignedKey) GetSignedKey() string
```

#### func (*SignedKey) ProtoMessage

```go
func (*SignedKey) ProtoMessage()
```

#### func (*SignedKey) Reset

```go
func (m *SignedKey) Reset()
```

#### func (*SignedKey) String

```go
func (m *SignedKey) String() string
```

#### func (*SignedKey) XXX_DiscardUnknown

```go
func (m *SignedKey) XXX_DiscardUnknown()
```

#### func (*SignedKey) XXX_Marshal

```go
func (m *SignedKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*SignedKey) XXX_Merge

```go
func (m *SignedKey) XXX_Merge(src proto.Message)
```

#### func (*SignedKey) XXX_Size

```go
func (m *SignedKey) XXX_Size() int
```

#### func (*SignedKey) XXX_Unmarshal

```go
func (m *SignedKey) XXX_Unmarshal(b []byte) error
```

#### type SigningMethod

```go
type SigningMethod int32
```


```go
const (
	SigningMethod_HMAC   SigningMethod = 0
	SigningMethod_ECDSA  SigningMethod = 1
	SigningMethod_RSA    SigningMethod = 2
	SigningMethod_RSAPPS SigningMethod = 3
)
```

#### func (SigningMethod) EnumDescriptor

```go
func (SigningMethod) EnumDescriptor() ([]byte, []int)
```

#### func (SigningMethod) String

```go
func (x SigningMethod) String() string
```

#### type StandardClaims

```go
type StandardClaims struct {
	Access               *Access  `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	Audience             string   `protobuf:"bytes,2,opt,name=audience,proto3" json:"audience,omitempty"`
	Subject              string   `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	ExpiresAt            int64    `protobuf:"varint,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Id                   string   `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	IssuedAt             int64    `protobuf:"varint,6,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	NotBefore            int64    `protobuf:"varint,7,opt,name=not_before,json=notBefore,proto3" json:"not_before,omitempty"`
	Grants               []Grant  `protobuf:"varint,8,rep,packed,name=grants,proto3,enum=api.Grant" json:"grants,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*StandardClaims) AddGrants

```go
func (t *StandardClaims) AddGrants(grants ...Grant)
```

#### func (*StandardClaims) Claims

```go
func (t *StandardClaims) Claims() *Claims
```

#### func (*StandardClaims) Descriptor

```go
func (*StandardClaims) Descriptor() ([]byte, []int)
```

#### func (*StandardClaims) GetAccess

```go
func (m *StandardClaims) GetAccess() *Access
```

#### func (*StandardClaims) GetAudience

```go
func (m *StandardClaims) GetAudience() string
```

#### func (*StandardClaims) GetExpiresAt

```go
func (m *StandardClaims) GetExpiresAt() int64
```

#### func (*StandardClaims) GetGrants

```go
func (m *StandardClaims) GetGrants() []Grant
```

#### func (*StandardClaims) GetId

```go
func (m *StandardClaims) GetId() string
```

#### func (*StandardClaims) GetIssuedAt

```go
func (m *StandardClaims) GetIssuedAt() int64
```

#### func (*StandardClaims) GetNotBefore

```go
func (m *StandardClaims) GetNotBefore() int64
```

#### func (*StandardClaims) GetSubject

```go
func (m *StandardClaims) GetSubject() string
```

#### func (*StandardClaims) ProtoMessage

```go
func (*StandardClaims) ProtoMessage()
```

#### func (*StandardClaims) Reset

```go
func (m *StandardClaims) Reset()
```

#### func (*StandardClaims) String

```go
func (m *StandardClaims) String() string
```

#### func (*StandardClaims) XXX_DiscardUnknown

```go
func (m *StandardClaims) XXX_DiscardUnknown()
```

#### func (*StandardClaims) XXX_Marshal

```go
func (m *StandardClaims) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*StandardClaims) XXX_Merge

```go
func (m *StandardClaims) XXX_Merge(src proto.Message)
```

#### func (*StandardClaims) XXX_Size

```go
func (m *StandardClaims) XXX_Size() int
```

#### func (*StandardClaims) XXX_Unmarshal

```go
func (m *StandardClaims) XXX_Unmarshal(b []byte) error
```

#### type Star

```go
type Star struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Item                 *ItemRef `protobuf:"bytes,4,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Star) Descriptor

```go
func (*Star) Descriptor() ([]byte, []int)
```

#### func (*Star) GetItem

```go
func (m *Star) GetItem() *ItemRef
```

#### func (*Star) GetText

```go
func (m *Star) GetText() string
```

#### func (*Star) ProtoMessage

```go
func (*Star) ProtoMessage()
```

#### func (*Star) Reset

```go
func (m *Star) Reset()
```

#### func (*Star) String

```go
func (m *Star) String() string
```

#### func (*Star) XXX_DiscardUnknown

```go
func (m *Star) XXX_DiscardUnknown()
```

#### func (*Star) XXX_Marshal

```go
func (m *Star) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Star) XXX_Merge

```go
func (m *Star) XXX_Merge(src proto.Message)
```

#### func (*Star) XXX_Size

```go
func (m *Star) XXX_Size() int
```

#### func (*Star) XXX_Unmarshal

```go
func (m *Star) XXX_Unmarshal(b []byte) error
```

#### type SubscribeCustomerRequest

```go
type SubscribeCustomerRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Plan                 string   `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	CardNumber           string   `protobuf:"bytes,3,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	ExpMonth             string   `protobuf:"bytes,4,opt,name=exp_month,json=expMonth,proto3" json:"exp_month,omitempty"`
	ExpYear              string   `protobuf:"bytes,5,opt,name=exp_year,json=expYear,proto3" json:"exp_year,omitempty"`
	Cvc                  string   `protobuf:"bytes,6,opt,name=cvc,proto3" json:"cvc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*SubscribeCustomerRequest) Descriptor

```go
func (*SubscribeCustomerRequest) Descriptor() ([]byte, []int)
```

#### func (*SubscribeCustomerRequest) GetCardNumber

```go
func (m *SubscribeCustomerRequest) GetCardNumber() string
```

#### func (*SubscribeCustomerRequest) GetCvc

```go
func (m *SubscribeCustomerRequest) GetCvc() string
```

#### func (*SubscribeCustomerRequest) GetEmail

```go
func (m *SubscribeCustomerRequest) GetEmail() string
```

#### func (*SubscribeCustomerRequest) GetExpMonth

```go
func (m *SubscribeCustomerRequest) GetExpMonth() string
```

#### func (*SubscribeCustomerRequest) GetExpYear

```go
func (m *SubscribeCustomerRequest) GetExpYear() string
```

#### func (*SubscribeCustomerRequest) GetPlan

```go
func (m *SubscribeCustomerRequest) GetPlan() string
```

#### func (*SubscribeCustomerRequest) ProtoMessage

```go
func (*SubscribeCustomerRequest) ProtoMessage()
```

#### func (*SubscribeCustomerRequest) Reset

```go
func (m *SubscribeCustomerRequest) Reset()
```

#### func (*SubscribeCustomerRequest) String

```go
func (m *SubscribeCustomerRequest) String() string
```

#### func (*SubscribeCustomerRequest) XXX_DiscardUnknown

```go
func (m *SubscribeCustomerRequest) XXX_DiscardUnknown()
```

#### func (*SubscribeCustomerRequest) XXX_Marshal

```go
func (m *SubscribeCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*SubscribeCustomerRequest) XXX_Merge

```go
func (m *SubscribeCustomerRequest) XXX_Merge(src proto.Message)
```

#### func (*SubscribeCustomerRequest) XXX_Size

```go
func (m *SubscribeCustomerRequest) XXX_Size() int
```

#### func (*SubscribeCustomerRequest) XXX_Unmarshal

```go
func (m *SubscribeCustomerRequest) XXX_Unmarshal(b []byte) error
```

#### type SubscribeCustomerResponse

```go
type SubscribeCustomerResponse struct {
	SubscriptionId       string   `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*SubscribeCustomerResponse) Descriptor

```go
func (*SubscribeCustomerResponse) Descriptor() ([]byte, []int)
```

#### func (*SubscribeCustomerResponse) GetSubscriptionId

```go
func (m *SubscribeCustomerResponse) GetSubscriptionId() string
```

#### func (*SubscribeCustomerResponse) ProtoMessage

```go
func (*SubscribeCustomerResponse) ProtoMessage()
```

#### func (*SubscribeCustomerResponse) Reset

```go
func (m *SubscribeCustomerResponse) Reset()
```

#### func (*SubscribeCustomerResponse) String

```go
func (m *SubscribeCustomerResponse) String() string
```

#### func (*SubscribeCustomerResponse) XXX_DiscardUnknown

```go
func (m *SubscribeCustomerResponse) XXX_DiscardUnknown()
```

#### func (*SubscribeCustomerResponse) XXX_Marshal

```go
func (m *SubscribeCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*SubscribeCustomerResponse) XXX_Merge

```go
func (m *SubscribeCustomerResponse) XXX_Merge(src proto.Message)
```

#### func (*SubscribeCustomerResponse) XXX_Size

```go
func (m *SubscribeCustomerResponse) XXX_Size() int
```

#### func (*SubscribeCustomerResponse) XXX_Unmarshal

```go
func (m *SubscribeCustomerResponse) XXX_Unmarshal(b []byte) error
```

#### type Token

```go
type Token struct {
	Raw                  string            `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	SigningMethod        SigningMethod     `protobuf:"varint,2,opt,name=signing_method,json=signingMethod,proto3,enum=api.SigningMethod" json:"signing_method,omitempty"`
	Valid                bool              `protobuf:"varint,3,opt,name=valid,proto3" json:"valid,omitempty"`
	Signature            string            `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	Header               map[string]string `protobuf:"bytes,5,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func (*Token) Descriptor

```go
func (*Token) Descriptor() ([]byte, []int)
```

#### func (*Token) GetHeader

```go
func (m *Token) GetHeader() map[string]string
```

#### func (*Token) GetRaw

```go
func (m *Token) GetRaw() string
```

#### func (*Token) GetSignature

```go
func (m *Token) GetSignature() string
```

#### func (*Token) GetSigningMethod

```go
func (m *Token) GetSigningMethod() SigningMethod
```

#### func (*Token) GetValid

```go
func (m *Token) GetValid() bool
```

#### func (*Token) ProtoMessage

```go
func (*Token) ProtoMessage()
```

#### func (*Token) Reset

```go
func (m *Token) Reset()
```

#### func (*Token) String

```go
func (m *Token) String() string
```

#### func (*Token) XXX_DiscardUnknown

```go
func (m *Token) XXX_DiscardUnknown()
```

#### func (*Token) XXX_Marshal

```go
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Token) XXX_Merge

```go
func (m *Token) XXX_Merge(src proto.Message)
```

#### func (*Token) XXX_Size

```go
func (m *Token) XXX_Size() int
```

#### func (*Token) XXX_Unmarshal

```go
func (m *Token) XXX_Unmarshal(b []byte) error
```

#### type UserReminder

```go
type UserReminder struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Time                 string   `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Item                 *ItemRef `protobuf:"bytes,4,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*UserReminder) Descriptor

```go
func (*UserReminder) Descriptor() ([]byte, []int)
```

#### func (*UserReminder) GetItem

```go
func (m *UserReminder) GetItem() *ItemRef
```

#### func (*UserReminder) GetText

```go
func (m *UserReminder) GetText() string
```

#### func (*UserReminder) GetTime

```go
func (m *UserReminder) GetTime() string
```

#### func (*UserReminder) GetUserId

```go
func (m *UserReminder) GetUserId() string
```

#### func (*UserReminder) ProtoMessage

```go
func (*UserReminder) ProtoMessage()
```

#### func (*UserReminder) Reset

```go
func (m *UserReminder) Reset()
```

#### func (*UserReminder) String

```go
func (m *UserReminder) String() string
```

#### func (*UserReminder) XXX_DiscardUnknown

```go
func (m *UserReminder) XXX_DiscardUnknown()
```

#### func (*UserReminder) XXX_Marshal

```go
func (m *UserReminder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*UserReminder) XXX_Merge

```go
func (m *UserReminder) XXX_Merge(src proto.Message)
```

#### func (*UserReminder) XXX_Size

```go
func (m *UserReminder) XXX_Size() int
```

#### func (*UserReminder) XXX_Unmarshal

```go
func (m *UserReminder) XXX_Unmarshal(b []byte) error
```

#### type UserServiceClient

```go
type UserServiceClient interface {
	CreateCustomer(ctx context.Context, in *AddCustomerRequest, opts ...grpc.CallOption) (*Customer, error)
	SubscribeCustomer(ctx context.Context, in *SubscribeCustomerRequest, opts ...grpc.CallOption) (*SubscribeCustomerResponse, error)
	UnSubscribeCustomer(ctx context.Context, in *CancelSubscriptionRequest, opts ...grpc.CallOption) (*Empty, error)
	CreateSubscriptionPlan(ctx context.Context, in *CreatePlanRequest, opts ...grpc.CallOption) (*CreatePlanResponse, error)
	SMSCustomer(ctx context.Context, in *SMSRequest, opts ...grpc.CallOption) (*Empty, error)
	CallCustomer(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*Empty, error)
	MMSCustomer(ctx context.Context, in *MMSRequest, opts ...grpc.CallOption) (*Empty, error)
	EmailCustomer(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*Empty, error)
	SMSUser(ctx context.Context, in *SMSRequest, opts ...grpc.CallOption) (*Empty, error)
	CallUser(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*Empty, error)
	MMSUser(ctx context.Context, in *MMSRequest, opts ...grpc.CallOption) (*Empty, error)
	EmailUser(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*Empty, error)
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
	CreateCustomer(context.Context, *AddCustomerRequest) (*Customer, error)
	SubscribeCustomer(context.Context, *SubscribeCustomerRequest) (*SubscribeCustomerResponse, error)
	UnSubscribeCustomer(context.Context, *CancelSubscriptionRequest) (*Empty, error)
	CreateSubscriptionPlan(context.Context, *CreatePlanRequest) (*CreatePlanResponse, error)
	SMSCustomer(context.Context, *SMSRequest) (*Empty, error)
	CallCustomer(context.Context, *CallRequest) (*Empty, error)
	MMSCustomer(context.Context, *MMSRequest) (*Empty, error)
	EmailCustomer(context.Context, *EmailRequest) (*Empty, error)
	SMSUser(context.Context, *SMSRequest) (*Empty, error)
	CallUser(context.Context, *CallRequest) (*Empty, error)
	MMSUser(context.Context, *MMSRequest) (*Empty, error)
	EmailUser(context.Context, *EmailRequest) (*Empty, error)
}
```

UserServiceServer is the server API for UserService service.
