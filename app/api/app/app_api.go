package app

import (
	"shadowCloud/app/models"
	"shadowCloud/app/response"
	"shadowCloud/internal/global"

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

func (a appApi) RedisTest(ctx *gin.Context) {
	global.Rdb.Set(ctx, "test", "test11111", 0).Err()
	str, err := global.Rdb.Get(ctx, "test").Result()
	if err != nil {
		response.ReturnError(ctx, 1, "redis设置失败")
		return
	}
	response.ReturnSuccess(ctx, 200, "redis测试成功", str)
}
