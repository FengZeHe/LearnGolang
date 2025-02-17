package consumer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	"mqttkafkaconsumer/config"
)

// StartConsumer 启动 Kafka 消费者
func StartConsumer(appConfig *config.AppConfig) {
	kafkaConfig, err := config.SetupKafkaConfig(&appConfig.Kafka)
	if err != nil {
		log.Fatalf("Failed to setup Kafka config: %v", err)
	}

	// 创建 Kafka 消费者组
	consumerGroup, err := sarama.NewConsumerGroup(appConfig.Kafka.Brokers, appConfig.Kafka.Consumer.GroupID, kafkaConfig)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer group: %v", err)
	}
	defer consumerGroup.Close()

	// 定义上下文
	ctx := context.Background()

	// 创建自定义消费者组程序
	handler := ConsumerHandler{
		BatchSize:   appConfig.Kafka.Consumer.BatchSize,
		MaxWaitTime: time.Duration(appConfig.Kafka.Consumer.MaxWaitTimeMs),
	}

	//开始消费
	for {
		if err := consumerGroup.Consume(ctx, appConfig.Kafka.Topic, &handler); err != nil {
			log.Fatalf("Error from consumer: %v", err)
		}
	}
}

// ConsumerHandler 实现 sarama.ConsumerGroupHandler 接口
type ConsumerHandler struct {
	//Collection  *mongo.Collection
	BatchSize   int
	MaxWaitTime time.Duration
}

func (h *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("Received message: Topic = %s, Partition = %d, Offset = %d, Key = %s, Value = %s\n",
			message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
		session.MarkMessage(message, "")
	}
	return nil
}

// Setup 在每个新会话开始时调用
func (h *ConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup 在每个会话结束时调用
func (h *ConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim 处理每个分区的消息
//func (h *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
//	messages := make([]interface{}, 0, h.BatchSize)
//	timer := time.NewTimer(h.MaxWaitTime)
//
//	for {
//		select {
//		case message, ok := <-claim.Messages():
//			if !ok {
//				return nil
//			}
//
//			messages = append(messages, bson.D{{"value", string(message.Value)}})
//
//			if len(messages) >= h.BatchSize {
//				if !timer.Stop() {
//					<-timer.C
//				}
//				h.processBatch(session, messages)
//				messages = make([]interface{}, 0, h.BatchSize)
//				timer.Reset(h.MaxWaitTime)
//			}
//		case <-timer.C:
//			if len(messages) > 0 {
//				h.processBatch(session, messages)
//				messages = make([]interface{}, 0, h.BatchSize)
//			}
//			timer.Reset(h.MaxWaitTime)
//		}
//	}
//}

//processBatch 处理批量消息
//func (h *ConsumerHandler) processBatch(session sarama.ConsumerGroupSession, messages []interface{}) {
//	_, err := h.Collection.InsertMany(context.Background(), messages)
//	if err != nil {
//		log.Printf("Error inserting messages into MongoDB: %v", err)
//	}
//
//	for _, msg := range messages {
//		fmt.Printf("Received message: %s\n", msg)
//	}
//
//	for _, msg := range messages {
//		session.MarkMessage(msg.(*sarama.ConsumerMessage), "")
//	}
//	session.Commit()
//}
