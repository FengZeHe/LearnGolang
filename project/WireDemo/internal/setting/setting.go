package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	*Mysql `mapstructure:"mysql"`
}

type Mysql struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func InitConf() (conf *AppConfig) {
	viper.SetConfigFile("./internal/config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Read Config failed err %v", err))
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed ,err %v", err))
	}
	log.Println("load config file success!")
	return conf
}

func InitMysqlConf() (mysqlconf *Mysql) {
	viper.SetConfigFile("./internal/config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Read Config failed err %v", err))
	}
	var appconf AppConfig
	if err := viper.Unmarshal(&appconf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed ,err %v", err))
	}
	mysqlconf = appconf.Mysql
	log.Println("load config file success!", mysqlconf)
	return mysqlconf
}
