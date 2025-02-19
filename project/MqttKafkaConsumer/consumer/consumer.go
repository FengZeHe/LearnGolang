package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mqttkafkaconsumer/config"
)

// StartConsumer 启动 Kafka 消费者
func StartConsumer(appConfig *config.AppConfig) {
	kafkaConfig, err := config.SetupKafkaConfig(&appConfig.Kafka)
	if err != nil {
		log.Fatalf("Failed to setup Kafka config: %v", err)
	}

	consumer, err := sarama.NewConsumer(appConfig.Kafka.Brokers, kafkaConfig)
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}
	defer consumer.Close()

	// 创建分区消费者
	partitionConsumer, err := consumer.ConsumePartition("read-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to create partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	// 连接 MongoDB
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(appConfig.MongoDB.URI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database(appConfig.MongoDB.DBName).Collection(appConfig.MongoDB.CollectionName)

	// 消费消息
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Received message: %s\n", string(msg.Value))

		var data struct {
			ArticleID string `json:"articleID"`
		}
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			log.Printf("Failed to parse message: %v", err)
			continue
		}

		// 定义 filter 和 update
		filter := bson.M{"articleID": data.ArticleID}
		update := bson.M{"$inc": bson.M{"read": 1}}

		// 更新 MongoDB 文档
		_, err = collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
		if err != nil {
			log.Printf("Failed to update MongoDB document: %v", err)
		}

	}

}
