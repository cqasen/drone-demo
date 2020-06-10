package config

import (
	"github.com/spf13/viper"
	"log"
)

type elasticsearch struct {
	Url      string
	Username string
	Password string
	Sniff    bool
}

func InitElasticsearch() *elasticsearch {
	es := new(elasticsearch)
	err := viper.UnmarshalKey("elasticsearch", &es)
	if err != nil {
		log.Fatalf("Elasticsearch Config Error:%s", err.Error())
		return es
	}
	return es
}
