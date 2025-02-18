package config

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/viper"
)

// KafkaConfig 定义Kafka配置结构体
type KafkaConfig struct {
	Brokers  []string `mapstructure:"brokers"`
	Topic    string   `mapstructure:"topic"`
	Security struct {
		Enabled  bool   `mapstructure:"enabled"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"security"`
}

// LoadConfig 加载配置文件
func LoadConfig() (*KafkaConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var kafkaConfig KafkaConfig
	if err := viper.UnmarshalKey("kafka", &kafkaConfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal kafka config: %w", err)
	}

	return &kafkaConfig, nil
}

// SetupKafkaConfig 设置Kafka配置
func SetupKafkaConfig(config *KafkaConfig) (*sarama.Config, error) {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForLocal // ACK=1
	kafkaConfig.Producer.Retry.Max = 5
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.Compression = sarama.CompressionSnappy

	if config.Security.Enabled {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = config.Security.Username
		kafkaConfig.Net.SASL.Password = config.Security.Password
		kafkaConfig.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	}

	return kafkaConfig, nil
}
