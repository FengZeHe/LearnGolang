package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	//user_service "github.com/basicprojectv2/user_service/
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

type App struct {
	grpcServer *grpc.Server
	gwServer   *http.Server
	userSvc    *UserService
}

func NewApp(userSvc *UserService) *App {
	return &App{
		grpcServer: grpc.NewServer(),
		userSvc:    userSvc,
	}
}

func (a *App) Start() error {
	RegisterUserServiceServer(a.grpcServer, a.userSvc)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	log.Println("start gRPC Server on 50051")

	go func() {
		if err := a.grpcServer.Serve(lis); err != nil {
			log.Println(err)
		}
	}()

	conn, err := grpc.NewClient(
		"0.0.0.0:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}
	defer conn.Close()

	gwmux := runtime.NewServeMux()
	err = RegisterUserServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatal("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	// 8090端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on in 8090...")
	log.Fatalln(gwServer.ListenAndServe())
	return a.gwServer.ListenAndServe()
}
