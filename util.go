//go:generate godocdown -o README.md

package api

import (
	"github.com/autom8ter/api/common"
	"google.golang.org/grpc"
)

var (
	DEFAULT_OAUTH_REDIRECT = common.ToString("http://localhost:8080/callback")
	DEFAULT_OAUTH_SCOPES   = []Scope{Scope_OPENID, Scope_PROFILE, Scope_EMAIL}
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

func (p *SubscriptionResponse) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *SMSResponse) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *CallResponse) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *User) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *PhoneNumberResource) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *PhoneNumber) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *JSONWebKeys) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *TokenQuery) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}
func (p *Identity) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}
func (p *Auth) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *Card) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}

func (p *Jwks) JSONString() (*common.String, error) {
	return common.MessageToJSONString(p)
}
