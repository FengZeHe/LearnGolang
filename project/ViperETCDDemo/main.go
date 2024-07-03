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
	defer cli.Close()

	if err = viper.AddRemoteProvider("etcd3", "http://127.0.0.1:2379", "/config/mysql"); err != nil {
		log.Fatal("Viper Add Remote Provider ERROR", err)
	}
	viper.SetConfigType("yaml")

	if err := viper.ReadRemoteConfig(); err != nil {
		log.Fatal("ReadRemoteConfig", err)
	}
	value := viper.GetStringMap("mysql")
	log.Println(value)

	// 监听etcd变化
	go watchEtcdChanges(cli)

	// 阻塞主线程，使程序保持运行
	select {}

}

// watchEtcdChanges 监听etcd中配置变化，更新viper的配置
func watchEtcdChanges(etcdclient *clientv3.Client) {
	watcher := clientv3.NewWatcher(etcdclient)

	// 设定etcd中监听的路径
	watchChan := watcher.Watch(context.Background(), "/config/mysql")
	for resp := range watchChan {
		for _, event := range resp.Events {
			if event.Type == clientv3.EventTypePut {
				// 更新Viper中的配置
				if err := viper.ReadRemoteConfig(); err != nil {
					log.Println("Viper Add Remote Provider ERROR", err)
					continue
				}
				value := viper.GetStringMap("mysql")
				fmt.Println("监听到值变化, value = ", value)
			}
		}
	}
}
