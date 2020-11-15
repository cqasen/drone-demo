package handler

import (
	"github.com/cqasen/gin-demo/pkg/service/dao"
	"github.com/ebar-go/ego/app"
	"github.com/gin-gonic/gin"
	"log"
	"regexp"
	"time"
)

var Route *gin.Engine

func InitRoute(route *gin.Engine) {
	Route = route
}

func Push(ctx *gin.Context) {
	routers := Route.Routes()
	db := dao.NewRouter(app.DB())
	for _, val := range routers {
		r, _ := regexp.Compile(":[a-z]+")
		path := val.Path
		v := r.FindString(val.Path)
		if v != "" {
			path = r.ReplaceAllString(val.Path, "*")
		}
		log.Println(path, val.Method)
		var route = db.GetOne(path, val.Method)
		route.Method = val.Method
		route.Path = path
		if route.ID < 1 {
			route.Createtime = time.Now()
		}
		route.Updatetime = time.Now()
		log.Println(route)
		db.Save(route)
	}
}
