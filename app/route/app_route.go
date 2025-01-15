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
		video.POST("/addVideoClass", app.VideoClassApi.AddVideoClass)
		video.GET("/getVideoClassList", app.VideoClassApi.GetVideoClassList)
		video.GET("/getSetting", app.VideoSettingApi.GetVideoSettingByKey)
		video.POST("/addSetting", app.VideoSettingApi.AddVideoSetting)
	}
	test := r.Group("/test")
	{
		test.GET("/user", app.AppApi.GetUserInfo)
		test.GET("/redis", app.AppApi.RedisTest)
		test.POST("/validator", app.AppApi.TestValidator)
		test.GET("/event", app.AppApi.TestEvent)
	}

}
