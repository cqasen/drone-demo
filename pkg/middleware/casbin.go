package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/ego/utils/secure"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var enforcer *casbin.Enforcer

func init() {
	log.Printf("加载配置权限")
	//dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	//	"root",
	//	"cqasen@qq.com",
	//	net.JoinHostPort("", strconv.Itoa(13306)),
	//	"casbin")
	//log.Printf(dns)
	//a, err := gormadapter.NewAdapter("mysql", dns, false)
	//secure.FatalError("加载数据库权限配置错误", err)
	//authEnforcer, err := casbin.NewEnforcer("./config/rbac_model.conf", a)
	authEnforcer, err := casbin.NewEnforcer("./config/rbac_model.conf", "./config/rbac_policy.csv")
	secure.FatalError("加载权限配置", err)
	enforcer = authEnforcer
}

func CheckPermission(ctx *gin.Context) {
	role := "anonymous"
	//log.Printf("配置加载")
	//enforcer, err := casbin.NewEnforcer("./config/rbac_model.conf", "./config/rbac_policy.csv")
	//if err != nil {
	//	log.Fatalf("配置加载错误:%s\n", err.Error())
	//}

	result, err := enforcer.Enforce(role, ctx.Request.RequestURI, ctx.Request.Method)

	if err != nil {
		log.Printf("权限加载错误:%s\n", err.Error())
		response.WrapContext(ctx).Error(500, err.Error())
		ctx.Abort()
	}
	if !result {
		log.Println(fmt.Sprintf("角色：%s 没有请求%s,%s的权限\n", role, ctx.Request.RequestURI, ctx.Request.Method))
		response.WrapContext(ctx).Error(401, "Unauthorized您无权查看此目录或页面")
		ctx.Abort()
	}
	ctx.Next()
}
