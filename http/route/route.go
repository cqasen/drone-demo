package route

import (
	"github.com/cqasen/gin-demo/http/handler"
	middleware2 "github.com/cqasen/gin-demo/pkg/middleware"
	"github.com/cqasen/gin-demo/pkg/service/data"
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
	post := route.Group("/post")
	{
		post.GET("", handler.GetPostList)
		post.GET("/:id", handler.GetPost)
	}

	route.GET("/login", handler.Login)
	route.GET("/search", handler.SearchPost)
	route.GET("/push", handler.PushPostList)
	route.GET("/push/:id", handler.PushPost)

	post1 := route.Group("/user-info").Use(middleware.JWT(&data.UseClaims{}))
	{
		post1.GET("", handler.GetInfo)
	}
}
