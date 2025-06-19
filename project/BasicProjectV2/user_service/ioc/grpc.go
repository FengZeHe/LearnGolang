package ioc

import "google.golang.org/grpc"

func NewGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	return s
}
