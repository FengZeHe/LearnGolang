package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"time"
)

func main() {
	config := sarama.NewConfig()
	// 设置 ACK 级别为 1 或 -1
	config.Producer.RequiredAcks = sarama.WaitForLocal // sarama.WaitForAll
	config.Producer.Timeout = 5 * time.Second          // 设置 5s 超时时间

	// 开启成功/错误消息返回
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating the producer: %s\n", err)
		return
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error closing the producer: %s\n", err)
		}
	}()

	message := &sarama.ProducerMessage{
		Topic: "test_topic",
		Value: sarama.StringEncoder("Hello,Kafka!"),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Error sending message: %s\n", err)
	}
	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)

}
