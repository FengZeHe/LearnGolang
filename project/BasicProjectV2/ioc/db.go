package ioc

import (
	"fmt"
	"github.com/basicprojectv2/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB(conf *settings.MysqlConfig) (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect mysql")
	}
	log.Println("connect to mysql success")
	return db
}
