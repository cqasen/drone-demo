package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/egu"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
	"strconv"
)

var enforcer *casbin.Enforcer

func init() {
	log.Printf("加载配置权限")

	dns := fmt.Sprintf("%s:%s@tcp(%s)/",
		"root",
		"cqasen@qq.com",
		net.JoinHostPort("111.229.103.26", strconv.Itoa(13306)))
	log.Printf(dns)
	//a, err := gormadapter.NewAdapter("mysql", dns, false)
	//secure.FatalError("加载数据库权限配置错误", err)
	//e, err := casbin.NewEnforcer("./config/rbac_model.conf", a)
	//读取csv
	//e, err := casbin.NewEnforcer("./config/rbac_model.conf", "./config/rbac_policy.csv")
	//egu.FatalError("加载权限配置", err)

	//读取数据库
	a, _ := gormadapter.NewAdapter("mysql", dns, "zblog")
	e, _ := casbin.NewEnforcer("./config/rbac_model.conf", a)
	//从DB加载策略
	err := e.LoadPolicy()
	egu.FatalError("加载权限配置", err)
	enforcer = e
}

func CheckPermission(ctx *gin.Context) {
	role := "anonymous"
	//log.Printf("配置加载")
	//enforcer, err := casbin.NewEnforcer("./config/rbac_model.conf", "./config/rbac_policy.csv")
	//if err != nil {
	//	log.Fatalf("配置加载错误:%s\n", err.Error())
	//}

	result, err := enforcer.Enforce(role, ctx.Request.URL.Path, ctx.Request.Method)

	if err != nil {
		log.Printf("权限加载错误:%s\n", err.Error())
		response.WrapContext(ctx).Error(500, err.Error())
		ctx.Abort()
	}
	if !result {
		log.Println(fmt.Sprintf("角色：%s 没有请求%s,%s的权限\n", role, ctx.Request.URL.Path, ctx.Request.Method))
		response.WrapContext(ctx).Error(401, "Unauthorized您无权查看此目录或页面")
		ctx.Abort()
	}
	ctx.Next()
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}
