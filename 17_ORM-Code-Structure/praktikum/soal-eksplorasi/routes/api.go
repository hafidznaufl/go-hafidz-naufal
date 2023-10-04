package routes

import (
	"appsv/controller"
	"appsv/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.NotFoundHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to RESTful API Services")
	})

	e.GET("/blogs", controller.Index)
	e.GET("/blogs/:id", controller.Show)
	e.POST("/blogs", controller.Store)
	e.PUT("/blogs/:id", controller.Update)
	e.DELETE("/blogs/:id", controller.Delete)

	return e

}
