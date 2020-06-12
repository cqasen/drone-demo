package main

import (
	"github.com/cqasen/gin-demo/http/router"
	"github.com/cqasen/gin-demo/pkg/config"
	"github.com/ebar-go/ego"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/utils/secure"
	"log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	//加载配置
	config.InitConfig()
	//获取http服务对象
	server := ego.HttpServer()
	secure.FatalError("Mysql Start", app.InitDB())
	secure.FatalError("Redis Start", app.InitRedis())
	//链接es
	//secure.FatalError("Elasticsearch Start", app2.InitElasticsearch())
	//加载路由
	router.InitRouter(server.Router)
	//启动服务
	secure.FatalError("StartServer", server.Start())
}
