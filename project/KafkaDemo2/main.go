package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

func main() {
	config := sarama.NewConfig()
	config.Net.SASL.Enable = true
	config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	config.Net.SASL.User = "admin"
	config.Net.SASL.Password = "123456"
	config.Net.KeepAlive = 0
	config.Version = sarama.V2_8_0_0
	brokers := []string{"localhost:9092"}
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	fmt.Println("Successfully connected to Kafka!")
}
