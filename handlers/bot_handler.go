package handlers

import (
	"api-botnet/database"
	"api-botnet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var bot models.Bot
	if err := c.ShouldBindJSON(&bot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := database.DB.Create(&bot).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bot"})
		return
	}
	c.JSON(http.StatusCreated, bot)
}

func GetAllBots(c *gin.Context) {
	var bots []models.Bot
	if err := database.DB.Find(&bots).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bots)
}
