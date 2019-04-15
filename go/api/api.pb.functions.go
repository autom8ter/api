package api

import (
	"context"
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"google.golang.org/grpc"
)

type EmailOption func(r *EmailRequest)
type SMSOption func(r *SMSRequest)
type MMSOption func(r *MMSRequest)
type CallOption func(r *CallRequest)
type ChargeOption func(r *ChargeRequest)
type SubscribeOption func(r *SubscribeCustomerRequest)
type RefundOption func(r *RefundRequest)
type PlanOption func(r *CreatePlanRequest)
type AccountOption func(r *CreateAccountRequest)
type MessageOption func(r *MessageUserRequest)

func NewEmailRequest(opts ...EmailOption) *EmailRequest {
	e := &EmailRequest{}
	for _, o := range opts {
		o(e)
	}
	return e
}

func NewSMSRequest(opts ...SMSOption) *SMSRequest {
	e := &SMSRequest{}
	for _, o := range opts {
		o(e)
	}
	return e
}

func NewCallRequest(opts ...CallOption) *CallRequest {
	e := &CallRequest{}
	for _, o := range opts {
		o(e)
	}
	return e
}

func NewMMSRequest(opts ...MMSOption) *MMSRequest {
	e := &MMSRequest{}
	for _, o := range opts {
		o(e)
	}
	return e
}

func NewSubscribeRequest(opts ...SubscribeOption) *SubscribeCustomerRequest {
	e := &SubscribeCustomerRequest{}
	for _, o := range opts {
		o(e)
	}
	return e
}

func NewChargeRequest(opts ...ChargeOption) *ChargeRequest {
	e := &ChargeRequest{}
	for _, o := range opts {
		o(e)
	}
	return e
}

func NewAccountRequest(opts ...AccountOption) *CreateAccountRequest {
	e := &CreateAccountRequest{}
	for _, o := range opts {
		o(e)
	}
	return e
}

func NewPlanRequest(opts ...PlanOption) *CreatePlanRequest {
	e := &CreatePlanRequest{}
	for _, o := range opts {
		o(e)
	}
	return e
}

func NewRefundRequest(opts ...RefundOption) *RefundRequest {
	e := &RefundRequest{}
	for _, o := range opts {
		o(e)
	}
	return e
}

func NewMessageRequest(opts ...MessageOption) *MessageUserRequest {
	e := &MessageUserRequest{}
	for _, o := range opts {
		o(e)
	}
	return e
}

/*
EmailUser(context.Context, *EmailRequest) (*JSON, error)
	MessageUser(ctx context.Context, r *UnImplemented)  (*JSON, error)
	CreateUser(ctx context.Context, r *UnImplemented)  (*JSON, error)
	UpdateUser(ctx context.Context, r *UnImplemented)  (*JSON, error)
	DeleteUser(ctx context.Context, r *UnImplemented)  (*JSON, error)
	ListUsers(ctx context.Context, r *UnImplemented)  (*JSON, error)
*/

type CustomerServiceServerFunctions struct {
	CreateCustomerFunc      func(context.Context, *CustomerRequest) (*JSON, error)
	UpdateCustomerFunc      func(context.Context, *UpdateCustomerRequest) (*JSON, error)
	DeleteCustomerFunc      func(context.Context, *Id) (*JSON, error)
	ListCustomersFunc       func(context.Context, *Empty) (*JSON, error)
	ChargeCustomerFunc      func(context.Context, *ChargeRequest) (*JSON, error)
	RefundCustomerFunc      func(context.Context, *RefundRequest) (*JSON, error)
	SubscribeCustomerFunc   func(context.Context, *SubscribeCustomerRequest) (*JSON, error)
	UnSubscribeCustomerFunc func(context.Context, *CancelSubscriptionRequest) (*JSON, error)
	SMSCustomerFunc         func(context.Context, *SMSRequest) (*JSON, error)
	CallCustomerFunc        func(context.Context, *CallRequest) (*JSON, error)
	MMSCustomerFunc         func(context.Context, *MMSRequest) (*JSON, error)
	EmailCustomerFunc       func(context.Context, *EmailRequest) (*JSON, error)
}

func NewCustomerServiceServerFunctions(createCustomerFunc func(context.Context, *CustomerRequest) (*JSON, error), updateCustomerFunc func(context.Context, *UpdateCustomerRequest) (*JSON, error), deleteCustomerFunc func(context.Context, *Id) (*JSON, error), listCustomersFunc func(context.Context, *Empty) (*JSON, error), chargeCustomerFunc func(context.Context, *ChargeRequest) (*JSON, error), refundCustomerFunc func(context.Context, *RefundRequest) (*JSON, error), subscribeCustomerFunc func(context.Context, *SubscribeCustomerRequest) (*JSON, error), unSubscribeCustomerFunc func(context.Context, *CancelSubscriptionRequest) (*JSON, error), SMSCustomerFunc func(context.Context, *SMSRequest) (*JSON, error), callCustomerFunc func(context.Context, *CallRequest) (*JSON, error), MMSCustomerFunc func(context.Context, *MMSRequest) (*JSON, error), emailCustomerFunc func(context.Context, *EmailRequest) (*JSON, error)) *CustomerServiceServerFunctions {
	return &CustomerServiceServerFunctions{CreateCustomerFunc: createCustomerFunc, UpdateCustomerFunc: updateCustomerFunc, DeleteCustomerFunc: deleteCustomerFunc, ListCustomersFunc: listCustomersFunc, ChargeCustomerFunc: chargeCustomerFunc, RefundCustomerFunc: refundCustomerFunc, SubscribeCustomerFunc: subscribeCustomerFunc, UnSubscribeCustomerFunc: unSubscribeCustomerFunc, SMSCustomerFunc: SMSCustomerFunc, CallCustomerFunc: callCustomerFunc, MMSCustomerFunc: MMSCustomerFunc, EmailCustomerFunc: emailCustomerFunc}
}

func (a CustomerServiceServerFunctions) RegisterWithServer(s *grpc.Server) {
	RegisterCustomerServiceServer(s, a)
}

func (c CustomerServiceServerFunctions) DeleteCustomer(ctx context.Context, r *Id) (*JSON, error) {
	return c.DeleteCustomerFunc(ctx, r)
}

func (c CustomerServiceServerFunctions) CreateCustomer(ctx context.Context, r *CustomerRequest) (*JSON, error) {
	return c.CreateCustomerFunc(ctx, r)
}

func (c CustomerServiceServerFunctions) UpdateCustomer(ctx context.Context, r *UpdateCustomerRequest) (*JSON, error) {
	return c.UpdateCustomerFunc(ctx, r)
}

func (c CustomerServiceServerFunctions) ListCustomers(ctx context.Context, r *Empty) (*JSON, error) {
	return c.ListCustomersFunc(ctx, r)
}

func (c CustomerServiceServerFunctions) ChargeCustomer(ctx context.Context, r *ChargeRequest) (*JSON, error) {
	return c.ChargeCustomerFunc(ctx, r)
}

func (c CustomerServiceServerFunctions) RefundCustomer(ctx context.Context, r *RefundRequest) (*JSON, error) {
	return c.RefundCustomerFunc(ctx, r)
}

func (c CustomerServiceServerFunctions) SubscribeCustomer(ctx context.Context, r *SubscribeCustomerRequest) (*JSON, error) {
	return c.SubscribeCustomerFunc(ctx, r)
}

func (c CustomerServiceServerFunctions) UnSubscribeCustomer(ctx context.Context, r *CancelSubscriptionRequest) (*JSON, error) {
	return c.UnSubscribeCustomerFunc(ctx, r)
}

func (c CustomerServiceServerFunctions) SMSCustomer(ctx context.Context, r *SMSRequest) (*JSON, error) {
	return c.SMSCustomerFunc(ctx, r)
}

func (c CustomerServiceServerFunctions) CallCustomer(ctx context.Context, r *CallRequest) (*JSON, error) {
	return c.CallCustomerFunc(ctx, r)
}

func (c CustomerServiceServerFunctions) MMSCustomer(ctx context.Context, r *MMSRequest) (*JSON, error) {
	return c.MMSCustomerFunc(ctx, r)
}

func (c CustomerServiceServerFunctions) EmailCustomer(ctx context.Context, r *EmailRequest) (*JSON, error) {
	return c.EmailCustomerFunc(ctx, r)
}

type UserServiceServerFunctions struct {
	EmailUserFunc   func(context.Context, *EmailRequest) (*JSON, error)
	MessageUserFunc func(ctx context.Context, r *MessageUserRequest) (*JSON, error)
	CreateUserFunc  func(ctx context.Context, r *User) (*JSON, error)
	UpdateUserFunc  func(ctx context.Context, r *User) (*JSON, error)
	DeleteUserFunc  func(ctx context.Context, r *Id) (*JSON, error)
	ListUsersFunc   func(ctx context.Context, r *Empty) (*JSON, error)
	SMSUserFunc     func(ctx context.Context, r *SMSRequest) (*JSON, error)
	CallUserFunc    func(ctx context.Context, r *CallRequest) (*JSON, error)
	MMSUserFunc     func(ctx context.Context, r *MMSRequest) (*JSON, error)
}

func NewUserServiceServerFunctions(emailUserFunc func(context.Context, *EmailRequest) (*JSON, error), messageUserFunc func(ctx context.Context, r *MessageUserRequest) (*JSON, error), createUserFunc func(ctx context.Context, r *User) (*JSON, error), updateUserFunc func(ctx context.Context, r *User) (*JSON, error), deleteUserFunc func(ctx context.Context, r *Id) (*JSON, error), listUsersFunc func(ctx context.Context, r *Empty) (*JSON, error), SMSUserFunc func(ctx context.Context, r *SMSRequest) (*JSON, error), callUserFunc func(ctx context.Context, r *CallRequest) (*JSON, error), MMSUserFunc func(ctx context.Context, r *MMSRequest) (*JSON, error)) *UserServiceServerFunctions {
	return &UserServiceServerFunctions{EmailUserFunc: emailUserFunc, MessageUserFunc: messageUserFunc, CreateUserFunc: createUserFunc, UpdateUserFunc: updateUserFunc, DeleteUserFunc: deleteUserFunc, ListUsersFunc: listUsersFunc, SMSUserFunc: SMSUserFunc, CallUserFunc: callUserFunc, MMSUserFunc: MMSUserFunc}
}

func (u UserServiceServerFunctions) SMSUser(ctx context.Context, r *SMSRequest) (*JSON, error) {
	return u.SMSUserFunc(ctx, r)
}

func (u UserServiceServerFunctions) CallUser(ctx context.Context, r *CallRequest) (*JSON, error) {
	return u.CallUserFunc(ctx, r)
}

func (u UserServiceServerFunctions) MMSUser(ctx context.Context, r *MMSRequest) (*JSON, error) {
	return u.MMSUserFunc(ctx, r)
}

func (u UserServiceServerFunctions) RegisterWithServer(s *grpc.Server) {
	RegisterUserServiceServer(s, u)
}

func (u UserServiceServerFunctions) EmailUser(ctx context.Context, r *EmailRequest) (*JSON, error) {
	return u.EmailUserFunc(ctx, r)
}

func (u UserServiceServerFunctions) MessageUser(ctx context.Context, r *MessageUserRequest) (*JSON, error) {
	return u.MessageUserFunc(ctx, r)
}

func (u UserServiceServerFunctions) CreateUser(ctx context.Context, r *User) (*JSON, error) {
	return u.CreateUserFunc(ctx, r)
}

func (u UserServiceServerFunctions) UpdateUser(ctx context.Context, r *User) (*JSON, error) {
	return u.UpdateUserFunc(ctx, r)
}

func (u UserServiceServerFunctions) DeleteUser(ctx context.Context, r *Id) (*JSON, error) {
	return u.DeleteUserFunc(ctx, r)
}

func (u UserServiceServerFunctions) ListUsers(ctx context.Context, r *Empty) (*JSON, error) {
	return u.ListUsersFunc(ctx, r)
}

type PlanServiceServerFunctions struct {
	CreateSubscriptionPlanFunc func(context.Context, *CreatePlanRequest) (*JSON, error)
}

func NewPlanServiceServerFunctions(createSubscriptionPlanFunc func(context.Context, *CreatePlanRequest) (*JSON, error)) *PlanServiceServerFunctions {
	return &PlanServiceServerFunctions{CreateSubscriptionPlanFunc: createSubscriptionPlanFunc}
}

func (p PlanServiceServerFunctions) RegisterWithServer(s *grpc.Server) {
	RegisterPlanServiceServer(s, p)
}

func (p PlanServiceServerFunctions) CreateSubscriptionPlan(ctx context.Context, r *CreatePlanRequest) (*JSON, error) {
	return p.CreateSubscriptionPlanFunc(ctx, r)
}

type AccountServiceServerFunctions struct {
	CreateAccountFunc func(context.Context, *CreateAccountRequest) (*JSON, error)
	UpdateAccountFunc func(ctx context.Context, r *Account) (*JSON, error)
	DeleteAccountFunc func(ctx context.Context, r *Id) (*JSON, error)
	ReadAccountFunc   func(ctx context.Context, r *Id) (*JSON, error)
	ListAccountsFunc  func(ctx context.Context, r *Empty) (*JSON, error)
}

func NewAccountServiceServerFunctions(createAccountFunc func(context.Context, *CreateAccountRequest) (*JSON, error), updateAccountFunc func(ctx context.Context, r *Account) (*JSON, error), deleteAccountFunc func(ctx context.Context, r *Id) (*JSON, error), readAccountFunc func(ctx context.Context, r *Id) (*JSON, error), listAccountsFunc func(ctx context.Context, r *Empty) (*JSON, error)) *AccountServiceServerFunctions {
	return &AccountServiceServerFunctions{CreateAccountFunc: createAccountFunc, UpdateAccountFunc: updateAccountFunc, DeleteAccountFunc: deleteAccountFunc, ReadAccountFunc: readAccountFunc, ListAccountsFunc: listAccountsFunc}
}

func (a AccountServiceServerFunctions) RegisterWithServer(s *grpc.Server) {
	RegisterAccountServiceServer(s, a)
}

func (a AccountServiceServerFunctions) CreateAccount(ctx context.Context, r *CreateAccountRequest) (*JSON, error) {
	return a.CreateAccountFunc(ctx, r)
}

func (a AccountServiceServerFunctions) UpdateAccount(ctx context.Context, r *Account) (*JSON, error) {
	return a.UpdateAccountFunc(ctx, r)
}

func (a AccountServiceServerFunctions) DeleteAccount(ctx context.Context, r *Id) (*JSON, error) {
	return a.DeleteAccountFunc(ctx, r)
}

func (a AccountServiceServerFunctions) ReadAccount(ctx context.Context, r *Id) (*JSON, error) {
	return a.ReadAccountFunc(ctx, r)
}

func (a AccountServiceServerFunctions) ListAccounts(ctx context.Context, r *Empty) (*JSON, error) {
	return a.ListAccountsFunc(ctx, r)
}

func ServeFunctions(addr string, debug bool, accounts AccountServiceServerFunctions, customer CustomerServiceServerFunctions, user UserServiceServerFunctions, plan PlanServiceServerFunctions) {
	if err := engine.Default("tcp", addr, debug).With(
		config.WithPlugins(
			accounts,
			customer,
			user,
			plan,
		),
	).Serve(); err != nil {
		Util.Fatalln("functions:", err.Error())
	}
}
