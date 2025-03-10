package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"prometheusdemo/internal/service"
	"time"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine, userCache gin.HandlerFunc) {
	ug := server.Group("/user")
	ug.GET("hi", h.Hi)
	ug.GET("/getUser", userCache, h.UserHandler)
}

func (h *UserHandler) Hi(c *gin.Context) {
	rand.NewSource(time.Now().UnixNano())
	// 生成一个 0 到 199 之间的随机整数，并将其转换为 time.Duration 类型
	randomDuration := time.Duration(rand.Intn(500)) * time.Millisecond
	fmt.Printf("即将休眠 %s\n", randomDuration)
	// 休眠指定的时间
	time.Sleep(randomDuration)
	c.JSON(200, gin.H{
		"message": "hi",
	})
}

func (h *UserHandler) UserHandler(c *gin.Context) {
	users, err := h.svc.GetUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"data": users,
	})
}
