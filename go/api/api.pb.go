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

type AccessToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessToken) Reset()         { *m = AccessToken{} }
func (m *AccessToken) String() string { return proto.CompactTextString(m) }
func (*AccessToken) ProtoMessage()    {}
func (*AccessToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *AccessToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessToken.Unmarshal(m, b)
}
func (m *AccessToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessToken.Marshal(b, m, deterministic)
}
func (m *AccessToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessToken.Merge(m, src)
}
func (m *AccessToken) XXX_Size() int {
	return xxx_messageInfo_AccessToken.Size(m)
}
func (m *AccessToken) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessToken.DiscardUnknown(m)
}

var xxx_messageInfo_AccessToken proto.InternalMessageInfo

func (m *AccessToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type IDToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDToken) Reset()         { *m = IDToken{} }
func (m *IDToken) String() string { return proto.CompactTextString(m) }
func (*IDToken) ProtoMessage()    {}
func (*IDToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *IDToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDToken.Unmarshal(m, b)
}
func (m *IDToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDToken.Marshal(b, m, deterministic)
}
func (m *IDToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDToken.Merge(m, src)
}
func (m *IDToken) XXX_Size() int {
	return xxx_messageInfo_IDToken.Size(m)
}
func (m *IDToken) XXX_DiscardUnknown() {
	xxx_messageInfo_IDToken.DiscardUnknown(m)
}

var xxx_messageInfo_IDToken proto.InternalMessageInfo

func (m *IDToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RefreshToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshToken) Reset()         { *m = RefreshToken{} }
func (m *RefreshToken) String() string { return proto.CompactTextString(m) }
func (*RefreshToken) ProtoMessage()    {}
func (*RefreshToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *RefreshToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshToken.Unmarshal(m, b)
}
func (m *RefreshToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshToken.Marshal(b, m, deterministic)
}
func (m *RefreshToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshToken.Merge(m, src)
}
func (m *RefreshToken) XXX_Size() int {
	return xxx_messageInfo_RefreshToken.Size(m)
}
func (m *RefreshToken) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshToken.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshToken proto.InternalMessageInfo

func (m *RefreshToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Tokens struct {
	Id                   *IDToken      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Access               *AccessToken  `protobuf:"bytes,2,opt,name=access,proto3" json:"access,omitempty"`
	Refresh              *RefreshToken `protobuf:"bytes,3,opt,name=refresh,proto3" json:"refresh,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Tokens) Reset()         { *m = Tokens{} }
func (m *Tokens) String() string { return proto.CompactTextString(m) }
func (*Tokens) ProtoMessage()    {}
func (*Tokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *Tokens) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tokens.Unmarshal(m, b)
}
func (m *Tokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tokens.Marshal(b, m, deterministic)
}
func (m *Tokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tokens.Merge(m, src)
}
func (m *Tokens) XXX_Size() int {
	return xxx_messageInfo_Tokens.Size(m)
}
func (m *Tokens) XXX_DiscardUnknown() {
	xxx_messageInfo_Tokens.DiscardUnknown(m)
}

var xxx_messageInfo_Tokens proto.InternalMessageInfo

func (m *Tokens) GetId() *IDToken {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Tokens) GetAccess() *AccessToken {
	if m != nil {
		return m.Access
	}
	return nil
}

func (m *Tokens) GetRefresh() *RefreshToken {
	if m != nil {
		return m.Refresh
	}
	return nil
}

type Paths struct {
	Home                 string   `protobuf:"bytes,1,opt,name=home,proto3" json:"home,omitempty"`
	Dashboard            string   `protobuf:"bytes,2,opt,name=dashboard,proto3" json:"dashboard,omitempty"`
	Logout               string   `protobuf:"bytes,4,opt,name=logout,proto3" json:"logout,omitempty"`
	Callback             string   `protobuf:"bytes,5,opt,name=callback,proto3" json:"callback,omitempty"`
	Login                string   `protobuf:"bytes,6,opt,name=login,proto3" json:"login,omitempty"`
	LoggedOutReturnTo    string   `protobuf:"bytes,7,opt,name=logged_out_return_to,json=loggedOutReturnTo,proto3" json:"logged_out_return_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Paths) Reset()         { *m = Paths{} }
func (m *Paths) String() string { return proto.CompactTextString(m) }
func (*Paths) ProtoMessage()    {}
func (*Paths) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *Paths) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Paths.Unmarshal(m, b)
}
func (m *Paths) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Paths.Marshal(b, m, deterministic)
}
func (m *Paths) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paths.Merge(m, src)
}
func (m *Paths) XXX_Size() int {
	return xxx_messageInfo_Paths.Size(m)
}
func (m *Paths) XXX_DiscardUnknown() {
	xxx_messageInfo_Paths.DiscardUnknown(m)
}

var xxx_messageInfo_Paths proto.InternalMessageInfo

func (m *Paths) GetHome() string {
	if m != nil {
		return m.Home
	}
	return ""
}

func (m *Paths) GetDashboard() string {
	if m != nil {
		return m.Dashboard
	}
	return ""
}

func (m *Paths) GetLogout() string {
	if m != nil {
		return m.Logout
	}
	return ""
}

func (m *Paths) GetCallback() string {
	if m != nil {
		return m.Callback
	}
	return ""
}

func (m *Paths) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *Paths) GetLoggedOutReturnTo() string {
	if m != nil {
		return m.LoggedOutReturnTo
	}
	return ""
}

type Auth0 struct {
	Domain               string   `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         string   `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	Scopes               []string `protobuf:"bytes,4,rep,name=scopes,proto3" json:"scopes,omitempty"`
	Redirect             string   `protobuf:"bytes,5,opt,name=redirect,proto3" json:"redirect,omitempty"`
	ResourceUrl          string   `protobuf:"bytes,6,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth0) Reset()         { *m = Auth0{} }
func (m *Auth0) String() string { return proto.CompactTextString(m) }
func (*Auth0) ProtoMessage()    {}
func (*Auth0) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *Auth0) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth0.Unmarshal(m, b)
}
func (m *Auth0) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth0.Marshal(b, m, deterministic)
}
func (m *Auth0) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth0.Merge(m, src)
}
func (m *Auth0) XXX_Size() int {
	return xxx_messageInfo_Auth0.Size(m)
}
func (m *Auth0) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth0.DiscardUnknown(m)
}

var xxx_messageInfo_Auth0 proto.InternalMessageInfo

func (m *Auth0) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Auth0) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Auth0) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *Auth0) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *Auth0) GetRedirect() string {
	if m != nil {
		return m.Redirect
	}
	return ""
}

func (m *Auth0) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "api.UserInfo")
	proto.RegisterType((*UserMetadata)(nil), "api.UserMetadata")
	proto.RegisterType((*AppMetadata)(nil), "api.AppMetadata")
	proto.RegisterType((*AccessToken)(nil), "api.AccessToken")
	proto.RegisterType((*IDToken)(nil), "api.IDToken")
	proto.RegisterType((*RefreshToken)(nil), "api.RefreshToken")
	proto.RegisterType((*Tokens)(nil), "api.Tokens")
	proto.RegisterType((*Paths)(nil), "api.Paths")
	proto.RegisterType((*Auth0)(nil), "api.Auth0")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4f, 0x6b, 0xdc, 0x3e,
	0x10, 0x65, 0x37, 0xd9, 0x3f, 0x1e, 0x3b, 0x3f, 0x12, 0x11, 0x7e, 0x88, 0x36, 0x25, 0xa9, 0xd3,
	0x43, 0x20, 0x90, 0x96, 0x04, 0x7a, 0x2c, 0x84, 0xf6, 0x92, 0x43, 0xff, 0xb9, 0xc9, 0xd9, 0x68,
	0xed, 0xd9, 0xb5, 0x88, 0xd7, 0x12, 0x92, 0x5c, 0xc8, 0xa1, 0x1f, 0xaa, 0xd0, 0x43, 0x3f, 0x5e,
	0xd1, 0x48, 0xde, 0xdd, 0x53, 0x6e, 0x7a, 0xef, 0x8d, 0xc6, 0x6f, 0x9e, 0x24, 0x43, 0x22, 0xb4,
	0xbc, 0xd2, 0x46, 0x39, 0xc5, 0xf6, 0x84, 0x96, 0xf9, 0xef, 0x31, 0xcc, 0x1f, 0x2c, 0x9a, 0xbb,
	0x6e, 0xa9, 0x18, 0x83, 0xfd, 0x4e, 0xac, 0x91, 0x4f, 0xcf, 0x46, 0x17, 0x49, 0x41, 0x6b, 0xf6,
	0x0a, 0x60, 0x25, 0x7f, 0x62, 0x57, 0x92, 0x32, 0x23, 0x25, 0x21, 0xe6, 0x8b, 0x97, 0x4f, 0x21,
	0x5d, 0x8a, 0xb5, 0x6c, 0x9f, 0x82, 0x3e, 0x27, 0x1d, 0x02, 0x45, 0x05, 0xff, 0xc3, 0x74, 0x85,
	0x5d, 0x8d, 0x86, 0x27, 0xa4, 0x45, 0xc4, 0x4e, 0x20, 0x59, 0x48, 0xe3, 0x9a, 0x5a, 0x38, 0xe4,
	0x10, 0xda, 0x6e, 0x08, 0x76, 0x0c, 0x13, 0x5c, 0x0b, 0xd9, 0xf2, 0x94, 0x94, 0x00, 0x18, 0x87,
	0x99, 0x96, 0x95, 0xeb, 0x0d, 0xf2, 0x8c, 0xf8, 0x01, 0xb2, 0xf7, 0x70, 0xd0, 0x5b, 0x34, 0xe5,
	0x1a, 0x9d, 0xa8, 0x85, 0x13, 0xfc, 0xe0, 0x6c, 0x74, 0x91, 0x5e, 0x1f, 0x5d, 0xf9, 0x71, 0xfd,
	0x7c, 0x9f, 0xa3, 0x50, 0x64, 0xfd, 0x0e, 0x62, 0x37, 0x90, 0x09, 0xad, 0xb7, 0xdb, 0xfe, 0xa3,
	0x6d, 0x87, 0xb4, 0xed, 0x56, 0xeb, 0xcd, 0xae, 0x54, 0x6c, 0x41, 0xfe, 0x1d, 0xb2, 0xdd, 0x96,
	0xde, 0xac, 0x6e, 0x54, 0x87, 0x7c, 0x14, 0xcc, 0x12, 0x60, 0x97, 0x70, 0xa4, 0x0d, 0x2e, 0xd1,
	0x18, 0xac, 0xcb, 0x4a, 0x75, 0x4e, 0x54, 0x8e, 0x8f, 0xa9, 0xe2, 0x70, 0x23, 0x7c, 0x0c, 0x7c,
	0xfe, 0x01, 0xd2, 0x9d, 0xcf, 0xf9, 0x83, 0xd0, 0xad, 0xe8, 0x62, 0x43, 0x5a, 0xb3, 0x97, 0x90,
	0x68, 0xf1, 0x54, 0x3a, 0xf5, 0x88, 0x5d, 0xec, 0x33, 0xd7, 0xe2, 0xe9, 0xde, 0xe3, 0xfc, 0x1c,
	0xd2, 0xdb, 0xaa, 0x42, 0x6b, 0x09, 0x7a, 0x47, 0xa1, 0x2e, 0x3a, 0x22, 0x90, 0x9f, 0xc2, 0xec,
	0xee, 0xd3, 0x73, 0x05, 0x6f, 0x20, 0x2b, 0x70, 0x69, 0xd0, 0x36, 0xcf, 0x55, 0xfd, 0x82, 0x29,
	0xc9, 0x96, 0x9d, 0xc0, 0x58, 0xd6, 0x24, 0xa6, 0xd7, 0x19, 0x65, 0x16, 0xfb, 0x17, 0x63, 0x59,
	0xb3, 0x0b, 0x98, 0x0a, 0xf2, 0x44, 0x6e, 0x37, 0xa9, 0x6e, 0x6d, 0x16, 0x51, 0x67, 0x97, 0x30,
	0x33, 0xe1, 0xbb, 0x7c, 0x6f, 0xe7, 0xdc, 0x76, 0xbd, 0x14, 0x43, 0x45, 0xfe, 0x67, 0x04, 0x93,
	0x6f, 0xc2, 0x35, 0xd6, 0xa7, 0xd4, 0xa8, 0xf5, 0x10, 0x3b, 0xad, 0xfd, 0xb5, 0xaa, 0x85, 0x6d,
	0x16, 0x4a, 0x98, 0x3a, 0xa6, 0xb4, 0x25, 0xfc, 0x65, 0x6c, 0xd5, 0x4a, 0xf5, 0x8e, 0xef, 0x87,
	0xcb, 0x18, 0x10, 0x7b, 0x01, 0xf3, 0x4a, 0xb4, 0xed, 0x42, 0x54, 0x8f, 0x7c, 0x12, 0xa2, 0x1d,
	0xb0, 0x0f, 0xa1, 0x55, 0x2b, 0xd9, 0xc5, 0x57, 0x11, 0x00, 0x7b, 0x0b, 0xc7, 0xad, 0x5a, 0xad,
	0xb0, 0x2e, 0x55, 0xef, 0x4a, 0x83, 0xae, 0x37, 0x5d, 0xe9, 0x54, 0x7c, 0x20, 0x47, 0x41, 0xfb,
	0xda, 0xbb, 0x82, 0x94, 0x7b, 0x95, 0xff, 0x1d, 0xc1, 0xe4, 0xb6, 0x77, 0xcd, 0x3b, 0x6f, 0xa2,
	0x56, 0x6b, 0x21, 0x87, 0x58, 0x23, 0xf2, 0x07, 0x5c, 0xb5, 0x12, 0x3b, 0x57, 0xca, 0xc1, 0xfa,
	0x3c, 0x10, 0x77, 0x35, 0x3b, 0x87, 0x83, 0x28, 0x5a, 0xac, 0x0c, 0x3a, 0x0a, 0x2a, 0x29, 0xb2,
	0x40, 0xfe, 0x20, 0xce, 0x77, 0xb6, 0x95, 0xd2, 0x68, 0xf9, 0xfe, 0xd9, 0x9e, 0xef, 0x1c, 0x90,
	0x1f, 0xcf, 0x60, 0x2d, 0x0d, 0x56, 0x6e, 0x18, 0x6f, 0xc0, 0xec, 0x35, 0x64, 0x06, 0xad, 0xea,
	0x4d, 0x85, 0x65, 0x6f, 0xda, 0x38, 0x65, 0x3a, 0x70, 0x0f, 0xa6, 0x5d, 0x4c, 0xe9, 0x7f, 0x71,
	0xf3, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x44, 0x7b, 0x8f, 0xf0, 0x3c, 0x04, 0x00, 0x00,
}
