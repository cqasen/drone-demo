package main

import (
	"github.com/ebar-go/ego/http"
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/ego/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	server := http.NewServer()
	server.Router.GET("/", func(context *gin.Context) {
		data := response.Data{
			"tips": "Hello World!",
		}
		response.WrapContext(context).Success(data)
	})
	utils.FatalError("StartServer", server.Start())
}
