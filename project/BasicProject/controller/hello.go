package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello!",
	})
}
