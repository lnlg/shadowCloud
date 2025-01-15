package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponseSuccess struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type JsonResponseError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 返回成功
func ReturnSuccess(c *gin.Context, code int, message string, data interface{}) {
	json := &JsonResponseSuccess{
		Code: code,
		Msg:  message,
		Data: data,
	}
	c.JSON(http.StatusOK, json)
}

// 返回错误
func ReturnError(c *gin.Context, code int, message string) {
	json := &JsonResponseError{
		Code: code,
		Msg:  message,
	}
	c.JSON(http.StatusOK, json)
}

// 验证器失败返回
func ReturnValidateFailed(c *gin.Context, msg string) {
	ReturnError(c, 1, msg)
}
