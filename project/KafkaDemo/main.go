package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// 配置 Kafka 生产者
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// 创建 Kafka 生产者
	producer, err := sarama.NewSyncProducer([]string{"30.116.184.138:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error closing Kafka producer: %v", err)
		}
	}()

	// 发送消息
	message := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}
	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)

	// 配置 Kafka 消费者
	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.Return.Errors = true

	// 创建 Kafka 消费者组
	consumer, err := sarama.NewConsumerGroup([]string{"30.116.184.138:9092"}, "test-group", consumerConfig)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer group: %v", err)
	}
	// 定义消费者处理函数
	handler := ConsumerHandler{}
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Kafka consumer error: %v", err)
		}
	}()

	// 开始消费消息
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			if err := consumer.Consume(context.Background(), []string{"test-topic"}, &handler); err != nil {
				log.Fatalf("Error from consumer: %v", err)
			}
		}
	}()

	// 等待信号中断
	<-signals
	// 确保在接收到信号后关闭消费者
	if err := consumer.Close(); err != nil {
		log.Fatalf("Error closing Kafka consumer group: %v", err)
	}
}

// ConsumerHandler 实现 sarama.ConsumerGroupHandler 接口
type ConsumerHandler struct{}

func (h *ConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message received at %v: value = %s, timestamp = %v, topic = %s, partition = %d, offset = %d\n",
			time.Now(), string(message.Value), message.Timestamp, message.Topic, message.Partition, message.Offset)
		session.MarkMessage(message, "")
	}
	return nil
}
