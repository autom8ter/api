// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sms/sms.proto

package sms

import (
	fmt "fmt"
	errors "github.com/autom8ter/api/sdk/go/errors"
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

type SendSMSRequest struct {
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	From                 string   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	StatusCallback       string   `protobuf:"bytes,5,opt,name=status_callback,json=statusCallback,proto3" json:"status_callback,omitempty"`
	ApplicationSid       string   `protobuf:"bytes,6,opt,name=application_sid,json=applicationSid,proto3" json:"application_sid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendSMSRequest) Reset()         { *m = SendSMSRequest{} }
func (m *SendSMSRequest) String() string { return proto.CompactTextString(m) }
func (*SendSMSRequest) ProtoMessage()    {}
func (*SendSMSRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5900ef846da80354, []int{0}
}

func (m *SendSMSRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendSMSRequest.Unmarshal(m, b)
}
func (m *SendSMSRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendSMSRequest.Marshal(b, m, deterministic)
}
func (m *SendSMSRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendSMSRequest.Merge(m, src)
}
func (m *SendSMSRequest) XXX_Size() int {
	return xxx_messageInfo_SendSMSRequest.Size(m)
}
func (m *SendSMSRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendSMSRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendSMSRequest proto.InternalMessageInfo

func (m *SendSMSRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SendSMSRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *SendSMSRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SendSMSRequest) GetStatusCallback() string {
	if m != nil {
		return m.StatusCallback
	}
	return ""
}

func (m *SendSMSRequest) GetApplicationSid() string {
	if m != nil {
		return m.ApplicationSid
	}
	return ""
}

type SendMMSRequest struct {
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	From                 string   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	MediaUrl             string   `protobuf:"bytes,5,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	StatusCallback       string   `protobuf:"bytes,6,opt,name=status_callback,json=statusCallback,proto3" json:"status_callback,omitempty"`
	ApplicationSid       string   `protobuf:"bytes,7,opt,name=application_sid,json=applicationSid,proto3" json:"application_sid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMMSRequest) Reset()         { *m = SendMMSRequest{} }
func (m *SendMMSRequest) String() string { return proto.CompactTextString(m) }
func (*SendMMSRequest) ProtoMessage()    {}
func (*SendMMSRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5900ef846da80354, []int{1}
}

func (m *SendMMSRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMMSRequest.Unmarshal(m, b)
}
func (m *SendMMSRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMMSRequest.Marshal(b, m, deterministic)
}
func (m *SendMMSRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMMSRequest.Merge(m, src)
}
func (m *SendMMSRequest) XXX_Size() int {
	return xxx_messageInfo_SendMMSRequest.Size(m)
}
func (m *SendMMSRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMMSRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMMSRequest proto.InternalMessageInfo

func (m *SendMMSRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SendMMSRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *SendMMSRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SendMMSRequest) GetMediaUrl() string {
	if m != nil {
		return m.MediaUrl
	}
	return ""
}

func (m *SendMMSRequest) GetStatusCallback() string {
	if m != nil {
		return m.StatusCallback
	}
	return ""
}

func (m *SendMMSRequest) GetApplicationSid() string {
	if m != nil {
		return m.ApplicationSid
	}
	return ""
}

type SendSMSWithCopilotRequest struct {
	MessagingServiceSid  string   `protobuf:"bytes,1,opt,name=messaging_service_sid,json=messagingServiceSid,proto3" json:"messaging_service_sid,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	StatusCallback       string   `protobuf:"bytes,4,opt,name=status_callback,json=statusCallback,proto3" json:"status_callback,omitempty"`
	ApplicationSid       string   `protobuf:"bytes,5,opt,name=application_sid,json=applicationSid,proto3" json:"application_sid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendSMSWithCopilotRequest) Reset()         { *m = SendSMSWithCopilotRequest{} }
func (m *SendSMSWithCopilotRequest) String() string { return proto.CompactTextString(m) }
func (*SendSMSWithCopilotRequest) ProtoMessage()    {}
func (*SendSMSWithCopilotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5900ef846da80354, []int{2}
}

func (m *SendSMSWithCopilotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendSMSWithCopilotRequest.Unmarshal(m, b)
}
func (m *SendSMSWithCopilotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendSMSWithCopilotRequest.Marshal(b, m, deterministic)
}
func (m *SendSMSWithCopilotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendSMSWithCopilotRequest.Merge(m, src)
}
func (m *SendSMSWithCopilotRequest) XXX_Size() int {
	return xxx_messageInfo_SendSMSWithCopilotRequest.Size(m)
}
func (m *SendSMSWithCopilotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendSMSWithCopilotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendSMSWithCopilotRequest proto.InternalMessageInfo

func (m *SendSMSWithCopilotRequest) GetMessagingServiceSid() string {
	if m != nil {
		return m.MessagingServiceSid
	}
	return ""
}

func (m *SendSMSWithCopilotRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SendSMSWithCopilotRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SendSMSWithCopilotRequest) GetStatusCallback() string {
	if m != nil {
		return m.StatusCallback
	}
	return ""
}

func (m *SendSMSWithCopilotRequest) GetApplicationSid() string {
	if m != nil {
		return m.ApplicationSid
	}
	return ""
}

type GetSMSRequest struct {
	Sid                  string   `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSMSRequest) Reset()         { *m = GetSMSRequest{} }
func (m *GetSMSRequest) String() string { return proto.CompactTextString(m) }
func (*GetSMSRequest) ProtoMessage()    {}
func (*GetSMSRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5900ef846da80354, []int{3}
}

func (m *GetSMSRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSMSRequest.Unmarshal(m, b)
}
func (m *GetSMSRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSMSRequest.Marshal(b, m, deterministic)
}
func (m *GetSMSRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSMSRequest.Merge(m, src)
}
func (m *GetSMSRequest) XXX_Size() int {
	return xxx_messageInfo_GetSMSRequest.Size(m)
}
func (m *GetSMSRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSMSRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSMSRequest proto.InternalMessageInfo

func (m *GetSMSRequest) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

type SMSResponse struct {
	Sid                  string        `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	DateCreated          string        `protobuf:"bytes,2,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	DateUpdated          string        `protobuf:"bytes,3,opt,name=date_updated,json=dateUpdated,proto3" json:"date_updated,omitempty"`
	DateSent             string        `protobuf:"bytes,4,opt,name=date_sent,json=dateSent,proto3" json:"date_sent,omitempty"`
	AccoundSid           string        `protobuf:"bytes,5,opt,name=accound_sid,json=accoundSid,proto3" json:"accound_sid,omitempty"`
	To                   string        `protobuf:"bytes,6,opt,name=to,proto3" json:"to,omitempty"`
	From                 string        `protobuf:"bytes,7,opt,name=from,proto3" json:"from,omitempty"`
	MediaUrl             string        `protobuf:"bytes,8,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	Body                 string        `protobuf:"bytes,10,opt,name=body,proto3" json:"body,omitempty"`
	Status               string        `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	Direction            string        `protobuf:"bytes,12,opt,name=direction,proto3" json:"direction,omitempty"`
	ApiVersion           string        `protobuf:"bytes,13,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	Price                string        `protobuf:"bytes,17,opt,name=price,proto3" json:"price,omitempty"`
	Uri                  string        `protobuf:"bytes,20,opt,name=uri,proto3" json:"uri,omitempty"`
	Exception            *errors.Error `protobuf:"bytes,21,opt,name=exception,proto3" json:"exception,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SMSResponse) Reset()         { *m = SMSResponse{} }
func (m *SMSResponse) String() string { return proto.CompactTextString(m) }
func (*SMSResponse) ProtoMessage()    {}
func (*SMSResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5900ef846da80354, []int{4}
}

func (m *SMSResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SMSResponse.Unmarshal(m, b)
}
func (m *SMSResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SMSResponse.Marshal(b, m, deterministic)
}
func (m *SMSResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SMSResponse.Merge(m, src)
}
func (m *SMSResponse) XXX_Size() int {
	return xxx_messageInfo_SMSResponse.Size(m)
}
func (m *SMSResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SMSResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SMSResponse proto.InternalMessageInfo

func (m *SMSResponse) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *SMSResponse) GetDateCreated() string {
	if m != nil {
		return m.DateCreated
	}
	return ""
}

func (m *SMSResponse) GetDateUpdated() string {
	if m != nil {
		return m.DateUpdated
	}
	return ""
}

func (m *SMSResponse) GetDateSent() string {
	if m != nil {
		return m.DateSent
	}
	return ""
}

func (m *SMSResponse) GetAccoundSid() string {
	if m != nil {
		return m.AccoundSid
	}
	return ""
}

func (m *SMSResponse) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SMSResponse) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *SMSResponse) GetMediaUrl() string {
	if m != nil {
		return m.MediaUrl
	}
	return ""
}

func (m *SMSResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SMSResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *SMSResponse) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *SMSResponse) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *SMSResponse) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *SMSResponse) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *SMSResponse) GetException() *errors.Error {
	if m != nil {
		return m.Exception
	}
	return nil
}

func init() {
	proto.RegisterType((*SendSMSRequest)(nil), "sms.SendSMSRequest")
	proto.RegisterType((*SendMMSRequest)(nil), "sms.SendMMSRequest")
	proto.RegisterType((*SendSMSWithCopilotRequest)(nil), "sms.SendSMSWithCopilotRequest")
	proto.RegisterType((*GetSMSRequest)(nil), "sms.GetSMSRequest")
	proto.RegisterType((*SMSResponse)(nil), "sms.SMSResponse")
}

func init() { proto.RegisterFile("sms/sms.proto", fileDescriptor_5900ef846da80354) }

var fileDescriptor_5900ef846da80354 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x55, 0x9a, 0x36, 0x6b, 0xbf, 0xac, 0x05, 0xbc, 0x0d, 0x85, 0x1f, 0x89, 0xad, 0x37, 0x4c,
	0x42, 0xda, 0xa4, 0xf1, 0x08, 0x15, 0xe2, 0x6a, 0x37, 0x8d, 0x06, 0x97, 0x91, 0x6b, 0x7f, 0x0c,
	0x8b, 0x24, 0x36, 0xb6, 0x33, 0xc1, 0xab, 0x70, 0xcd, 0x3b, 0xf0, 0x04, 0xbc, 0x17, 0xf2, 0xcf,
	0x96, 0xc0, 0x2a, 0x54, 0x69, 0x57, 0xb1, 0xcf, 0x39, 0x72, 0xce, 0x39, 0xfe, 0x0c, 0x73, 0xd3,
	0x98, 0x73, 0xd3, 0x98, 0x33, 0xa5, 0xa5, 0x95, 0x24, 0x35, 0x8d, 0x79, 0x7e, 0x80, 0x5a, 0x4b,
	0x6d, 0xce, 0xc3, 0x27, 0x30, 0xcb, 0x1f, 0x09, 0x2c, 0x4a, 0x6c, 0x79, 0x79, 0x59, 0xae, 0xf1,
	0x6b, 0x87, 0xc6, 0x92, 0x05, 0x8c, 0xac, 0x2c, 0x46, 0xc7, 0xc9, 0xe9, 0x6c, 0x3d, 0xb2, 0x92,
	0x10, 0x18, 0x7f, 0xd2, 0xb2, 0x29, 0x52, 0x8f, 0xf8, 0xb5, 0xc3, 0x36, 0x92, 0x7f, 0x2f, 0xc6,
	0x01, 0x73, 0x6b, 0xf2, 0x1a, 0x1e, 0x19, 0x4b, 0x6d, 0x67, 0x2a, 0x46, 0xeb, 0x7a, 0x43, 0xd9,
	0x97, 0x62, 0xe2, 0xe9, 0x45, 0x80, 0x57, 0x11, 0x75, 0x42, 0xaa, 0x54, 0x2d, 0x18, 0xb5, 0x42,
	0xb6, 0x95, 0x11, 0xbc, 0xc8, 0x82, 0x70, 0x00, 0x97, 0x82, 0x2f, 0x7f, 0x45, 0x73, 0x97, 0x0f,
	0x37, 0xf7, 0x02, 0x66, 0x0d, 0x72, 0x41, 0xab, 0x4e, 0xd7, 0xd1, 0xd6, 0xd4, 0x03, 0x57, 0xba,
	0xde, 0xe6, 0x3c, 0xdb, 0xd5, 0xf9, 0xde, 0x56, 0xe7, 0xbf, 0x13, 0x78, 0x16, 0x6b, 0xfd, 0x28,
	0xec, 0xe7, 0x95, 0x54, 0xa2, 0x96, 0xf6, 0x36, 0xc4, 0x05, 0x1c, 0x35, 0x68, 0x0c, 0xbd, 0x16,
	0xed, 0x75, 0x65, 0x50, 0xdf, 0x08, 0x86, 0xfe, 0xb0, 0xc4, 0x1f, 0x76, 0x70, 0x47, 0x96, 0x81,
	0x2b, 0x05, 0xdf, 0x16, 0xdc, 0x87, 0x4c, 0xff, 0x7f, 0x03, 0xe3, 0x5d, 0x73, 0x4c, 0xb6, 0xe6,
	0x38, 0x81, 0xf9, 0x7b, 0xb4, 0x83, 0xe1, 0x78, 0x0c, 0x69, 0x6f, 0xd4, 0x2d, 0x97, 0x3f, 0x53,
	0xc8, 0xbd, 0xc0, 0x28, 0xd9, 0x1a, 0xbc, 0xaf, 0x20, 0x27, 0xb0, 0xcf, 0xa9, 0xc5, 0x8a, 0x69,
	0xa4, 0x16, 0x79, 0x0c, 0x91, 0x3b, 0x6c, 0x15, 0xa0, 0x3b, 0x49, 0xa7, 0xb8, 0x97, 0xa4, 0xbd,
	0xe4, 0x2a, 0x40, 0xee, 0x06, 0xbd, 0xc4, 0x60, 0x6b, 0x63, 0xac, 0xa9, 0x03, 0x4a, 0x6c, 0x2d,
	0x79, 0x05, 0x39, 0x65, 0x4c, 0x76, 0x2d, 0x1f, 0x84, 0x81, 0x08, 0xf5, 0xf5, 0x65, 0xf7, 0xe6,
	0x66, 0x6f, 0x30, 0x37, 0x7f, 0xcd, 0xc8, 0xf4, 0x9f, 0x19, 0xb9, 0xed, 0x1b, 0x06, 0x7d, 0x3f,
	0x85, 0x2c, 0x14, 0x5b, 0xe4, 0x1e, 0x8d, 0x3b, 0xf2, 0x12, 0x66, 0x5c, 0x68, 0x64, 0xae, 0xc5,
	0x62, 0xdf, 0x53, 0x3d, 0xe0, 0xbd, 0x2a, 0x51, 0xdd, 0xa0, 0x36, 0x8e, 0x9f, 0x47, 0xaf, 0x4a,
	0x7c, 0x08, 0x08, 0x39, 0x84, 0x89, 0xd2, 0x82, 0x61, 0xf1, 0xc4, 0x53, 0x61, 0xe3, 0x7a, 0xed,
	0xb4, 0x28, 0x0e, 0x43, 0xaf, 0x9d, 0x16, 0xe4, 0x0d, 0xcc, 0xf0, 0x1b, 0x43, 0xe5, 0x7f, 0x73,
	0x74, 0x9c, 0x9c, 0xe6, 0x17, 0xf3, 0xb3, 0xf8, 0xba, 0xdf, 0xb9, 0xcf, 0xba, 0xe7, 0x37, 0x99,
	0x7f, 0xef, 0x6f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xf4, 0x84, 0xf1, 0x1a, 0x04, 0x00,
	0x00,
}
