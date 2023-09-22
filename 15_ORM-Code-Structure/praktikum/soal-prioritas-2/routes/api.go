package routes

import (
	"apps/controller"
	"apps/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.NotFoundHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to RESTful API Services")
	})

	e.GET("/books", controller.Index)
	e.GET("/books/:id", controller.Show)
	e.POST("/books", controller.Store)
	e.PUT("/books/:id", controller.Update)
	e.DELETE("/books/:id", controller.Delete)

	return e

}
