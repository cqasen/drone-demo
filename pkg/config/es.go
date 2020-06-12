package config

import (
	"github.com/spf13/viper"
	"log"
)

const ELASTICSEARCH_CONFIG_KEY = "elasticsearch"

type elasticsearch struct {
	Url      string
	Username string
	Password string
	Sniff    bool
}

func InitElasticsearch() *elasticsearch {
	conf := new(elasticsearch)
	err := viper.UnmarshalKey(ELASTICSEARCH_CONFIG_KEY, &conf)
	if err != nil {
		log.Fatalf("Elasticsearch Config Error:%s", err.Error())
		return conf
	}
	return conf
}
