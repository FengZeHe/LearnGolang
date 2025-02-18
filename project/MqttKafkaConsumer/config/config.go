package config

import (
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
)

// KafkaConfig 定义 Kafka 配置结构体
type KafkaConfig struct {
	Brokers  []string `mapstructure:"brokers"`
	Topic    []string `mapstructure:"topic"`
	Security struct {
		Enabled  bool   `mapstructure:"enabled"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"security"`
	Consumer struct {
		MinSize     int32 `mapstructure:"min_size"`
		MaxSize     int32 `mapstructure:"max_size"`
		MaxWaitTime int   `mapstructure:"max_wait_time"`
	} `mapstructure:"consumer"`
}

// MongoDBConfig 定义 MongoDB 配置结构体
type MongoDBConfig struct {
	URI            string `mapstructure:"uri"`
	DBName         string `mapstructure:"db_name"`
	CollectionName string `mapstructure:"collection_name"`
}

// AppConfig 定义应用配置结构体
type AppConfig struct {
	Kafka   KafkaConfig   `mapstructure:"kafka"`
	MongoDB MongoDBConfig `mapstructure:"mongodb"`
}

// LoadConfig 加载配置文件
func LoadConfig() (*AppConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var appConfig AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &appConfig, nil
}

// SetupKafkaConfig 设置 Kafka 配置
func SetupKafkaConfig(config *KafkaConfig) (*sarama.Config, error) {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	kafkaConfig.Consumer.Fetch.Min = config.Consumer.MinSize
	kafkaConfig.Consumer.Fetch.Max = config.Consumer.MaxSize
	kafkaConfig.Consumer.MaxWaitTime = time.Duration(config.Consumer.MaxWaitTime) * time.Millisecond

	kafkaConfig.Consumer.Return.Errors = true
	if config.Security.Enabled {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = config.Security.Username
		kafkaConfig.Net.SASL.Password = config.Security.Password
		kafkaConfig.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	}
	return kafkaConfig, nil
}
