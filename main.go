package main

import (
	"fmt"
	"github.com/cqasen/gin-demo/http/route"
	app2 "github.com/cqasen/gin-demo/pkg/app"
	"github.com/cqasen/gin-demo/pkg/config"
	"github.com/ebar-go/ego"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/component/event"
	"github.com/ebar-go/egu"

	"log"
)

func init() {
	event.Listen(event.BeforeHttpShutdown, func(ev event.Event) {
		log.Printf("close database")
		_ = app.Redis().Close()
	})
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	//加载配置
	env := config.GetEnv()
	log.Println("获取的环境变量：" + env)
	egu.SecurePanic(app.Config().LoadFile(fmt.Sprintf("./config/config_%s.yaml", env)))
	egu.SecurePanic(app.InitDB())
	egu.SecurePanic(app.Redis().Connect())
	//链接es
	egu.SecurePanic(app2.InitElasticsearch())

	//获取http服务对象
	server := ego.HttpServer()
	//加载路由
	route.Load(server.Router)
	//启动服务
	egu.SecurePanic(server.Start())
}
