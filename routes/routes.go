package routes

import (
	"clean-code/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	AuthController users.AuthController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("api/v1/users")

	users.POST("/register", cl.AuthController.CreateUser)
	users.POST("/login", cl.AuthController.Login)

	auth := e.Group("api/v1/users", middleware.JWTWithConfig(cl.JWTMiddleware))
	auth.GET("", cl.AuthController.GetAllUsers)
}
