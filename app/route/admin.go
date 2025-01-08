package route

import (
	"github.com/gin-gonic/gin"
)

func RegisterAdminRouter(r *gin.RouterGroup) {
	r.GET("/admin", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "admin"})
	})
}
