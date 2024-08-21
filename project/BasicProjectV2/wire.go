//go:build wireinject

package main

import (
	"github.com/basicprojectv2/internal/repository"
	"github.com/basicprojectv2/internal/repository/cache"
	"github.com/basicprojectv2/internal/repository/dao"
	"github.com/basicprojectv2/internal/service"
	"github.com/basicprojectv2/internal/web"
	"github.com/basicprojectv2/internal/web/middleware"
	"github.com/basicprojectv2/ioc"
	"github.com/basicprojectv2/settings"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp() *gin.Engine {
	wire.Build(
		// 读取配置
		settings.InitMysqlConfig, settings.InitRedisConfig,

		// 第三方依赖部分
		ioc.InitDB, ioc.InitRedis, ioc.InitMysqlCasbinEnforcer,

		// 测试Enforcer

		// Cache部分
		cache.NewCodeCache,
		cache.NewUserCache,

		// DAO部分
		dao.NewUserDAO,
		dao.NewSysDAO,
		dao.NewMenuDAO,
		dao.NewRoleDAO,

		// repository部分
		repository.NewCacheUserRepository,
		repository.NewCodeRepository,
		repository.NewSysRepository,
		repository.NewMenuRepository,
		repository.NewRoleRepository,

		// service部分
		ioc.InitSMSService,
		service.NewCodeService,
		service.NewUserService,
		service.NewSysService,
		service.NewMenuService,
		service.NewRoleService,

		//handler部分
		web.NewUserHandler,
		web.NewSysHandler,
		web.NewMenuHandler,
		web.NewRoleHandler,

		// 中间件和路由
		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
		middleware.NewCasbinRoleCheck,
	)
	return gin.Default()

}
