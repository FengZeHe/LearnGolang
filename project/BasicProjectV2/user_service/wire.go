//go:build wireinject

package main

import (
	"github.com/basicprojectv2/user_service/ioc"
	"github.com/basicprojectv2/user_service/repository"
	"github.com/basicprojectv2/user_service/repository/dao"
	"github.com/basicprojectv2/user_service/service"
	"github.com/basicprojectv2/user_service/settings"
	"github.com/google/wire"
)

func InitializeApp() *App {
	wire.Build(
		// 服务注册
		ioc.NewEtcdConfig,

		settings.InitMysqlConfig,
		ioc.InitDB,
		//ioc.NewGrpcServer,
		dao.NewUserDAO,
		service.NewUserService,
		repository.NewUserRepository,
		NewUerService,
		NewApp,
	)
	return &App{}
}
