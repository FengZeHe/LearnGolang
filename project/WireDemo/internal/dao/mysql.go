package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"wiretes/internal/setting"
)

func InitDB(conf *setting.Mysql) (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Mysql init success")
	return db
}
