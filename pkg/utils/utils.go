package utils

import (
	"github.com/cqasen/gin-demo/pkg/dto/response"
	"github.com/cqasen/gin-demo/pkg/service/data"
	"github.com/cqasen/gin-demo/pkg/service/entity"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/egu"
	"github.com/gin-gonic/gin"
)

//GetLoginUserFromContext通过Context获取登录用户信息
func GetLoginUserFromContext(ctx *gin.Context) data.User {
	claims, _ := ctx.Get(app.Jwt().ClaimsKey)
	if claims == nil {
		egu.SecurePanic(errors.Unauthorized("please login first"))
	}
	useClaims, ok := claims.(*data.UseClaims)
	if !ok {
		egu.SecurePanic(errors.Unauthorized("please login first"))
	}
	return useClaims.User
}

//获取token
func GetAuthToken(member *entity.ZbpMember) (*response.UserAuthResponse, error) {
	useClaims := new(data.UseClaims)
	useClaims.ExpiresAt = egu.GetTimeStamp() + 3600
	useClaims.User.Id = int(member.MemID)
	useClaims.User.Name = member.MemName
	token, err := app.Jwt().GenerateToken(useClaims)
	if err != nil {
		return nil, err
	}
	res := new(response.UserAuthResponse)
	res.Token = token
	return res, nil
}
