//go:build wireinject

package main

import (
	"github.com/google/wire"
	"prometheusdemo/internal/repository"
	"prometheusdemo/internal/repository/dao"
	"prometheusdemo/internal/service"
	"prometheusdemo/internal/web"
	"prometheusdemo/ioc"
)

func InitializeApp() *App {
	wire.Build(
		ioc.InitMysqlConfig,
		ioc.InitDB,

		dao.NewUserDAO,
		repository.NewCacheUserRepository,
		service.NewUserService,

		web.NewUserHandler,

		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
		wire.Struct(new(App), "*"),
	)
	return new(App)
}
