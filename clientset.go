package api

import (
	"google.golang.org/grpc"
)

type ClientSet struct {
	Echoer EchoServiceClient
}

func NewClientSet(conn *grpc.ClientConn) *ClientSet {
	return &ClientSet{
		Echoer: NewEchoServiceClient(conn),
	}
}
