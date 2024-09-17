package web

import (
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/service"
	"github.com/basicprojectv2/pkg/jwt"
	"github.com/basicprojectv2/pkg/snowflake"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	svc     service.UserService
	codeSvc service.CodeService
}

const (
	//emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	//passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	bizLogin = "login"
)

func NewUserHandler(svc service.UserService, codeSvc service.CodeService) *UserHandler {
	return &UserHandler{svc: svc, codeSvc: codeSvc}
}

// 注册路由
func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/v2/users/")
	ug.GET("/hi", h.Hi)
	ug.POST("/signin", h.SignIn)
	ug.POST("/login", h.Login)
	ug.POST("/loginsms/code/send", h.SendSMS)
	ug.POST("/loginsms", h.VerifySMS)
	ug.POST("/userList", h.HandleUserList)
	ug.POST("/updateUser", h.updateUser)
}

func (h *UserHandler) updateUser(ctx *gin.Context) {
	var form domain.User
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "bad request"})
	}

	//交给service层
	if err := h.svc.UpdateUser(ctx, form); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "系统错误"})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

// 处理获取用户列表
func (h *UserHandler) HandleUserList(ctx *gin.Context) {
	// 验证请求参数是否正确
	var form domain.UserListRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "bad request"})
	}

	// 交给service层
	data, err := h.svc.GetUserList(ctx, form)
	if err != nil {
		log.Println("user service get user list error", err)
	}
	log.Println(data)
	ctx.JSON(http.StatusOK, gin.H{"data": data})

}

func (h *UserHandler) Hi(ctx *gin.Context) {
	ctx.JSON(200, "msg:hello")
}

// 处理注册请求
func (h *UserHandler) SignIn(ctx *gin.Context) {

	// 验证请求参数是否正确
	var form domain.SignInRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
	}
	// 检验两次密码输入是否正确
	if form.Password != form.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "两次密码输入不一致",
		})
	}

	err := h.svc.Signup(ctx, domain.User{Email: form.Email, Password: form.Password})
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, "msg:注册成功")
	default:
		ctx.JSON(http.StatusNotFound, "msg:注册失败")
	}

}

// 处理登录请求
func (h *UserHandler) Login(ctx *gin.Context) {
	var form domain.LoginRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, "参数错误")
	}
	u, err := h.svc.Login(ctx, form.Email, form.Password)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "msg: 登录失败")
		return
	}
	// 登录成功
	log.Println(u)
	token, err := jwt.GenToken(u.ID)
	if err != nil {
		// 生成token失败
		ctx.JSON(http.StatusInternalServerError, "系统错误")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "登录成功",
		"token": token,
	})
}

// 处理sms请求
func (h *UserHandler) SendSMS(ctx *gin.Context) {
	var form domain.SMSRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求参数错误",
		})
		return
	}

	if form.Phone == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号为空",
		})
		return
	}

	// 开始交给code svc
	if err := h.codeSvc.SendCode(ctx, bizLogin, form.Phone); err != nil {
		log.Println("发送短信失败", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})

}

// 处理验证sms登录
func (h *UserHandler) VerifySMS(ctx *gin.Context) {
	var form domain.SMSLogin
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求参数错误",
		})
		// 可以交给code svc了
	}
	ok, err := h.codeSvc.VerifyCode(ctx, bizLogin, form.Phone, form.Code)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "系统错误",
		})
	}
	if ok != true {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "验证失败",
		})
	} else {
		/*
			数据库中查询是否有该手机号，如果没有的话帮用户注册
		*/

		// 返回一个token
		id := snowflake.GenId()
		u, err := h.svc.FindOrCreate(ctx, form.Phone, strconv.Itoa(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"msg": "系统错误",
			})
			return
		}
		token, err := jwt.GenToken(u.ID)

		ctx.JSON(http.StatusOK, gin.H{
			"msg":   "登录成功",
			"token": token,
		})
	}

}
