package main

import (
	"github.com/cqasen/drone-demo/http/router"
	"github.com/cqasen/drone-demo/pkg/config"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/http"
	"github.com/ebar-go/ego/utils/secure"
)

func main() {
	//加载配置
	config.InitConfig()
	//获取http服务对象
	server := http.NewServer()
	secure.FatalError("Mysql Start", app.InitDB())
	secure.FatalError("Redis Start", app.InitRedis())
	//链接es
	//secure.FatalError("Elasticsearch Start", app2.InitElasticsearch())
	//加载路由
	router.InitRouter(server.Router)
	//启动服务
	secure.FatalError("StartServer", server.Start())
}
