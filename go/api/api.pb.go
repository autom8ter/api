// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type GetProfileByEmail struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProfileByEmail) Reset()         { *m = GetProfileByEmail{} }
func (m *GetProfileByEmail) String() string { return proto.CompactTextString(m) }
func (*GetProfileByEmail) ProtoMessage()    {}
func (*GetProfileByEmail) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *GetProfileByEmail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileByEmail.Unmarshal(m, b)
}
func (m *GetProfileByEmail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileByEmail.Marshal(b, m, deterministic)
}
func (m *GetProfileByEmail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileByEmail.Merge(m, src)
}
func (m *GetProfileByEmail) XXX_Size() int {
	return xxx_messageInfo_GetProfileByEmail.Size(m)
}
func (m *GetProfileByEmail) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileByEmail.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileByEmail proto.InternalMessageInfo

func (m *GetProfileByEmail) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Profile struct {
	Email                string        `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	EmailVerified        bool          `protobuf:"varint,2,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	Name                 string        `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	GivenName            string        `protobuf:"bytes,4,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	FamilyName           string        `protobuf:"bytes,5,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	Picture              string        `protobuf:"bytes,6,opt,name=picture,proto3" json:"picture,omitempty"`
	Locale               string        `protobuf:"bytes,7,opt,name=locale,proto3" json:"locale,omitempty"`
	UserId               string        `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Nickname             string        `protobuf:"bytes,9,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Connection           string        `protobuf:"bytes,10,opt,name=connection,proto3" json:"connection,omitempty"`
	Identities           []*Identity   `protobuf:"bytes,11,rep,name=identities,proto3" json:"identities,omitempty"`
	LastIp               string        `protobuf:"bytes,12,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
	LoginCount           int64         `protobuf:"varint,13,opt,name=login_count,json=loginCount,proto3" json:"login_count,omitempty"`
	UpdatedAt            string        `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt            string        `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Sub                  string        `protobuf:"bytes,16,opt,name=sub,proto3" json:"sub,omitempty"`
	Iss                  string        `protobuf:"bytes,17,opt,name=iss,proto3" json:"iss,omitempty"`
	Aud                  string        `protobuf:"bytes,18,opt,name=aud,proto3" json:"aud,omitempty"`
	Iat                  string        `protobuf:"bytes,19,opt,name=iat,proto3" json:"iat,omitempty"`
	UserMetadata         *UserMetadata `protobuf:"bytes,20,opt,name=user_metadata,json=userMetadata,proto3" json:"user_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Profile) GetEmailVerified() bool {
	if m != nil {
		return m.EmailVerified
	}
	return false
}

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetGivenName() string {
	if m != nil {
		return m.GivenName
	}
	return ""
}

func (m *Profile) GetFamilyName() string {
	if m != nil {
		return m.FamilyName
	}
	return ""
}

func (m *Profile) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *Profile) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *Profile) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Profile) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Profile) GetConnection() string {
	if m != nil {
		return m.Connection
	}
	return ""
}

func (m *Profile) GetIdentities() []*Identity {
	if m != nil {
		return m.Identities
	}
	return nil
}

func (m *Profile) GetLastIp() string {
	if m != nil {
		return m.LastIp
	}
	return ""
}

func (m *Profile) GetLoginCount() int64 {
	if m != nil {
		return m.LoginCount
	}
	return 0
}

func (m *Profile) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Profile) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Profile) GetSub() string {
	if m != nil {
		return m.Sub
	}
	return ""
}

func (m *Profile) GetIss() string {
	if m != nil {
		return m.Iss
	}
	return ""
}

func (m *Profile) GetAud() string {
	if m != nil {
		return m.Aud
	}
	return ""
}

func (m *Profile) GetIat() string {
	if m != nil {
		return m.Iat
	}
	return ""
}

func (m *Profile) GetUserMetadata() *UserMetadata {
	if m != nil {
		return m.UserMetadata
	}
	return nil
}

type UserMetadata struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Plan                 string   `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	PayToken             string   `protobuf:"bytes,3,opt,name=pay_token,json=payToken,proto3" json:"pay_token,omitempty"`
	LastContact          string   `protobuf:"bytes,4,opt,name=last_contact,json=lastContact,proto3" json:"last_contact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserMetadata) Reset()         { *m = UserMetadata{} }
func (m *UserMetadata) String() string { return proto.CompactTextString(m) }
func (*UserMetadata) ProtoMessage()    {}
func (*UserMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
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

func (m *UserMetadata) GetPlan() string {
	if m != nil {
		return m.Plan
	}
	return ""
}

func (m *UserMetadata) GetPayToken() string {
	if m != nil {
		return m.PayToken
	}
	return ""
}

func (m *UserMetadata) GetLastContact() string {
	if m != nil {
		return m.LastContact
	}
	return ""
}

type Identity struct {
	Provider             string   `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Connection           string   `protobuf:"bytes,3,opt,name=connection,proto3" json:"connection,omitempty"`
	IsSocial             bool     `protobuf:"varint,4,opt,name=isSocial,proto3" json:"isSocial,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identity) Reset()         { *m = Identity{} }
func (m *Identity) String() string { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()    {}
func (*Identity) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *Identity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity.Unmarshal(m, b)
}
func (m *Identity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity.Marshal(b, m, deterministic)
}
func (m *Identity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity.Merge(m, src)
}
func (m *Identity) XXX_Size() int {
	return xxx_messageInfo_Identity.Size(m)
}
func (m *Identity) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity.DiscardUnknown(m)
}

var xxx_messageInfo_Identity proto.InternalMessageInfo

func (m *Identity) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Identity) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Identity) GetConnection() string {
	if m != nil {
		return m.Connection
	}
	return ""
}

func (m *Identity) GetIsSocial() bool {
	if m != nil {
		return m.IsSocial
	}
	return false
}

func init() {
	proto.RegisterType((*GetProfileByEmail)(nil), "api.GetProfileByEmail")
	proto.RegisterType((*Profile)(nil), "api.Profile")
	proto.RegisterType((*UserMetadata)(nil), "api.UserMetadata")
	proto.RegisterType((*Identity)(nil), "api.Identity")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x4f, 0x6f, 0x13, 0x3d,
	0x10, 0xc6, 0x95, 0xa6, 0x6d, 0x36, 0x93, 0xa4, 0x6f, 0xeb, 0xb7, 0x2a, 0x26, 0xfc, 0x69, 0x88,
	0x84, 0x14, 0x90, 0x68, 0xa5, 0x22, 0x71, 0xe0, 0x56, 0x2a, 0x84, 0x7a, 0x00, 0xa1, 0x2d, 0x70,
	0x5d, 0x4d, 0x77, 0xdd, 0x30, 0xea, 0xc6, 0xb6, 0x76, 0xbd, 0x11, 0x51, 0xc5, 0x85, 0xaf, 0xc0,
	0x97, 0xe2, 0xce, 0x57, 0xe0, 0x83, 0x20, 0x8f, 0x9d, 0x36, 0x55, 0xb9, 0xcd, 0xf3, 0x7b, 0x66,
	0x65, 0xef, 0xf8, 0x19, 0xe8, 0xa2, 0xa5, 0x03, 0x5b, 0x19, 0x67, 0x44, 0x1b, 0x2d, 0x0d, 0x1f,
	0x4e, 0x8d, 0x99, 0x96, 0xea, 0x10, 0x2d, 0x1d, 0xa2, 0xd6, 0xc6, 0xa1, 0x23, 0xa3, 0xeb, 0xd0,
	0x32, 0x7e, 0x06, 0x3b, 0xef, 0x94, 0xfb, 0x58, 0x99, 0x0b, 0x2a, 0xd5, 0x9b, 0xc5, 0xdb, 0x19,
	0x52, 0x29, 0x76, 0x61, 0x43, 0xf9, 0x42, 0xb6, 0x46, 0xad, 0x49, 0x37, 0x0d, 0x62, 0xfc, 0x6b,
	0x1d, 0x3a, 0xb1, 0xf1, 0xdf, 0x1d, 0xe2, 0x29, 0x6c, 0x71, 0x91, 0xcd, 0x55, 0x45, 0x17, 0xa4,
	0x0a, 0xb9, 0x36, 0x6a, 0x4d, 0x92, 0x74, 0xc0, 0xf4, 0x4b, 0x84, 0x42, 0xc0, 0xba, 0xc6, 0x99,
	0x92, 0x6d, 0xfe, 0x96, 0x6b, 0xf1, 0x08, 0x60, 0x4a, 0x73, 0xa5, 0x33, 0x76, 0xd6, 0xd9, 0xe9,
	0x32, 0xf9, 0xe0, 0xed, 0x7d, 0xe8, 0x5d, 0xe0, 0x8c, 0xca, 0x45, 0xf0, 0x37, 0xd8, 0x87, 0x80,
	0xb8, 0x41, 0x42, 0xc7, 0x52, 0xee, 0x9a, 0x4a, 0xc9, 0x4d, 0x36, 0x97, 0x52, 0xec, 0xc1, 0x66,
	0x69, 0x72, 0x2c, 0x95, 0xec, 0xb0, 0x11, 0x95, 0xb8, 0x07, 0x9d, 0xa6, 0x56, 0x55, 0x46, 0x85,
	0x4c, 0x82, 0xe1, 0xe5, 0x69, 0x21, 0x86, 0x90, 0x68, 0xca, 0x2f, 0xf9, 0xa0, 0x2e, 0x3b, 0xd7,
	0x5a, 0x3c, 0x06, 0xc8, 0x8d, 0xd6, 0x2a, 0xf7, 0x33, 0x94, 0x10, 0xae, 0x71, 0x43, 0xc4, 0x0b,
	0x00, 0x2a, 0x94, 0x76, 0xe4, 0x48, 0xd5, 0xb2, 0x37, 0x6a, 0x4f, 0x7a, 0x47, 0x83, 0x03, 0xff,
	0x22, 0xa7, 0x01, 0x2f, 0xd2, 0x95, 0x06, 0x7f, 0x87, 0x12, 0x6b, 0x97, 0x91, 0x95, 0xfd, 0x78,
	0x39, 0xac, 0xdd, 0xa9, 0xf5, 0xff, 0x5b, 0x9a, 0x29, 0xe9, 0x2c, 0x37, 0x8d, 0x76, 0x72, 0x30,
	0x6a, 0x4d, 0xda, 0x29, 0x30, 0x3a, 0xf1, 0xc4, 0xcf, 0xab, 0xb1, 0x05, 0x3a, 0x55, 0x64, 0xe8,
	0xe4, 0x56, 0x98, 0x57, 0x24, 0xc7, 0x6c, 0xe7, 0x95, 0x5a, 0xda, 0xff, 0x05, 0x3b, 0x92, 0x63,
	0x27, 0xb6, 0xa1, 0x5d, 0x37, 0xe7, 0x72, 0x9b, 0xb9, 0x2f, 0x3d, 0xa1, 0xba, 0x96, 0x3b, 0x81,
	0x50, 0x5d, 0x7b, 0x82, 0x4d, 0x21, 0x45, 0x20, 0xd8, 0x14, 0xdc, 0x83, 0x4e, 0xfe, 0x1f, 0x7b,
	0xd0, 0x89, 0x57, 0x30, 0xe0, 0x19, 0xce, 0x94, 0xc3, 0x02, 0x1d, 0xca, 0xdd, 0x51, 0x6b, 0xd2,
	0x3b, 0xda, 0xe1, 0x3f, 0xfe, 0x5c, 0xab, 0xea, 0x7d, 0x34, 0xd2, 0x7e, 0xb3, 0xa2, 0xc6, 0xdf,
	0xa0, 0xbf, 0xea, 0xfa, 0x38, 0xd9, 0xaf, 0x46, 0xab, 0x65, 0x9c, 0x58, 0xf8, 0x9c, 0xd8, 0x12,
	0x35, 0x87, 0xa8, 0x9b, 0x72, 0x2d, 0x1e, 0x40, 0xd7, 0xe2, 0x22, 0x73, 0xe6, 0x52, 0xe9, 0x18,
	0xa0, 0xc4, 0xe2, 0xe2, 0x93, 0xd7, 0xe2, 0x09, 0xf4, 0x79, 0x9c, 0xb9, 0xd1, 0x0e, 0x73, 0x17,
	0x63, 0xd4, 0xf3, 0xec, 0x24, 0xa0, 0xf1, 0x15, 0x24, 0xcb, 0x97, 0xf0, 0x0f, 0x6d, 0x2b, 0x33,
	0xa7, 0x42, 0x55, 0xf1, 0xe0, 0x6b, 0xbd, 0x9a, 0x8e, 0xb5, 0x5b, 0xe9, 0xb8, 0x9d, 0x80, 0xf6,
	0x9d, 0x04, 0x0c, 0x21, 0xa1, 0xfa, 0xcc, 0xe4, 0x84, 0x25, 0x9f, 0x9f, 0xa4, 0xd7, 0xfa, 0x48,
	0xc1, 0x56, 0x5c, 0xa0, 0x33, 0x55, 0xcd, 0x29, 0x57, 0xe2, 0x0c, 0xe0, 0x66, 0xfd, 0xc4, 0x1e,
	0xcf, 0xed, 0xce, 0x3e, 0x0e, 0xfb, 0xcc, 0x23, 0x1c, 0xef, 0xff, 0xf8, 0xfd, 0xe7, 0xe7, 0xda,
	0xfd, 0xf1, 0x2e, 0xaf, 0xb4, 0x0d, 0xf4, 0xf0, 0x8a, 0x77, 0xec, 0xfb, 0xeb, 0xd6, 0xf3, 0xf3,
	0x4d, 0x5e, 0xed, 0x97, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xeb, 0x63, 0xe7, 0x0a, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileServiceClient interface {
	GetProfile(ctx context.Context, in *GetProfileByEmail, opts ...grpc.CallOption) (*Profile, error)
}

type profileServiceClient struct {
	cc *grpc.ClientConn
}

func NewProfileServiceClient(cc *grpc.ClientConn) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) GetProfile(ctx context.Context, in *GetProfileByEmail, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/api.ProfileService/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
type ProfileServiceServer interface {
	GetProfile(context.Context, *GetProfileByEmail) (*Profile, error)
}

func RegisterProfileServiceServer(s *grpc.Server, srv ProfileServiceServer) {
	s.RegisterService(&_ProfileService_serviceDesc, srv)
}

func _ProfileService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileByEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ProfileService/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfile(ctx, req.(*GetProfileByEmail))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfile",
			Handler:    _ProfileService_GetProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
