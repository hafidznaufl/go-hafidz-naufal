package routes

import (
	"os"
	"clean_architecture/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewUserRoutes(e *echo.Echo, userController controller.UserController) {
	usersGroup := e.Group("users")

	usersGroup.POST("", userController.CreateUserController())

	usersGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	usersGroup.GET("", userController.GetAllUserController())
}