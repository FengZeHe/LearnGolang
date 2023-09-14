package controller

import (
	"BasicProject/logic"
	"BasicProject/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// 可以通过邮箱注册，需要做的步骤是首先在数据库查询是否已经有这个邮箱，有的话返回错误
func HandleUserSiginIn(ctx *gin.Context) {
	// 1.获取请求参数
	var fo *models.RegisterForm

	// 2.校验数据的有效性
	if err := ctx.ShouldBindJSON(&fo); err != nil {
		zap.L().Error("Sign In with invalid params", zap.Error(err))
		return
	}
	// 3. 交给logic层
	if err := logic.SignIn(fo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Sign In Error",
		})
	} else {
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
	result, _ := logic.Login(fo)
	if result == true {
		// 登录成功
		ctx.JSON(http.StatusOK, gin.H{
			"message": "login success",
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "login error",
		})
	}
}

// 处理获取用户信息请求
func HandlerUserProfile(ctx *gin.Context) {
	email := ctx.Query("email")
	if len(email) <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Get User Profile Error",
		})
	}
	userinfo, _ := logic.GetUserProfile(email)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    userinfo,
	})
}

func HandleEditProfile(ctx *gin.Context) {

}
