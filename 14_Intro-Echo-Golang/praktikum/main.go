package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users = []User{
	{Id: 1, Name: "Naufal", Email: "hafidznaufl@gmail.com", Password: "12345"},
	{Id: 2, Name: "Barik", Email: "barik@example.com", Password: "12345"},
	{Id: 3, Name: "Hafidz", Email: "hafidze@example.com", Password: "12345"},
}

// -------------------- controller --------------------
// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id := c.Param("id")
	for _, result := range users {
		if id == strconv.Itoa(result.Id) {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success get user by id",
				"user":    result,
			})
		} else {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}
	}

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": "internal server srror",
	})

}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	for i, result := range users {
		if id == strconv.Itoa(result.Id) {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success delete user",
				"user":    result,
			})
		} else {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}
	}

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": "internal server error",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	userID := c.Param("id")
	var updatedUser User
	c.Bind(&updatedUser)

	for i, result := range users {
		if userID == strconv.Itoa(result.Id) {
			users[i] = updatedUser
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success update user",
				"user":    updatedUser,
			})
		} else {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Naufal")
	})
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.POST("/users", CreateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
