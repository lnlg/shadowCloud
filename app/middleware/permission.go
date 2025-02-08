package middleware

import (
	"shadowCloud/app/response"
	"shadowCloud/app/service"

	"github.com/gin-gonic/gin"
)

// 权限中间件
func PermissionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 判断Authorization是否存在
		if ctx.GetHeader("Authorization") == "" {
			// 未登录，返回401
			response.ReturnError(ctx, 401, "Unauthorized")
			ctx.Abort() //终止后接口调用，不加的话recover到异常后，还会继续执行接口后面的代码
			return
		}
		// 验证token是否有效
		token := ctx.GetHeader("Authorization")
		if !isPermissionToken(token) {
			// 无效token，返回401
			response.ReturnError(ctx, 401, "token invalid")
			ctx.Abort() //终止后接口调用，不加的话recover到异常后，还会继续执行接口后面的代码
			return
		}
		// 验证通过，继续执行接口
		ctx.Next()
	}
	// 获取路由
	//path := c.FullPath()

}

// 验证token是否有效
func isPermissionToken(token string) bool {
	isTrue, _ := service.AdminUsersService.GetUserInfoByToken(token)
	return isTrue
}
