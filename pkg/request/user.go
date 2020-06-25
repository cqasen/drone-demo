package request

type UserLogin struct {
	Name string `json:"name" binding:"required" comment:"用户名"` //验证用户名
}
