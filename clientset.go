package api

import (
	"google.golang.org/grpc"
)

type ClientSet struct {
	Utility UtilityServiceClient
	Contact ContactServiceClient
	User    UserServiceClient
	Payment PaymentServiceClient
}

func NewClientSet(conn *grpc.ClientConn) *ClientSet {
	return &ClientSet{
		Utility: NewUtilityServiceClient(conn),
		Contact: NewContactServiceClient(conn),
		User:    NewUserServiceClient(conn),
	}
}
