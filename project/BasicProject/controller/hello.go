package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleHello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello, 欢迎!!")
}
