package ioc

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
	"log"
	"prometheusdemo/internal/web/middlewares"
)

func InitDB(conf *MysqlConfig) (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect mysql")
	}
	// Prometheus Plugin
	err = db.Use(prometheus.New(prometheus.Config{
		DBName:          conf.DB,
		RefreshInterval: 10,
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				VariableNames: []string{"Threads_running"},
			},
		},
	}))
	if err != nil {
		panic(err)
	}
	// 注册gorm查询时间回调函数
	middlewares.GormQueryCallback(db)
	log.Println("connect to mysql success")
	return db
}
