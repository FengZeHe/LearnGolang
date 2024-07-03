package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	clientv3 "go.etcd.io/etcd/client/v3"

	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect to etcd success")
	//defer cli.Close()

	get, err := cli.Get(context.Background(), "/config/mysql")
	if err != nil {
		return
	}
	fmt.Println("get=", get)

	v := viper.New()
	v.SetConfigType("yaml")
	v.AddRemoteProvider("etcd3", "https://127.0.0.1:2379", "/config/mysql")

	if err := v.ReadRemoteConfig(); err != nil {
		log.Fatal("ReadRemoteConfig", err)
	}
	value := v.GetStringMap("mysql")
	log.Println(value)
}
