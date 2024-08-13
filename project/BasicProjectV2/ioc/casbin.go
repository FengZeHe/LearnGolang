package ioc

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"log"
)

// 初始化enforcer
func InitMysqlCasbinEnforcer(db *gorm.DB) *casbin.Enforcer {
	ad, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic("failed to init adapter")
	}

	enforcer, err := casbin.NewEnforcer("./config/rbac_model.conf", ad)
	if err != nil {
		log.Println("failed to init enforcer", err)
		panic("failed to init enforcer")
	}
	//加载策略
	if err := enforcer.LoadPolicy(); err != nil {
		panic("failed to load policy")
	}
	return enforcer

}
