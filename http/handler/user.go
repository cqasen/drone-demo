package handler

import (
	"fmt"
	"github.com/cqasen/gin-demo/pkg/request"
	"github.com/cqasen/gin-demo/pkg/service"
	"github.com/cqasen/gin-demo/pkg/utils"
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/ego/utils/secure"
	"github.com/gin-gonic/gin"
)

func GetInfo(ctx *gin.Context) {
	//keyword := ctx.Query("keyword")
	user := utils.GetLoginUserFromContext(ctx)
	response.WrapContext(ctx).Success(user)
}

func Login(ctx *gin.Context) {
	var req request.UserLogin
	if err := ctx.ShouldBindJSON(&req); err != nil {
		secure.Panic(err)
	}
	fmt.Println("name:", req)
	res, err := service.User().Auth(req)
	fmt.Println(res)
	if err != nil {
		secure.Panic(err)
	}
	response.WrapContext(ctx).Success(res)
}
