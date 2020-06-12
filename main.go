package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/cqasen/gin-demo/http/router"
	"github.com/cqasen/gin-demo/pkg/config"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/http"
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
	//加载权限
	mycasbin()

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

func mycasbin()  {
	log.Println("加载权限配置BEGIN")
	//e, _:= casbin.NewEnforcer("path/to/model.conf", "path/to/policy.csv")
	e := casbin.NewEnforcer("./config/rbac_model.conf", "./config/rbac_policy.csv")
	fmt.Println(e)
	log.Println("加载权限配置END")
}