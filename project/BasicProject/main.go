package main

import (
	"BasicProject/dao/mysql"
	"BasicProject/middlewares/cache"
	"BasicProject/router"
	"BasicProject/setting"
	"fmt"
	"log"
	"strconv"
)

func main() {
	if err := setting.Init(); err != nil {
		log.Println("setting Init ERROR:", err)
	}

	// 初始化MySQL
	if err := mysql.Init(setting.Conf.Mysql); err != nil {
		log.Println("init Mysql DB error")
	}

	if err := cache.RedisInit(); err != nil {
		log.Println("init Redis DB error")
	}

	// 初始化路由
	r := router.SetupRouter(setting.Conf.Mode)
	port := fmt.Sprintf(":%s", strconv.Itoa(setting.Conf.Port))
	err := r.Run(port)
	if err != nil {
		return
	}

}
