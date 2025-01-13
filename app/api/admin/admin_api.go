package admin

import (
	"shadowCloud/app/response"
	"shadowCloud/app/service"

	"github.com/gin-gonic/gin"
)

type adminApi struct{}

func (a adminApi) Profile(ctx *gin.Context) {
	user, err := service.AdminService.Profile()
	if err != nil {
		response.ReturnError(ctx, 1, "获取用户信息失败")
		return
	}
	response.ReturnSuccess(ctx, 200, "获取用户信息成功", user)
}
