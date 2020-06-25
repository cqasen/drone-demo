package service

import (
	"fmt"
	"github.com/cqasen/gin-demo/pkg/dto/request"
	"github.com/cqasen/gin-demo/pkg/dto/response"
	"github.com/cqasen/gin-demo/pkg/service/dao"
	"github.com/cqasen/gin-demo/pkg/utils"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/ego/utils/strings"
)

type userService struct {
}

func User() *userService {
	return &userService{}
}

func (service *userService) Auth(req request.UserLogin) (*response.UserAuthResponse, error) {
	user, err := dao.User(app.DB()).Get(req.Name)
	if err != nil {
		return nil, errors.New(-1, fmt.Sprintf("获取用户信息失败：%s", err.Error()))
	}
	pass := strings.Md5(strings.Md5(req.Pass) + string(user.MemGUID))
	if pass != user.MemPassword {
		return nil, errors.New(-1, fmt.Sprintf("密码错误"))
	}
	//验证密码
	res, err := utils.GetAuthToken(user)
	if err != nil {
		return nil, errors.New(-1, fmt.Sprintf("Token生成失败：%s", err.Error()))
	}
	return res, nil
}
