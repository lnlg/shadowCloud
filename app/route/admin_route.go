package route

import (
	"github.com/gin-gonic/gin"
)

func RegisterAdminRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "admin pong"})
	})
}
