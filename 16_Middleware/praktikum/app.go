package main

import (
	"appmid/config"
	"appmid/routes"
)

func main() {

	config.ConnectDB()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
