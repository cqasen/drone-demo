package main

import (
	"github.com/cqasen/drone-demo/http/router"
	"github.com/cqasen/drone-demo/pkg/config"
	"github.com/ebar-go/ego/http"
	"github.com/ebar-go/ego/utils"
)

func main() {
	//加载配置
	config.InitConfig()
	//获取http服务对象
	server := http.NewServer()
	//链接es
	//app.SetElasticsearchContent()
	//加载路由
	router.InitRouter(server.Router)
	//启动服务
	utils.FatalError("StartServer", server.Start())
}
