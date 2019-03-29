// Code generated by protoc-gen-go. DO NOT EDIT.
// source: documents/documents.proto

package documents

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type CreateDocRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDocRequest) Reset()         { *m = CreateDocRequest{} }
func (m *CreateDocRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDocRequest) ProtoMessage()    {}
func (*CreateDocRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{0}
}

func (m *CreateDocRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDocRequest.Unmarshal(m, b)
}
func (m *CreateDocRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDocRequest.Marshal(b, m, deterministic)
}
func (m *CreateDocRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDocRequest.Merge(m, src)
}
func (m *CreateDocRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDocRequest.Size(m)
}
func (m *CreateDocRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDocRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDocRequest proto.InternalMessageInfo

func (m *CreateDocRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type Document struct {
	Path                 string      `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Id                   string      `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Parent               *Collection `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{1}
}

func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Document) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Document) GetParent() *Collection {
	if m != nil {
		return m.Parent
	}
	return nil
}

type Collection struct {
	Path                 string    `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Id                   string    `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Parent               *Document `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Collection) Reset()         { *m = Collection{} }
func (m *Collection) String() string { return proto.CompactTextString(m) }
func (*Collection) ProtoMessage()    {}
func (*Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{2}
}

func (m *Collection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collection.Unmarshal(m, b)
}
func (m *Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collection.Marshal(b, m, deterministic)
}
func (m *Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collection.Merge(m, src)
}
func (m *Collection) XXX_Size() int {
	return xxx_messageInfo_Collection.Size(m)
}
func (m *Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_Collection proto.InternalMessageInfo

func (m *Collection) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Collection) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Collection) GetParent() *Document {
	if m != nil {
		return m.Parent
	}
	return nil
}

type GetDocRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDocRequest) Reset()         { *m = GetDocRequest{} }
func (m *GetDocRequest) String() string { return proto.CompactTextString(m) }
func (*GetDocRequest) ProtoMessage()    {}
func (*GetDocRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{3}
}

func (m *GetDocRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocRequest.Unmarshal(m, b)
}
func (m *GetDocRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocRequest.Marshal(b, m, deterministic)
}
func (m *GetDocRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocRequest.Merge(m, src)
}
func (m *GetDocRequest) XXX_Size() int {
	return xxx_messageInfo_GetDocRequest.Size(m)
}
func (m *GetDocRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocRequest proto.InternalMessageInfo

type GetDocResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDocResponse) Reset()         { *m = GetDocResponse{} }
func (m *GetDocResponse) String() string { return proto.CompactTextString(m) }
func (*GetDocResponse) ProtoMessage()    {}
func (*GetDocResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{4}
}

func (m *GetDocResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocResponse.Unmarshal(m, b)
}
func (m *GetDocResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocResponse.Marshal(b, m, deterministic)
}
func (m *GetDocResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocResponse.Merge(m, src)
}
func (m *GetDocResponse) XXX_Size() int {
	return xxx_messageInfo_GetDocResponse.Size(m)
}
func (m *GetDocResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocResponse proto.InternalMessageInfo

type GetAllDocRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllDocRequest) Reset()         { *m = GetAllDocRequest{} }
func (m *GetAllDocRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllDocRequest) ProtoMessage()    {}
func (*GetAllDocRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{5}
}

func (m *GetAllDocRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllDocRequest.Unmarshal(m, b)
}
func (m *GetAllDocRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllDocRequest.Marshal(b, m, deterministic)
}
func (m *GetAllDocRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllDocRequest.Merge(m, src)
}
func (m *GetAllDocRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllDocRequest.Size(m)
}
func (m *GetAllDocRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllDocRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllDocRequest proto.InternalMessageInfo

type GetAllDocResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllDocResponse) Reset()         { *m = GetAllDocResponse{} }
func (m *GetAllDocResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllDocResponse) ProtoMessage()    {}
func (*GetAllDocResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{6}
}

func (m *GetAllDocResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllDocResponse.Unmarshal(m, b)
}
func (m *GetAllDocResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllDocResponse.Marshal(b, m, deterministic)
}
func (m *GetAllDocResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllDocResponse.Merge(m, src)
}
func (m *GetAllDocResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllDocResponse.Size(m)
}
func (m *GetAllDocResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllDocResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllDocResponse proto.InternalMessageInfo

type UpdateDocRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateDocRequest) Reset()         { *m = UpdateDocRequest{} }
func (m *UpdateDocRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDocRequest) ProtoMessage()    {}
func (*UpdateDocRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{7}
}

func (m *UpdateDocRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDocRequest.Unmarshal(m, b)
}
func (m *UpdateDocRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDocRequest.Marshal(b, m, deterministic)
}
func (m *UpdateDocRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDocRequest.Merge(m, src)
}
func (m *UpdateDocRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDocRequest.Size(m)
}
func (m *UpdateDocRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDocRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDocRequest proto.InternalMessageInfo

type UpdateDocResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateDocResponse) Reset()         { *m = UpdateDocResponse{} }
func (m *UpdateDocResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateDocResponse) ProtoMessage()    {}
func (*UpdateDocResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{8}
}

func (m *UpdateDocResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDocResponse.Unmarshal(m, b)
}
func (m *UpdateDocResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDocResponse.Marshal(b, m, deterministic)
}
func (m *UpdateDocResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDocResponse.Merge(m, src)
}
func (m *UpdateDocResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateDocResponse.Size(m)
}
func (m *UpdateDocResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDocResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDocResponse proto.InternalMessageInfo

type DeleteDocRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDocRequest) Reset()         { *m = DeleteDocRequest{} }
func (m *DeleteDocRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDocRequest) ProtoMessage()    {}
func (*DeleteDocRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{9}
}

func (m *DeleteDocRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDocRequest.Unmarshal(m, b)
}
func (m *DeleteDocRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDocRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDocRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDocRequest.Merge(m, src)
}
func (m *DeleteDocRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDocRequest.Size(m)
}
func (m *DeleteDocRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDocRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDocRequest proto.InternalMessageInfo

type DeleteDocResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDocResponse) Reset()         { *m = DeleteDocResponse{} }
func (m *DeleteDocResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteDocResponse) ProtoMessage()    {}
func (*DeleteDocResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{10}
}

func (m *DeleteDocResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDocResponse.Unmarshal(m, b)
}
func (m *DeleteDocResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDocResponse.Marshal(b, m, deterministic)
}
func (m *DeleteDocResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDocResponse.Merge(m, src)
}
func (m *DeleteDocResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteDocResponse.Size(m)
}
func (m *DeleteDocResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDocResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDocResponse proto.InternalMessageInfo

type CreateCollectionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCollectionRequest) Reset()         { *m = CreateCollectionRequest{} }
func (m *CreateCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCollectionRequest) ProtoMessage()    {}
func (*CreateCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{11}
}

func (m *CreateCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCollectionRequest.Unmarshal(m, b)
}
func (m *CreateCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCollectionRequest.Marshal(b, m, deterministic)
}
func (m *CreateCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCollectionRequest.Merge(m, src)
}
func (m *CreateCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCollectionRequest.Size(m)
}
func (m *CreateCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCollectionRequest proto.InternalMessageInfo

type CreateCollectionResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCollectionResponse) Reset()         { *m = CreateCollectionResponse{} }
func (m *CreateCollectionResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCollectionResponse) ProtoMessage()    {}
func (*CreateCollectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{12}
}

func (m *CreateCollectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCollectionResponse.Unmarshal(m, b)
}
func (m *CreateCollectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCollectionResponse.Marshal(b, m, deterministic)
}
func (m *CreateCollectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCollectionResponse.Merge(m, src)
}
func (m *CreateCollectionResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCollectionResponse.Size(m)
}
func (m *CreateCollectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCollectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCollectionResponse proto.InternalMessageInfo

type GetCollectionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCollectionRequest) Reset()         { *m = GetCollectionRequest{} }
func (m *GetCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*GetCollectionRequest) ProtoMessage()    {}
func (*GetCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{13}
}

func (m *GetCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCollectionRequest.Unmarshal(m, b)
}
func (m *GetCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCollectionRequest.Marshal(b, m, deterministic)
}
func (m *GetCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCollectionRequest.Merge(m, src)
}
func (m *GetCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_GetCollectionRequest.Size(m)
}
func (m *GetCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCollectionRequest proto.InternalMessageInfo

type GetCollectionResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCollectionResponse) Reset()         { *m = GetCollectionResponse{} }
func (m *GetCollectionResponse) String() string { return proto.CompactTextString(m) }
func (*GetCollectionResponse) ProtoMessage()    {}
func (*GetCollectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{14}
}

func (m *GetCollectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCollectionResponse.Unmarshal(m, b)
}
func (m *GetCollectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCollectionResponse.Marshal(b, m, deterministic)
}
func (m *GetCollectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCollectionResponse.Merge(m, src)
}
func (m *GetCollectionResponse) XXX_Size() int {
	return xxx_messageInfo_GetCollectionResponse.Size(m)
}
func (m *GetCollectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCollectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCollectionResponse proto.InternalMessageInfo

type GetAllCollectionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllCollectionRequest) Reset()         { *m = GetAllCollectionRequest{} }
func (m *GetAllCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllCollectionRequest) ProtoMessage()    {}
func (*GetAllCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{15}
}

func (m *GetAllCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllCollectionRequest.Unmarshal(m, b)
}
func (m *GetAllCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllCollectionRequest.Marshal(b, m, deterministic)
}
func (m *GetAllCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllCollectionRequest.Merge(m, src)
}
func (m *GetAllCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllCollectionRequest.Size(m)
}
func (m *GetAllCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllCollectionRequest proto.InternalMessageInfo

type GetAllCollectionResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllCollectionResponse) Reset()         { *m = GetAllCollectionResponse{} }
func (m *GetAllCollectionResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllCollectionResponse) ProtoMessage()    {}
func (*GetAllCollectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{16}
}

func (m *GetAllCollectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllCollectionResponse.Unmarshal(m, b)
}
func (m *GetAllCollectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllCollectionResponse.Marshal(b, m, deterministic)
}
func (m *GetAllCollectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllCollectionResponse.Merge(m, src)
}
func (m *GetAllCollectionResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllCollectionResponse.Size(m)
}
func (m *GetAllCollectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllCollectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllCollectionResponse proto.InternalMessageInfo

type UpdateCollectionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCollectionRequest) Reset()         { *m = UpdateCollectionRequest{} }
func (m *UpdateCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCollectionRequest) ProtoMessage()    {}
func (*UpdateCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{17}
}

func (m *UpdateCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCollectionRequest.Unmarshal(m, b)
}
func (m *UpdateCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCollectionRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCollectionRequest.Merge(m, src)
}
func (m *UpdateCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCollectionRequest.Size(m)
}
func (m *UpdateCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCollectionRequest proto.InternalMessageInfo

type UpdateCollectionResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCollectionResponse) Reset()         { *m = UpdateCollectionResponse{} }
func (m *UpdateCollectionResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateCollectionResponse) ProtoMessage()    {}
func (*UpdateCollectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{18}
}

func (m *UpdateCollectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCollectionResponse.Unmarshal(m, b)
}
func (m *UpdateCollectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCollectionResponse.Marshal(b, m, deterministic)
}
func (m *UpdateCollectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCollectionResponse.Merge(m, src)
}
func (m *UpdateCollectionResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateCollectionResponse.Size(m)
}
func (m *UpdateCollectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCollectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCollectionResponse proto.InternalMessageInfo

type DeleteCollectionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCollectionRequest) Reset()         { *m = DeleteCollectionRequest{} }
func (m *DeleteCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCollectionRequest) ProtoMessage()    {}
func (*DeleteCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{19}
}

func (m *DeleteCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCollectionRequest.Unmarshal(m, b)
}
func (m *DeleteCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCollectionRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCollectionRequest.Merge(m, src)
}
func (m *DeleteCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCollectionRequest.Size(m)
}
func (m *DeleteCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCollectionRequest proto.InternalMessageInfo

type DeleteCollectionResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCollectionResponse) Reset()         { *m = DeleteCollectionResponse{} }
func (m *DeleteCollectionResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCollectionResponse) ProtoMessage()    {}
func (*DeleteCollectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d307ff86905cec7, []int{20}
}

func (m *DeleteCollectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCollectionResponse.Unmarshal(m, b)
}
func (m *DeleteCollectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCollectionResponse.Marshal(b, m, deterministic)
}
func (m *DeleteCollectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCollectionResponse.Merge(m, src)
}
func (m *DeleteCollectionResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCollectionResponse.Size(m)
}
func (m *DeleteCollectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCollectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCollectionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateDocRequest)(nil), "documents.CreateDocRequest")
	proto.RegisterType((*Document)(nil), "documents.Document")
	proto.RegisterType((*Collection)(nil), "documents.Collection")
	proto.RegisterType((*GetDocRequest)(nil), "documents.GetDocRequest")
	proto.RegisterType((*GetDocResponse)(nil), "documents.GetDocResponse")
	proto.RegisterType((*GetAllDocRequest)(nil), "documents.GetAllDocRequest")
	proto.RegisterType((*GetAllDocResponse)(nil), "documents.GetAllDocResponse")
	proto.RegisterType((*UpdateDocRequest)(nil), "documents.UpdateDocRequest")
	proto.RegisterType((*UpdateDocResponse)(nil), "documents.UpdateDocResponse")
	proto.RegisterType((*DeleteDocRequest)(nil), "documents.DeleteDocRequest")
	proto.RegisterType((*DeleteDocResponse)(nil), "documents.DeleteDocResponse")
	proto.RegisterType((*CreateCollectionRequest)(nil), "documents.CreateCollectionRequest")
	proto.RegisterType((*CreateCollectionResponse)(nil), "documents.CreateCollectionResponse")
	proto.RegisterType((*GetCollectionRequest)(nil), "documents.GetCollectionRequest")
	proto.RegisterType((*GetCollectionResponse)(nil), "documents.GetCollectionResponse")
	proto.RegisterType((*GetAllCollectionRequest)(nil), "documents.GetAllCollectionRequest")
	proto.RegisterType((*GetAllCollectionResponse)(nil), "documents.GetAllCollectionResponse")
	proto.RegisterType((*UpdateCollectionRequest)(nil), "documents.UpdateCollectionRequest")
	proto.RegisterType((*UpdateCollectionResponse)(nil), "documents.UpdateCollectionResponse")
	proto.RegisterType((*DeleteCollectionRequest)(nil), "documents.DeleteCollectionRequest")
	proto.RegisterType((*DeleteCollectionResponse)(nil), "documents.DeleteCollectionResponse")
}

func init() { proto.RegisterFile("documents/documents.proto", fileDescriptor_5d307ff86905cec7) }

var fileDescriptor_5d307ff86905cec7 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4f, 0x83, 0x40,
	0x10, 0x85, 0x03, 0x9a, 0xc6, 0x8e, 0xb1, 0xe2, 0xd6, 0x5a, 0xda, 0x8b, 0x0d, 0x07, 0xd3, 0xc4,
	0x58, 0x12, 0xfd, 0x05, 0xa6, 0x24, 0xdc, 0x49, 0x3c, 0xf6, 0x40, 0x61, 0xac, 0x24, 0x0b, 0xbb,
	0xc2, 0xf0, 0xff, 0x8d, 0x2c, 0x65, 0x57, 0xe0, 0xe0, 0x6d, 0xf3, 0xde, 0xf7, 0xf6, 0xed, 0x0c,
	0xc0, 0x2a, 0x15, 0x49, 0x9d, 0x63, 0x41, 0x95, 0xdf, 0x9d, 0x76, 0xb2, 0x14, 0x24, 0xd8, 0xb4,
	0x13, 0xd6, 0x8f, 0x27, 0x21, 0x4e, 0x1c, 0xfd, 0xc6, 0x38, 0xd6, 0x9f, 0x3e, 0x65, 0x39, 0x56,
	0x14, 0xe7, 0x52, 0xb1, 0xde, 0x13, 0x38, 0xfb, 0x12, 0x63, 0xc2, 0x40, 0x24, 0x11, 0x7e, 0xd7,
	0x58, 0x11, 0x63, 0x70, 0x29, 0x63, 0xfa, 0x72, 0xad, 0x8d, 0xb5, 0x9d, 0x46, 0xcd, 0xd9, 0x3b,
	0xc0, 0x55, 0xd0, 0xde, 0x3a, 0xe6, 0xb3, 0x19, 0xd8, 0x59, 0xea, 0xda, 0x8d, 0x62, 0x67, 0x29,
	0x7b, 0x81, 0x89, 0x8c, 0x4b, 0x2c, 0xc8, 0xbd, 0xd8, 0x58, 0xdb, 0xeb, 0xd7, 0xc5, 0x4e, 0xbf,
	0x72, 0x2f, 0x38, 0xc7, 0x84, 0x32, 0x51, 0x44, 0x2d, 0xe4, 0x1d, 0x00, 0xb4, 0xfa, 0xaf, 0x82,
	0xe7, 0x5e, 0xc1, 0xdc, 0x28, 0x38, 0xbf, 0xb4, 0xbb, 0xfe, 0x16, 0x6e, 0x42, 0x24, 0x3d, 0xa2,
	0xe7, 0xc0, 0xec, 0x2c, 0x54, 0x52, 0x14, 0x15, 0x7a, 0x0c, 0x9c, 0x10, 0xe9, 0x9d, 0x73, 0x83,
	0x9a, 0xc3, 0x9d, 0xa1, 0x69, 0xf0, 0x43, 0xa6, 0x7f, 0x36, 0xf6, 0x0b, 0x1a, 0x9a, 0x06, 0x03,
	0xe4, 0xd8, 0x07, 0x0d, 0xad, 0x05, 0x57, 0xb0, 0x54, 0xdf, 0xc0, 0x58, 0x4c, 0xcb, 0xaf, 0xc1,
	0x1d, 0x5a, 0x6d, 0xec, 0x01, 0xee, 0x43, 0xa4, 0x61, 0x66, 0x09, 0x8b, 0x9e, 0xae, 0x7b, 0xd4,
	0x38, 0xa3, 0x3d, 0x43, 0x4b, 0xc7, 0xd4, 0x70, 0xa3, 0xb1, 0xa1, 0xa5, 0x63, 0x6a, 0xd4, 0xd1,
	0xd8, 0xd0, 0x52, 0xb1, 0xe3, 0xa4, 0xf9, 0x2f, 0xdf, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2b,
	0x09, 0x72, 0xe4, 0xe0, 0x02, 0x00, 0x00,
}
