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
	"time"
)

var Route *gin.Engine

func InitRoute(route *gin.Engine) {
	Route = route
}

func PushRoute(ctx *gin.Context) {
	routers := Route.Routes()
	db := dao.NewRoute(app.DB())
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
		if route.ID < 1 {
			route.Createtime = time.Now()
		}
		route.Updatetime = time.Now()
		db.Save(route)
	}
}

//
func SetJurisdiction(ctx *gin.Context) {
	db := dao.NewRoute(app.DB())
	e := middleware.GetEnforcer()
	var id []int
	id = append(id, 10)
	id = append(id, 3)

	res := db.GetById(id)
	p := "anonymous"

	for _, val := range res {
		log.Println(val)
		state, err := e.AddPolicy(p, val.Path, val.Method)
		egu.SecurePanic(err)
		log.Println(state)
	}
	log.Println(p)
	log.Println(res)
	egu.SecurePanic(e.SavePolicy())
	response.WrapContext(ctx).Success(res)
}
