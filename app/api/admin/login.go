package admin

import (
	"shadowCloud/app/request"
	"shadowCloud/app/response"
	"shadowCloud/app/service"

	"github.com/gin-gonic/gin"
)

type loginAdmin struct{}

// Admin 登录
func (l *loginAdmin) Login(ctx *gin.Context) {
	var form request.LoginAdminRequest
	if err := ctx.ShouldBind(&form); err != nil {
		response.ReturnValidateFailed(ctx, request.GetErrorMsg(form, err))
		return
	}
	isTrue, token := service.AdminUsersService.Login(form.Username, form.Password, ctx.ClientIP())
	if isTrue {
		response.ReturnSuccess(ctx, 200, "登录成功！", token)
	} else {
		response.ReturnError(ctx, 1001, "用户名或密码错误！"+token)
	}
}

// 根据token获取用户信息
func (l *loginAdmin) GetUserInfoByToken(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		response.ReturnError(ctx, 1, "token不能为空！")
		return
	}
	isTrue, userinfo := service.AdminUsersService.GetUserInfoByToken(token)
	if !isTrue {
		response.ReturnError(ctx, 10000, "请先登录！")
		return
	}
	response.ReturnSuccess(ctx, 200, "获取成功！", userinfo)
}
