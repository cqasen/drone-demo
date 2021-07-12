package handler

import (
	"github.com/ebar-go/ego/http/response"
	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	data := response.Data{
		"tips": "Hello World!!!",
	}
	response.WrapContext(ctx).Success(data)
}
