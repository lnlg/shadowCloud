package route

import (
	"shadowCloud/app/api/admin"
	"shadowCloud/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "admin pong"})
	})
	r.POST("/login", admin.AdminLogin.Login)
	// ------ 以上的接口不走权限校验 ------
	// 权限验证中间件
	r.Use(middleware.PermissionMiddleware())

	r.GET("/token", admin.AdminLogin.GetUserInfoByToken)
	user := r.Group("/user")
	{
		user.GET("/profile", admin.AdminApi.Profile)
	}

}
