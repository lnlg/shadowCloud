package route

import (
	"github.com/gin-gonic/gin"
)

func RegisterAppRouter(r *gin.RouterGroup) {
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "api"})
	})
}
