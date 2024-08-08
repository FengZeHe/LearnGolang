package settings

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type AppConfig struct {
	DB
}
type DB struct {
	DSN string `mapstructure:"dsn"`
}

func InitMysqlConfig() (mysqlConfig *DB) {
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read Config failed ERROR %v", err))
	}
	var appConfig AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		panic(fmt.Errorf("unmarshal Config failed ERROR %v", err))
	}
	mysqlConfig = &appConfig.DB
	return mysqlConfig

}

func InitDB(conf *DB) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(conf.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect mysql")
	}
	log.Println("connect to mysql success")
	return db, nil
}
