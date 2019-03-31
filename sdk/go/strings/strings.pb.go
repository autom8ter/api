// Code generated by protoc-gen-go. DO NOT EDIT.
// source: strings/strings.proto

package strings

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Format int32

const (
	Format_TEXT  Format = 0
	Format_HTML  Format = 1
	Format_BYTES Format = 2
	Format_JSON  Format = 3
	Format_YAML  Format = 4
	Format_XML   Format = 5
	Format_HCL   Format = 6
)

var Format_name = map[int32]string{
	0: "TEXT",
	1: "HTML",
	2: "BYTES",
	3: "JSON",
	4: "YAML",
	5: "XML",
	6: "HCL",
}

var Format_value = map[string]int32{
	"TEXT":  0,
	"HTML":  1,
	"BYTES": 2,
	"JSON":  3,
	"YAML":  4,
	"XML":   5,
	"HCL":   6,
}

func (x Format) String() string {
	return proto.EnumName(Format_name, int32(x))
}

func (Format) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{0}
}

type String struct {
	Val                  string   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{0}
}

func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (m *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(m, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type Bytes struct {
	Bits                 []byte   `protobuf:"bytes,1,opt,name=bits,proto3" json:"bits,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bytes) Reset()         { *m = Bytes{} }
func (m *Bytes) String() string { return proto.CompactTextString(m) }
func (*Bytes) ProtoMessage()    {}
func (*Bytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{1}
}

func (m *Bytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bytes.Unmarshal(m, b)
}
func (m *Bytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bytes.Marshal(b, m, deterministic)
}
func (m *Bytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bytes.Merge(m, src)
}
func (m *Bytes) XXX_Size() int {
	return xxx_messageInfo_Bytes.Size(m)
}
func (m *Bytes) XXX_DiscardUnknown() {
	xxx_messageInfo_Bytes.DiscardUnknown(m)
}

var xxx_messageInfo_Bytes proto.InternalMessageInfo

func (m *Bytes) GetBits() []byte {
	if m != nil {
		return m.Bits
	}
	return nil
}

type Text struct {
	Format               Format   `protobuf:"varint,1,opt,name=format,proto3,enum=strings.Format" json:"format,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{2}
}

func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (m *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(m, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetFormat() Format {
	if m != nil {
		return m.Format
	}
	return Format_TEXT
}

func (m *Text) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type InputString struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputString) Reset()         { *m = InputString{} }
func (m *InputString) String() string { return proto.CompactTextString(m) }
func (*InputString) ProtoMessage()    {}
func (*InputString) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{4}
}

func (m *InputString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputString.Unmarshal(m, b)
}
func (m *InputString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputString.Marshal(b, m, deterministic)
}
func (m *InputString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputString.Merge(m, src)
}
func (m *InputString) XXX_Size() int {
	return xxx_messageInfo_InputString.Size(m)
}
func (m *InputString) XXX_DiscardUnknown() {
	xxx_messageInfo_InputString.DiscardUnknown(m)
}

var xxx_messageInfo_InputString proto.InternalMessageInfo

func (m *InputString) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type RenderResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Text                 *Text    `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenderResponse) Reset()         { *m = RenderResponse{} }
func (m *RenderResponse) String() string { return proto.CompactTextString(m) }
func (*RenderResponse) ProtoMessage()    {}
func (*RenderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{5}
}

func (m *RenderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenderResponse.Unmarshal(m, b)
}
func (m *RenderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenderResponse.Marshal(b, m, deterministic)
}
func (m *RenderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenderResponse.Merge(m, src)
}
func (m *RenderResponse) XXX_Size() int {
	return xxx_messageInfo_RenderResponse.Size(m)
}
func (m *RenderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RenderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RenderResponse proto.InternalMessageInfo

func (m *RenderResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RenderResponse) GetText() *Text {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *RenderResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type RenderRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Text                 *Text    `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenderRequest) Reset()         { *m = RenderRequest{} }
func (m *RenderRequest) String() string { return proto.CompactTextString(m) }
func (*RenderRequest) ProtoMessage()    {}
func (*RenderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{6}
}

func (m *RenderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenderRequest.Unmarshal(m, b)
}
func (m *RenderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenderRequest.Marshal(b, m, deterministic)
}
func (m *RenderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenderRequest.Merge(m, src)
}
func (m *RenderRequest) XXX_Size() int {
	return xxx_messageInfo_RenderRequest.Size(m)
}
func (m *RenderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RenderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RenderRequest proto.InternalMessageInfo

func (m *RenderRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RenderRequest) GetText() *Text {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *RenderRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ReplaceRequest struct {
	Text                 *Text             `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Replacements         map[string]string `protobuf:"bytes,2,rep,name=replacements,proto3" json:"replacements,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReplaceRequest) Reset()         { *m = ReplaceRequest{} }
func (m *ReplaceRequest) String() string { return proto.CompactTextString(m) }
func (*ReplaceRequest) ProtoMessage()    {}
func (*ReplaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{7}
}

func (m *ReplaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplaceRequest.Unmarshal(m, b)
}
func (m *ReplaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplaceRequest.Marshal(b, m, deterministic)
}
func (m *ReplaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceRequest.Merge(m, src)
}
func (m *ReplaceRequest) XXX_Size() int {
	return xxx_messageInfo_ReplaceRequest.Size(m)
}
func (m *ReplaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceRequest proto.InternalMessageInfo

func (m *ReplaceRequest) GetText() *Text {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *ReplaceRequest) GetReplacements() map[string]string {
	if m != nil {
		return m.Replacements
	}
	return nil
}

type ReplaceResponse struct {
	Text                 *Text             `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Replacements         map[string]string `protobuf:"bytes,2,rep,name=replacements,proto3" json:"replacements,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReplaceResponse) Reset()         { *m = ReplaceResponse{} }
func (m *ReplaceResponse) String() string { return proto.CompactTextString(m) }
func (*ReplaceResponse) ProtoMessage()    {}
func (*ReplaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{8}
}

func (m *ReplaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplaceResponse.Unmarshal(m, b)
}
func (m *ReplaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplaceResponse.Marshal(b, m, deterministic)
}
func (m *ReplaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceResponse.Merge(m, src)
}
func (m *ReplaceResponse) XXX_Size() int {
	return xxx_messageInfo_ReplaceResponse.Size(m)
}
func (m *ReplaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceResponse proto.InternalMessageInfo

func (m *ReplaceResponse) GetText() *Text {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *ReplaceResponse) GetReplacements() map[string]string {
	if m != nil {
		return m.Replacements
	}
	return nil
}

type ExtractRequest struct {
	Text                 *Text    `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Keywords             []string `protobuf:"bytes,2,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtractRequest) Reset()         { *m = ExtractRequest{} }
func (m *ExtractRequest) String() string { return proto.CompactTextString(m) }
func (*ExtractRequest) ProtoMessage()    {}
func (*ExtractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{9}
}

func (m *ExtractRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtractRequest.Unmarshal(m, b)
}
func (m *ExtractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtractRequest.Marshal(b, m, deterministic)
}
func (m *ExtractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtractRequest.Merge(m, src)
}
func (m *ExtractRequest) XXX_Size() int {
	return xxx_messageInfo_ExtractRequest.Size(m)
}
func (m *ExtractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExtractRequest proto.InternalMessageInfo

func (m *ExtractRequest) GetText() *Text {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *ExtractRequest) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

type ExtractResponse struct {
	Text                 *Text    `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Keywords             []string `protobuf:"bytes,2,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtractResponse) Reset()         { *m = ExtractResponse{} }
func (m *ExtractResponse) String() string { return proto.CompactTextString(m) }
func (*ExtractResponse) ProtoMessage()    {}
func (*ExtractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{10}
}

func (m *ExtractResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtractResponse.Unmarshal(m, b)
}
func (m *ExtractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtractResponse.Marshal(b, m, deterministic)
}
func (m *ExtractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtractResponse.Merge(m, src)
}
func (m *ExtractResponse) XXX_Size() int {
	return xxx_messageInfo_ExtractResponse.Size(m)
}
func (m *ExtractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExtractResponse proto.InternalMessageInfo

func (m *ExtractResponse) GetText() *Text {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *ExtractResponse) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

type GetString struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetString) Reset()         { *m = GetString{} }
func (m *GetString) String() string { return proto.CompactTextString(m) }
func (*GetString) ProtoMessage()    {}
func (*GetString) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{11}
}

func (m *GetString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetString.Unmarshal(m, b)
}
func (m *GetString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetString.Marshal(b, m, deterministic)
}
func (m *GetString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetString.Merge(m, src)
}
func (m *GetString) XXX_Size() int {
	return xxx_messageInfo_GetString.Size(m)
}
func (m *GetString) XXX_DiscardUnknown() {
	xxx_messageInfo_GetString.DiscardUnknown(m)
}

var xxx_messageInfo_GetString proto.InternalMessageInfo

func (m *GetString) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetBool struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBool) Reset()         { *m = GetBool{} }
func (m *GetBool) String() string { return proto.CompactTextString(m) }
func (*GetBool) ProtoMessage()    {}
func (*GetBool) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{12}
}

func (m *GetBool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBool.Unmarshal(m, b)
}
func (m *GetBool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBool.Marshal(b, m, deterministic)
}
func (m *GetBool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBool.Merge(m, src)
}
func (m *GetBool) XXX_Size() int {
	return xxx_messageInfo_GetBool.Size(m)
}
func (m *GetBool) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBool.DiscardUnknown(m)
}

var xxx_messageInfo_GetBool proto.InternalMessageInfo

func (m *GetBool) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Bool struct {
	This                 bool     `protobuf:"varint,1,opt,name=this,proto3" json:"this,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bool) Reset()         { *m = Bool{} }
func (m *Bool) String() string { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()    {}
func (*Bool) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{13}
}

func (m *Bool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bool.Unmarshal(m, b)
}
func (m *Bool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bool.Marshal(b, m, deterministic)
}
func (m *Bool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bool.Merge(m, src)
}
func (m *Bool) XXX_Size() int {
	return xxx_messageInfo_Bool.Size(m)
}
func (m *Bool) XXX_DiscardUnknown() {
	xxx_messageInfo_Bool.DiscardUnknown(m)
}

var xxx_messageInfo_Bool proto.InternalMessageInfo

func (m *Bool) GetThis() bool {
	if m != nil {
		return m.This
	}
	return false
}

type GetStrings struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStrings) Reset()         { *m = GetStrings{} }
func (m *GetStrings) String() string { return proto.CompactTextString(m) }
func (*GetStrings) ProtoMessage()    {}
func (*GetStrings) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{14}
}

func (m *GetStrings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStrings.Unmarshal(m, b)
}
func (m *GetStrings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStrings.Marshal(b, m, deterministic)
}
func (m *GetStrings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStrings.Merge(m, src)
}
func (m *GetStrings) XXX_Size() int {
	return xxx_messageInfo_GetStrings.Size(m)
}
func (m *GetStrings) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStrings.DiscardUnknown(m)
}

var xxx_messageInfo_GetStrings proto.InternalMessageInfo

func (m *GetStrings) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Strings struct {
	This                 []string `protobuf:"bytes,1,rep,name=this,proto3" json:"this,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Strings) Reset()         { *m = Strings{} }
func (m *Strings) String() string { return proto.CompactTextString(m) }
func (*Strings) ProtoMessage()    {}
func (*Strings) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{15}
}

func (m *Strings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Strings.Unmarshal(m, b)
}
func (m *Strings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Strings.Marshal(b, m, deterministic)
}
func (m *Strings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Strings.Merge(m, src)
}
func (m *Strings) XXX_Size() int {
	return xxx_messageInfo_Strings.Size(m)
}
func (m *Strings) XXX_DiscardUnknown() {
	xxx_messageInfo_Strings.DiscardUnknown(m)
}

var xxx_messageInfo_Strings proto.InternalMessageInfo

func (m *Strings) GetThis() []string {
	if m != nil {
		return m.This
	}
	return nil
}

type GetStringMap struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStringMap) Reset()         { *m = GetStringMap{} }
func (m *GetStringMap) String() string { return proto.CompactTextString(m) }
func (*GetStringMap) ProtoMessage()    {}
func (*GetStringMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{16}
}

func (m *GetStringMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStringMap.Unmarshal(m, b)
}
func (m *GetStringMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStringMap.Marshal(b, m, deterministic)
}
func (m *GetStringMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStringMap.Merge(m, src)
}
func (m *GetStringMap) XXX_Size() int {
	return xxx_messageInfo_GetStringMap.Size(m)
}
func (m *GetStringMap) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStringMap.DiscardUnknown(m)
}

var xxx_messageInfo_GetStringMap proto.InternalMessageInfo

func (m *GetStringMap) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type StringMap struct {
	This                 map[string]string `protobuf:"bytes,1,rep,name=this,proto3" json:"this,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StringMap) Reset()         { *m = StringMap{} }
func (m *StringMap) String() string { return proto.CompactTextString(m) }
func (*StringMap) ProtoMessage()    {}
func (*StringMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{17}
}

func (m *StringMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMap.Unmarshal(m, b)
}
func (m *StringMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMap.Marshal(b, m, deterministic)
}
func (m *StringMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMap.Merge(m, src)
}
func (m *StringMap) XXX_Size() int {
	return xxx_messageInfo_StringMap.Size(m)
}
func (m *StringMap) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMap.DiscardUnknown(m)
}

var xxx_messageInfo_StringMap proto.InternalMessageInfo

func (m *StringMap) GetThis() map[string]string {
	if m != nil {
		return m.This
	}
	return nil
}

type GetBoolMap struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBoolMap) Reset()         { *m = GetBoolMap{} }
func (m *GetBoolMap) String() string { return proto.CompactTextString(m) }
func (*GetBoolMap) ProtoMessage()    {}
func (*GetBoolMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{18}
}

func (m *GetBoolMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBoolMap.Unmarshal(m, b)
}
func (m *GetBoolMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBoolMap.Marshal(b, m, deterministic)
}
func (m *GetBoolMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBoolMap.Merge(m, src)
}
func (m *GetBoolMap) XXX_Size() int {
	return xxx_messageInfo_GetBoolMap.Size(m)
}
func (m *GetBoolMap) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBoolMap.DiscardUnknown(m)
}

var xxx_messageInfo_GetBoolMap proto.InternalMessageInfo

func (m *GetBoolMap) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type BoolMap struct {
	This                 map[string]bool `protobuf:"bytes,1,rep,name=this,proto3" json:"this,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BoolMap) Reset()         { *m = BoolMap{} }
func (m *BoolMap) String() string { return proto.CompactTextString(m) }
func (*BoolMap) ProtoMessage()    {}
func (*BoolMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{19}
}

func (m *BoolMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolMap.Unmarshal(m, b)
}
func (m *BoolMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolMap.Marshal(b, m, deterministic)
}
func (m *BoolMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolMap.Merge(m, src)
}
func (m *BoolMap) XXX_Size() int {
	return xxx_messageInfo_BoolMap.Size(m)
}
func (m *BoolMap) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolMap.DiscardUnknown(m)
}

var xxx_messageInfo_BoolMap proto.InternalMessageInfo

func (m *BoolMap) GetThis() map[string]bool {
	if m != nil {
		return m.This
	}
	return nil
}

type GetInt64 struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInt64) Reset()         { *m = GetInt64{} }
func (m *GetInt64) String() string { return proto.CompactTextString(m) }
func (*GetInt64) ProtoMessage()    {}
func (*GetInt64) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{20}
}

func (m *GetInt64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInt64.Unmarshal(m, b)
}
func (m *GetInt64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInt64.Marshal(b, m, deterministic)
}
func (m *GetInt64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInt64.Merge(m, src)
}
func (m *GetInt64) XXX_Size() int {
	return xxx_messageInfo_GetInt64.Size(m)
}
func (m *GetInt64) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInt64.DiscardUnknown(m)
}

var xxx_messageInfo_GetInt64 proto.InternalMessageInfo

func (m *GetInt64) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Int64 struct {
	This                 int64    `protobuf:"varint,1,opt,name=this,proto3" json:"this,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64) Reset()         { *m = Int64{} }
func (m *Int64) String() string { return proto.CompactTextString(m) }
func (*Int64) ProtoMessage()    {}
func (*Int64) Descriptor() ([]byte, []int) {
	return fileDescriptor_44b471cba9524d5b, []int{21}
}

func (m *Int64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64.Unmarshal(m, b)
}
func (m *Int64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64.Marshal(b, m, deterministic)
}
func (m *Int64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64.Merge(m, src)
}
func (m *Int64) XXX_Size() int {
	return xxx_messageInfo_Int64.Size(m)
}
func (m *Int64) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64.DiscardUnknown(m)
}

var xxx_messageInfo_Int64 proto.InternalMessageInfo

func (m *Int64) GetThis() int64 {
	if m != nil {
		return m.This
	}
	return 0
}

func init() {
	proto.RegisterEnum("strings.Format", Format_name, Format_value)
	proto.RegisterType((*String)(nil), "strings.String")
	proto.RegisterType((*Bytes)(nil), "strings.Bytes")
	proto.RegisterType((*Text)(nil), "strings.Text")
	proto.RegisterType((*Empty)(nil), "strings.Empty")
	proto.RegisterType((*InputString)(nil), "strings.InputString")
	proto.RegisterType((*RenderResponse)(nil), "strings.RenderResponse")
	proto.RegisterType((*RenderRequest)(nil), "strings.RenderRequest")
	proto.RegisterType((*ReplaceRequest)(nil), "strings.ReplaceRequest")
	proto.RegisterMapType((map[string]string)(nil), "strings.ReplaceRequest.ReplacementsEntry")
	proto.RegisterType((*ReplaceResponse)(nil), "strings.ReplaceResponse")
	proto.RegisterMapType((map[string]string)(nil), "strings.ReplaceResponse.ReplacementsEntry")
	proto.RegisterType((*ExtractRequest)(nil), "strings.ExtractRequest")
	proto.RegisterType((*ExtractResponse)(nil), "strings.ExtractResponse")
	proto.RegisterType((*GetString)(nil), "strings.GetString")
	proto.RegisterType((*GetBool)(nil), "strings.GetBool")
	proto.RegisterType((*Bool)(nil), "strings.Bool")
	proto.RegisterType((*GetStrings)(nil), "strings.GetStrings")
	proto.RegisterType((*Strings)(nil), "strings.Strings")
	proto.RegisterType((*GetStringMap)(nil), "strings.GetStringMap")
	proto.RegisterType((*StringMap)(nil), "strings.StringMap")
	proto.RegisterMapType((map[string]string)(nil), "strings.StringMap.ThisEntry")
	proto.RegisterType((*GetBoolMap)(nil), "strings.GetBoolMap")
	proto.RegisterType((*BoolMap)(nil), "strings.BoolMap")
	proto.RegisterMapType((map[string]bool)(nil), "strings.BoolMap.ThisEntry")
	proto.RegisterType((*GetInt64)(nil), "strings.GetInt64")
	proto.RegisterType((*Int64)(nil), "strings.Int64")
}

func init() { proto.RegisterFile("strings/strings.proto", fileDescriptor_44b471cba9524d5b) }

var fileDescriptor_44b471cba9524d5b = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0x6f, 0xd2, 0x5e,
	0x14, 0xfe, 0x95, 0xb6, 0x94, 0x1e, 0x36, 0xe8, 0xef, 0x46, 0x13, 0xd2, 0x6d, 0x06, 0xfa, 0xe2,
	0xdc, 0x03, 0x1a, 0x34, 0x6a, 0x7c, 0x31, 0x62, 0x90, 0x61, 0xe8, 0x66, 0x4a, 0x1f, 0x36, 0x7d,
	0xea, 0xe0, 0xea, 0x08, 0xd0, 0xd6, 0xf6, 0x80, 0xf4, 0xef, 0x33, 0xfe, 0x5f, 0xe6, 0xde, 0xb6,
	0x17, 0xc8, 0xaa, 0xd9, 0xa2, 0x3e, 0xf1, 0xdd, 0xfb, 0x9d, 0xf3, 0x9d, 0xef, 0xe3, 0xe4, 0x16,
	0xee, 0xc7, 0x18, 0x4d, 0xfd, 0x2f, 0xf1, 0xe3, 0xec, 0xb7, 0x1d, 0x46, 0x01, 0x06, 0x44, 0xcb,
	0x8e, 0x96, 0x09, 0xe5, 0x11, 0x87, 0xc4, 0x00, 0x79, 0xe5, 0xcd, 0x1b, 0x52, 0x53, 0x3a, 0xd6,
	0x1d, 0x06, 0xad, 0x03, 0x50, 0xbb, 0x09, 0xd2, 0x98, 0x10, 0x50, 0xae, 0xa6, 0x18, 0x73, 0x6e,
	0xcf, 0xe1, 0xd8, 0x1a, 0x80, 0xe2, 0xd2, 0x35, 0x92, 0x87, 0x50, 0xfe, 0x1c, 0x44, 0x0b, 0x0f,
	0x39, 0x5b, 0xeb, 0xd4, 0xdb, 0xf9, 0xa4, 0x77, 0xfc, 0xda, 0xc9, 0x68, 0xd2, 0x00, 0x6d, 0x1c,
	0xf8, 0x48, 0x7d, 0x6c, 0x94, 0xf8, 0x8c, 0xfc, 0x68, 0x69, 0xa0, 0xf6, 0x16, 0x21, 0x26, 0x56,
	0x0b, 0xaa, 0x03, 0x3f, 0x5c, 0x62, 0xe6, 0x88, 0x80, 0x82, 0x74, 0x8d, 0x99, 0x25, 0x8e, 0xad,
	0x4f, 0x50, 0x73, 0xa8, 0x3f, 0xa1, 0x91, 0x43, 0xe3, 0x30, 0xf0, 0x63, 0xca, 0xaa, 0x7c, 0x6f,
	0x41, 0xf3, 0x2a, 0x86, 0x49, 0x2b, 0xeb, 0x64, 0x83, 0xaa, 0x9d, 0x7d, 0x61, 0x89, 0x39, 0x4e,
	0x85, 0x58, 0xdb, 0xc4, 0x43, 0xaf, 0x21, 0xa7, 0x99, 0x18, 0xb6, 0x3e, 0xc2, 0x7e, 0x2e, 0xfe,
	0x75, 0x49, 0x63, 0xfc, 0x9b, 0xda, 0xdf, 0x25, 0xe6, 0x3c, 0x9c, 0x7b, 0x63, 0x9a, 0xab, 0xb7,
	0xb6, 0xf2, 0xfd, 0x42, 0xc9, 0x86, 0xbd, 0x28, 0x6d, 0x5a, 0x50, 0x1f, 0xe3, 0x46, 0xa9, 0x29,
	0x1f, 0x57, 0x3b, 0x8f, 0x44, 0xe9, 0xae, 0x62, 0x7e, 0xe4, 0xb5, 0x3d, 0x1f, 0xa3, 0xc4, 0xd9,
	0x69, 0x37, 0x5f, 0xc3, 0xff, 0x37, 0x4a, 0xd8, 0xe2, 0x67, 0x34, 0xc9, 0x17, 0x3f, 0xa3, 0x09,
	0xb9, 0x07, 0xea, 0xca, 0x9b, 0x2f, 0x69, 0xb6, 0xa8, 0xf4, 0xf0, 0xaa, 0xf4, 0x52, 0xb2, 0x7e,
	0x48, 0x50, 0x17, 0x33, 0xb3, 0x05, 0xdc, 0x22, 0xc6, 0x59, 0x61, 0x8c, 0x93, 0x9b, 0x31, 0x52,
	0xc9, 0x7f, 0x9f, 0xe3, 0x1c, 0x6a, 0xbd, 0x35, 0x46, 0xde, 0x18, 0xef, 0xb0, 0x0c, 0x13, 0x2a,
	0x33, 0x9a, 0x7c, 0x0b, 0xa2, 0x49, 0x9a, 0x40, 0x77, 0xc4, 0xd9, 0xfa, 0x00, 0x75, 0x21, 0x78,
	0xfb, 0xff, 0xe5, 0x77, 0x8a, 0x47, 0xa0, 0xf7, 0x29, 0x6e, 0x1e, 0xe7, 0x6e, 0x36, 0xeb, 0x00,
	0xb4, 0x3e, 0xc5, 0x6e, 0x10, 0xcc, 0x0b, 0x48, 0x13, 0x14, 0xce, 0xb0, 0x17, 0x74, 0x3d, 0x4d,
	0x1f, 0x6e, 0xc5, 0xe1, 0xd8, 0x7a, 0x00, 0x20, 0x74, 0xe3, 0x82, 0xde, 0x23, 0xd0, 0x72, 0x72,
	0xd3, 0x2e, 0xf3, 0x07, 0xc8, 0xda, 0x9b, 0xb0, 0x27, 0xda, 0x6d, 0x2f, 0x2c, 0x10, 0x58, 0x81,
	0xbe, 0xa1, 0x9f, 0x6c, 0x49, 0x54, 0x3b, 0x87, 0xe2, 0x4f, 0x10, 0x15, 0x6d, 0xf7, 0x7a, 0x9a,
	0xed, 0x98, 0x57, 0x9a, 0x2f, 0x40, 0x17, 0x57, 0x77, 0xda, 0x69, 0x1a, 0x8c, 0xe5, 0x2e, 0xf6,
	0x15, 0x81, 0x96, 0x93, 0xed, 0x1d, 0x57, 0xa6, 0x70, 0x95, 0xf1, 0x7f, 0xe4, 0xa9, 0xb2, 0xed,
	0xe9, 0x10, 0x2a, 0x7d, 0x8a, 0x03, 0x1f, 0x9f, 0x3f, 0x2b, 0xdc, 0xa1, 0x9a, 0x52, 0xdb, 0x7b,
	0x92, 0xd3, 0x99, 0x27, 0x36, 0x94, 0xd3, 0x2f, 0x28, 0xa9, 0x80, 0xe2, 0xf6, 0x2e, 0x5c, 0xe3,
	0x3f, 0x86, 0x4e, 0x5d, 0x7b, 0x68, 0x48, 0x44, 0x07, 0xb5, 0x7b, 0xe9, 0xf6, 0x46, 0x46, 0x89,
	0x5d, 0xbe, 0x1f, 0x9d, 0x9f, 0x19, 0x32, 0x43, 0x97, 0x6f, 0xec, 0xa1, 0xa1, 0x10, 0x0d, 0xe4,
	0x0b, 0x7b, 0x68, 0xa8, 0x0c, 0x9c, 0xbe, 0x1d, 0x1a, 0xe5, 0xab, 0x32, 0xff, 0xf0, 0x3f, 0xfd,
	0x19, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xab, 0x26, 0x21, 0x11, 0x06, 0x00, 0x00,
}
