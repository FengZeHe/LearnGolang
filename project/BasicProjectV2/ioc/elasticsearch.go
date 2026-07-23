package ioc

import (
	"log"

	"github.com/basicprojectv2/settings"
	"github.com/elastic/go-elasticsearch/v8"
)

func InitES(conf *settings.ElasticsearchConfig) (client *elasticsearch.Client) {
	cfg := elasticsearch.Config{
		Addresses: []string{conf.Addr},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal("create ES client ERROR", err)
	}
	return client
}
