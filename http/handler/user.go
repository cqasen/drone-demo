package handler

import (
	"github.com/cqasen/gin-demo/pkg/dto/request"
	"github.com/cqasen/gin-demo/pkg/service"
	"github.com/cqasen/gin-demo/pkg/utils"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/ego/http/validator"
	"github.com/ebar-go/egu"
	"github.com/gin-gonic/gin"
)

func GetInfo(ctx *gin.Context) {
	user := utils.GetLoginUserFromContext(ctx)
	response.WrapContext(ctx).Success(user)
}

func Login(ctx *gin.Context) {
	var req request.UserLogin
	if err := ctx.ShouldBindJSON(&req); err != nil {
		egu.SecurePanic(errors.New(-1, err.Error()))
	}
	vd := new(validator.Validator)
	err := vd.ValidateStruct(req)
	if err != nil {
		egu.SecurePanic(errors.New(-1, err.Error()))
	}
	res, err := service.User().Auth(req)
	if err != nil {
		egu.SecurePanic(err)
	}
	response.WrapContext(ctx).Success(res)
}
