//go:generate godocdown -o README.md

package api

import (
	"encoding/json"
	"errors"
	"github.com/autom8ter/api/common"
	"github.com/dgrijalva/jwt-go"
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

func (p *SMS) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Call) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Fax) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *FaxBlast) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *FaxResponse) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *Email) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *EmailBlastRequest) JSONString() *common.String {
	return common.MessageToJSONString(p)
}

func (p *EmailRequest) JSONString() *common.String {
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

func (c *Jwks) TokenCert(token *jwt.Token) (string, error) {
	var cert string
	for k, _ := range c.Keys {
		if token.Header["kid"] == c.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + c.Keys[k].X5C.Strings[0].Text + "\n-----END CERTIFICATE-----"
		}
	}
	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}
	return cert, nil
}
