package config

import (
	"github.com/ebar-go/ego/app"
	"log"
)

const ElasticsearchConfigKey = "elasticsearch"

type elasticsearch struct {
	Url      string
	Username string
	Password string
	Sniff    bool
}

func InitElasticsearch() *elasticsearch {
	conf := new(elasticsearch)
	err := app.Config().Viper.UnmarshalKey(ElasticsearchConfigKey, &conf)
	if err != nil {
		log.Fatalf("Elasticsearch Config Error:%s", err.Error())
		return conf
	}
	return conf
}
