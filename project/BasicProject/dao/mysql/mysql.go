package mysql

import (
	"BasicProject/models"
	"BasicProject/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var user models.User

func Init(cfg *setting.Mysql) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return err
	}
	// 自动建表
	if err = db.AutoMigrate(&user); err != nil {
		return err
	}
	log.Println("Mysql init success")
	return nil
}
