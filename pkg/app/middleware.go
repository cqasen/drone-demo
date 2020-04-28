package app

import "github.com/gin-gonic/gin"

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(ctx *gin.Context) bool

// EmptyMiddleware 不执行业务处理的中间件
func EmptyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

func SkipHandler(ctx *gin.Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(ctx) {
			return true
		}
	}
	return false
}
