package init

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitMysql() (db *gorm.DB) {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/dbtask?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("connect to Mysql success")
	return db
}
