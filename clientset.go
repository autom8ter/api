package api

import (
	"google.golang.org/grpc"
)

type ClientSet struct {
	Echoer    EchoServiceClient
	Marshaler MarshalServiceClient
}

func NewClientSet(conn *grpc.ClientConn) *ClientSet {
	return &ClientSet{
		Echoer:    NewEchoServiceClient(conn),
		Marshaler: NewMarshalServiceClient(conn),
	}
}
