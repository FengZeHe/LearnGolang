package main

import (
	"log"

	"mqttkafkaconsumer/config"
	"mqttkafkaconsumer/consumer"
)

func main() {
	// 加载配置
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 启动 Kafka 消费者
	consumer.StartConsumer(appConfig)
}
