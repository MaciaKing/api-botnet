package main

import (
	"api-botnet/cmd/router"
	"api-botnet/database"
)

func initDatabase() {
	database.Connect()
	database.Migrate()
}

func main() {
	initDatabase()
	r := router.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
