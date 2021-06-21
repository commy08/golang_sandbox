package routes

import (
	"github.com/commy08/golang_sandbox.git/routes/users"
	"github.com/labstack/echo"
)

func Routes(g *echo.Group) {
	users.UsersRouter{}.Init(g.Group("/users"))
}
