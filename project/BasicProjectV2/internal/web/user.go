package web

import (
	"database/sql"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/service"
	"github.com/basicprojectv2/pkg/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// 注册路由
func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/v2/users/")
	ug.GET("/hi", h.Hi)
	ug.POST("/signin", h.SignIn)
	ug.POST("/loginin", h.Login)
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
	email := sql.NullString{String: form.Email, Valid: true}
	err := h.svc.Signup(ctx, domain.User{Email: email, Password: form.Password})
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
