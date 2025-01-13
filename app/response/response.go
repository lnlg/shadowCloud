package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type JsonResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 返回成功
func ReturnSuccess(c *gin.Context, code int, message string, data interface{}) {
	json := &JsonResponseSuccess{
		Code:    code,
		Message: message,
		Data:    data,
	}
	c.JSON(http.StatusOK, json)
}

// 返回错误
func ReturnError(c *gin.Context, code int, message string) {
	json := &JsonResponseError{
		Code:    code,
		Message: message,
	}
	c.JSON(http.StatusOK, json)
}
