package app

import (
	"shadowCloud/app/models"
	"shadowCloud/app/request"
	"shadowCloud/app/response"

	"github.com/gin-gonic/gin"
)

type videoSetting struct{}

// 获取视频配置信息
func (v *videoSetting) GetVideoSettingByKey(ctx *gin.Context) {
	key := ctx.Query("key")
	if key == "" {
		response.ReturnError(ctx, 1, "key不能为空")
		return
	}
	data := models.GetVideoSettingByKey(key)
	response.ReturnSuccess(ctx, 200, "获取成功", data)
}

// 添加视频配置信息
func (v *videoSetting) AddVideoSetting(ctx *gin.Context) {
	var params request.VideoSetting
	if err := ctx.ShouldBind(&params); err != nil {
		response.ReturnValidateFailed(ctx, request.GetErrorMsg(params, err))
		return
	}

	res, err := models.AddVideoSetting(params.Key, params.Value, params.Notes)
	if err != nil {
		response.ReturnError(ctx, 1, "添加失败")
		return
	}
	response.ReturnSuccess(ctx, 200, "添加成功", res)
}
