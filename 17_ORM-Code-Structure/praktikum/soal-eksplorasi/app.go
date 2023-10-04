package main

import (
	"appsv/config"
	"appsv/routes"
)

func main() {

	config.ConnectDB()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
