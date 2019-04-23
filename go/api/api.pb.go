// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

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

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetGivenName() string {
	if m != nil {
		return m.GivenName
	}
	return ""
}

func (m *UserInfo) GetFamilyName() string {
	if m != nil {
		return m.FamilyName
	}
	return ""
}

func (m *UserInfo) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UserInfo) GetBirthdate() string {
	if m != nil {
		return m.Birthdate
	}
	return ""
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfo) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *UserInfo) GetUserMetadata() *UserMetadata {
	if m != nil {
		return m.UserMetadata
	}
	return nil
}

func (m *UserInfo) GetAppMetadata() *AppMetadata {
	if m != nil {
		return m.AppMetadata
	}
	return nil
}

type UserMetadata struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	PreferredContact     string   `protobuf:"bytes,2,opt,name=preferred_contact,json=preferredContact,proto3" json:"preferred_contact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserMetadata) Reset()         { *m = UserMetadata{} }
func (m *UserMetadata) String() string { return proto.CompactTextString(m) }
func (*UserMetadata) ProtoMessage()    {}
func (*UserMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *UserMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMetadata.Unmarshal(m, b)
}
func (m *UserMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMetadata.Marshal(b, m, deterministic)
}
func (m *UserMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMetadata.Merge(m, src)
}
func (m *UserMetadata) XXX_Size() int {
	return xxx_messageInfo_UserMetadata.Size(m)
}
func (m *UserMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_UserMetadata proto.InternalMessageInfo

func (m *UserMetadata) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserMetadata) GetPreferredContact() string {
	if m != nil {
		return m.PreferredContact
	}
	return ""
}

type AppMetadata struct {
	Plan                 string   `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	PayToken             string   `protobuf:"bytes,2,opt,name=pay_token,json=payToken,proto3" json:"pay_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppMetadata) Reset()         { *m = AppMetadata{} }
func (m *AppMetadata) String() string { return proto.CompactTextString(m) }
func (*AppMetadata) ProtoMessage()    {}
func (*AppMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *AppMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppMetadata.Unmarshal(m, b)
}
func (m *AppMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppMetadata.Marshal(b, m, deterministic)
}
func (m *AppMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppMetadata.Merge(m, src)
}
func (m *AppMetadata) XXX_Size() int {
	return xxx_messageInfo_AppMetadata.Size(m)
}
func (m *AppMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_AppMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_AppMetadata proto.InternalMessageInfo

func (m *AppMetadata) GetPlan() string {
	if m != nil {
		return m.Plan
	}
	return ""
}

func (m *AppMetadata) GetPayToken() string {
	if m != nil {
		return m.PayToken
	}
	return ""
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "api.UserInfo")
	proto.RegisterType((*UserMetadata)(nil), "api.UserMetadata")
	proto.RegisterType((*AppMetadata)(nil), "api.AppMetadata")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x41, 0x4b, 0x3b, 0x31,
	0x10, 0xc5, 0x69, 0xff, 0x7f, 0xdb, 0xee, 0xec, 0x56, 0xda, 0x20, 0x12, 0x50, 0xb1, 0xf4, 0x54,
	0x10, 0x7a, 0xb0, 0xe0, 0x51, 0x10, 0x4f, 0x1e, 0x14, 0x2c, 0x7a, 0x5e, 0xa6, 0xdd, 0x69, 0x1b,
	0xec, 0x66, 0x87, 0x34, 0x15, 0xfa, 0xb1, 0xfc, 0x86, 0x92, 0xc9, 0xda, 0xee, 0x2d, 0xef, 0xfd,
	0xf2, 0x5e, 0x32, 0x0c, 0x24, 0xc8, 0x66, 0xca, 0xae, 0xf2, 0x95, 0xfa, 0x87, 0x6c, 0xc6, 0x3f,
	0x6d, 0xe8, 0x7d, 0xee, 0xc8, 0xbd, 0xd8, 0x55, 0xa5, 0x14, 0xfc, 0xb7, 0x58, 0x92, 0xee, 0x8c,
	0x5a, 0x93, 0x64, 0x2e, 0x67, 0x75, 0x03, 0xb0, 0x36, 0xdf, 0x64, 0x73, 0x21, 0x5d, 0x21, 0x89,
	0x38, 0x6f, 0x01, 0xdf, 0x42, 0xba, 0xc2, 0xd2, 0x6c, 0x0f, 0x91, 0xf7, 0x84, 0x43, 0xb4, 0xe4,
	0xc2, 0x25, 0x74, 0xd6, 0x64, 0x0b, 0x72, 0x3a, 0x11, 0x56, 0x2b, 0x75, 0x0d, 0xc9, 0xc2, 0x38,
	0xbf, 0x29, 0xd0, 0x93, 0x86, 0x58, 0x7b, 0x34, 0xd4, 0x05, 0x9c, 0x51, 0x89, 0x66, 0xab, 0x53,
	0x21, 0x51, 0x28, 0x0d, 0x5d, 0x36, 0x4b, 0xbf, 0x77, 0xa4, 0x33, 0xf1, 0xff, 0xa4, 0x7a, 0x80,
	0xfe, 0x7e, 0x47, 0x2e, 0x2f, 0xc9, 0x63, 0x81, 0x1e, 0x75, 0x7f, 0xd4, 0x9a, 0xa4, 0xf7, 0xc3,
	0x69, 0x18, 0x37, 0xcc, 0xf7, 0x5a, 0x83, 0x79, 0xb6, 0x6f, 0x28, 0x35, 0x83, 0x0c, 0x99, 0x4f,
	0xb1, 0x73, 0x89, 0x0d, 0x24, 0xf6, 0xc4, 0x7c, 0x4c, 0xa5, 0x78, 0x12, 0xe3, 0x77, 0xc8, 0x9a,
	0x95, 0xe1, 0xb3, 0xbc, 0xa9, 0x2c, 0xe9, 0x56, 0xfc, 0xac, 0x08, 0x75, 0x07, 0x43, 0x76, 0xb4,
	0x22, 0xe7, 0xa8, 0xc8, 0x97, 0x95, 0xf5, 0xb8, 0xf4, 0xba, 0x2d, 0x37, 0x06, 0x47, 0xf0, 0x1c,
	0xfd, 0xf1, 0x23, 0xa4, 0x8d, 0xe7, 0xc2, 0x22, 0x78, 0x8b, 0xb6, 0x2e, 0x94, 0xb3, 0xba, 0x82,
	0x84, 0xf1, 0x90, 0xfb, 0xea, 0x8b, 0x6c, 0xdd, 0xd3, 0x63, 0x3c, 0x7c, 0x04, 0xbd, 0xe8, 0xc8,
	0x4a, 0x67, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x95, 0x55, 0xb8, 0xdf, 0x01, 0x00, 0x00,
}
