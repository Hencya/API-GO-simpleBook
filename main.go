package main

import (
	"Book_CRUD/config"
	Routes "Book_CRUD/routes"
)

func main() {
	config.InitDB()
	echoAPP := Routes.New()
	echoAPP.Logger.Fatal(echoAPP.Start(":8080"))
}
