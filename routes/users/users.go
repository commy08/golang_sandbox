package users

import (
	"github.com/labstack/echo"
)

type UsersRouter struct{}

// Init : Init Router
func (controller UsersRouter) Init(g *echo.Group) {
	g.POST("", controller.register)
}
