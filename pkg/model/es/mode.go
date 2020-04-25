package es

import (
	"context"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/utils"
	"github.com/ebar-go/event"
	"github.com/olivere/elastic"
	"github.com/spf13/viper"
	"log"
)

const ElasticsearchConnectEvent = "ELASTICSEARCH_CONNECT_EVENT"

func Setup() {
	listener := event.Listener{
		Async: false,
		Handle: func(ev event.Event) {
			utils.FatalError("Elasticsearch Register", connectElasticsearch())
		},
	}
	event.DefaultDispatcher().Register(ElasticsearchConnectEvent, listener)
	_ = event.DefaultDispatcher().Trigger(ElasticsearchConnectEvent, nil)
}

func connectElasticsearch() error {
	return app.Container.Provide(func() (*elastic.Client, error) {
		ctx := context.Background()
		url := viper.GetString("ELASTICSEARCH_URL")
		username := viper.GetString("ELASTICSEARCH_USERNAME")
		password := viper.GetString("ELASTICSEARCH_PASSWORD")
		sniff := viper.GetBool("ELASTICSEARCH_SNIFF")
		client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetBasicAuth(username, password), elastic.SetSniff(sniff))
		utils.FatalError("Elasticsearch  Connect", err)
		info, code, err := client.Ping(url).Do(ctx)
		utils.FatalError("Elasticsearch  Ping", err)
		log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
		return client, nil
	})
}

func GetContection() (connection *elastic.Client) {
	_ = app.Container.Invoke(func(conn *elastic.Client) {
		connection = conn
	})
	return
}
