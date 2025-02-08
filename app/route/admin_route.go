package route

import (
	"shadowCloud/app/api/admin"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "admin pong"})
	})
	r.POST("/login", admin.AdminLogin.Login)
	r.GET("/token", admin.AdminLogin.GetUserInfoByToken)

	user := r.Group("/user")
	{
		user.GET("/profile", admin.AdminApi.Profile)
	}

}
