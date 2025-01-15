package app

import (
	"shadowCloud/app/models"
	"shadowCloud/app/request"
	"shadowCloud/app/response"

	"github.com/gin-gonic/gin"
)

type videoClass struct{}

// 添加视频分类
func (v *videoClass) AddVideoClass(ctx *gin.Context) {
	var param request.AddVideoClass
	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.ReturnValidateFailed(ctx, request.GetErrorMsg(param, err))
		return
	}
	id, err := models.AddVideoClass(param.Name, param.Sort)
	if err != nil {
		response.ReturnError(ctx, 1, "添加视频分类失败")
		return
	}
	response.ReturnSuccess(ctx, 200, "添加视频分类成功", id)
}

// 获取视频分类列表
func (v *videoClass) GetVideoClassList(ctx *gin.Context) {
	list, err := models.GetVideoClassList()
	if err != nil {
		response.ReturnError(ctx, 1, "获取视频分类列表失败")
		return
	}
	response.ReturnSuccess(ctx, 200, "获取视频分类列表成功", list)
}
