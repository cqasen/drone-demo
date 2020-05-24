package router

import (
	"github.com/cqasen/drone-demo/http/handler"
	"github.com/ebar-go/ego/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.Use(middleware.CORS)
	router.Use(middleware.RequestLog)
	router.Use(handler.Recover)
	router.StaticFile("favicon.ico", "./resources/favicon.ico")

	router.GET("/", handler.IndexHandler)
	router.GET("/post/:id", handler.GetPost)
	router.GET("/post", handler.GetPostList)
}
