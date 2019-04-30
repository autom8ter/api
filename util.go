//go:generate godocdown -o README.md

package api

import (
	"encoding/json"
	"github.com/autom8ter/api/common"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

type ClientSet struct {
	Utility UtilityServiceClient
	Contact ContactServiceClient
	Payment PaymentServiceClient
	User    UserServiceClient
}

func NewClientSet(conn *grpc.ClientConn) *ClientSet {
	return &ClientSet{
		Utility: NewUtilityServiceClient(conn),
		Contact: NewContactServiceClient(conn),
		Payment: NewPaymentServiceClient(conn),
		User:    NewUserServiceClient(conn),
	}
}

func (p *SubscriptionResponse) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *SMSResponse) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *CallResponse) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *User) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *PhoneNumberResource) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *PhoneNumber) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *JSONWebKeys) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *TokenQuery) JSONString() *common.String {
	return common.MessageToJSONString(p)
}
func (p *Identity) JSONString() *common.String {
	return common.MessageToJSONString(p)
}
func (p *Auth) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Card) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Jwks) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Jwks) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *Card) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *Identity) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *User) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *SMSResponse) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *CallResponse) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *FaxResponse) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *SubscriptionResponse) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *PhoneNumberResource) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *PhoneNumber) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *JSONWebKeys) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *TokenQuery) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *Auth) UnmarshalJSONFrom(bits []byte) error {
	return json.Unmarshal(bits, p)
}

func (p *TokenQuery) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *Auth) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *JSONWebKeys) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *Jwks) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *FaxResponse) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *CallResponse) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *SMS) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *Call) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *Fax) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *FaxBlast) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *SMS) JSONString(bits []byte) *common.String {
	return common.MessageToJSONString(p)
}

func (p *Call) JSONString(bits []byte) *common.String {
	return common.MessageToJSONString(p)
}

func (p *Fax) JSONString(bits []byte) *common.String {
	return common.MessageToJSONString(p)
}

func (p *FaxBlast) JSONString(bits []byte) *common.String {
	return common.MessageToJSONString(p)
}

func (p *FaxResponse) JSONString(bits []byte) *common.String {
	return common.MessageToJSONString(p)
}

func (p *Email) JSONString(bits []byte) *common.String {
	return common.MessageToJSONString(p)
}

func (p *EmailBlastRequest) JSONString(bits []byte) *common.String {
	return common.MessageToJSONString(p)
}

func (p *EmailRequest) JSONString(bits []byte) *common.String {
	return common.MessageToJSONString(p)
}
func (p *Email) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *EmailBlastRequest) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}

func (p *EmailRequest) UnmarshalProtoFrom(bits []byte) error {
	return proto.Unmarshal(bits, p)
}
