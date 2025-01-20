package ioc

import (
	"github.com/IBM/sarama"
	"github.com/basicprojectv2/settings"
)

func InitSaramaClient(conf *settings.SaramaConfig) sarama.Client {
	scfg := sarama.NewConfig()
	scfg.Producer.Return.Successes = true
	client, err := sarama.NewClient(conf.Addr, scfg)

	if err != nil {
		panic(err)
	}
	return client
}
