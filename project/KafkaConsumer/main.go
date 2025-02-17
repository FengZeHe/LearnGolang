package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

func main() {
	// 配置 Kafka 生产者
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.NoResponse
	producerConfig.Producer.Retry.Max = 1
	producerConfig.Producer.Return.Successes = true

	// 创建 Kafka 生产者
	producer, err := sarama.NewSyncProducer([]string{
		"30.116.184.138:9092",
		"30.116.184.138:9093",
		"30.116.184.138:9094",
	}, producerConfig)
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
		Value: sarama.StringEncoder("2025-02-17-0959"),
	}
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}
	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)

	// 配置 Kafka 消费者
	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	// 创建 Kafka 消费者
	consumer, err := sarama.NewConsumer([]string{
		"30.116.184.138:9092",
		"30.116.184.138:9093",
		"30.116.184.138:9094",
	}, consumerConfig)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalf("Error closing Kafka consumer: %v", err)
		}
	}()

	// 获取主题的分区信息
	partitions, err := consumer.Partitions("test-topic")
	if err != nil {
		log.Fatalf("Error getting partitions: %v", err)
	}

	// 为每个分区创建一个消费者
	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition("test-topic", partition, sarama.OffsetOldest)
		if err != nil {
			log.Fatalf("Error consuming partition: %v", err)
		}
		defer pc.AsyncClose()

		// 异步消费消息
		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				fmt.Printf("Received message: topic=%s, partition=%d, offset=%d, value=%s\n",
					message.Topic, message.Partition, message.Offset, string(message.Value))
			}
		}(pc)
	}

	// 保持主 goroutine 运行，持续消费消息
	select {}

}
