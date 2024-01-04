package controller

import (
	"BasicProject/logic"
	"BasicProject/middlewares/JWT"
	"BasicProject/middlewares/cache"
	"BasicProject/models"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
)

// 可以通过邮箱注册，需要做的步骤是首先在数据库查询是否已经有这个邮箱，有的话返回错误
func HandleUserSiginIn(ctx *gin.Context) {
	// 1.获取请求参数
	var fo *models.RegisterForm

	// 2.校验数据的有效性
	if err := ctx.ShouldBindJSON(&fo); err != nil {
		zap.L().Error("Sign In with invalid params", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "请求参数错误",
		})
		return
	}
	// 3. 交给logic层
	if err := logic.SignIn(fo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		//生成token
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}

}

// 处理登录业务
func HanlerUserLogin(ctx *gin.Context) {
	// 1.获取请求参数
	var fo *models.LoginForm

	// 2.校验数据的有效性
	if err := ctx.ShouldBindJSON(&fo); err != nil {
		zap.L().Error("Sign In with invalid params", zap.Error(err))
		return
	}
	// 3. 交给Logic层
	result, tempuser, _ := logic.Login(fo)
	if result == true {
		// 登录成功
		strToken, _ := JWT.GenToken(tempuser.Id)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "login success",
			"token":   strToken,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "login error",
		})
	}
}

// 请求发送验证码
func HandlerSendSMSForLogin(ctx *gin.Context) {
	var fo *models.SMS
	if err := ctx.ShouldBindJSON(&fo); err != nil {
		zap.L().Error("Sign In with invalid params", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}

	// 没有问题的话发送验证码，交给Logic层处理
	if err := logic.SMSLogin(fo.Phone); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "send code error",
		})
	} else {
		// 发送验证码没问题，将验证码存储到redis中，设置过期时间5分钟

		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}

}

/*
处理使用验证码登录的请求
*/
func HandlerUserSMSLogin(ctx *gin.Context) {
	/*
		1. 看看请求参数是否正确
		2. 验证验证码
	*/
	var fo *models.VerifySMSLogin
	if err := ctx.ShouldBindJSON(&fo); err != nil {
		fmt.Println("请求参数错误")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求错误",
		})
	}

}

// 处理获取用户信息请求
func HandlerUserProfile(ctx *gin.Context) {
	userId, _ := ctx.Get("userid")
	userIdStr := fmt.Sprintf("%v", userId)
	if len(userIdStr) <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Get User Profile Error",
		})
	}
	userinfo, _ := logic.GetUserProfileById(userIdStr)
	// 设置userinfo到redis中缓存
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "success",
		"userprofile": userinfo,
	})
	if err := cache.SetCacheByUserId(&userinfo, userinfo.Id); err != nil {
		log.Println("Set User Profile Cache ERROR", err)
	}

}

func HandleEditProfile(ctx *gin.Context) {

	// 1.获取请求参数
	var fo *models.EditUserProfile
	// 2.校验数据的有效性
	if err := ctx.ShouldBindJSON(&fo); err != nil {
		zap.L().Error("Sign In with invalid params", zap.Error(err))
		return
	}
	// 3.logic层
	userId, _ := ctx.Get("userid")
	userStr, _ := userId.(string)
	if err := logic.EditUserProfile(userStr, fo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func HandleTestSession(c *gin.Context) {
	c.JSON(http.StatusOK, "Welcome ")
}

func HandleGetSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("user-session", "username")
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
