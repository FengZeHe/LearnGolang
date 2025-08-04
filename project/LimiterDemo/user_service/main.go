package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"limiterdemo/user_service/limiter"
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

	//每秒生成100个令牌，桶容量200（允许突发200个请求）
	li := limiter.NewTokenBucketLimiter(rate.Limit(1), 2)

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			li.LimitInterceptor()),
	)
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
