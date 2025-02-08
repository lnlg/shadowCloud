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
	r.Use(middleware.PermissionMiddleware()) // 权限验证中间件

	r.GET("/token", admin.AdminLogin.GetUserInfoByToken)
	user := r.Group("/user")
	{
		user.POST("/create_user", admin.AdminUsers.CreateUsers) //创建用户信息
		//user.GET("/profile", admin.AdminApi.Profile)
	}

}
