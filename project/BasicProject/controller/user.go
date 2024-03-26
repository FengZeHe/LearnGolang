package controller

import (
	"BasicProject/logic"
	"BasicProject/middlewares/JWT"
	"BasicProject/middlewares/cache"
	"BasicProject/models"
	"BasicProject/pkg/localcache"
	_ "embed"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
)

var (

	//go:embed lua/sms_get_degree.lua
	luaSMS string

	//go:embed lua/sms_set_code.lua
	luaSetCode string

	//go:embed lua/sms_verify_code.lua
	luaVerifyCode string
)

type UserController struct {
	userLogic *logic.UserLogic
}

func NewUserController(userLogic *logic.UserLogic) *UserController {
	return &UserController{userLogic: userLogic}
}

func (userController *UserController) Hello(c *gin.Context) {
	c.JSON(200, "hello")
}

// 可以通过邮箱注册，需要做的步骤是首先在数据库查询是否已经有这个邮箱，有的话返回错误
func HandleUserSignIn(ctx *gin.Context) {
	// 1.获取请求参数
	var form models.RegisterFormByEmail
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 3. 交给logic层
	if err := logic.SignIn(&form); err != nil {
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
	var fo models.LoginForm

	// 2.校验数据的有效性
	if err := ctx.ShouldBindJSON(&fo); err != nil {
		zap.L().Error("Sign In with invalid params", zap.Error(err))
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "请求参数有误",
		})
		return
	}

	// 3. 交给Logic层
	result, tempuser, _ := logic.Login(&fo)

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

	// 检查一下该用户是否还有剩余发送短信的次数
	status, err := cache.CheckSMSResidualDegree(fo.Phone)
	if err != nil {
		log.Println("检查短信发送次数错误", err)
		return
	}
	if status != true {
		log.Println("发送短信验证码次数用完")
		ctx.JSON(http.StatusTooManyRequests, gin.H{
			"code": http.StatusTooManyRequests,
			"msg":  "操作过于频繁，请稍后再试",
		})
		return
	}

	// 没有问题的话发送验证码，交给Logic层处理
	code, err := logic.SMSLogin(fo.Phone)
	if err != nil {
		// 发送验证码失败
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "send code error",
		})
	}
	// 发送验证码没问题，将验证码存储到redis中，设置过期时间5分钟
	if err = cache.SetCodeForUserSMSLogin(fo.Phone, code); err != nil {
		log.Println("set code for user sms in redis error", err)
		// 如果redis写不进去，就要写进其他数据库或本地存储
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "send sms success",
	})
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		ctx.Abort()
	}

	/*
		在从redis中取出验证码，此时会有两种情况：
		1. redis中没有这个key
		2. redis中key对应的value不正确
	*/
	key, verify, err := cache.VerifyCodeForUserSMSLogin(fo.Phone, fo.Code)
	if err != nil {
		fmt.Println("系统化错误", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "系统错误",
		})
	}

	user, err := logic.GetUserProfileByPhone(fo.Phone)
	if err != nil {
		fmt.Println("查询Mysql数据库错误")
	}
	if user.Phone == "" {
		// 创建用户
		if err := logic.CreateUserByPhone(fo.Phone); err != nil {
			fmt.Println("创建用户失败 返回系统错误")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "系统错误",
			})
			ctx.Abort()
		}
	}
	// 如果验证码正确且用户手机不为空
	if verify == true && user.Phone != "" {
		strToken, _ := JWT.GenToken(user.Id)
		fmt.Println("SMS验证通过,清除redis中的sms cache")
		_ = cache.DeleteKey(key)
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"msg":   "登录成功",
			"token": strToken,
		})
	} else if verify == false {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "登录失败,请重试",
		})
	}
}

/*
使用lua脚本处理发送验证码的请求
*/
func HandlerUserSMSForLoginV2(ctx *gin.Context) {
	var fo *models.SMS
	if err := ctx.ShouldBindJSON(&fo); err != nil {
		zap.L().Error("Sign In with invalid params", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}
	// 使用lua脚本
	result, err := cache.EvalLuaScript(fo.Phone, "", luaSMS)
	if err != nil {
		log.Println("Eval Lua Script ERROR", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "系统错误",
		})
		ctx.Abort()
		return
	}
	intres := result.(int64)
	switch intres {
	case -1:
		ctx.JSON(http.StatusTooManyRequests, gin.H{
			"code": http.StatusTooManyRequests,
			"msg":  "操作太频繁，请稍后再试",
		})
		ctx.Abort()
		return
	case 1:
		// 没有问题的话发送验证码，交给Logic层处理
		code, err := logic.SMSLogin(fo.Phone)
		if err != nil {
			// 发送验证码失败
			ctx.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "send code error",
			})
			ctx.Abort()
			return
		}
		// 使用lua脚本 将验证码保存到redis中
		_, err = cache.EvalLuaScript(fo.Phone, code, luaSetCode)
		if err != nil {
			// 存储验证码失败
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "send code error",
			})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "send sms success",
		})
	}

}

func HandlerUserSMSLoginV2(ctx *gin.Context) {
	var fo *models.VerifySMSLogin
	if err := ctx.ShouldBindJSON(&fo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		ctx.Abort()
		return
	}

	result, err := cache.EvalLuaScript(fo.Phone, fo.Code, luaVerifyCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "系统错误",
		})
	}
	intres := result.(int64)
	switch intres {
	case -1:
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "验证码错误",
		})
		ctx.Abort()
		return
	case 1:
		user, err := logic.GetUserProfileByPhone(fo.Phone)
		if err != nil {
			fmt.Println("查询Mysql数据库错误")
		}
		if user.Phone == "" {
			// 创建用户
			if err := logic.CreateUserByPhone(fo.Phone); err != nil {
				fmt.Println("创建用户失败 返回系统错误")
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code": http.StatusBadRequest,
					"msg":  "系统错误",
				})
				ctx.Abort()
				return
			}
		}
		strToken, _ := JWT.GenToken(user.Id)
		ctx.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"msg":   "登录成功",
			"token": strToken,
		})
		ctx.Abort()
		return
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

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "success",
		"userprofile": userinfo,
	})
	// 设置userinfo到redis中缓存
	if err := cache.SetCacheByUserId(&userinfo, userinfo.Id); err != nil {
		log.Println("Set User Profile Cache ERROR", err)
	}

}

// 处理获取用户信息 V2
func HandleUserProfileV2(ctx *gin.Context) {
	userId, _ := ctx.Get("userid")
	userIdStr := fmt.Sprintf("%v", userId)
	if len(userIdStr) <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Get User Profile Error",
		})
	}
	/*
		1. 请求一级缓存 LocalCache；如果LocalCache中不存在数据，那么向redis请求
		2. 请求二级缓存 Redis； 如果Redis中不存在数据，那么向MySQL中发送请求
		3. Mysql中如果获取数据成功，那么将数据设置到本地缓存和redis中去
	*/
	userinfo, _ := logic.GetUserProfileById(userIdStr)

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "success",
		"userprofile": userinfo,
	})
	// 设置userinfo到本地缓存
	if err := localcache.SetLocalCacheByUserId(&userinfo, userinfo.Id); err != nil {
		log.Println("Set User Profile Local Cache ERROR", err)
	}

	// 设置userinfo到redis中缓存
	if err := cache.SetCacheByUserId(&userinfo, userinfo.Id); err != nil {
		log.Println("Set User Profile Redis Cache ERROR", err)
	}

}

// 处理用户编辑信息请求
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

// 处理用户编辑信息请求 V2
func HandleEditProfileV2(ctx *gin.Context) {
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

	//修改完成后清除LocalCache
	if err := localcache.DelLocalCacheByUserId(userStr); err != nil {
		log.Println("Del Local Cache By UserId ERROR ", err)
	}

	//删除redis缓存
	if err := cache.DelCacheByUserId(userStr); err != nil {
		log.Println("Del Redis Cache By UserId ERROR ", err)
	}
	// 返回成功
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
