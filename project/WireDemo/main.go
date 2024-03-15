package main

import (
	"github.com/gin-gonic/gin"
	"wiretes/internal/wire"
)

func main() {
	//conf := setting.InitConf()
	//db := dao.InitDB(conf.Mysql)
	//userDao := dao.NewUserDao(db)
	//userLogic := logic.NewUserLogic(userDao)
	//userController := controller.NewUserController(userLogic)

	userController, err := wire.InitializeApp()
	if err != nil {
		panic("wire Initialize App error")
	}

	router := gin.Default()
	router.POST("/user", userController.CreateUser)

	err = router.Run(":8555")
	if err != nil {
		return
	}
}
