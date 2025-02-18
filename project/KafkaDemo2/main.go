package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
)

func main() {
	//config := sarama.NewConfig()
	//config.Net.SASL.Enable = true
	//config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	//config.Net.SASL.User = "user"
	//config.Net.SASL.Password = "123456"
	//
	//config.Net.KeepAlive = 0
	//config.Version = sarama.V2_8_0_0
	//brokers := []string{"localhost:9092"}
	//client, err := sarama.NewClient(brokers, config)
	//if err != nil {
	//	log.Fatalf("Failed to create client: %v", err)
	//}
	//defer client.Close()
	//
	//fmt.Println("Successfully connected to Kafka!")

	brokers := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Net.TLS.Enable = true

	// 加载 truststore
	caCert, err := os.ReadFile("./truststore/kafka.truststore.pem")
	if err != nil {
		log.Fatalf("Failed to read truststore: %v", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	// 禁用证书验证
	config.Net.TLS.Config = &tls.Config{
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}

	// 加载 keystore
	cert, err := tls.LoadX509KeyPair("./keystore/kafka.keystore.pem", "./keystore/kafka.keystore.pem")
	if err != nil {
		log.Fatalf("Failed to load keystore: %v", err)
	}
	config.Net.TLS.Config.Certificates = []tls.Certificate{cert}

	// 配置SASL认证
	config.Net.SASL.Enable = true
	config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	config.Net.SASL.User = "user"
	config.Net.SASL.Password = "123456"

	// 创建 Kafka 客户端
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka client: %v", err)
	}
	defer client.Close()
	fmt.Printf("Connected to Kafka brokers : %v successfully", brokers)
}
