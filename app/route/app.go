package route

import (
	"shadowCloud/app/api/app"

	"github.com/gin-gonic/gin"
)

func RegisterAppRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "app pong"})
	})
	test := r.Group("/test")
	{
		test.GET("/user", app.AppApi.GetUserInfo)
	}
}
