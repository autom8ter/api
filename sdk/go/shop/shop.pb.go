// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shop/shop.proto

package shop

import (
	fmt "fmt"
	geo "github.com/autom8ter/api/sdk/go/geo"
	time "github.com/autom8ter/api/sdk/go/time"
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

type ProductType int32

const (
	ProductType_SERVICE  ProductType = 0
	ProductType_CONTRACT ProductType = 1
	ProductType_PRODUCT  ProductType = 2
	ProductType_DIGITAL  ProductType = 3
)

var ProductType_name = map[int32]string{
	0: "SERVICE",
	1: "CONTRACT",
	2: "PRODUCT",
	3: "DIGITAL",
}

var ProductType_value = map[string]int32{
	"SERVICE":  0,
	"CONTRACT": 1,
	"PRODUCT":  2,
	"DIGITAL":  3,
}

func (x ProductType) String() string {
	return proto.EnumName(ProductType_name, int32(x))
}

func (ProductType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cbe66af8c4005172, []int{0}
}

type Rating int32

const (
	Rating_EXCELLENT Rating = 0
	Rating_GOOD      Rating = 1
	Rating_AVERAGE   Rating = 2
	Rating_POOR      Rating = 3
	Rating_BAD       Rating = 4
)

var Rating_name = map[int32]string{
	0: "EXCELLENT",
	1: "GOOD",
	2: "AVERAGE",
	3: "POOR",
	4: "BAD",
}

var Rating_value = map[string]int32{
	"EXCELLENT": 0,
	"GOOD":      1,
	"AVERAGE":   2,
	"POOR":      3,
	"BAD":       4,
}

func (x Rating) String() string {
	return proto.EnumName(Rating_name, int32(x))
}

func (Rating) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cbe66af8c4005172, []int{1}
}

type CardType int32

const (
	CardType_VISA       CardType = 0
	CardType_MASTERCARD CardType = 1
)

var CardType_name = map[int32]string{
	0: "VISA",
	1: "MASTERCARD",
}

var CardType_value = map[string]int32{
	"VISA":       0,
	"MASTERCARD": 1,
}

func (x CardType) String() string {
	return proto.EnumName(CardType_name, int32(x))
}

func (CardType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cbe66af8c4005172, []int{2}
}

type Card struct {
	CardType             CardType           `protobuf:"varint,1,opt,name=card_type,json=cardType,proto3,enum=shop.CardType" json:"card_type,omitempty"`
	FullName             string             `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Number               int64              `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	Code                 int64              `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	Expiration           string             `protobuf:"bytes,5,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Postal               *geo.PostalAddress `protobuf:"bytes,6,opt,name=postal,proto3" json:"postal,omitempty"`
	Owner                string             `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbe66af8c4005172, []int{0}
}

func (m *Card) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Card.Unmarshal(m, b)
}
func (m *Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Card.Marshal(b, m, deterministic)
}
func (m *Card) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Card.Merge(m, src)
}
func (m *Card) XXX_Size() int {
	return xxx_messageInfo_Card.Size(m)
}
func (m *Card) XXX_DiscardUnknown() {
	xxx_messageInfo_Card.DiscardUnknown(m)
}

var xxx_messageInfo_Card proto.InternalMessageInfo

func (m *Card) GetCardType() CardType {
	if m != nil {
		return m.CardType
	}
	return CardType_VISA
}

func (m *Card) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *Card) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Card) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Card) GetExpiration() string {
	if m != nil {
		return m.Expiration
	}
	return ""
}

func (m *Card) GetPostal() *geo.PostalAddress {
	if m != nil {
		return m.Postal
	}
	return nil
}

func (m *Card) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type Review struct {
	Author               string        `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Comment              string        `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Rating               int64         `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
	Tracker              *time.Tracker `protobuf:"bytes,4,opt,name=tracker,proto3" json:"tracker,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Review) Reset()         { *m = Review{} }
func (m *Review) String() string { return proto.CompactTextString(m) }
func (*Review) ProtoMessage()    {}
func (*Review) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbe66af8c4005172, []int{1}
}

func (m *Review) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Review.Unmarshal(m, b)
}
func (m *Review) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Review.Marshal(b, m, deterministic)
}
func (m *Review) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Review.Merge(m, src)
}
func (m *Review) XXX_Size() int {
	return xxx_messageInfo_Review.Size(m)
}
func (m *Review) XXX_DiscardUnknown() {
	xxx_messageInfo_Review.DiscardUnknown(m)
}

var xxx_messageInfo_Review proto.InternalMessageInfo

func (m *Review) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Review) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Review) GetRating() int64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Review) GetTracker() *time.Tracker {
	if m != nil {
		return m.Tracker
	}
	return nil
}

type Picture struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Picture) Reset()         { *m = Picture{} }
func (m *Picture) String() string { return proto.CompactTextString(m) }
func (*Picture) ProtoMessage()    {}
func (*Picture) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbe66af8c4005172, []int{2}
}

func (m *Picture) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Picture.Unmarshal(m, b)
}
func (m *Picture) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Picture.Marshal(b, m, deterministic)
}
func (m *Picture) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Picture.Merge(m, src)
}
func (m *Picture) XXX_Size() int {
	return xxx_messageInfo_Picture.Size(m)
}
func (m *Picture) XXX_DiscardUnknown() {
	xxx_messageInfo_Picture.DiscardUnknown(m)
}

var xxx_messageInfo_Picture proto.InternalMessageInfo

func (m *Picture) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Picture) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type Slideshow struct {
	Pictures             []*Picture `protobuf:"bytes,2,rep,name=pictures,proto3" json:"pictures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Slideshow) Reset()         { *m = Slideshow{} }
func (m *Slideshow) String() string { return proto.CompactTextString(m) }
func (*Slideshow) ProtoMessage()    {}
func (*Slideshow) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbe66af8c4005172, []int{3}
}

func (m *Slideshow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Slideshow.Unmarshal(m, b)
}
func (m *Slideshow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Slideshow.Marshal(b, m, deterministic)
}
func (m *Slideshow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Slideshow.Merge(m, src)
}
func (m *Slideshow) XXX_Size() int {
	return xxx_messageInfo_Slideshow.Size(m)
}
func (m *Slideshow) XXX_DiscardUnknown() {
	xxx_messageInfo_Slideshow.DiscardUnknown(m)
}

var xxx_messageInfo_Slideshow proto.InternalMessageInfo

func (m *Slideshow) GetPictures() []*Picture {
	if m != nil {
		return m.Pictures
	}
	return nil
}

type Inventory struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Inventory) Reset()         { *m = Inventory{} }
func (m *Inventory) String() string { return proto.CompactTextString(m) }
func (*Inventory) ProtoMessage()    {}
func (*Inventory) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbe66af8c4005172, []int{4}
}

func (m *Inventory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Inventory.Unmarshal(m, b)
}
func (m *Inventory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Inventory.Marshal(b, m, deterministic)
}
func (m *Inventory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inventory.Merge(m, src)
}
func (m *Inventory) XXX_Size() int {
	return xxx_messageInfo_Inventory.Size(m)
}
func (m *Inventory) XXX_DiscardUnknown() {
	xxx_messageInfo_Inventory.DiscardUnknown(m)
}

var xxx_messageInfo_Inventory proto.InternalMessageInfo

func (m *Inventory) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type Product struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Info                 string            `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Slideshow            *Slideshow        `protobuf:"bytes,3,opt,name=slideshow,proto3" json:"slideshow,omitempty"`
	Price                float32           `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Remaining            int64             `protobuf:"varint,5,opt,name=remaining,proto3" json:"remaining,omitempty"`
	Rating               int64             `protobuf:"varint,6,opt,name=rating,proto3" json:"rating,omitempty"`
	Reviews              []*Review         `protobuf:"bytes,7,rep,name=reviews,proto3" json:"reviews,omitempty"`
	DeliveryDate         string            `protobuf:"bytes,8,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	RefundPolicy         string            `protobuf:"bytes,9,opt,name=refund_policy,json=refundPolicy,proto3" json:"refund_policy,omitempty"`
	Cogs                 float32           `protobuf:"fixed32,10,opt,name=cogs,proto3" json:"cogs,omitempty"`
	Labels               map[string]string `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ProdType             ProductType       `protobuf:"varint,12,opt,name=prod_type,json=prodType,proto3,enum=shop.ProductType" json:"prod_type,omitempty"`
	Tracker              *time.Tracker     `protobuf:"bytes,13,opt,name=tracker,proto3" json:"tracker,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbe66af8c4005172, []int{5}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *Product) GetSlideshow() *Slideshow {
	if m != nil {
		return m.Slideshow
	}
	return nil
}

func (m *Product) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Product) GetRemaining() int64 {
	if m != nil {
		return m.Remaining
	}
	return 0
}

func (m *Product) GetRating() int64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Product) GetReviews() []*Review {
	if m != nil {
		return m.Reviews
	}
	return nil
}

func (m *Product) GetDeliveryDate() string {
	if m != nil {
		return m.DeliveryDate
	}
	return ""
}

func (m *Product) GetRefundPolicy() string {
	if m != nil {
		return m.RefundPolicy
	}
	return ""
}

func (m *Product) GetCogs() float32 {
	if m != nil {
		return m.Cogs
	}
	return 0
}

func (m *Product) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Product) GetProdType() ProductType {
	if m != nil {
		return m.ProdType
	}
	return ProductType_SERVICE
}

func (m *Product) GetTracker() *time.Tracker {
	if m != nil {
		return m.Tracker
	}
	return nil
}

type Receipt struct {
	LastFourDigits       string   `protobuf:"bytes,1,opt,name=last_four_digits,json=lastFourDigits,proto3" json:"last_four_digits,omitempty"`
	Cart                 *Cart    `protobuf:"bytes,2,opt,name=cart,proto3" json:"cart,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Receipt) Reset()         { *m = Receipt{} }
func (m *Receipt) String() string { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()    {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbe66af8c4005172, []int{6}
}

func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Receipt.Unmarshal(m, b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Receipt.Marshal(b, m, deterministic)
}
func (m *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(m, src)
}
func (m *Receipt) XXX_Size() int {
	return xxx_messageInfo_Receipt.Size(m)
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetLastFourDigits() string {
	if m != nil {
		return m.LastFourDigits
	}
	return ""
}

func (m *Receipt) GetCart() *Cart {
	if m != nil {
		return m.Cart
	}
	return nil
}

type Cart struct {
	Products             []*Product    `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	Total                float32       `protobuf:"fixed32,2,opt,name=total,proto3" json:"total,omitempty"`
	Reminder             *Reminder     `protobuf:"bytes,3,opt,name=reminder,proto3" json:"reminder,omitempty"`
	Tracker              *time.Tracker `protobuf:"bytes,4,opt,name=tracker,proto3" json:"tracker,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Cart) Reset()         { *m = Cart{} }
func (m *Cart) String() string { return proto.CompactTextString(m) }
func (*Cart) ProtoMessage()    {}
func (*Cart) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbe66af8c4005172, []int{7}
}

func (m *Cart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cart.Unmarshal(m, b)
}
func (m *Cart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cart.Marshal(b, m, deterministic)
}
func (m *Cart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cart.Merge(m, src)
}
func (m *Cart) XXX_Size() int {
	return xxx_messageInfo_Cart.Size(m)
}
func (m *Cart) XXX_DiscardUnknown() {
	xxx_messageInfo_Cart.DiscardUnknown(m)
}

var xxx_messageInfo_Cart proto.InternalMessageInfo

func (m *Cart) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *Cart) GetTotal() float32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Cart) GetReminder() *Reminder {
	if m != nil {
		return m.Reminder
	}
	return nil
}

func (m *Cart) GetTracker() *time.Tracker {
	if m != nil {
		return m.Tracker
	}
	return nil
}

type Reminder struct {
	EmailAfter           int64    `protobuf:"varint,1,opt,name=email_after,json=emailAfter,proto3" json:"email_after,omitempty"`
	SmsAfter             int64    `protobuf:"varint,2,opt,name=sms_after,json=smsAfter,proto3" json:"sms_after,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reminder) Reset()         { *m = Reminder{} }
func (m *Reminder) String() string { return proto.CompactTextString(m) }
func (*Reminder) ProtoMessage()    {}
func (*Reminder) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbe66af8c4005172, []int{8}
}

func (m *Reminder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reminder.Unmarshal(m, b)
}
func (m *Reminder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reminder.Marshal(b, m, deterministic)
}
func (m *Reminder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reminder.Merge(m, src)
}
func (m *Reminder) XXX_Size() int {
	return xxx_messageInfo_Reminder.Size(m)
}
func (m *Reminder) XXX_DiscardUnknown() {
	xxx_messageInfo_Reminder.DiscardUnknown(m)
}

var xxx_messageInfo_Reminder proto.InternalMessageInfo

func (m *Reminder) GetEmailAfter() int64 {
	if m != nil {
		return m.EmailAfter
	}
	return 0
}

func (m *Reminder) GetSmsAfter() int64 {
	if m != nil {
		return m.SmsAfter
	}
	return 0
}

func init() {
	proto.RegisterEnum("shop.ProductType", ProductType_name, ProductType_value)
	proto.RegisterEnum("shop.Rating", Rating_name, Rating_value)
	proto.RegisterEnum("shop.CardType", CardType_name, CardType_value)
	proto.RegisterType((*Card)(nil), "shop.Card")
	proto.RegisterType((*Review)(nil), "shop.Review")
	proto.RegisterType((*Picture)(nil), "shop.Picture")
	proto.RegisterType((*Slideshow)(nil), "shop.Slideshow")
	proto.RegisterType((*Inventory)(nil), "shop.Inventory")
	proto.RegisterType((*Product)(nil), "shop.Product")
	proto.RegisterMapType((map[string]string)(nil), "shop.Product.LabelsEntry")
	proto.RegisterType((*Receipt)(nil), "shop.Receipt")
	proto.RegisterType((*Cart)(nil), "shop.Cart")
	proto.RegisterType((*Reminder)(nil), "shop.Reminder")
}

func init() { proto.RegisterFile("shop/shop.proto", fileDescriptor_cbe66af8c4005172) }

var fileDescriptor_cbe66af8c4005172 = []byte{
	// 836 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xeb, 0x44,
	0x10, 0xae, 0x93, 0x34, 0xb6, 0xc7, 0x4d, 0x6b, 0x56, 0x08, 0x99, 0x82, 0x0e, 0x51, 0x40, 0x10,
	0x82, 0x48, 0x45, 0x90, 0x10, 0x70, 0xe7, 0x93, 0x98, 0x12, 0xa9, 0x34, 0xd1, 0x26, 0x54, 0xdc,
	0x45, 0xae, 0xbd, 0x49, 0xad, 0xda, 0x5e, 0x6b, 0xbd, 0x6e, 0x89, 0x78, 0x15, 0x9e, 0x8d, 0xf7,
	0xe0, 0x0e, 0xed, 0xac, 0x9d, 0xa6, 0x12, 0x42, 0x9c, 0x9b, 0xd5, 0xcc, 0x37, 0x33, 0xbb, 0xdf,
	0xce, 0x1f, 0x5c, 0x94, 0x0f, 0xbc, 0xb8, 0x52, 0xc7, 0xb8, 0x10, 0x5c, 0x72, 0xd2, 0x51, 0xf2,
	0xe5, 0x85, 0x4c, 0x32, 0x76, 0xa5, 0x0e, 0x0d, 0x5f, 0xf6, 0x76, 0x8c, 0x5f, 0xed, 0x18, 0xd7,
	0xea, 0xe0, 0x2f, 0x03, 0x3a, 0xd3, 0x50, 0xc4, 0xe4, 0x2b, 0xb0, 0xa3, 0x50, 0xc4, 0x1b, 0xb9,
	0x2f, 0x98, 0x67, 0xf4, 0x8d, 0xe1, 0xf9, 0xe4, 0x7c, 0x8c, 0xd7, 0x29, 0xf3, 0x7a, 0x5f, 0x30,
	0x6a, 0x45, 0xb5, 0x44, 0x3e, 0x02, 0x7b, 0x5b, 0xa5, 0xe9, 0x26, 0x0f, 0x33, 0xe6, 0xb5, 0xfa,
	0xc6, 0xd0, 0xa6, 0x96, 0x02, 0x6e, 0xc3, 0x8c, 0x91, 0x0f, 0xa0, 0x9b, 0x57, 0xd9, 0x3d, 0x13,
	0x5e, 0xbb, 0x6f, 0x0c, 0xdb, 0xb4, 0xd6, 0x08, 0x81, 0x4e, 0xc4, 0x63, 0xe6, 0x75, 0x10, 0x45,
	0x99, 0xbc, 0x01, 0x60, 0xbf, 0x17, 0x89, 0x08, 0x65, 0xc2, 0x73, 0xef, 0x14, 0x6f, 0x3a, 0x42,
	0xc8, 0x08, 0xba, 0x05, 0x2f, 0x65, 0x98, 0x7a, 0xdd, 0xbe, 0x31, 0x74, 0x26, 0x64, 0xac, 0xa8,
	0x2f, 0x11, 0xf2, 0xe3, 0x58, 0xb0, 0xb2, 0xa4, 0xb5, 0x07, 0x79, 0x1f, 0x4e, 0xf9, 0x73, 0xce,
	0x84, 0x67, 0xe2, 0x35, 0x5a, 0x19, 0xfc, 0x01, 0x5d, 0xca, 0x9e, 0x12, 0xf6, 0xac, 0x78, 0x85,
	0x95, 0x7c, 0xe0, 0x02, 0xbf, 0x67, 0xd3, 0x5a, 0x23, 0x1e, 0x98, 0x11, 0xcf, 0x32, 0x96, 0xcb,
	0xfa, 0x2b, 0x8d, 0xaa, 0x22, 0x14, 0x8f, 0x7c, 0xd7, 0xfc, 0x44, 0x6b, 0xe4, 0x0b, 0x30, 0xa5,
	0x08, 0xa3, 0x47, 0x26, 0xf0, 0x33, 0xce, 0xa4, 0x37, 0xc6, 0x0c, 0xaf, 0x35, 0x48, 0x1b, 0xeb,
	0xe0, 0x0a, 0xcc, 0x65, 0x12, 0xc9, 0x4a, 0x30, 0xe2, 0x42, 0xbb, 0x12, 0x69, 0xfd, 0xb4, 0x12,
	0x55, 0x3e, 0x92, 0x7c, 0xcb, 0xeb, 0x47, 0x51, 0x1e, 0x7c, 0x07, 0xf6, 0x2a, 0x4d, 0x62, 0x56,
	0x3e, 0xf0, 0x67, 0xf2, 0x25, 0x58, 0x85, 0x8e, 0x2e, 0xbd, 0x56, 0xbf, 0x8d, 0xef, 0x60, 0x45,
	0xea, 0x3b, 0xe9, 0xc1, 0xac, 0xe2, 0xe6, 0xf9, 0x13, 0xcb, 0x25, 0x17, 0x7b, 0x8c, 0x13, 0x3c,
	0xae, 0x22, 0x59, 0x7a, 0xc6, 0xab, 0x38, 0x8d, 0xd2, 0x83, 0x79, 0xf0, 0x77, 0x1b, 0xcc, 0x1a,
	0x55, 0x7c, 0xb0, 0x9e, 0x9a, 0x22, 0xca, 0xff, 0xc6, 0x91, 0x7c, 0x0d, 0x76, 0xd9, 0x70, 0xc4,
	0xc4, 0x38, 0x93, 0x0b, 0x7d, 0xff, 0x81, 0x3a, 0x7d, 0xf1, 0x50, 0x65, 0x29, 0x44, 0x12, 0xe9,
	0xba, 0xb7, 0xa8, 0x56, 0xc8, 0xc7, 0x60, 0x0b, 0x96, 0x85, 0x49, 0xae, 0xb2, 0x7b, 0x8a, 0xd9,
	0x7d, 0x01, 0x8e, 0x12, 0xdf, 0x7d, 0x95, 0xf8, 0xcf, 0xc1, 0x14, 0x58, 0xcc, 0xd2, 0x33, 0xf1,
	0x63, 0x67, 0xfa, 0x61, 0x5d, 0x61, 0xda, 0x18, 0xc9, 0xa7, 0xd0, 0x8b, 0x59, 0x9a, 0x3c, 0x31,
	0xb1, 0xdf, 0xc4, 0xa1, 0x64, 0x9e, 0x85, 0xfc, 0xcf, 0x1a, 0x70, 0x16, 0x4a, 0xa6, 0x9c, 0x04,
	0xdb, 0x56, 0x79, 0xbc, 0x29, 0x78, 0x9a, 0x44, 0x7b, 0xcf, 0xd6, 0x4e, 0x1a, 0x5c, 0x22, 0xa6,
	0x9b, 0x76, 0x57, 0x7a, 0x80, 0xe4, 0x51, 0x26, 0xdf, 0x40, 0x37, 0x0d, 0xef, 0x59, 0x5a, 0x7a,
	0x0e, 0x92, 0xf8, 0xf0, 0x55, 0x76, 0xc7, 0x37, 0x68, 0x0b, 0x72, 0x29, 0xf6, 0xb4, 0x76, 0x24,
	0x63, 0xb0, 0x55, 0xce, 0xf5, 0x74, 0x9d, 0xe1, 0x74, 0xbd, 0xf7, 0x2a, 0x4a, 0x0f, 0x98, 0xf2,
	0xc1, 0x01, 0x3b, 0xea, 0xb0, 0xde, 0x7f, 0x75, 0xd8, 0xe5, 0x0f, 0xe0, 0x1c, 0xbd, 0xa7, 0xba,
	0xec, 0x91, 0xed, 0x9b, 0x2e, 0x7b, 0x64, 0x7b, 0x95, 0xfe, 0xa7, 0x30, 0xad, 0x9a, 0x31, 0xd5,
	0xca, 0x8f, 0xad, 0xef, 0x8d, 0xc1, 0x0a, 0x4c, 0xca, 0x22, 0x96, 0x14, 0x92, 0x0c, 0xc1, 0x4d,
	0xc3, 0x52, 0x6e, 0xb6, 0xbc, 0x12, 0x9b, 0x38, 0xd9, 0x25, 0xd8, 0x39, 0xca, 0xff, 0x5c, 0xe1,
	0x3f, 0xf1, 0x4a, 0xcc, 0x10, 0x25, 0x6f, 0xa0, 0x13, 0x85, 0x42, 0x4f, 0x8a, 0x33, 0x81, 0xc3,
	0x86, 0x90, 0x14, 0xf1, 0xc1, 0x9f, 0x7a, 0x9f, 0xc8, 0x77, 0x68, 0x42, 0x45, 0x51, 0x72, 0x35,
	0xe3, 0x2d, 0xdd, 0x21, 0xa8, 0x90, 0x11, 0x58, 0x82, 0x65, 0x49, 0x1e, 0xd7, 0x8b, 0xc4, 0x69,
	0xf6, 0x11, 0xad, 0x51, 0x7a, 0xb0, 0xff, 0xff, 0x81, 0xfc, 0x19, 0xac, 0x26, 0x9c, 0x7c, 0x02,
	0x8e, 0x6a, 0xb8, 0x74, 0x13, 0x6e, 0x25, 0xd3, 0x4b, 0xa1, 0x4d, 0x01, 0x21, 0x5f, 0x21, 0x6a,
	0xcb, 0x95, 0x59, 0x59, 0x9b, 0x5b, 0x68, 0xb6, 0xca, 0xac, 0x44, 0xe3, 0xe8, 0x2d, 0x38, 0x47,
	0xa5, 0x23, 0x0e, 0x98, 0xab, 0x80, 0xde, 0xcd, 0xa7, 0x81, 0x7b, 0x42, 0xce, 0xc0, 0x9a, 0x2e,
	0x6e, 0xd7, 0xd4, 0x9f, 0xae, 0x5d, 0x43, 0x99, 0x96, 0x74, 0x31, 0xfb, 0x75, 0xba, 0x76, 0x5b,
	0x4a, 0x99, 0xcd, 0xaf, 0xe7, 0x6b, 0xff, 0xc6, 0x6d, 0x8f, 0x7c, 0xe8, 0x52, 0xdd, 0xd8, 0x3d,
	0xb0, 0x83, 0xdf, 0xa6, 0xc1, 0xcd, 0x4d, 0x70, 0xbb, 0x76, 0x4f, 0x88, 0x05, 0x9d, 0xeb, 0xc5,
	0x62, 0xa6, 0x83, 0xfd, 0xbb, 0x80, 0xfa, 0xd7, 0x81, 0xdb, 0x52, 0xf0, 0x72, 0xb1, 0xa0, 0x6e,
	0x9b, 0x98, 0xd0, 0x7e, 0xeb, 0xcf, 0xdc, 0xce, 0xe8, 0x33, 0xb0, 0x9a, 0xfd, 0xac, 0xcc, 0x77,
	0xf3, 0x95, 0xef, 0x9e, 0x90, 0x73, 0x80, 0x5f, 0xfc, 0xd5, 0x3a, 0xa0, 0x53, 0x9f, 0xce, 0x5c,
	0xe3, 0xbe, 0x8b, 0xcb, 0xfe, 0xdb, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x55, 0xe0, 0x25,
	0x25, 0x06, 0x00, 0x00,
}
