// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage/storage.proto

package storage

import (
	fmt "fmt"
	auth "github.com/autom8ter/api/sdk/go/auth"
	os "github.com/autom8ter/api/sdk/go/os"
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

type CreateObjectRequest struct {
	Bucket               string   `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	File                 *os.File `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateObjectRequest) Reset()         { *m = CreateObjectRequest{} }
func (m *CreateObjectRequest) String() string { return proto.CompactTextString(m) }
func (*CreateObjectRequest) ProtoMessage()    {}
func (*CreateObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_803a4779d728664c, []int{0}
}

func (m *CreateObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateObjectRequest.Unmarshal(m, b)
}
func (m *CreateObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateObjectRequest.Marshal(b, m, deterministic)
}
func (m *CreateObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateObjectRequest.Merge(m, src)
}
func (m *CreateObjectRequest) XXX_Size() int {
	return xxx_messageInfo_CreateObjectRequest.Size(m)
}
func (m *CreateObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateObjectRequest proto.InternalMessageInfo

func (m *CreateObjectRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *CreateObjectRequest) GetFile() *os.File {
	if m != nil {
		return m.File
	}
	return nil
}

type Object struct {
	Bucket               string            `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ContentType          string            `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	ContentLanguage      string            `protobuf:"bytes,4,opt,name=content_language,json=contentLanguage,proto3" json:"content_language,omitempty"`
	CacheControl         string            `protobuf:"bytes,5,opt,name=cache_control,json=cacheControl,proto3" json:"cache_control,omitempty"`
	Owner                string            `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Size                 int64             `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	Meta                 map[string]string `protobuf:"bytes,8,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Created              int64             `protobuf:"varint,9,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64             `protobuf:"varint,10,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted              int64             `protobuf:"varint,11,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_803a4779d728664c, []int{1}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *Object) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Object) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Object) GetContentLanguage() string {
	if m != nil {
		return m.ContentLanguage
	}
	return ""
}

func (m *Object) GetCacheControl() string {
	if m != nil {
		return m.CacheControl
	}
	return ""
}

func (m *Object) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Object) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Object) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Object) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Object) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *Object) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

type Bucket struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location             string            `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	EncryptionKey        string            `protobuf:"bytes,3,opt,name=encryption_key,json=encryptionKey,proto3" json:"encryption_key,omitempty"`
	Versioning           bool              `protobuf:"varint,4,opt,name=versioning,proto3" json:"versioning,omitempty"`
	Class                string            `protobuf:"bytes,5,opt,name=class,proto3" json:"class,omitempty"`
	Cors                 *auth.CORS        `protobuf:"bytes,6,opt,name=cors,proto3" json:"cors,omitempty"`
	Rules                []*auth.ACLRule   `protobuf:"bytes,7,rep,name=rules,proto3" json:"rules,omitempty"`
	Meta                 map[string]string `protobuf:"bytes,8,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Created              int64             `protobuf:"varint,9,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64             `protobuf:"varint,10,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted              int64             `protobuf:"varint,11,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_803a4779d728664c, []int{2}
}

func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Bucket) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Bucket) GetEncryptionKey() string {
	if m != nil {
		return m.EncryptionKey
	}
	return ""
}

func (m *Bucket) GetVersioning() bool {
	if m != nil {
		return m.Versioning
	}
	return false
}

func (m *Bucket) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *Bucket) GetCors() *auth.CORS {
	if m != nil {
		return m.Cors
	}
	return nil
}

func (m *Bucket) GetRules() []*auth.ACLRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *Bucket) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Bucket) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Bucket) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *Bucket) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateObjectRequest)(nil), "storage.CreateObjectRequest")
	proto.RegisterType((*Object)(nil), "storage.Object")
	proto.RegisterMapType((map[string]string)(nil), "storage.Object.MetaEntry")
	proto.RegisterType((*Bucket)(nil), "storage.Bucket")
	proto.RegisterMapType((map[string]string)(nil), "storage.Bucket.MetaEntry")
}

func init() { proto.RegisterFile("storage/storage.proto", fileDescriptor_803a4779d728664c) }

var fileDescriptor_803a4779d728664c = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0x4f, 0x6f, 0x13, 0x3d,
	0x10, 0xc6, 0xb5, 0xdd, 0x64, 0x93, 0xcc, 0x36, 0x6f, 0x2b, 0xbf, 0x80, 0x4c, 0x84, 0xaa, 0x90,
	0x0a, 0x29, 0x1c, 0xd8, 0x48, 0xe1, 0x00, 0xe2, 0x06, 0x11, 0x5c, 0x5a, 0x54, 0xc9, 0x70, 0x8f,
	0x1c, 0x67, 0x48, 0x97, 0xba, 0xf6, 0x62, 0x7b, 0x8b, 0x96, 0x2f, 0xc0, 0x95, 0x8f, 0x8c, 0xfc,
	0x27, 0xa5, 0x1c, 0x7a, 0xe7, 0x92, 0xf8, 0xf9, 0x3d, 0xb3, 0x63, 0xfb, 0x99, 0x5d, 0x78, 0x68,
	0x9d, 0x36, 0x7c, 0x87, 0x8b, 0xf4, 0x5f, 0x35, 0x46, 0x3b, 0x4d, 0x06, 0x49, 0x4e, 0x4a, 0x6d,
	0x17, 0xda, 0x46, 0x3a, 0x39, 0xe2, 0xad, 0xbb, 0x5c, 0xf8, 0x9f, 0x08, 0x66, 0x67, 0xf0, 0xff,
	0xca, 0x20, 0x77, 0x78, 0xb1, 0xf9, 0x8a, 0xc2, 0x31, 0xfc, 0xd6, 0xa2, 0x75, 0xe4, 0x11, 0x14,
	0x9b, 0x56, 0x5c, 0xa1, 0xa3, 0xd9, 0x34, 0x9b, 0x8f, 0x58, 0x52, 0xe4, 0x09, 0xf4, 0xbe, 0xd4,
	0x12, 0xe9, 0xc1, 0x34, 0x9b, 0x97, 0xcb, 0x61, 0xa5, 0x6d, 0xf5, 0xa1, 0x96, 0xc8, 0x02, 0x9d,
	0xfd, 0xcc, 0xa1, 0x88, 0x7d, 0xee, 0x6d, 0x40, 0xa0, 0xa7, 0xf8, 0x75, 0x6c, 0x30, 0x62, 0x61,
	0x4d, 0x9e, 0xc2, 0xa1, 0xd0, 0xca, 0xa1, 0x72, 0x6b, 0xd7, 0x35, 0x48, 0xf3, 0xe0, 0x95, 0x89,
	0x7d, 0xee, 0x1a, 0x24, 0xcf, 0xe1, 0x78, 0x5f, 0x22, 0xb9, 0xda, 0xb5, 0x7c, 0x87, 0xb4, 0x17,
	0xca, 0x8e, 0x12, 0x3f, 0x4f, 0x98, 0x9c, 0xc2, 0x58, 0x70, 0x71, 0x89, 0x6b, 0x6f, 0x18, 0x2d,
	0x69, 0x3f, 0xd4, 0x1d, 0x06, 0xb8, 0x8a, 0x8c, 0x3c, 0x80, 0xbe, 0xfe, 0xae, 0xd0, 0xd0, 0x22,
	0x98, 0x51, 0xf8, 0xc3, 0xd9, 0xfa, 0x07, 0xd2, 0xc1, 0x34, 0x9b, 0xe7, 0x2c, 0xac, 0xc9, 0x0b,
	0xe8, 0x5d, 0xa3, 0xe3, 0x74, 0x38, 0xcd, 0xe7, 0xe5, 0xf2, 0x71, 0xb5, 0x4f, 0x39, 0xde, 0xb3,
	0xfa, 0x88, 0x8e, 0xbf, 0x57, 0xce, 0x74, 0x2c, 0x94, 0x11, 0x0a, 0x03, 0x11, 0xf2, 0xdc, 0xd2,
	0x51, 0xe8, 0xb2, 0x97, 0xde, 0x69, 0x9b, 0x6d, 0x70, 0x20, 0x3a, 0x49, 0x7a, 0x67, 0x8b, 0x12,
	0xbd, 0x53, 0x46, 0x27, 0xc9, 0xc9, 0x2b, 0x18, 0xdd, 0x6e, 0x40, 0x8e, 0x21, 0xbf, 0xc2, 0x2e,
	0xe5, 0xe9, 0x97, 0xfe, 0x16, 0x37, 0x5c, 0xb6, 0xfb, 0x34, 0xa3, 0x78, 0x73, 0xf0, 0x3a, 0x9b,
	0xfd, 0xca, 0xa1, 0x78, 0xf7, 0x77, 0xe2, 0xd9, 0x9d, 0xc4, 0x27, 0x30, 0x94, 0x5a, 0x70, 0x57,
	0x6b, 0x95, 0x9e, 0xbd, 0xd5, 0xe4, 0x19, 0xfc, 0x87, 0x4a, 0x98, 0xae, 0xf1, 0x6a, 0xed, 0x77,
	0x8c, 0xf3, 0x18, 0xff, 0xa1, 0x67, 0xd8, 0x91, 0x13, 0x80, 0x1b, 0x34, 0xb6, 0xd6, 0xaa, 0x56,
	0xbb, 0x30, 0x8b, 0x21, 0xbb, 0x43, 0xfc, 0xd9, 0x84, 0xe4, 0xd6, 0xa6, 0xf8, 0xa3, 0x20, 0x27,
	0xd0, 0x13, 0xda, 0xd8, 0x10, 0x7b, 0xb9, 0x84, 0x2a, 0xbc, 0x89, 0xab, 0x0b, 0xf6, 0x89, 0x05,
	0x4e, 0x4e, 0xa1, 0x6f, 0x5a, 0x89, 0x96, 0x0e, 0x42, 0xdc, 0xe3, 0x58, 0xf0, 0x76, 0x75, 0xce,
	0x5a, 0x89, 0x2c, 0x7a, 0xf7, 0x8e, 0x24, 0x5e, 0xf8, 0x1f, 0x1d, 0xc9, 0xa6, 0x08, 0x1f, 0xdc,
	0xcb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0x3a, 0x31, 0x50, 0xb0, 0x03, 0x00, 0x00,
}
