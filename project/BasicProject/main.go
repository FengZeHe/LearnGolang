package main

import (
	"BasicProject/controller"
	"BasicProject/dao/mysql"
	"BasicProject/logic"
	"BasicProject/router"
	"BasicProject/setting"
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	env := flag.String("env", "", "set env")
	flag.Parse()

	if err := setting.Init(*env); err != nil {
		log.Println("setting Init ERROR:", err)
	}

	conf := setting.InitMysqlConf()
	db := mysql.InitMysqlDB(conf)
	userDao := mysql.NewUserDao(db)
	userLogic := logic.NewUserLogic(userDao)
	userController := controller.NewUserController(userLogic)

	// 初始化MySQL
	//if err := mysql.Init(setting.Conf.Mysql); err != nil {
	//	log.Println("init Mysql DB error", err)
	//}

	//if err := cache.RedisInit(setting.Conf.Redis); err != nil {
	//	log.Println("init Redis DB error", err)
	//}

	// 初始化路由
	r := router.SetupRouter(setting.Conf.Mode)
	r.GET("/hello", userController.Hello)
	port := fmt.Sprintf(":%s", strconv.Itoa(setting.Conf.Port))
	err := r.Run(port)
	if err != nil {
		return
	}

}
