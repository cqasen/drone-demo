package app

import (
	"context"
	"github.com/cqasen/gin-demo/pkg/config"
	"github.com/ebar-go/ego/app"
	"github.com/olivere/elastic"
	"log"
)

//创建Elasticsearch链接
func InitElasticsearch() error {
	conf := config.InitElasticsearch()
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(conf.Url), elastic.SetBasicAuth(conf.Username, conf.Password), elastic.SetSniff(conf.Sniff))
	if err != nil {
		log.Println("Elasticsearch Connect Error:" + err.Error())
		return err
	}
	info, code, err := client.Ping(conf.Url).Timeout("10ms").Do(ctx)
	if err != nil {
		log.Println("Elasticsearch Ping Error:" + err.Error())
		return err
	}
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	return app.Container().Provide(func() *elastic.Client {
		return client
	})
}

//获取elasticsearch链接对象
func Elasticsearch() (connection *elastic.Client) {
	_ = app.Container().Invoke(func(conn *elastic.Client) {
		connection = conn
	})
	return
}
