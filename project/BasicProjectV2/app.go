package main

import (
	"github.com/basicprojectv2/internal/events/article"
	"github.com/gin-gonic/gin"
)

type App struct {
	server         *gin.Engine
	saramaConsumer article.Consumer
}
