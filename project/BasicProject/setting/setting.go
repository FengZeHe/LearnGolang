package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Mode    string `mapstructure:"mode"`
	Port    int    `mapstructure:"port"`
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	*Mysql  `mapstructure:"mysql"`
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

func Init() error {
	// 设置viper读取的配置文件路径
	viper.SetConfigFile("./config/dev.config.yaml")
	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("配置文件被修改了")
		if err := viper.Unmarshal(&Conf); err != nil {
			return
		}
	})

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Read Config failed err %v", err))
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed ,err %v", err))
	}
	log.Println("load config file success!")

	return nil
}
