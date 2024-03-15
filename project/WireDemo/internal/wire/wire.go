//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"wiretes/internal/controller"
	"wiretes/internal/dao"
	"wiretes/internal/logic"
	"wiretes/internal/setting"
)

func InitializeApp() (*controller.UserController, error) {
	// 传入各种组件的初始化方法
	wire.Build(
		setting.InitMysqlConf,
		dao.InitDB,
		controller.NewUserController,
		logic.NewUserLogic,
		dao.NewUserDao)

	return &controller.UserController{}, nil
}
