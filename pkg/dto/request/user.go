package request

type UserLogin struct {
	Name string `json:"name" binding:"required" comment:"用户名"` //验证用户名
	Pass string `json:"pass" binding:"required,min=6" comment:"密码"` //验证用户名
}
