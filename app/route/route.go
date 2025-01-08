package route

import "github.com/gin-gonic/gin"

type AppRouter struct{}

// AddRoutes implements route.RouterGeneratorInterface.
func (a *AppRouter) AddRoutes(server *gin.Engine) {
	panic("unimplemented")
}

func (*AppRouter) AddRoute(e *gin.Engine) {
	RegisterAdminRouter(e.Group("/admin"))
	RegisterAppRouter(e.Group("/app"))
}

func New() *AppRouter {
	return &AppRouter{}
}
