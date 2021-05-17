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
	//route.Use(middleware2.CheckPermission)
	//资源
	route.StaticFile("favicon.ico", "./resources/favicon.ico")
	//路由
	route.GET("/", handler.IndexHandler)
	post := route.Group("/post")
	{
		post.GET("", handler.GetPostList)
		post.GET("/:id", handler.GetPost)
	}

	//诗歌
	poetry := route.Group("/poetry")
	{
		poetry.GET("/push", handler.PushPoetry)
		poetry.GET("/search", handler.SearchPoetry)
	}

	route.GET("/login", handler.Login)
	route.GET("/search", handler.SearchPost)
	route.GET("/push", handler.PushPostList)
	route.GET("/push/:id", handler.PushPost)

	user := route.Group("/user").Use(middleware.JWT(&data.UseClaims{}))
	{
		user.GET("/info", handler.GetInfo)
	}
	route.GET("/route/push", handler.PushRoute)
	route.GET("/quanxian", handler.SetJurisdiction)

	role := route.Group("/role", middleware2.CheckPermission)
	{
		role.GET("/add", handler.AddRoleForUser)
		role.GET("/del", handler.DeleteRoleForUser)
		role.GET("/list", handler.ListRole)
		role.GET("/get", handler.GetRole)
	}
	handler.InitRoute(route)
}
