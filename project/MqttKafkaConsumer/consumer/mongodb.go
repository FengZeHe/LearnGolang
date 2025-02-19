package consumer

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mqttkafkaconsumer/config"
)

// UpdateMany 执行 MongoDB 的 updateMany 操作
func UpdateMany(appConfig *config.AppConfig, filter bson.M, update bson.M) error {
	// 创建 MongoDB 客户端
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(appConfig.MongoDB.URI))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	defer client.Disconnect(context.Background())

	// 获取数据库和集合
	database := client.Database(appConfig.MongoDB.DBName)
	collection := database.Collection(appConfig.MongoDB.CollectionName)

	// 执行 updateMany 操作
	result, err := collection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update documents: %w", err)
	}

	log.Printf("Updated %d documents", result.ModifiedCount)
	return nil
}
