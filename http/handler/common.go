package handler

import (
	"github.com/ebar-go/ego/component/trace"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/ego/utils"
	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	data := response.Data{
		"tips": "Hello World!",
	}
	response.WrapContext(ctx).Success(data)
}

func Recover(ctx *gin.Context) {
	defer func(traceId string) {
		if err := recover(); err != nil {
			trace.SetTraceId(traceId)
			defer trace.DeleteTraceId()
			utils.Debug(utils.Trace())
			wrapper := response.WrapContext(ctx)
			if err, ok := err.(*errors.Error); ok {
				wrapper.Error(err.Code, err.Message)
			} else {
				wrapper.Error(500, "system error")
			}
		}
	}(trace.GetTraceId())
	ctx.Next()
}
