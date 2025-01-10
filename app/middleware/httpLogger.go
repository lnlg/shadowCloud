package middleware

import (
	"shadowCloud/internal/global"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HttpLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		elapsed := time.Since(start)
		global.Logger.Info("Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("client_iP", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.Duration("elapsed", elapsed),
		)
	}
}
