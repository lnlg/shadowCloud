package admin

import (
	"shadowCloud/app/request"
	"shadowCloud/app/response"
	"shadowCloud/app/service"

	"github.com/gin-gonic/gin"
)

type adminUsers struct{}

// 创建用户信息
func (a *adminUsers) CreateUsers(ctx *gin.Context) {
	var form request.AddUserRequest
	if err := ctx.ShouldBind(&form); err != nil {
		response.ReturnValidateFailed(ctx, request.GetErrorMsg(form, err))
		return
	}
	id, err := service.AdminUsersService.CreateAdminUser(form.Username, form.Password, form.Nickname, form.Email, form.Mobile, form.Avatar)
	if err != nil {
		response.ReturnError(ctx, 1, err.Error())
		return
	}
	response.ReturnSuccess(ctx, 200, "创建成功", id)
}
