package handler

import (
	"github.com/cqasen/drone-demo/pkg/model/dao"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/ego/http/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPost(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	post, err := dao.Post(app.DB()).Get(int32(id))
	if err != nil {
		panic(errors.New(404, "not found"))
	}
	response.WrapContext(ctx).Success(post)
}

func GetPostList(ctx *gin.Context) {
	post, err := dao.Post(app.DB()).GetList()
	if err != nil {
		panic(errors.New(404, "not found"))
	}
	response.WrapContext(ctx).Success(post)
}