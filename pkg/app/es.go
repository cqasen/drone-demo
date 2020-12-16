package app

import (
	"github.com/cqasen/gin-demo/pkg/config"
	"github.com/ebar-go/ego/app"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

//创建Elasticsearch链接
func InitElasticsearch() error {
	infolog := log.New(os.Stdout, "Info ", log.LstdFlags)
	errorlog := log.New(os.Stdout, "Error ", log.LstdFlags)
	traceLog := log.New(os.Stdout, "Trace ", log.LstdFlags)

	conf := config.InitElasticsearch()
	retrier := elastic.NewBackoffRetrier(elastic.NewSimpleBackoff(0, 50, 200, 500, 1000, 2000, 4000, 8000))
	client, err := elastic.NewClient(
		elastic.SetURL(conf.Url),
		elastic.SetBasicAuth(conf.Username, conf.Password),
		elastic.SetSniff(conf.Sniff),
		elastic.SetInfoLog(infolog),
		elastic.SetErrorLog(errorlog),
		elastic.SetTraceLog(traceLog),
		elastic.SetRetrier(retrier))

	if err != nil {
		log.Println("Elasticsearch Connect Error:" + err.Error())
		return err
	}
	version, _ := client.ElasticsearchVersion(conf.Url)
	log.Printf("Elasticsearch version %s\n", version)

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
