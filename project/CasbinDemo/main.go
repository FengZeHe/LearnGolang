package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/casbindemo/settings"
	"log"
)

func main() {
	//enforcer, err := casbin.NewEnforcer("./rbac_model.conf", "./rbac_policy.csv")
	//if err != nil {
	//	log.Fatal(err)
	//}
	mysqlConf := settings.InitMysqlConfig()

	db, err := settings.InitDB(mysqlConf)
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic("failed to initialize adapter")
	}

	// 初始化Enforcer
	enforcer, err := casbin.NewEnforcer("./rbac_model.conf", adapter)
	if err != nil {
		panic("failed to create enforcer")
	}

	// 加载策略
	if err := enforcer.LoadPolicy(); err != nil {
		panic("failed to load policy failed")
	}

	err = enforcer.SavePolicy()
	if err != nil {
		panic(err)
	}

	res, err := enforcer.RemovePolicy("alice", "data1", "read")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	// 权限检查，创建请求
	sub := "alice"
	obj := "data1"
	act := "read"
	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		log.Println("err:", err)
	}
	if ok == true {
		log.Println("true")
	} else {
		log.Println("false")
	}

}

// Initialize the model from a string.
var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`
