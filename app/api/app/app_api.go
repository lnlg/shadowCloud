package app

import (
	"shadowCloud/app/models"
	"shadowCloud/app/response"

	"github.com/gin-gonic/gin"
)

type appApi struct{}

func (a appApi) GetUserInfo(ctx *gin.Context) {
	user, err := models.GetTestOne(1)
	if err != nil {
		response.ReturnError(ctx, 1, "获取用户信息失败")
		return
	}
	response.ReturnSuccess(ctx, 200, "获取用户信息成功", user)
}
func (a appApi) GetUserList(ctx *gin.Context) {

}
func (a appApi) CreateUser(ctx *gin.Context) {

}
func (a appApi) UpdateUser(ctx *gin.Context) {

}
func (a appApi) DeleteUser(ctx *gin.Context) {

}
