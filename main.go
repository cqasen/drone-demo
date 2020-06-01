package main

import (
	"github.com/cqasen/drone-demo/http/router"
	app2 "github.com/cqasen/drone-demo/pkg/app"
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
	secure.FatalError("Mysql Start",app.InitDB())
	secure.FatalError("Redis Start",app.InitRedis())
	//链接es
	app2.SetElasticsearchContent()
	//加载路由
	router.InitRouter(server.Router)
	//启动服务
	secure.FatalError("StartServer", server.Start())
}
