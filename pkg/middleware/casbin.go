package middleware

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/ebar-go/ego/http/response"
	"github.com/gin-gonic/gin"
	"log"
)

var enforcer *casbin.Enforcer

func init() {
	log.Printf("加载配置权限")
	authEnforcer, err := casbin.NewEnforcerSafe("./config/rbac_model.conf", "./config/rbac_policy.csv")
	if err != nil {
		log.Fatalf("加载配置权限错误:%s\n", err.Error())
	}
	enforcer = authEnforcer
}

func CheckPermission(ctx *gin.Context) {
	role := "superAdmin"
	//log.Printf("配置加载")
	//enforcer, err := casbin.NewEnforcerSafe("./config/rbac_model.conf", "./config/rbac_policy.csv")
	//if err != nil {
	//	log.Fatalf("配置加载错误:%s\n", err.Error())
	//}

	result, err := enforcer.EnforceSafe(role, ctx.Request.RequestURI, ctx.Request.Method)

	if err != nil {
		log.Printf("权限加载错误:%s\n", err.Error())
		response.WrapContext(ctx).Error(500, err.Error())
		ctx.Abort()
	}
	if !result {
		log.Println(fmt.Sprintf("角色：superAdmin 没有请求%s,%s的权限\n", ctx.Request.RequestURI, ctx.Request.Method))
		response.WrapContext(ctx).Error(401, "Unauthorized您无权查看此目录或页面")
		ctx.Abort()
	}
	ctx.Next()
}
