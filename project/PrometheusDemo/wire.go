//go:build wireinject

package main

import (
	"github.com/google/wire"
	"prometheusdemo/internal/repository"
	"prometheusdemo/internal/repository/cache"
	"prometheusdemo/internal/repository/dao"
	"prometheusdemo/internal/service"
	"prometheusdemo/internal/web"
	"prometheusdemo/internal/web/middlewares"
	"prometheusdemo/ioc"
)

func InitializeApp() *App {
	wire.Build(
		ioc.InitMysqlConfig,
		ioc.InitRedisConfig,

		ioc.InitDB,
		ioc.InitRedis,
		ioc.InitRedisClient,

		// cache部分
		cache.NewUserCache,
		middlewares.NewMiddleUserCache,

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
