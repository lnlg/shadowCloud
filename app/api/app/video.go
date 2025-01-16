package app

import (
	"shadowCloud/app/response"
	"shadowCloud/app/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type videoApi struct {
}

// 根据type_id获取视频列表
func (v *videoApi) GetVideoList(c *gin.Context) {
	typeId, _ := strconv.Atoi(c.DefaultQuery("type_id", "0"))  // 分类id
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))       // 页码
	pageSize, _ := strconv.Atoi(c.DefaultQuery("limit", "10")) // 每页数量

	data := service.VideoService.GetVideoList(typeId, page, pageSize)
	response.ReturnSuccess(c, 200, "获取视频列表成功", data)
}

// 根据id获取视频信息
func (v *videoApi) GetVideoInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	data := service.VideoService.GetVideoInfo(id)
	response.ReturnSuccess(c, 200, "获取视频信息成功", data)
}

// 重新下载视频
func (v *videoApi) ReDownloadVideo(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	if id == 0 {
		response.ReturnError(c, 400, "id不能为空")
		return
	}
	data := service.VideoService.DownloadVideo(id, 0)
	response.ReturnSuccess(c, 200, "重新下载视频成功", data)
}

// 获取未下载视频信息
func (v *videoApi) GetNoDownloadVideo(c *gin.Context) {
	data := service.VideoService.GetNoDownloadVideo()
	response.ReturnSuccess(c, 200, "获取未下载视频信息成功", data)
}

// 根据id标记成功
func (v *videoApi) MarkDownloadSuccess(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	if id == 0 {
		response.ReturnError(c, 400, "id不能为空")
		return
	}
	service.VideoService.MarkDownloadSuccess(id)
	response.ReturnSuccess(c, 200, "标记成功", nil)
}

// 把缓存数据存入数据库
func (v *videoApi) CacheVideoTODB(ctx *gin.Context) {
	list := service.VideoService.CacheVideoTODB()
	response.ReturnSuccess(ctx, 200, "添加成功!", list)
}

// 定时删除大于当前时间的下载中状态的文件
func (v *videoApi) DeleteDownloadVideo(ctx *gin.Context) {
	res := service.VideoService.DeleteDownloadVideo()
	response.ReturnSuccess(ctx, 200, "处理成功!", res)
}
