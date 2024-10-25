// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"wiretes/internal/controller"
	"wiretes/internal/dao"
	"wiretes/internal/logic"
	"wiretes/internal/setting"
)

// Injectors from wire.go:

func InitializeApp() (*controller.UserController, error) {
	mysql := setting.InitMysqlConf()
	db := dao.InitDB(mysql)
	userDao := dao.NewUserDao(db)
	userLogic := logic.NewUserLogic(userDao)
	userController := controller.NewUserController(userLogic)
	return userController, nil
}