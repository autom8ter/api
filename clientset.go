package api

import (
	"google.golang.org/grpc"
)

type ClientSet struct {
	Utility    UtilityServiceClient
}

func NewClientSet(conn *grpc.ClientConn) *ClientSet {
	return &ClientSet{
		Utility: NewUtilityServiceClient(conn),
	}
}
