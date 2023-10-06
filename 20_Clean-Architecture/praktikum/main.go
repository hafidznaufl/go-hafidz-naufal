package main

import (
	"clean_architecture/config"
	"clean_architecture/controller"
	"clean_architecture/model/domain"
	"clean_architecture/repository"
	"clean_architecture/routes"
	"clean_architecture/service"

	"github.com/labstack/echo/v4"
)

func main() {
	clean_architecture := echo.New()

	db, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(domain.User{})

	userRepository := repository.NewUserRepository(db)
	userServcice := service.NewUserService(userRepository)
	userController := controller.NewUserController(userServcice)

	routes.NewUserRoutes(clean_architecture, userController)


	clean_architecture.Logger.Fatal(clean_architecture.Start(":8080"))
}
