package main

import (
	"github.com/basicprojectv2/internal/repository"
	"github.com/basicprojectv2/internal/repository/cache"
	"github.com/basicprojectv2/internal/repository/dao"
	"github.com/basicprojectv2/internal/service"
	"github.com/basicprojectv2/internal/service/web"
	"github.com/basicprojectv2/ioc"
	"github.com/basicprojectv2/settings"
	"github.com/gin-gonic/gin"
)

func main() {
	redisConf := settings.InitRedisConfig()
	cmdable := ioc.InitRedis(redisConf)
	v := ioc.InitGinMiddlewares()
	dbconf := settings.InitMysqlConfig()
	db := ioc.InitDB(dbconf)
	userDAO := dao.NewUserDAO(db)
	userCache := cache.NewUserCache(cmdable)
	userRepository := repository.NewCacheUserRepository(userDAO, userCache)
	userService := service.NewUserService(userRepository)
	userHandler := web.NewUserHandler(userService)
	engine := ioc.InitWebServer(v, userHandler)

	engine.GET("/hi", func(c *gin.Context) {
		c.JSON(200, "hello")
	})
	engine.Run(":8088")

}
