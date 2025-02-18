package ioc

import (
	"github.com/IBM/sarama"
	"github.com/basicprojectv2/settings"
)

func InitSaramaClient(conf *settings.KafkaConfig) sarama.Client {
	scfg := sarama.NewConfig()
	scfg.Producer.Return.Successes = true
	client, err := sarama.NewClient(conf.Addr, scfg)

	if err != nil {
		panic(err)
	}
	return client
}

func InitSyncProducer(client sarama.Client) sarama.SyncProducer {
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}
	return producer
}

func InitConsumer(client sarama.Client) sarama.Consumer {
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
	}
	return consumer
}
