package router

import (
	"BasicProject/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(cors.Default())
	v1 := r.Group("/api/v1")
	{
		v1.POST("/signin", controller.HandleUserSiginIn)
		v1.POST("/login", controller.HanlerUserLogin)
		v1.GET("/user/profile", controller.HandlerUserProfile)
		v1.POST("/user/edit", controller.HandleEditProfile)
	}
	return r

}
