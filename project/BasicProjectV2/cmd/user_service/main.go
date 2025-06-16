package main

import (
	"context"
	"github.com/basicprojectv2/internal/repository"
	"github.com/basicprojectv2/internal/repository/cache"
	"github.com/basicprojectv2/internal/repository/dao"
	"github.com/basicprojectv2/internal/service"
	"github.com/basicprojectv2/ioc"
	"github.com/basicprojectv2/settings"
	"github.com/basicprojectv2/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func main() {

	mysqlConfig := settings.InitMysqlConfig()
	db := ioc.InitDB(mysqlConfig)
	userDAO := dao.NewUserDAO(db)
	redisConfig := settings.InitRedisConfig()
	cmdable := ioc.InitRedis(redisConfig)
	userCache := cache.NewUserCache(cmdable)

	repo := repository.NewCacheUserRepository(userDAO, userCache)
	svc := service.NewUserService(repo)

	userSvc := user_service.NewUerService(svc)

	// 创建gRPC服务器
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	user_service.RegisterUserServiceServer(s, userSvc)
	log.Println("Starting gRPC Server in 50051... ")

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	conn, err := grpc.NewClient(
		"0.0.0.0:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = user_service.RegisterUserServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	// 8090端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on in 8090...")
	log.Fatalln(gwServer.ListenAndServe())
}
