# common
--
    import "github.com/autom8ter/api/common"


## Usage

#### func  FuncMap

```go
func FuncMap() template.FuncMap
```

#### func  ToError

```go
func ToError(err error, msg string) error
```

#### type AuthToken

```go
type AuthToken struct {
	Token                *String  `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```


#### func (*AuthToken) Bearer

```go
func (a *AuthToken) Bearer(req *http.Request)
```

#### func (*AuthToken) Descriptor

```go
func (*AuthToken) Descriptor() ([]byte, []int)
```

#### func (*AuthToken) GetToken

```go
func (m *AuthToken) GetToken() *String
```

#### func (*AuthToken) ProtoMessage

```go
func (*AuthToken) ProtoMessage()
```

#### func (*AuthToken) Reset

```go
func (m *AuthToken) Reset()
```

#### func (*AuthToken) String

```go
func (m *AuthToken) String() string
```

#### func (*AuthToken) XXX_DiscardUnknown

```go
func (m *AuthToken) XXX_DiscardUnknown()
```

#### func (*AuthToken) XXX_Marshal

```go
func (m *AuthToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*AuthToken) XXX_Merge

```go
func (m *AuthToken) XXX_Merge(src proto.Message)
```

#### func (*AuthToken) XXX_Size

```go
func (m *AuthToken) XXX_Size() int
```

#### func (*AuthToken) XXX_Unmarshal

```go
func (m *AuthToken) XXX_Unmarshal(b []byte) error
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
