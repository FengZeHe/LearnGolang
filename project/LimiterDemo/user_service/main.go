package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"limiterdemo/user_service/proto/user_service"
	"log"
	"net"
	"net/http"
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
	log.Println("Staring user service in 50051...")

	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	conn, err := grpc.NewClient("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux()
	ctx := context.Background()
	err = service.RegisterUserServiceHandler(ctx, mux, conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Staring gRPC gateway on :8081")
	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal("Failed to serve HTTP:", err)
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
