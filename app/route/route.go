package route

import (
	"shadowCloud/app/middleware"
	"shadowCloud/internal/global"

	"github.com/gin-gonic/gin"
)

type AppRouter struct{}

// 添加路由
func (a *AppRouter) AddRoutes(server *gin.Engine) {
	panic("unimplemented")
}

func (*AppRouter) AddRoute(e *gin.Engine) {
	// 全局跨域中间件
	e.Use(middleware.CorsMiddleware())
	// 全局http请求异常处理中间件
	e.Use(middleware.HttpExceptionRecover())
	if global.Config.App.Debug {
		// 记录访问日志全局中间件
		e.Use(middleware.HttpLogger())
	}
	// 注册管理员路由
	RegisterAdminRouter(e.Group("/admin"))
	// 注册应用路由
	RegisterAppRouter(e.Group("/app"))
}

func New() *AppRouter {
	return &AppRouter{}
}
