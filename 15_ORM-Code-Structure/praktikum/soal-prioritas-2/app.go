package main

import (
	"apps/config"
	"apps/routes"
)

func main() {

	config.ConnectDB()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
