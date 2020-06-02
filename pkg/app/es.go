package app

import (
	"context"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/utils/secure"
	"github.com/olivere/elastic"
	"github.com/spf13/viper"
	"log"
)

//创建Elasticsearch链接
func InitElasticsearch() error {
	ctx := context.Background()
	url := viper.GetString("elasticsearch.url")
	username := viper.GetString("elasticsearch.username")
	password := viper.GetString("elasticsearch.password")
	sniff := viper.GetBool("elasticsearch.sniff")
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetBasicAuth(username, password), elastic.SetSniff(sniff))
	secure.FatalError("Elasticsearch Connect", err)
	info, code, err := client.Ping(url).Do(ctx)
	secure.FatalError("Elasticsearch Ping", err)
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	return app.Container.Provide(func() (*elastic.Client) {
		return client
	})
}

//获取elasticsearch链接对象
func Elasticsearch() (connection *elastic.Client) {
	_ = app.Container.Invoke(func(conn *elastic.Client) {
		connection = conn
	})
	return
}
