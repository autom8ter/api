# common
--
    import "github.com/autom8ter/api/common"


## Usage

```go
var (
	Util    *objectify.Handler
	Context context.Context

	AUTH_SESSION = "auth-session"
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

#### func  RenderFile

```go
func RenderFile(name string, data []byte) http.HandlerFunc
```

#### func  SaveSession

```go
func SaveSession(w http.ResponseWriter, r *http.Request)
```

#### func  WriteFile

```go
func WriteFile(name string) http.HandlerFunc
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
func (m *Bytes) Compile(w io.Writer, t *Template) error
```

#### func (*Bytes) CompileHTTP

```go
func (m *Bytes) CompileHTTP(t *Template) http.HandlerFunc
```

#### func (*Bytes) Contains

```go
func (m *Bytes) Contains(str string) bool
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


#### func (*Float64) Descriptor

```go
func (*Float64) Descriptor() ([]byte, []int)
```

#### func (*Float64) GetNum

```go
func (m *Float64) GetNum() float64
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

#### type Int32

```go
type Int32 struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Int32) Descriptor

```go
func (*Int32) Descriptor() ([]byte, []int)
```

#### func (*Int32) GetNum

```go
func (m *Int32) GetNum() int32
```

#### func (*Int32) ProtoMessage

```go
func (*Int32) ProtoMessage()
```

#### func (*Int32) Reset

```go
func (m *Int32) Reset()
```

#### func (*Int32) String

```go
func (m *Int32) String() string
```

#### func (*Int32) XXX_DiscardUnknown

```go
func (m *Int32) XXX_DiscardUnknown()
```

#### func (*Int32) XXX_Marshal

```go
func (m *Int32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Int32) XXX_Merge

```go
func (m *Int32) XXX_Merge(src proto.Message)
```

#### func (*Int32) XXX_Size

```go
func (m *Int32) XXX_Size() int
```

#### func (*Int32) XXX_Unmarshal

```go
func (m *Int32) XXX_Unmarshal(b []byte) error
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


#### func (*Int64) Descriptor

```go
func (*Int64) Descriptor() ([]byte, []int)
```

#### func (*Int64) GetNum

```go
func (m *Int64) GetNum() int64
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

#### type Password

```go
type Password struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Password) Descriptor

```go
func (*Password) Descriptor() ([]byte, []int)
```

#### func (*Password) GetText

```go
func (m *Password) GetText() string
```

#### func (*Password) ProtoMessage

```go
func (*Password) ProtoMessage()
```

#### func (*Password) Reset

```go
func (m *Password) Reset()
```

#### func (*Password) String

```go
func (m *Password) String() string
```

#### func (*Password) XXX_DiscardUnknown

```go
func (m *Password) XXX_DiscardUnknown()
```

#### func (*Password) XXX_Marshal

```go
func (m *Password) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Password) XXX_Merge

```go
func (m *Password) XXX_Merge(src proto.Message)
```

#### func (*Password) XXX_Size

```go
func (m *Password) XXX_Size() int
```

#### func (*Password) XXX_Unmarshal

```go
func (m *Password) XXX_Unmarshal(b []byte) error
```

#### type Query

```go
type Query struct {
	Lucene               string   `protobuf:"bytes,1,opt,name=lucene,proto3" json:"lucene,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*Query) Descriptor

```go
func (*Query) Descriptor() ([]byte, []int)
```

#### func (*Query) GetLucene

```go
func (m *Query) GetLucene() string
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

#### type Secret

```go
type Secret struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  SecretFromEnv

```go
func SecretFromEnv() *Secret
```

#### func (*Secret) Descriptor

```go
func (*Secret) Descriptor() ([]byte, []int)
```

#### func (*Secret) GetText

```go
func (m *Secret) GetText() string
```

#### func (*Secret) InitSessions

```go
func (s *Secret) InitSessions() error
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

#### type String

```go
type String struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  Random

```go
func Random() *String
```

#### func  ToString

```go
func ToString(s string) *String
```

#### func (*String) Contains

```go
func (s *String) Contains(sub string) *Bool
```

#### func (*String) Descriptor

```go
func (*String) Descriptor() ([]byte, []int)
```

#### func (*String) GetText

```go
func (m *String) GetText() string
```

#### func (*String) Index

```go
func (s *String) Index(sub string) *Int64
```

#### func (*String) IsEmpty

```go
func (s *String) IsEmpty() bool
```

#### func (*String) Println

```go
func (s *String) Println()
```

#### func (*String) ProtoMessage

```go
func (*String) ProtoMessage()
```

#### func (*String) Replace

```go
func (s *String) Replace(oldNew ...string)
```

#### func (*String) Reset

```go
func (m *String) Reset()
```

#### func (*String) String

```go
func (m *String) String() string
```

#### func (*String) ToLower

```go
func (s *String) ToLower()
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
	Strings              []string `protobuf:"bytes,1,rep,name=strings,proto3" json:"strings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  ENVFromContext

```go
func ENVFromContext(ctx context.Context) *StringArray
```

#### func  ToStringArray

```go
func ToStringArray(s []string) *StringArray
```

#### func (*StringArray) Append

```go
func (m *StringArray) Append(vals ...*String)
```

#### func (*StringArray) Descriptor

```go
func (*StringArray) Descriptor() ([]byte, []int)
```

#### func (*StringArray) GetStrings

```go
func (m *StringArray) GetStrings() []string
```

#### func (*StringArray) ProtoMessage

```go
func (*StringArray) ProtoMessage()
```

#### func (*StringArray) Reset

```go
func (m *StringArray) Reset()
```

#### func (*StringArray) String

```go
func (m *StringArray) String() string
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
	StringMap            map[string]string `protobuf:"bytes,1,rep,name=string_map,json=stringMap,proto3" json:"string_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
```


#### func  ToStringMap

```go
func ToStringMap(s map[string]string) *StringMap
```

#### func (*StringMap) Clear

```go
func (s *StringMap) Clear(key string)
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
func (s *StringMap) Get(key string) string
```

#### func (*StringMap) GetStringMap

```go
func (m *StringMap) GetStringMap() map[string]string
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
func (s *StringMap) Put(key string, val string)
```

#### func (*StringMap) Reset

```go
func (m *StringMap) Reset()
```

#### func (*StringMap) String

```go
func (m *StringMap) String() string
```

#### func (*StringMap) TotalKeys

```go
func (s *StringMap) TotalKeys() int
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

#### type Template

```go
type Template struct {
	Name                 *String  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Text                 *String  `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func  NewTemplateFromFile

```go
func NewTemplateFromFile(filename string) (*Template, error)
```

#### func (*Template) AsTemplate

```go
func (m *Template) AsTemplate() (*template.Template, error)
```

#### func (*Template) Descriptor

```go
func (*Template) Descriptor() ([]byte, []int)
```

#### func (*Template) GetName

```go
func (m *Template) GetName() *String
```

#### func (*Template) GetText

```go
func (m *Template) GetText() *String
```

#### func (*Template) IsTemplate

```go
func (m *Template) IsTemplate() *Bool
```

#### func (*Template) ProtoMessage

```go
func (*Template) ProtoMessage()
```

#### func (*Template) Render

```go
func (m *Template) Render(w io.Writer, data interface{}) error
```

#### func (*Template) RenderBytes

```go
func (m *Template) RenderBytes(w io.Writer, bits *Bytes) error
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

#### func (*Token) ToSession

```go
func (t *Token) ToSession(session *sessions.Session)
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
