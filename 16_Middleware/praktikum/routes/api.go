package routes

import (
	"appmid/controller/authenticate"
	"appmid/controller/book"
	"appmid/controller/user"
	auth "appmid/middleware"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Use(auth.NotFoundHandler)
	auth.Logger(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to RESTful API Services")
	})

	e.POST("/login", authenticate.LoginController)

	e.GET("/users", user.Index)
	e.GET("/users/:id", user.Show)
	e.POST("/users", user.Store)
	e.PUT("/users/:id", user.Update)
	e.DELETE("/users/:id", user.Delete)

	e.GET("/books", book.Index)
	e.GET("/books/:id", book.Show)
	e.POST("/books", book.Store)
	e.PUT("/books/:id", book.Update)
	e.DELETE("/books/:id", book.Delete)

	eAuth := e.Group("/auth")
	eAuth.Use(echojwt.JWT([]byte("1234")))

	eAuth.GET("/users", user.Index)
	eAuth.GET("/users/:id", user.Show)
	eAuth.POST("/users", user.Store)
	eAuth.PUT("/users/:id", user.Update)
	eAuth.DELETE("/users/:id", user.Delete)

	eAuth.GET("/books", book.Index)
	eAuth.GET("/books/:id", book.Show)
	eAuth.POST("/books", book.Store)
	eAuth.PUT("/books/:id", book.Update)
	eAuth.DELETE("/books/:id", book.Delete)

	return e

}
