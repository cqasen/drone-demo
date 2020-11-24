package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/egu"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
	"net/http"
	"strconv"
)

var enforcer *casbin.Enforcer

func init() {
	log.Printf("加载配置权限")
	dns := fmt.Sprintf("%s:%s@tcp(%s)/",
		"root",
		"cqasen@qq.com",
		net.JoinHostPort("111.229.103.26", strconv.Itoa(13306)))
	log.Println(dns)

	//读取csv
	//e, err := casbin.NewEnforcer("./config/rbac_model.conf", "./config/rbac_policy.csv")
	//egu.FatalError("加载权限配置", err)

	//读取数据库
	a, _ := gormadapter.NewAdapter("mysql", dns, "zblog")
	e, _ := casbin.NewEnforcer("./config/rbac_model.conf", a)
	//从DB加载策略
	//e.EnableLog(true)
	err := e.LoadPolicy()
	egu.FatalError("加载权限配置", err)
	enforcer = e
}

func CheckPermission(ctx *gin.Context) {
	user := ctx.Query("user")
	roles, _ := enforcer.GetRolesForUser(user)
	log.Println(fmt.Sprintf("用户:%s 角色：%s", "demo1", roles))
	role := "anonymous"
	if len(roles) == 0 {
		ctx.Abort()
		panic(errors.New(http.StatusForbidden, fmt.Sprintf("用户:%s 未设置角色", user)))
	}
	role = roles[0]
	result, err := enforcer.Enforce(role, ctx.Request.URL.Path, ctx.Request.Method)

	if err != nil {
		log.Printf("权限加载错误:%s\n", err.Error())
		ctx.Abort()
		panic(errors.New(http.StatusInternalServerError, err.Error()))
	}
	if !result {
		log.Println(fmt.Sprintf("角色：%s 没有请求%s,%s的权限\n", role, ctx.Request.URL.Path, ctx.Request.Method))
		ctx.Abort()
		panic(errors.New(http.StatusUnauthorized, "Unauthorized您无权查看此目录或页面"))
	}
	ctx.Next()
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}
