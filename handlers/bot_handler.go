package handlers

import (
	"api-botnet/database"
	"api-botnet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBots(c *gin.Context) {
	var bots []models.Bot
	if err := database.DB.Find(&bots).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bots)
}
