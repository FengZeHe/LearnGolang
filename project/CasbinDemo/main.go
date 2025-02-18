package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/casbindemo/settings"
	"log"
)

func main() {
	//enforcer, err := casbin.NewEnforcer("./rbac_model.conf", "./rbac_policy.csv")
	//if err != nil {
	//	log.Fatal(err)
	//}
	// 初始化 Mysql
	mysqlConf := settings.InitMysqlConfig()
	db, err := settings.InitDB(mysqlConf)

	db.AutoMigrate()
	//
	//adapter, err := gormadapter.NewAdapterByDBUseTableName(db, "", "casbin_rules")
	//if err != nil {
	//	panic("failed to initialize adapter")
	//}

	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic("failed to initialize adapter")
	}

	//读取字符串model
	m, _ := model.NewModelFromString(textModel)

	// 初始化Enforcer
	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		panic("failed to create enforcer")
	}

	// 加载策略
	if err := enforcer.LoadPolicy(); err != nil {
		panic("failed to load policy failed")
	}

	res, err := enforcer.AddPolicy("alice", "data", "read")
	if err != nil {
		return
	}
	res2, err := enforcer.RemovePolicy("bob", "data", "read")

	fmt.Println("res->", res, res2)

	// 权限检查，创建请求
	sub := "bob"
	obj := "data"
	act := "read"
	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		log.Println("err:", err)
	}
	if ok == true {
		log.Println("true!")
	} else {
		log.Println("false!")
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

var textModel = `
[request_definition]
r = sub,obj,act


[policy_definition]
p = sub,obj,act

[role_definition]
g = _,_
g2 = _,_

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && g2(r.obj, p.obj) && r.act == p.act

`
