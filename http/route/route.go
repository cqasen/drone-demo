package route

import (
	"github.com/cqasen/gin-demo/http/handler"
	middleware2 "github.com/cqasen/gin-demo/pkg/middleware"
	"github.com/ebar-go/ego/http/middleware"
	"github.com/gin-gonic/gin"
)

func Load(route *gin.Engine) {
	//中间件
	route.Use(middleware.CORS)
	route.Use(middleware.RequestLog)
	route.Use(middleware.Recover)
	route.Use(middleware.Favicon)
	route.Use(middleware2.CheckPermission)
	//资源
	route.StaticFile("favicon.ico", "./resources/favicon.ico")
	//路由
	route.GET("/", handler.IndexHandler)
	//route.GET("/post/:id", handler.GetPost)
	//route.GET("/post", handler.GetPostList)
	post := route.Group("/post")
	{
		post.GET("/:id", handler.GetPost)
		post.GET("", handler.GetPostList)
	}
}
