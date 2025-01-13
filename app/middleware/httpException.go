package middleware

import (
	"fmt"
	"net/http"
	"shadowCloud/internal/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// http异常处理中间件
func HttpExceptionRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.Error("http异常", zap.Any("error", err))
				c.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  fmt.Sprintf("%v", err),
				})
				//终止后接口调用，不加的话recover到异常后，还会继续执行接口后面的代码
				c.Abort()
			}
		}()
		c.Next()
	}
}
