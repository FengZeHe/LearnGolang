package main

import (
	"log"
	"timerDemo/dao"
	idb "timerDemo/init"
	"timerDemo/router"
	"timerDemo/service"
)

func main() {

	db := idb.InitMysqlDB()
	repo := dao.NewTaskDAO(db)
	srv := service.NewTaskService(repo)

	r := router.RegisterTaskRouter(srv)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}

}
