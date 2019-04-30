# common
--
    import "github.com/autom8ter/api/common"


## Usage

```go
var (
	EnvContext context.Context
)
```

#### func  FuncMap

```go
func FuncMap() template.FuncMap
```

#### func  ToError

```go
func ToError(err error, msg string) error
```

#### type Config

```go
type Config struct {
	ClientId             *String      `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         *String      `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	TokenUrl             *String      `protobuf:"bytes,3,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	AuthUrl              *String      `protobuf:"bytes,4,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
	Scopes               *StringArray `protobuf:"bytes,5,opt,name=scopes,proto3" json:"scopes,omitempty"`
	Redirect             *String      `protobuf:"bytes,6,opt,name=redirect,proto3" json:"redirect,omitempty"`
	EndpointParams       *StringMap   `protobuf:"bytes,7,opt,name=endpoint_params,json=endpointParams,proto3" json:"endpoint_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}
```


#### func (*Config) ClientCredentials

```go
func (c *Config) ClientCredentials() *clientcredentials.Config
```

#### func (*Config) Debugf

```go
func (s *Config) Debugf(format string)
```

#### func (*Config) Descriptor

```go
func (*Config) Descriptor() ([]byte, []int)
```

#### func (*Config) GetAuthUrl

```go
func (m *Config) GetAuthUrl() *String
```

#### func (*Config) GetClientId

```go
func (m *Config) GetClientId() *String
```

#### func (*Config) GetClientSecret

```go
func (m *Config) GetClientSecret() *String
```

#### func (*Config) GetEndpointParams

```go
func (m *Config) GetEndpointParams() *StringMap
```

#### func (*Config) GetRedirect

```go
func (m *Config) GetRedirect() *String
```

#### func (*Config) GetScopes

```go
func (m *Config) GetScopes() *StringArray
```

#### func (*Config) GetToken

```go
func (c *Config) GetToken(tok *oauth2.Token) *Token
```

#### func (*Config) GetTokenUrl

```go
func (m *Config) GetTokenUrl() *String
```

#### func (*Config) JSONString

```go
func (s *Config) JSONString() *String
```

#### func (*Config) ProtoMessage

```go
func (*Config) ProtoMessage()
```

#### func (*Config) Reset

```go
func (m *Config) Reset()
```

#### func (*Config) String

```go
func (m *Config) String() string
```

#### func (*Config) ToContext

```go
func (s *Config) ToContext(ctx context.Context, key string) context.Context
```

#### func (*Config) Token

```go
func (c *Config) Token(code string) (*Token, error)
```

#### func (*Config) UnmarshalJSONFrom

```go
func (s *Config) UnmarshalJSONFrom(b []byte) error
```

#### func (*Config) UnmarshalProtoFrom

```go
func (s *Config) UnmarshalProtoFrom(b []byte) error
```

#### func (*Config) XXX_DiscardUnknown

```go
func (m *Config) XXX_DiscardUnknown()
```

#### func (*Config) XXX_Marshal

```go
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Config) XXX_Merge

```go
func (m *Config) XXX_Merge(src proto.Message)
```

#### func (*Config) XXX_Size

```go
func (m *Config) XXX_Size() int
```

#### func (*Config) XXX_Unmarshal

```go
func (m *Config) XXX_Unmarshal(b []byte) error
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

#### type Event

```go
type Event struct {
	Date                 *String  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Type                 *String  `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ClientId             *String  `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientName           *String  `protobuf:"bytes,4,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	Ip                   *String  `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	LocationInfo         *String  `protobuf:"bytes,6,opt,name=location_info,json=locationInfo,proto3" json:"location_info,omitempty"`
	Details              *String  `protobuf:"bytes,7,opt,name=details,proto3" json:"details,omitempty"`
	UserId               *String  `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Event) Descriptor

```go
func (*Event) Descriptor() ([]byte, []int)
```

#### func (*Event) GetClientId

```go
func (m *Event) GetClientId() *String
```

#### func (*Event) GetClientName

```go
func (m *Event) GetClientName() *String
```

#### func (*Event) GetDate

```go
func (m *Event) GetDate() *String
```

#### func (*Event) GetDetails

```go
func (m *Event) GetDetails() *String
```

#### func (*Event) GetIp

```go
func (m *Event) GetIp() *String
```

#### func (*Event) GetLocationInfo

```go
func (m *Event) GetLocationInfo() *String
```

#### func (*Event) GetType

```go
func (m *Event) GetType() *String
```

#### func (*Event) GetUserId

```go
func (m *Event) GetUserId() *String
```

#### func (*Event) JSONString

```go
func (s *Event) JSONString() *String
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

#### type Float64

```go
type Float64 struct {
	Num                  float64  `protobuf:"fixed64,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  FloatFromThePowerOf10

```go
func FloatFromThePowerOf10(i *Int64) *Float64
```

#### func  ToFloat64

```go
func ToFloat64(i float64) *Float64
```

#### func  ToFloats

```go
func ToFloats(floats ...float64) []*Float64
```

#### func (*Float64) Abs

```go
func (m *Float64) Abs() *Float64
```
Abs returns the absolute value of Float64.

#### func (*Float64) BinaryExponent

```go
func (m *Float64) BinaryExponent() *Int64
```

#### func (*Float64) Ceiling

```go
func (m *Float64) Ceiling() *Float64
```
Ceil returns the least integer value greater than or equal to x.

#### func (*Float64) CubeRoot

```go
func (m *Float64) CubeRoot() *Float64
```
CubeRoot returns the cube root of x.

#### func (*Float64) Cubed

```go
func (s *Float64) Cubed() *Float64
```

#### func (*Float64) Debugf

```go
func (s *Float64) Debugf(format string)
```

#### func (*Float64) DeepEqual

```go
func (s *Float64) DeepEqual(y interface{}) bool
```

#### func (*Float64) Descriptor

```go
func (*Float64) Descriptor() ([]byte, []int)
```

#### func (*Float64) DividedBy

```go
func (m *Float64) DividedBy(n *Float64) *Float64
```

#### func (*Float64) Exp

```go
func (m *Float64) Exp() *Float64
```

#### func (*Float64) Exp2

```go
func (m *Float64) Exp2() *Float64
```

#### func (*Float64) Expm1

```go
func (m *Float64) Expm1() *Float64
```

#### func (*Float64) Gamma

```go
func (m *Float64) Gamma() *Float64
```

#### func (*Float64) GetNum

```go
func (m *Float64) GetNum() float64
```

#### func (*Float64) Humanized

```go
func (m *Float64) Humanized() *String
```

#### func (*Float64) JSONString

```go
func (s *Float64) JSONString() *String
```

#### func (*Float64) Log

```go
func (m *Float64) Log() *Float64
```

#### func (*Float64) Log10

```go
func (m *Float64) Log10() *Float64
```

#### func (*Float64) Log1p

```go
func (m *Float64) Log1p() *Float64
```

#### func (*Float64) Log2

```go
func (m *Float64) Log2() *Float64
```

#### func (*Float64) Logb

```go
func (m *Float64) Logb() *Float64
```

#### func (*Float64) Max

```go
func (m *Float64) Max(f *Float64) *Float64
```

#### func (*Float64) Minus

```go
func (m *Float64) Minus(n *Float64) *Float64
```

#### func (*Float64) Negative

```go
func (m *Float64) Negative() *Float64
```

#### func (*Float64) NotANumber

```go
func (m *Float64) NotANumber() bool
```

#### func (*Float64) Plus

```go
func (m *Float64) Plus(n *Float64) *Float64
```

#### func (*Float64) Pointer

```go
func (s *Float64) Pointer() *float64
```

#### func (*Float64) ProtoMessage

```go
func (*Float64) ProtoMessage()
```

#### func (*Float64) Remainder

```go
func (m *Float64) Remainder(f *Float64) *Float64
```

#### func (*Float64) Reset

```go
func (m *Float64) Reset()
```

#### func (*Float64) Round

```go
func (m *Float64) Round() *Float64
```

#### func (*Float64) RoundToEven

```go
func (m *Float64) RoundToEven() *Float64
```

#### func (*Float64) Sin

```go
func (m *Float64) Sin() *Float64
```

#### func (*Float64) SinCos

```go
func (m *Float64) SinCos() (*Float64, *Float64)
```

#### func (*Float64) SqrtofCombinedSquared

```go
func (m *Float64) SqrtofCombinedSquared(f *Float64) *Float64
```

#### func (*Float64) SquareRoot

```go
func (m *Float64) SquareRoot() *Float64
```

#### func (*Float64) Squared

```go
func (s *Float64) Squared() *Float64
```

#### func (*Float64) String

```go
func (m *Float64) String() string
```

#### func (*Float64) StringWithUnit

```go
func (s *Float64) StringWithUnit(unit *String) *String
```

#### func (*Float64) Tan

```go
func (m *Float64) Tan() *Float64
```

#### func (*Float64) Times

```go
func (m *Float64) Times(n *Float64) *Float64
```

#### func (*Float64) ToContext

```go
func (s *Float64) ToContext(ctx context.Context, key string) context.Context
```

#### func (*Float64) ToString

```go
func (s *Float64) ToString() *String
```

#### func (*Float64) ToThePowerOf

```go
func (m *Float64) ToThePowerOf(f *Float64) *Float64
```

#### func (*Float64) TypeMatches

```go
func (s *Float64) TypeMatches(src interface{}) bool
```

#### func (*Float64) UnmarshalJSONFrom

```go
func (s *Float64) UnmarshalJSONFrom(bits []byte) error
```

#### func (*Float64) UnmarshalProtoFrom

```go
func (s *Float64) UnmarshalProtoFrom(bits []byte) error
```

#### func (*Float64) Validate

```go
func (s *Float64) Validate(fn func(s *Float64) error) error
```

#### func (*Float64) XXX_DiscardUnknown

```go
func (m *Float64) XXX_DiscardUnknown()
```

#### func (*Float64) XXX_Marshal

```go
func (m *Float64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Float64) XXX_Merge

```go
func (m *Float64) XXX_Merge(src proto.Message)
```

#### func (*Float64) XXX_Size

```go
func (m *Float64) XXX_Size() int
```

#### func (*Float64) XXX_Unmarshal

```go
func (m *Float64) XXX_Unmarshal(b []byte) error
```

#### func (*Float64) Zero

```go
func (m *Float64) Zero()
```

#### type Identifier

```go
type Identifier struct {
	Id                   *String  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  EmailIdentifier

```go
func EmailIdentifier(email string) *Identifier
```

#### func  GenericIdentifier

```go
func GenericIdentifier(id string) *Identifier
```

#### func  NickNameIdentifier

```go
func NickNameIdentifier(id string) *Identifier
```

#### func  PhoneIdentifier

```go
func PhoneIdentifier(phone string) *Identifier
```

#### func (*Identifier) Descriptor

```go
func (*Identifier) Descriptor() ([]byte, []int)
```

#### func (*Identifier) GetId

```go
func (m *Identifier) GetId() *String
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

#### type Int64

```go
type Int64 struct {
	Num                  int64    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  ToInt64

```go
func ToInt64(i int) *Int64
```

#### func (*Int64) Debugf

```go
func (s *Int64) Debugf(format string)
```

#### func (*Int64) DeepEqual

```go
func (s *Int64) DeepEqual(y interface{}) bool
```

#### func (*Int64) Descriptor

```go
func (*Int64) Descriptor() ([]byte, []int)
```

#### func (*Int64) DividedBy

```go
func (m *Int64) DividedBy(n *Int64) *Int64
```

#### func (*Int64) GetNum

```go
func (m *Int64) GetNum() int64
```

#### func (*Int64) Humanized

```go
func (m *Int64) Humanized() *String
```

#### func (*Int64) Int

```go
func (i *Int64) Int() int
```

#### func (*Int64) JSONString

```go
func (s *Int64) JSONString() *String
```

#### func (*Int64) Minus

```go
func (m *Int64) Minus(n *Int64) *Int64
```

#### func (*Int64) Plus

```go
func (m *Int64) Plus(n *Int64) *Int64
```

#### func (*Int64) Pointer

```go
func (s *Int64) Pointer() *int64
```

#### func (*Int64) ProtoMessage

```go
func (*Int64) ProtoMessage()
```

#### func (*Int64) Remainder

```go
func (m *Int64) Remainder(n *Int64) *Int64
```

#### func (*Int64) Reset

```go
func (m *Int64) Reset()
```

#### func (*Int64) String

```go
func (m *Int64) String() string
```

#### func (*Int64) Times

```go
func (m *Int64) Times(n *Int64) *Int64
```

#### func (*Int64) ToContext

```go
func (s *Int64) ToContext(ctx context.Context, key string) context.Context
```

#### func (*Int64) ToString

```go
func (s *Int64) ToString() *String
```

#### func (*Int64) TypeMatches

```go
func (s *Int64) TypeMatches(src interface{}) bool
```

#### func (*Int64) UnmarshalJSONFrom

```go
func (s *Int64) UnmarshalJSONFrom(bits []byte) error
```

#### func (*Int64) UnmarshalProtoFrom

```go
func (s *Int64) UnmarshalProtoFrom(bits []byte) error
```

#### func (*Int64) Validate

```go
func (s *Int64) Validate(fn func(s *Int64) error) error
```

#### func (*Int64) XXX_DiscardUnknown

```go
func (m *Int64) XXX_DiscardUnknown()
```

#### func (*Int64) XXX_Marshal

```go
func (m *Int64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Int64) XXX_Merge

```go
func (m *Int64) XXX_Merge(src proto.Message)
```

#### func (*Int64) XXX_Size

```go
func (m *Int64) XXX_Size() int
```

#### func (*Int64) XXX_Unmarshal

```go
func (m *Int64) XXX_Unmarshal(b []byte) error
```

#### type Query

```go
type Query struct {
	Query                *String  `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	Fields               *String  `protobuf:"bytes,5,opt,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Query) Descriptor

```go
func (*Query) Descriptor() ([]byte, []int)
```

#### func (*Query) GetFields

```go
func (m *Query) GetFields() *String
```

#### func (*Query) GetQuery

```go
func (m *Query) GetQuery() *String
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

#### type String

```go
type String struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  JSONObjToString

```go
func JSONObjToString(b interface{}) *String
```

#### func  MessageToJSONString

```go
func MessageToJSONString(msg proto.Message) *String
```

#### func  RandomString

```go
func RandomString() *String
```

#### func  StringFromEnv

```go
func StringFromEnv(key string) *String
```

#### func  StringFromFile

```go
func StringFromFile(filename string) (*String, error)
```

#### func  StringFromPrompt

```go
func StringFromPrompt(prompt string) *String
```

#### func  ToString

```go
func ToString(s string) *String
```

#### func (*String) AddLine

```go
func (s *String) AddLine(text string)
```

#### func (*String) Adler32

```go
func (s *String) Adler32() string
```

#### func (*String) Append

```go
func (s *String) Append(text string)
```

#### func (*String) AsAdler32

```go
func (s *String) AsAdler32()
```

#### func (*String) AsBase64Decode

```go
func (m *String) AsBase64Decode()
```

#### func (*String) AsBase64Encoded

```go
func (m *String) AsBase64Encoded()
```

#### func (*String) AsHash

```go
func (m *String) AsHash() error
```

#### func (*String) AsSha1

```go
func (s *String) AsSha1()
```

#### func (*String) AsSha256

```go
func (s *String) AsSha256()
```

#### func (*String) AsTemplate

```go
func (m *String) AsTemplate(name string) (*template.Template, error)
```

#### func (*String) Base64Encode

```go
func (m *String) Base64Encode() string
```

#### func (*String) Buffer

```go
func (b *String) Buffer() *bytes.Buffer
```

#### func (*String) BufferedReader

```go
func (b *String) BufferedReader() *bufio.Reader
```

#### func (*String) BufferedWriter

```go
func (b *String) BufferedWriter() *bufio.Writer
```

#### func (*String) Contains

```go
func (s *String) Contains(sub string) bool
```

#### func (*String) Debugf

```go
func (s *String) Debugf(format string)
```

#### func (*String) Descriptor

```go
func (*String) Descriptor() ([]byte, []int)
```

#### func (*String) ExecuteAsBashCMD

```go
func (m *String) ExecuteAsBashCMD() *String
```

#### func (*String) ExecuteAsPython3

```go
func (m *String) ExecuteAsPython3() *String
```

#### func (*String) ExecuteAsShellCMD

```go
func (m *String) ExecuteAsShellCMD() *String
```

#### func (*String) GetText

```go
func (m *String) GetText() string
```

#### func (*String) HashMatchesPassword

```go
func (m *String) HashMatchesPassword(hash string) error
```

#### func (*String) Hashed

```go
func (m *String) Hashed() (string, error)
```

#### func (*String) Indent

```go
func (s *String) Indent(spaces *Int64) string
```

#### func (*String) Index

```go
func (s *String) Index(sub string) *Int64
```

#### func (*String) IsEmpty

```go
func (s *String) IsEmpty() bool
```

#### func (*String) IsTemplate

```go
func (m *String) IsTemplate() bool
```

#### func (*String) JSONString

```go
func (s *String) JSONString() *String
```

#### func (*String) Matches

```go
func (s *String) Matches(this string) bool
```

#### func (*String) ParseDuration

```go
func (b *String) ParseDuration() time.Duration
```

#### func (*String) ParseLanguage

```go
func (s *String) ParseLanguage() (language.Tag, error)
```

#### func (*String) ParseRegion

```go
func (s *String) ParseRegion() (language.Region, error)
```

#### func (*String) ParseScientificUnits

```go
func (s *String) ParseScientificUnits() (*Float64, *String, error)
```

#### func (*String) ParseTime

```go
func (b *String) ParseTime(str *String) time.Time
```

#### func (*String) ParseURL

```go
func (s *String) ParseURL() (*url.URL, error)
```

#### func (*String) PasswordMatchesHashed

```go
func (m *String) PasswordMatchesHashed(pass string) error
```

#### func (*String) Pointer

```go
func (s *String) Pointer() *string
```

#### func (*String) PrePend

```go
func (s *String) PrePend(text string)
```

#### func (*String) Println

```go
func (s *String) Println()
```

#### func (*String) ProtoMessage

```go
func (*String) ProtoMessage()
```

#### func (*String) RegExReplaceAll

```go
func (s *String) RegExReplaceAll(reg string, replaceWith string) *String
```

#### func (*String) RegExReplaceAllLiteral

```go
func (s *String) RegExReplaceAllLiteral(reg string, replaceWith string) *String
```

#### func (*String) RegExSplit

```go
func (s *String) RegExSplit(reg string, num int) *StringArray
```

#### func (*String) RegexFind

```go
func (s *String) RegexFind(reg string) *String
```

#### func (*String) RegexFindAll

```go
func (s *String) RegexFindAll(reg string, num int) *StringArray
```

#### func (*String) RegexMatches

```go
func (s *String) RegexMatches(reg string) bool
```

#### func (*String) Render

```go
func (m *String) Render(name string, w io.Writer, data interface{}) error
```

#### func (*String) Replace

```go
func (s *String) Replace(oldNew ...string)
```

#### func (*String) Reset

```go
func (m *String) Reset()
```

#### func (*String) SetEnv

```go
func (s *String) SetEnv(key string) error
```

#### func (*String) Sha1

```go
func (s *String) Sha1() string
```

#### func (*String) Sha256

```go
func (s *String) Sha256() string
```

#### func (*String) Split

```go
func (s *String) Split(sep *String) *StringArray
```

#### func (*String) StartArray

```go
func (s *String) StartArray() *StringArray
```

#### func (*String) String

```go
func (m *String) String() string
```

#### func (*String) StringReader

```go
func (b *String) StringReader() *strings.Reader
```

#### func (*String) ToContext

```go
func (s *String) ToContext(ctx context.Context, key string) context.Context
```

#### func (*String) ToInt64

```go
func (s *String) ToInt64() (*Int64, error)
```

#### func (*String) ToLower

```go
func (s *String) ToLower()
```

#### func (*String) ToStringArray

```go
func (s *String) ToStringArray() (*StringArray, error)
```

#### func (*String) ToTitle

```go
func (s *String) ToTitle()
```

#### func (*String) ToUpper

```go
func (s *String) ToUpper()
```

#### func (*String) TrimPrefix

```go
func (s *String) TrimPrefix(sub string)
```

#### func (*String) TrimSuffix

```go
func (s *String) TrimSuffix(sub string)
```

#### func (*String) TrimWhiteSpace

```go
func (s *String) TrimWhiteSpace()
```

#### func (*String) TypeMatches

```go
func (s *String) TypeMatches(src interface{}) bool
```

#### func (*String) UnmarshalJSONFrom

```go
func (s *String) UnmarshalJSONFrom(bits []byte) error
```

#### func (*String) Validate

```go
func (s *String) Validate(fn func(s *String) error) error
```

#### func (*String) WriteString

```go
func (m *String) WriteString() http.HandlerFunc
```

#### func (*String) XXX_DiscardUnknown

```go
func (m *String) XXX_DiscardUnknown()
```

#### func (*String) XXX_Marshal

```go
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*String) XXX_Merge

```go
func (m *String) XXX_Merge(src proto.Message)
```

#### func (*String) XXX_Size

```go
func (m *String) XXX_Size() int
```

#### func (*String) XXX_Unmarshal

```go
func (m *String) XXX_Unmarshal(b []byte) error
```

#### type StringArray

```go
type StringArray struct {
	Strings              []*String `protobuf:"bytes,1,rep,name=strings,proto3" json:"strings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}
```


#### func  ENVFromContext

```go
func ENVFromContext(ctx context.Context) *StringArray
```

#### func  StringArrayFromEnv

```go
func StringArrayFromEnv() *StringArray
```

#### func  ToStringArray

```go
func ToStringArray(s []string) *StringArray
```

#### func (*StringArray) Append

```go
func (m *StringArray) Append(vals ...*String)
```

#### func (*StringArray) Array

```go
func (s *StringArray) Array() []string
```

#### func (*StringArray) Debugf

```go
func (s *StringArray) Debugf(format string)
```

#### func (*StringArray) DeepEqual

```go
func (s *StringArray) DeepEqual(y interface{}) bool
```

#### func (*StringArray) Descriptor

```go
func (*StringArray) Descriptor() ([]byte, []int)
```

#### func (*StringArray) First

```go
func (m *StringArray) First() *String
```

#### func (*StringArray) Get

```go
func (m *StringArray) Get(index int) *String
```

#### func (*StringArray) GetStrings

```go
func (m *StringArray) GetStrings() []*String
```

#### func (*StringArray) IsEmpty

```go
func (m *StringArray) IsEmpty() bool
```

#### func (*StringArray) JSONString

```go
func (s *StringArray) JSONString() *String
```

#### func (*StringArray) Last

```go
func (m *StringArray) Last() *String
```

#### func (*StringArray) Length

```go
func (m *StringArray) Length() int
```

#### func (*StringArray) Pointer

```go
func (s *StringArray) Pointer() []*string
```

#### func (*StringArray) ProtoMessage

```go
func (*StringArray) ProtoMessage()
```

#### func (*StringArray) Reset

```go
func (m *StringArray) Reset()
```

#### func (*StringArray) SelectRandom

```go
func (m *StringArray) SelectRandom() *String
```

#### func (*StringArray) String

```go
func (m *StringArray) String() string
```

#### func (*StringArray) StringMap

```go
func (s *StringArray) StringMap(y interface{}) bool
```

#### func (*StringArray) ToContext

```go
func (s *StringArray) ToContext(ctx context.Context, key string) context.Context
```

#### func (*StringArray) TypeMatches

```go
func (s *StringArray) TypeMatches(src interface{}) bool
```

#### func (*StringArray) UnmarshalJSONFrom

```go
func (s *StringArray) UnmarshalJSONFrom(bits []byte) error
```

#### func (*StringArray) UnmarshalProtoFrom

```go
func (s *StringArray) UnmarshalProtoFrom(bits []byte) error
```

#### func (*StringArray) Validate

```go
func (s *StringArray) Validate(fn func(s *StringArray) error) error
```

#### func (*StringArray) XXX_DiscardUnknown

```go
func (m *StringArray) XXX_DiscardUnknown()
```

#### func (*StringArray) XXX_Marshal

```go
func (m *StringArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*StringArray) XXX_Merge

```go
func (m *StringArray) XXX_Merge(src proto.Message)
```

#### func (*StringArray) XXX_Size

```go
func (m *StringArray) XXX_Size() int
```

#### func (*StringArray) XXX_Unmarshal

```go
func (m *StringArray) XXX_Unmarshal(b []byte) error
```

#### type StringMap

```go
type StringMap struct {
	StringMap            map[string]*String `protobuf:"bytes,1,rep,name=string_map,json=stringMap,proto3" json:"string_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
```


#### func  ToMap

```go
func ToMap(obj interface{}) *StringMap
```

#### func  ToStringMap

```go
func ToStringMap(s map[string]string) *StringMap
```

#### func (*StringMap) Clear

```go
func (s *StringMap) Clear(key string)
```

#### func (*StringMap) Debugf

```go
func (s *StringMap) Debugf(format string)
```

#### func (*StringMap) Descriptor

```go
func (*StringMap) Descriptor() ([]byte, []int)
```

#### func (*StringMap) Exists

```go
func (s *StringMap) Exists(key string) bool
```

#### func (*StringMap) Get

```go
func (s *StringMap) Get(key string) *String
```

#### func (*StringMap) GetStringMap

```go
func (m *StringMap) GetStringMap() map[string]*String
```

#### func (*StringMap) JSONString

```go
func (s *StringMap) JSONString() *String
```

#### func (*StringMap) Keys

```go
func (s *StringMap) Keys() *StringArray
```

#### func (*StringMap) ProtoMessage

```go
func (*StringMap) ProtoMessage()
```

#### func (*StringMap) Put

```go
func (s *StringMap) Put(key string, val *String)
```

#### func (*StringMap) Reset

```go
func (m *StringMap) Reset()
```

#### func (*StringMap) String

```go
func (m *StringMap) String() string
```

#### func (*StringMap) ToContext

```go
func (s *StringMap) ToContext(ctx context.Context, key string) context.Context
```

#### func (*StringMap) TotalKeys

```go
func (s *StringMap) TotalKeys() *Int64
```

#### func (*StringMap) TypeMatches

```go
func (s *StringMap) TypeMatches(src interface{}) bool
```

#### func (*StringMap) UnmarshalJSONFrom

```go
func (s *StringMap) UnmarshalJSONFrom(bits []byte) error
```

#### func (*StringMap) UnmarshalProtoFrom

```go
func (s *StringMap) UnmarshalProtoFrom(bits []byte) error
```

#### func (*StringMap) Validate

```go
func (s *StringMap) Validate(fn func(s *StringMap) error) error
```

#### func (*StringMap) XXX_DiscardUnknown

```go
func (m *StringMap) XXX_DiscardUnknown()
```

#### func (*StringMap) XXX_Marshal

```go
func (m *StringMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*StringMap) XXX_Merge

```go
func (m *StringMap) XXX_Merge(src proto.Message)
```

#### func (*StringMap) XXX_Size

```go
func (m *StringMap) XXX_Size() int
```

#### func (*StringMap) XXX_Unmarshal

```go
func (m *StringMap) XXX_Unmarshal(b []byte) error
```

#### type Token

```go
type Token struct {
	AccessToken          *String  `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	TokenType            *String  `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	RefreshToken         *String  `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Expiry               *String  `protobuf:"bytes,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	IdToken              *String  `protobuf:"bytes,5,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  TokenFromOAuthToken

```go
func TokenFromOAuthToken(tok *oauth2.Token) *Token
```

#### func (*Token) Debugf

```go
func (s *Token) Debugf(format string)
```

#### func (*Token) DeepEqual

```go
func (s *Token) DeepEqual(y interface{}) bool
```

#### func (*Token) Descriptor

```go
func (*Token) Descriptor() ([]byte, []int)
```

#### func (*Token) GetAccessToken

```go
func (m *Token) GetAccessToken() *String
```

#### func (*Token) GetExpiry

```go
func (m *Token) GetExpiry() *String
```

#### func (*Token) GetIdToken

```go
func (m *Token) GetIdToken() *String
```

#### func (*Token) GetRefreshToken

```go
func (m *Token) GetRefreshToken() *String
```

#### func (*Token) GetTokenType

```go
func (m *Token) GetTokenType() *String
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

#### func (*Token) ToContext

```go
func (s *Token) ToContext(ctx context.Context, key string) context.Context
```

#### func (*Token) TypeMatches

```go
func (s *Token) TypeMatches(src interface{}) bool
```

#### func (*Token) UnmarshalJSONFrom

```go
func (s *Token) UnmarshalJSONFrom(bits []byte) error
```

#### func (*Token) UnmarshalProtoFrom

```go
func (s *Token) UnmarshalProtoFrom(bits []byte) error
```

#### func (*Token) Validate

```go
func (s *Token) Validate(fn func(s *Token) error) error
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

#### type TokenSet

```go
type TokenSet struct {
	Tokens               map[string]*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func (*TokenSet) Debugf

```go
func (s *TokenSet) Debugf(format string)
```

#### func (*TokenSet) DeepEqual

```go
func (s *TokenSet) DeepEqual(y interface{}) bool
```

#### func (*TokenSet) Descriptor

```go
func (*TokenSet) Descriptor() ([]byte, []int)
```

#### func (*TokenSet) Exists

```go
func (s *TokenSet) Exists(key string) bool
```

#### func (*TokenSet) Get

```go
func (s *TokenSet) Get(key string) *Token
```

#### func (*TokenSet) GetTokens

```go
func (m *TokenSet) GetTokens() map[string]*Token
```

#### func (*TokenSet) JSONString

```go
func (s *TokenSet) JSONString() *String
```

#### func (*TokenSet) Length

```go
func (s *TokenSet) Length() int
```

#### func (*TokenSet) ProtoMessage

```go
func (*TokenSet) ProtoMessage()
```

#### func (*TokenSet) Put

```go
func (s *TokenSet) Put(key string, tok *Token)
```

#### func (*TokenSet) Reset

```go
func (m *TokenSet) Reset()
```

#### func (*TokenSet) String

```go
func (m *TokenSet) String() string
```

#### func (*TokenSet) ToContext

```go
func (s *TokenSet) ToContext(ctx context.Context, key string) context.Context
```

#### func (*TokenSet) UnmarshalJSONFrom

```go
func (s *TokenSet) UnmarshalJSONFrom(b []byte) error
```

#### func (*TokenSet) UnmarshalProtoFrom

```go
func (s *TokenSet) UnmarshalProtoFrom(b []byte) error
```

#### func (*TokenSet) Validate

```go
func (s *TokenSet) Validate(fn func(set *TokenSet) error) error
```

#### func (*TokenSet) XXX_DiscardUnknown

```go
func (m *TokenSet) XXX_DiscardUnknown()
```

#### func (*TokenSet) XXX_Marshal

```go
func (m *TokenSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*TokenSet) XXX_Merge

```go
func (m *TokenSet) XXX_Merge(src proto.Message)
```

#### func (*TokenSet) XXX_Size

```go
func (m *TokenSet) XXX_Size() int
```

#### func (*TokenSet) XXX_Unmarshal

```go
func (m *TokenSet) XXX_Unmarshal(b []byte) error
```
