package route

import (
	"shadowCloud/app/middleware"

	"github.com/gin-gonic/gin"
)

type AppRouter struct{}

// AddRoutes implements route.RouterGeneratorInterface.
func (a *AppRouter) AddRoutes(server *gin.Engine) {
	panic("unimplemented")
}

func (*AppRouter) AddRoute(e *gin.Engine) {
	//记录访问日志中间件
	e.Use(middleware.HttpLogger())
	RegisterAdminRouter(e.Group("/admin"))
	RegisterAppRouter(e.Group("/app"))
}

func New() *AppRouter {
	return &AppRouter{}
}
