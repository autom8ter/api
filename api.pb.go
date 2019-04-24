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
	Scopes               []string `protobuf:"bytes,6,rep,name=scopes,proto3" json:"scopes,omitempty"`
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

func (m *Config) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "api.UserInfo")
	proto.RegisterType((*UserMetadata)(nil), "api.UserMetadata")
	proto.RegisterType((*AppMetadata)(nil), "api.AppMetadata")
	proto.RegisterType((*Config)(nil), "api.Config")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x4f, 0x6b, 0xdc, 0x30,
	0x10, 0xc5, 0xd9, 0x6c, 0xe2, 0xac, 0xc7, 0xde, 0x92, 0x88, 0x52, 0x44, 0xff, 0xd0, 0x65, 0x7b,
	0x09, 0x14, 0x72, 0x68, 0xa0, 0xc7, 0x42, 0xc9, 0x29, 0x87, 0x16, 0x9a, 0xb6, 0x67, 0x33, 0x91,
	0xc6, 0x1b, 0x51, 0x5b, 0x12, 0xb2, 0x5c, 0xd8, 0x8f, 0xd5, 0x7e, 0xc2, 0xa0, 0x91, 0xbd, 0xd9,
	0x9b, 0xde, 0xfb, 0xcd, 0x7b, 0x1e, 0xd9, 0x86, 0x12, 0xbd, 0xb9, 0xf6, 0xc1, 0x45, 0x27, 0x96,
	0xe8, 0xcd, 0xf6, 0xdf, 0x09, 0xac, 0x7e, 0x0f, 0x14, 0xee, 0x6c, 0xeb, 0x84, 0x80, 0x53, 0x8b,
	0x3d, 0xc9, 0x62, 0xb3, 0xb8, 0x2a, 0xef, 0xf9, 0x2c, 0xde, 0x01, 0xec, 0xcc, 0x5f, 0xb2, 0x0d,
	0x93, 0x73, 0x26, 0x25, 0x3b, 0xdf, 0x13, 0x7e, 0x0f, 0x55, 0x8b, 0xbd, 0xe9, 0xf6, 0x99, 0xaf,
	0x98, 0x43, 0xb6, 0x78, 0xe0, 0x15, 0x14, 0x3b, 0xb2, 0x9a, 0x82, 0x2c, 0x99, 0x4d, 0x4a, 0xbc,
	0x85, 0xf2, 0xc1, 0x84, 0xf8, 0xa8, 0x31, 0x92, 0x84, 0x5c, 0x7b, 0x30, 0xc4, 0x4b, 0x38, 0xa3,
	0x1e, 0x4d, 0x27, 0x2b, 0x26, 0x59, 0x08, 0x09, 0xe7, 0xde, 0xa8, 0x38, 0x06, 0x92, 0x35, 0xfb,
	0xb3, 0x14, 0x9f, 0x61, 0x3d, 0x0e, 0x14, 0x9a, 0x9e, 0x22, 0x6a, 0x8c, 0x28, 0xd7, 0x9b, 0xc5,
	0x55, 0xf5, 0xe9, 0xf2, 0x3a, 0x5d, 0x37, 0xdd, 0xef, 0xdb, 0x04, 0xee, 0xeb, 0xf1, 0x48, 0x89,
	0x1b, 0xa8, 0xd1, 0xfb, 0xe7, 0xd8, 0x0b, 0x8e, 0x5d, 0x70, 0xec, 0xab, 0xf7, 0x87, 0x54, 0x85,
	0xcf, 0x62, 0xfb, 0x03, 0xea, 0xe3, 0xca, 0xb4, 0xac, 0x7f, 0x74, 0x96, 0xe4, 0x22, 0x2f, 0xcb,
	0x42, 0x7c, 0x84, 0x4b, 0x1f, 0xa8, 0xa5, 0x10, 0x48, 0x37, 0xca, 0xd9, 0x88, 0x2a, 0xca, 0x13,
	0x9e, 0xb8, 0x38, 0x80, 0xdb, 0xec, 0x6f, 0xbf, 0x40, 0x75, 0xf4, 0xb8, 0xf4, 0x21, 0x7c, 0x87,
	0x76, 0x2a, 0xe4, 0xb3, 0x78, 0x03, 0xa5, 0xc7, 0x7d, 0x13, 0xdd, 0x1f, 0xb2, 0x53, 0xcf, 0xca,
	0xe3, 0xfe, 0x57, 0xd2, 0xdb, 0xff, 0x0b, 0x28, 0x6e, 0x9d, 0x6d, 0xcd, 0x2e, 0xbd, 0x70, 0xed,
	0x7a, 0x34, 0x73, 0x7a, 0x52, 0x29, 0xaf, 0x3a, 0x43, 0x36, 0x36, 0x46, 0xcf, 0xf9, 0x6c, 0xdc,
	0x69, 0xf1, 0x01, 0xd6, 0x13, 0x1c, 0x48, 0x05, 0x8a, 0x72, 0xc9, 0x03, 0x75, 0x36, 0x7f, 0xb2,
	0x27, 0x5e, 0xc3, 0x2a, 0x90, 0x36, 0x81, 0x54, 0x94, 0xa7, 0xb9, 0x60, 0xd6, 0x89, 0xe1, 0xa8,
	0x0d, 0x59, 0x45, 0xf2, 0x2c, 0xb3, 0x59, 0xa7, 0x8d, 0x06, 0xe5, 0x3c, 0x0d, 0xb2, 0xd8, 0x2c,
	0xd3, 0x46, 0x59, 0x3d, 0x14, 0xfc, 0x1f, 0xde, 0x3c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x56, 0x1d,
	0xb0, 0x1a, 0x94, 0x02, 0x00, 0x00,
}
