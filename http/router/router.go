package router

import (
	"github.com/cqasen/gin-demo/http/handler"
	middleware2 "github.com/cqasen/gin-demo/pkg/middleware"
	"github.com/ebar-go/ego/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	//中间件
	router.Use(middleware.CORS)
	router.Use(middleware.RequestLog)
	router.Use(middleware.Recover)
	router.Use(middleware.Favicon)
	router.Use(middleware2.CheckPermission)
	//资源
	router.StaticFile("favicon.ico", "./resources/favicon.ico")
	//路由
	router.GET("/", handler.IndexHandler)
	//router.GET("/post/:id", handler.GetPost)
	//router.GET("/post", handler.GetPostList)
	post := router.Group("/post")
	{
		post.GET("/:id", handler.GetPost)
		post.GET("", handler.GetPostList)
	}
}
