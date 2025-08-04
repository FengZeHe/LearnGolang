package main

import (
	"context"
	"google.golang.org/grpc"
	"limiterdemo/user_service/proto/user_service"
	"net"
)

type UserService struct {
	service.UnimplementedUserServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	service.RegisterUserServiceServer(s, &UserService{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *UserService) GetUserById(ctx context.Context, req *service.GetUserReq) (*service.User, error) {
	userId := req.Id
	user := &service.User{
		Userid:   userId,
		Username: "熊二",
		Age:      "8",
	}
	return user, nil
}
