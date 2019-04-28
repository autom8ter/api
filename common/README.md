# common
--
    import "github.com/autom8ter/api/common"


## Usage

```go
var (
	Util    *objectify.Handler
	Context context.Context

	AUTH_SESSION           = "auth-session"
	SESSION_SECRET_ENV_KEY = "SECRET"
)
```

```go
var HTTPMethod_name = map[int32]string{
	0: "GET",
	1: "POST",
	2: "PATCH",
}
```

```go
var HTTPMethod_value = map[string]int32{
	"GET":   0,
	"POST":  1,
	"PATCH": 2,
}
```

#### func  AuthSessionValues

```go
func AuthSessionValues(session *sessions.Session, key string, data map[string]interface{})
```

#### func  FromSession

```go
func FromSession(key string, session *sessions.Session) interface{}
```

#### func  FuncMap

```go
func FuncMap() template.FuncMap
```

#### func  GetAuthSession

```go
func GetAuthSession(r *http.Request) (*sessions.Session, error)
```

#### func  GetStateSession

```go
func GetStateSession(r *http.Request) (*sessions.Session, error)
```

#### func  SaveSession

```go
func SaveSession(w http.ResponseWriter, r *http.Request)
```

#### type Bool

```go
type Bool struct {
	Answer               bool     `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  ToBool

```go
func ToBool(b bool) *Bool
```

#### func (*Bool) Descriptor

```go
func (*Bool) Descriptor() ([]byte, []int)
```

#### func (*Bool) GetAnswer

```go
func (m *Bool) GetAnswer() bool
```

#### func (*Bool) Pointer

```go
func (s *Bool) Pointer() *bool
```

#### func (*Bool) ProtoMessage

```go
func (*Bool) ProtoMessage()
```

#### func (*Bool) Reset

```go
func (m *Bool) Reset()
```

#### func (*Bool) String

```go
func (m *Bool) String() string
```

#### func (*Bool) XXX_DiscardUnknown

```go
func (m *Bool) XXX_DiscardUnknown()
```

#### func (*Bool) XXX_Marshal

```go
func (m *Bool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Bool) XXX_Merge

```go
func (m *Bool) XXX_Merge(src proto.Message)
```

#### func (*Bool) XXX_Size

```go
func (m *Bool) XXX_Size() int
```

#### func (*Bool) XXX_Unmarshal

```go
func (m *Bool) XXX_Unmarshal(b []byte) error
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

#### func  ObjToBytes

```go
func ObjToBytes(b interface{}) *Bytes
```

#### func  ToBytes

```go
func ToBytes(b []byte) *Bytes
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
func (m *Bytes) Compile(name string, w io.Writer, t *String) error
```

#### func (*Bytes) CompileHTTP

```go
func (m *Bytes) CompileHTTP(name string, t *String) http.HandlerFunc
```

#### func (*Bytes) Contains

```go
func (m *Bytes) Contains(str string) bool
```

#### func (*Bytes) Debugln

```go
func (s *Bytes) Debugln()
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

#### func (*Bytes) ToString

```go
func (b *Bytes) ToString() *String
```

#### func (*Bytes) UnMarshalJSON

```go
func (m *Bytes) UnMarshalJSON(obj interface{}) error
```

#### func (*Bytes) UnMarshalProto

```go
func (m *Bytes) UnMarshalProto(obj interface{}) error
```

#### func (*Bytes) UnmarshalJSON

```go
func (b *Bytes) UnmarshalJSON(obj interface{}) error
```

#### func (*Bytes) UnmarshalProto

```go
func (b *Bytes) UnmarshalProto(obj interface{}) error
```

#### func (*Bytes) Write

```go
func (m *Bytes) Write(p []byte) (n int, err error)
```

#### func (*Bytes) WriteString

```go
func (m *Bytes) WriteString() http.HandlerFunc
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

#### type Float64

```go
type Float64 struct {
	Num                  float64  `protobuf:"fixed64,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  ToFloat64

```go
func ToFloat64(i float64) *Float64
```

#### func (*Float64) Descriptor

```go
func (*Float64) Descriptor() ([]byte, []int)
```

#### func (*Float64) GetNum

```go
func (m *Float64) GetNum() float64
```

#### func (*Float64) Pointer

```go
func (s *Float64) Pointer() *float64
```

#### func (*Float64) ProtoMessage

```go
func (*Float64) ProtoMessage()
```

#### func (*Float64) Reset

```go
func (m *Float64) Reset()
```

#### func (*Float64) String

```go
func (m *Float64) String() string
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

#### type HTTPMethod

```go
type HTTPMethod int32
```


```go
const (
	HTTPMethod_GET   HTTPMethod = 0
	HTTPMethod_POST  HTTPMethod = 1
	HTTPMethod_PATCH HTTPMethod = 2
)
```

#### func (HTTPMethod) EnumDescriptor

```go
func (HTTPMethod) EnumDescriptor() ([]byte, []int)
```

#### func (HTTPMethod) Normalize

```go
func (h HTTPMethod) Normalize() *String
```

#### func (HTTPMethod) String

```go
func (x HTTPMethod) String() string
```

#### type HTTPRequest

```go
type HTTPRequest struct {
	Method               HTTPMethod `protobuf:"varint,1,opt,name=method,proto3,enum=common.HTTPMethod" json:"method,omitempty"`
	Url                  *String    `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Form                 *StringMap `protobuf:"bytes,3,opt,name=form,proto3" json:"form,omitempty"`
	Body                 *Bytes     `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}
```


#### func (*HTTPRequest) Descriptor

```go
func (*HTTPRequest) Descriptor() ([]byte, []int)
```

#### func (*HTTPRequest) Do

```go
func (h *HTTPRequest) Do(token *Token) (*Bytes, error)
```

#### func (*HTTPRequest) GetBody

```go
func (m *HTTPRequest) GetBody() *Bytes
```

#### func (*HTTPRequest) GetForm

```go
func (m *HTTPRequest) GetForm() *StringMap
```

#### func (*HTTPRequest) GetMethod

```go
func (m *HTTPRequest) GetMethod() HTTPMethod
```

#### func (*HTTPRequest) GetUrl

```go
func (m *HTTPRequest) GetUrl() *String
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
	Id                   *String  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
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

#### func (*Int64) Descriptor

```go
func (*Int64) Descriptor() ([]byte, []int)
```

#### func (*Int64) GetNum

```go
func (m *Int64) GetNum() int64
```

#### func (*Int64) Pointer

```go
func (s *Int64) Pointer() *int64
```

#### func (*Int64) ProtoMessage

```go
func (*Int64) ProtoMessage()
```

#### func (*Int64) Reset

```go
func (m *Int64) Reset()
```

#### func (*Int64) String

```go
func (m *Int64) String() string
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

#### type String

```go
type String struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  ObjToString

```go
func ObjToString(b interface{}) *String
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

#### func (*String) Contains

```go
func (s *String) Contains(sub string) *Bool
```

#### func (*String) Debugln

```go
func (s *String) Debugln()
```

#### func (*String) Descriptor

```go
func (*String) Descriptor() ([]byte, []int)
```

#### func (*String) ExecuteAsBashCMD

```go
func (m *String) ExecuteAsBashCMD() *Bytes
```

#### func (*String) ExecuteAsPython3

```go
func (m *String) ExecuteAsPython3() *Bytes
```

#### func (*String) ExecuteAsShellCMD

```go
func (m *String) ExecuteAsShellCMD() *Bytes
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
func (m *String) IsTemplate() *Bool
```

#### func (*String) JSON

```go
func (m *String) JSON() []byte
```

#### func (*String) Matches

```go
func (s *String) Matches(this string) bool
```

#### func (*String) ParseLanguage

```go
func (s *String) ParseLanguage() (language.Tag, error)
```

#### func (*String) ParseRegion

```go
func (s *String) ParseRegion() (language.Region, error)
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
func (s *String) RegexMatches(reg string) *Bool
```

#### func (*String) Render

```go
func (m *String) Render(name string, w io.Writer, data interface{}) error
```

#### func (*String) RenderBytes

```go
func (m *String) RenderBytes(name string, w io.Writer, bits *Bytes) error
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

#### func (*String) String

```go
func (m *String) String() string
```

#### func (*String) ToBytes

```go
func (b *String) ToBytes() *Bytes
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

#### func (*String) WriteString

```go
func (m *String) WriteString() http.HandlerFunc
```

#### func (*String) XML

```go
func (m *String) XML() []byte
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

#### func (*String) YAML

```go
func (m *String) YAML() []byte
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

#### func (*StringArray) Debugln

```go
func (s *StringArray) Debugln()
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

#### func (*StringArray) JSON

```go
func (t *StringArray) JSON() []byte
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

#### func (*StringArray) ToString

```go
func (s *StringArray) ToString() *String
```

#### func (*StringArray) XML

```go
func (t *StringArray) XML() []byte
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

#### func (*StringArray) YAML

```go
func (t *StringArray) YAML() []byte
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


#### func  StructToMap

```go
func StructToMap(obj interface{}) *StringMap
```

#### func  ToStringMap

```go
func ToStringMap(s map[string]string) *StringMap
```

#### func (*StringMap) Clear

```go
func (s *StringMap) Clear(key string)
```

#### func (*StringMap) Debugln

```go
func (s *StringMap) Debugln()
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

#### func (*StringMap) JSON

```go
func (t *StringMap) JSON() []byte
```

#### func (*StringMap) Keys

```go
func (s *StringMap) Keys() []string
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

#### func (*StringMap) ToString

```go
func (s *StringMap) ToString() *String
```

#### func (*StringMap) TotalKeys

```go
func (s *StringMap) TotalKeys() int
```

#### func (*StringMap) XML

```go
func (t *StringMap) XML() []byte
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

#### func (*StringMap) YAML

```go
func (t *StringMap) YAML() []byte
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


#### func  TokenFromAuthSession

```go
func TokenFromAuthSession(session *sessions.Session) (*Token, error)
```

#### func  TokenFromOAuthToken

```go
func TokenFromOAuthToken(tok *oauth2.Token) *Token
```

#### func (*Token) Debugln

```go
func (s *Token) Debugln()
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

#### func (*Token) JSON

```go
func (t *Token) JSON() []byte
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

#### func (*Token) ToSession

```go
func (t *Token) ToSession(session *sessions.Session)
```

#### func (*Token) XML

```go
func (t *Token) XML() []byte
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

#### func (*Token) YAML

```go
func (t *Token) YAML() []byte
```
