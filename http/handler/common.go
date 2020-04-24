package handler

import (
	"github.com/ebar-go/ego/http/response"
	"github.com/gin-gonic/gin"
)

func hello(context *gin.Context) {
	data := response.Data{
		"tips": "Hello World!",
	}
	response.WrapContext(context).Success(data)
}
