package main

import (
	"context"
	"log"
	"timerDemo/dao"
	idb "timerDemo/init"
	"timerDemo/job/fun"
	"timerDemo/router"
	"timerDemo/scheduler"
	"timerDemo/service"
	"timerDemo/tasks"
)

func main() {
	tr := tasks.NewTaskRegistry()
	tr.Register("sayhi", fun.ExecTimeKeeping)

	db := idb.InitMysqlDB()
	repo := dao.NewTaskDAO(db)
	srv := service.NewTaskService(repo)

	// 初始化cron调度器
	sch := scheduler.NewCronScheduler(repo, tr)

	// 启动Cron调度器
	ctx := context.Background()
	if err := sch.Start(ctx); err != nil {
		log.Println("启动cron失败", err)
	}

	r := router.RegisterTaskRouter(srv)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}

}
