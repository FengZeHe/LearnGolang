package settings

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	Name         string `mapstructure:"name"`
	Version      string `mapstructure:"version"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
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

type RedisConfig struct {
	Addr string `mapstructure:"addr"`
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

func InitRedisConfig() (redisConfig *RedisConfig) {
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Read Config failed %v", err))
	}
	var appConfig AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		panic(fmt.Errorf("Unmarshal failed %v", err))
	}
	redisConfig = appConfig.RedisConfig
	log.Println("Get Redis Config success!")
	return redisConfig
}
