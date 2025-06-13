package main

import (
	"github.com/basicprojectv2/internal/repository"
	"github.com/basicprojectv2/internal/repository/cache"
	"github.com/basicprojectv2/internal/repository/dao"
	"github.com/basicprojectv2/internal/service"
	"github.com/basicprojectv2/ioc"
	"github.com/basicprojectv2/settings"
	"github.com/basicprojectv2/user_service"
	"google.golang.org/grpc"
	"log"
	"net"
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
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
