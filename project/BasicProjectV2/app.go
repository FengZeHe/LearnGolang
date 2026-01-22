package main

import (
	"github.com/basicprojectv2/jobs"
	"github.com/basicprojectv2/jobs/scheduler"
	"github.com/gin-gonic/gin"
)

type App struct {
	server *gin.Engine
	//saramaConsumer article.Consumer
	Registry  *jobs.TaskRegistry
	Scheduler *scheduler.CronScheduler
}
