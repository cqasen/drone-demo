package router

import (
	"github.com/ebar-go/ego/http/middleware"
	"github.com/ebar-go/ego/http/response"
	"github.com/gin-gonic/gin"
)

func InitRouter(router * gin.Engine)  {
	router.Use(middleware.CORS)
	router.Use(middleware.RequestLog)
	router.Use(middleware.Recover)

	router.GET("/", func(context *gin.Context) {
		data := response.Data{
			"tips": "Hello World!",
		}
		response.WrapContext(context).Success(data)
	})
}