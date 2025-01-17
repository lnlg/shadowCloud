package route

import (
	"shadowCloud/app/api/app"

	"github.com/gin-gonic/gin"
)

func RegisterAppRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "app pong"})
	})
	video := r.Group("/video")
	{
		video.POST("/addSetting", app.VideoSettingApi.AddVideoSetting)
		video.GET("/getSetting", app.VideoSettingApi.GetVideoSettingByKey)

		video.POST("/addVideoClass", app.VideoClassApi.AddVideoClass)
		video.GET("/getVideoClassList", app.VideoClassApi.GetVideoClassList)

		video.GET("/getVideoList", app.VideoApi.GetVideoList)          //获取视频列表
		video.GET("/getVideoInfo", app.VideoApi.GetVideoInfo)          //根据id获取视频信息
		video.GET("/downloadOK", app.VideoApi.MarkDownloadSuccess)     //标记下载成功
		video.POST("/reDownload", app.VideoApi.ReDownloadVideo)        //重新下载视频
		video.GET("/getDownload", app.VideoApi.GetNoDownloadVideo)     //获取未下载视频信息
		video.GET("/cacheToDb", app.VideoApi.CacheVideoTODB)           //缓存视频到数据库
		video.GET("/initVideoInfo", app.VideoApi.InitVideoDirPath)     //初始化视频信息
		video.GET("/delDelayVideos", app.VideoApi.DeleteDownloadVideo) //删除大于当前时间的下载中状态的文件
	}

}
