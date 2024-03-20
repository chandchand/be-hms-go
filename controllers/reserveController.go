package controllers

import (
	"net/http"

	"github.com/chandchand/be-hms-go/database"
	"github.com/chandchand/be-hms-go/database/models"
	"github.com/gin-gonic/gin"
)

func CreateReservation(c *gin.Context) {
	var rs models.Reservation
	if err := c.ShouldBindJSON(&rs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var g models.Guest
	if err := database.DB.First(&g, "id = ?", rs.GuestID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid guest ID"})
		return
	}
	var r models.Room
	if err := database.DB.First(&r, "id = ?", rs.RoomID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid guest ID"})
		return
	}
	var s models.Staff
	if err := database.DB.First(&s, "id = ?", rs.StaffID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid guest ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Success Created Reservations", "data": rs})
}