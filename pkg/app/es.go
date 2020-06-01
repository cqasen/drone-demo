package app

import (
	"context"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/utils/secure"
	"github.com/ebar-go/event"
	"github.com/olivere/elastic"
	"github.com/spf13/viper"
	"log"
)

const ELASTICSEARCH_CONNECT_EVENT = "ELASTICSEARCH_CONNECT_EVENT"

//创建Elasticsearch链接
func SetElasticsearchContent() {
	listener := event.Listener{
		Async: false,
		Handle: func(ev event.Event) {
			secure.FatalError("Elasticsearch Register", connectElasticsearch())
		},
	}
	event.DefaultDispatcher().Register(ELASTICSEARCH_CONNECT_EVENT, listener)
	_ = event.DefaultDispatcher().Trigger(ELASTICSEARCH_CONNECT_EVENT, nil)
}

//Elasticsearch链接
func connectElasticsearch() error {
	return app.Container.Provide(func() (*elastic.Client, error) {
		ctx := context.Background()
		url := viper.GetString("ELASTICSEARCH_URL")
		username := viper.GetString("ELASTICSEARCH_USERNAME")
		password := viper.GetString("ELASTICSEARCH_PASSWORD")
		sniff := viper.GetBool("ELASTICSEARCH_SNIFF")
		client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetBasicAuth(username, password), elastic.SetSniff(sniff))
		secure.FatalError("Elasticsearch  Connect", err)
		info, code, err := client.Ping(url).Do(ctx)
		secure.FatalError("Elasticsearch  Ping", err)
		log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
		return client, nil
	})
}

//获取elasticsearch链接对象
func Elasticsearch() (connection *elastic.Client) {
	_ = app.Container.Invoke(func(conn *elastic.Client) {
		connection = conn
	})
	return
}
