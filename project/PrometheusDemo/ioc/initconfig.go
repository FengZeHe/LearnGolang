package ioc

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	MysqlConfig *MysqlConfig `mapstructure:"mysql"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func InitMysqlConfig() (mysqlConfig *MysqlConfig) {
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Read Config failed %v", err))
	}
	var appConfig AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		panic(fmt.Errorf("Unmarshal failed %v", err))
	}
	mysqlConfig = appConfig.MysqlConfig
	log.Println("Get Mysql Config success!")
	return mysqlConfig
}
