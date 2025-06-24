package main

import (
	"context"
	pb "github.com/basicprojectv2/proto/user_service"
	"github.com/basicprojectv2/user_service/interceptors/jwt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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
	jwtInterceptor := jwt.NewJWTInterceptor([]string{
		// 免检路径
		/*
			    在 gRPC 中，每个服务方法都有一个唯一的全限定路径（Full Method Name），
				格式为：/包名.服务名/方法名。这个路径用于客户端与服务器之间的通信，也是拦截器中配置免检路径的依据。
						package声明/ proto中对应的service关键字  / Userlogin
		*/
		"/user_service.UserService/UserLogin",
	})

	grpcSerer := grpc.NewServer(
		grpc.UnaryInterceptor(jwtInterceptor.UnaryInterceptor()))

	return &App{
		grpcServer: grpcSerer,
		userSvc:    userSvc,
	}
}

func (a *App) Start() error {
	pb.RegisterUserServiceServer(a.grpcServer, a.userSvc)

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
	err = pb.RegisterUserServiceHandler(context.Background(), gwmux, conn)
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
