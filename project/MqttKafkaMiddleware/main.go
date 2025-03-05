package main

import (
	"log"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"mqttkafkamiddleware/config"
)

var (
	producer sarama.AsyncProducer
	topic    string
)

func main() {
	// 加载配置
	kafkaConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 设置Kafka配置
	saramaConfig, err := config.SetupKafkaConfig(kafkaConfig)
	if err != nil {
		log.Fatalf("Failed to setup Kafka config: %v", err)
	}

	// 创建异步生产者
	producer, err = sarama.NewAsyncProducer(kafkaConfig.Brokers, saramaConfig)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	topic = kafkaConfig.Topic

	// 处理生产者成功和错误事件
	go handleProducerEvents()

	r := gin.Default()

	r.POST("/webhook", handleWebhook)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// handleWebhook 处理EMQ X的WebHook请求
func handleWebhook(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.JSON(405, gin.H{"error": "Invalid request method"})
		return
	}

	var message []byte
	var err error
	if message, err = c.GetRawData(); err != nil {
		c.JSON(400, gin.H{"error": "Failed to read request body"})
		return
	}

	// 发送消息到Kafka
	producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string(message)),
	}

	c.JSON(200, gin.H{"message": "Message sent to Kafka"})
}

// handleProducerEvents 处理Kafka生产者的成功和错误事件
func handleProducerEvents() {
	for {
		select {
		case success := <-producer.Successes():
			// 处理消息发送成功事件
			log.Printf("Message sent successfully to partition %d, offset %d", success.Partition, success.Offset)
		case err := <-producer.Errors():
			// 处理消息发送失败事件
			log.Printf("Failed to send message: %v", err.Err)
		}
	}
}
