package router

import (
	"BasicProject/controller"
	"BasicProject/middlewares/JWT"
	"BasicProject/middlewares/cache"
	"BasicProject/middlewares/cors"
	"BasicProject/middlewares/session"
	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(cors.Cors())
	//r.Use(cors.Default())
	session.InitSession(r)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/signin", controller.HandleUserSignIn)
		v1.POST("/login", controller.HanlerUserLogin)
		v1.GET("/user/profile", JWT.JWTAuth(), cache.CacheMiddleWare(cache.KeyUserIdSet), controller.HandlerUserProfile)
		v1.POST("/user/edit", JWT.JWTAuth(), controller.HandleEditProfile)
		v1.POST("/sendSMS", controller.HandlerSendSMSForLogin) // 发送验证码
		v1.POST("/smsLogin", controller.HandlerUserSMSLogin)   // 使用验证码登录
	}
	v2 := r.Group("/api/v2")
	{
		v2.GET("/getsession", controller.HandleGetSession)
		v2.GET("/login", session.SessionMiddleware(), controller.HandleTestSession)
		v2.POST("/sendSMS", controller.HandlerUserSMSForLoginV2)
		v2.POST("/smsLogin", controller.HandlerUserSMSLoginV2)
	}
	v3 := r.Group("/api/v3")
	{
		/*
			设置本地缓存为一个中间件
		*/
		v3.POST("/user/profile", controller.HandleUserProfileV2)
		v3.POST("/user/edit", controller.HandleEditProfileV2)
	}
	r.GET("/hello", controller.HandleHello)

	return r

}
