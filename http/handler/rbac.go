package handler

import (
	"github.com/cqasen/gin-demo/pkg/middleware"
	"github.com/cqasen/gin-demo/pkg/service/dao"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/egu"
	"github.com/gin-gonic/gin"
	"log"
	"regexp"
	"strconv"
	"time"
)

var Route *gin.Engine

func InitRoute(route *gin.Engine) {
	Route = route
}

func PushRoute(ctx *gin.Context) {
	routers := Route.Routes()
	db := dao.NewRoute(app.DB())
	db.SetDelFlag()
	for _, val := range routers {
		r, _ := regexp.Compile(":[a-z]+")
		path := val.Path
		v := r.FindString(val.Path)
		if v != "" {
			path = r.ReplaceAllString(val.Path, "*")
		}
		var route = db.GetOne(path, val.Method)
		route.Method = val.Method
		route.Path = path
		route.Del = 0
		if route.ID < 1 {
			route.Createtime = time.Now()
		}
		route.Updatetime = time.Now()
		db.Save(route)
	}
}

//授权
func SetJurisdiction(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	db := dao.NewRoute(app.DB())
	e := middleware.GetEnforcer()
	var ids []int
	ids = append(ids, id)
	res := db.GetById(ids)
	p := "anonymous"
	var rules [][]string
	for _, val := range res {
		rules = append(rules, []string{p, val.Path, val.Method})
	}
	if len(rules) > 0 {
		_, err := e.AddPolicies(rules)
		egu.SecurePanic(err)
	}
	egu.SecurePanic(e.SavePolicy())
	response.WrapContext(ctx).Success(res)
}

func AddRoleForUser(ctx *gin.Context) {
	e := middleware.GetEnforcer()
	state, _ := e.AddRoleForUser("demo1", "anonymous")
	log.Println(state)
}

func DeleteRoleForUser(ctx *gin.Context) {
	e := middleware.GetEnforcer()
	state, _ := e.DeleteRoleForUser("demo1", "anonymous")
	log.Println(state)
}

func ListRole(ctx *gin.Context) {
	e := middleware.GetEnforcer()
	roles := e.GetAllRoles()
	response.WrapContext(ctx).Success(roles)
}

func GetRole(ctx *gin.Context) {
	e := middleware.GetEnforcer()
	role, _ := e.GetRolesForUser("demo1")
	response.WrapContext(ctx).Success(role)
}
