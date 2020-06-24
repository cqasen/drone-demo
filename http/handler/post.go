package handler

import (
	"github.com/cqasen/gin-demo/pkg/model/dao"
	"github.com/cqasen/gin-demo/pkg/model/data"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/ego/utils/date"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPost(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	post, err := dao.Post(app.DB()).Get(int32(id))
	if err != nil {
		panic(errors.NotFound("404 Not Found"))
	}
	response.WrapContext(ctx).Success(post)
}

func GetPostList(ctx *gin.Context) {
	post, err := dao.Post(app.DB()).GetList()
	if err != nil {
		panic(errors.NotFound("404 Not Found"))
	}
	response.WrapContext(ctx).Success(post)
}

func SeaarchPost(ctx *gin.Context) {
	//keyword := ctx.Query("keyword")
	claims, _ := ctx.Get(app.Jwt().ClaimsKey)
	userClaims, _ := claims.(*data.UseClaims)
	response.WrapContext(ctx).Success(userClaims.User)
}

func Login(ctx *gin.Context) {
	useClaims := new(data.UseClaims)
	useClaims.ExpiresAt = date.GetTimeStamp() + 3600
	useClaims.Subject = "测试"
	useClaims.User.Id = 1
	useClaims.User.Name = "张三"
	token, err := app.Jwt().GenerateToken(useClaims)
	if err != nil {
		response.WrapContext(ctx).Error(-1, err.Error())
	}
	response.WrapContext(ctx).Success(token)
}
