package app

import (
	"shadowCloud/app/event/event"
	"shadowCloud/app/models"
	"shadowCloud/app/request"
	"shadowCloud/app/response"
	"shadowCloud/internal/global"

	"github.com/gin-gonic/gin"
)

type appApi struct{}

func (a *appApi) GetUserInfo(ctx *gin.Context) {
	user, err := models.GetTestOne(1)
	if err != nil {
		response.ReturnError(ctx, 1, "获取用户信息失败")
		return
	}
	response.ReturnSuccess(ctx, 200, "获取用户信息成功", user)
}

// 测试redis
func (a *appApi) RedisTest(ctx *gin.Context) {
	//测试异常请求
	global.Rdb.Set(ctx, "test", "test11111", 0).Err()
	str, err := global.Rdb.Get(ctx, "test").Result()
	if err != nil {
		response.ReturnError(ctx, 1, "redis设置失败")
		return
	}
	response.ReturnSuccess(ctx, 200, "redis测试成功", str)
}

// 这里故意触发一个panic，测试全局异常处理中间件
func (a *appApi) PanicTest(ctx *gin.Context) {
	panic("出现了一个异常")
}

// 测试验证器
func (a *appApi) TestValidator(ctx *gin.Context) {
	var param request.TestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.ReturnValidateFailed(ctx, request.GetErrorMsg(param, err))
		return
	}
	response.ReturnSuccess(ctx, 200, "验证器测试成功", param)
}

// 测试事件
func (a *appApi) TestEvent(ctx *gin.Context) {
	// 触发事件
	global.Event.Dispatch(&event.FooEvent{Id: 1, Username: "小明"})
	global.Event.Dispatch(&event.FooEvent{Id: 2, Username: "小红"})
	response.ReturnSuccess(ctx, 200, "事件测试成功", "test")
}
