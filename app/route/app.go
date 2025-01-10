package route

import (
	"shadowCloud/internal/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterAppRouter(r *gin.RouterGroup) {
	r.GET("/api", func(c *gin.Context) {
		global.Logger.Error("test", zap.String("test", "test"))
		c.JSON(200, gin.H{"message": "api"})
	})
}
