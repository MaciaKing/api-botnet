package main

import (
	"api-botnet/database"
	"api-botnet/handlers"

	"github.com/gin-gonic/gin"
)

func initDatabase() {
	database.Connect()
	database.Migrate()
}

func main() {
	initDatabase()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/bots", handlers.GetAllBots)
	r.POST("/bot/create", handlers.Create)
	r.Run() // listen and serve on 0.0.0.0:8080
}
