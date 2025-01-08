package route

import "github.com/gin-gonic/gin"

type AppRouter struct{}

// AddRoutes implements route.RouterGeneratorInterface.
func (a *AppRouter) AddRoutes(server *gin.Engine) {
	panic("unimplemented")
}

func (*AppRouter) AddRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}

func New() *AppRouter {
	return &AppRouter{}
}
