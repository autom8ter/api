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

type Config struct {
	Domain               string   `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         string   `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	Redirect             string   `protobuf:"bytes,4,opt,name=redirect,proto3" json:"redirect,omitempty"`
	Audience             string   `protobuf:"bytes,5,opt,name=audience,proto3" json:"audience,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Config) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Config) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *Config) GetRedirect() string {
	if m != nil {
		return m.Redirect
	}
	return ""
}

func (m *Config) GetAudience() string {
	if m != nil {
		return m.Audience
	}
	return ""
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "api.UserInfo")
	proto.RegisterType((*UserMetadata)(nil), "api.UserMetadata")
	proto.RegisterType((*AppMetadata)(nil), "api.AppMetadata")
	proto.RegisterType((*Config)(nil), "api.Config")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x4d, 0x6b, 0xdc, 0x30,
	0x10, 0x86, 0x71, 0x3e, 0x1c, 0x7b, 0xec, 0x2d, 0x89, 0x28, 0x45, 0xf4, 0x83, 0x06, 0xf7, 0x12,
	0x28, 0xe4, 0xd0, 0x40, 0x8f, 0x85, 0x92, 0x53, 0x0e, 0x2d, 0x34, 0x6d, 0xcf, 0x66, 0x62, 0x8d,
	0x77, 0x45, 0x6d, 0x49, 0x68, 0xe5, 0xc2, 0xfe, 0x95, 0xfe, 0x8b, 0xfe, 0xc3, 0xa0, 0x91, 0xbd,
	0xbb, 0x37, 0x3d, 0xef, 0xa3, 0x77, 0x2c, 0xd9, 0x86, 0x12, 0x9d, 0xbe, 0x75, 0xde, 0x06, 0x2b,
	0x4e, 0xd1, 0xe9, 0xe6, 0xff, 0x09, 0x14, 0xbf, 0xb7, 0xe4, 0x1f, 0x4c, 0x6f, 0x85, 0x80, 0x33,
	0x83, 0x23, 0xc9, 0xfc, 0x3a, 0xbb, 0x29, 0x1f, 0x79, 0x2d, 0xde, 0x01, 0xac, 0xf5, 0x5f, 0x32,
	0x2d, 0x9b, 0x0b, 0x36, 0x25, 0x27, 0xdf, 0xa3, 0x7e, 0x0f, 0x55, 0x8f, 0xa3, 0x1e, 0x76, 0xc9,
	0x17, 0xec, 0x21, 0x45, 0xbc, 0xe1, 0x15, 0xe4, 0x6b, 0x32, 0x8a, 0xbc, 0x2c, 0xd9, 0xcd, 0x24,
	0xde, 0x42, 0xf9, 0xa4, 0x7d, 0xd8, 0x28, 0x0c, 0x24, 0x21, 0x8d, 0xdd, 0x07, 0xe2, 0x25, 0x9c,
	0xd3, 0x88, 0x7a, 0x90, 0x15, 0x9b, 0x04, 0x42, 0xc2, 0x85, 0xd3, 0x5d, 0x98, 0x3c, 0xc9, 0x9a,
	0xf3, 0x05, 0xc5, 0x67, 0x58, 0x4d, 0x5b, 0xf2, 0xed, 0x48, 0x01, 0x15, 0x06, 0x94, 0xab, 0xeb,
	0xec, 0xa6, 0xfa, 0x74, 0x75, 0x1b, 0xaf, 0x1b, 0xef, 0xf7, 0x6d, 0x16, 0x8f, 0xf5, 0x74, 0x44,
	0xe2, 0x0e, 0x6a, 0x74, 0xee, 0x50, 0x7b, 0xc1, 0xb5, 0x4b, 0xae, 0x7d, 0x75, 0x6e, 0xdf, 0xaa,
	0xf0, 0x00, 0xcd, 0x0f, 0xa8, 0x8f, 0x47, 0xc6, 0xc3, 0xba, 0x8d, 0x35, 0x24, 0xb3, 0x74, 0x58,
	0x06, 0xf1, 0x11, 0xae, 0x9c, 0xa7, 0x9e, 0xbc, 0x27, 0xd5, 0x76, 0xd6, 0x04, 0xec, 0x82, 0x3c,
	0xe1, 0x1d, 0x97, 0x7b, 0x71, 0x9f, 0xf2, 0xe6, 0x0b, 0x54, 0x47, 0x8f, 0x8b, 0x1f, 0xc2, 0x0d,
	0x68, 0xe6, 0x81, 0xbc, 0x16, 0x6f, 0xa0, 0x74, 0xb8, 0x6b, 0x83, 0xfd, 0x43, 0x66, 0x9e, 0x53,
	0x38, 0xdc, 0xfd, 0x8a, 0xdc, 0xfc, 0xcb, 0x20, 0xbf, 0xb7, 0xa6, 0xd7, 0xeb, 0xf8, 0xc2, 0x95,
	0x1d, 0x51, 0x2f, 0xed, 0x99, 0x62, 0xbf, 0x1b, 0x34, 0x99, 0xd0, 0x6a, 0xb5, 0xf4, 0x53, 0xf0,
	0xa0, 0xc4, 0x07, 0x58, 0xcd, 0x72, 0x4b, 0x9d, 0xa7, 0x20, 0x4f, 0x79, 0x43, 0x9d, 0xc2, 0x9f,
	0x9c, 0x89, 0xd7, 0x50, 0x78, 0x52, 0xda, 0x53, 0x17, 0xe4, 0x59, 0x1a, 0xb0, 0x70, 0x74, 0x38,
	0x29, 0x4d, 0xa6, 0x23, 0x79, 0x9e, 0xdc, 0xc2, 0x4f, 0x39, 0xff, 0x6f, 0x77, 0xcf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xfc, 0xbb, 0x38, 0x54, 0x7c, 0x02, 0x00, 0x00,
}
