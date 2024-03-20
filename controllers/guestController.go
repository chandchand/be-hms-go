package controllers

import (
	"net/http"

	"github.com/chandchand/be-hms-go/database"
	"github.com/chandchand/be-hms-go/database/models"
	"github.com/chandchand/be-hms-go/services"
	"github.com/gin-gonic/gin"
)

func GuestCreate(c *gin.Context) {
	var guest models.Guest
	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&guest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false,"message": "Failed to create data","error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true,"message": "Created data success", "data": guest})
}

func GetGuests(c *gin.Context)  {
	var guest []models.Guest

	if err := database.DB.Find(&guest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if guest == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "no guest data found!!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Get Data Success", "data": guest})
}

func GetGuest(c *gin.Context) {
	id := c.Param("id")
	guest, err := services.GetGuestByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "guest with id " + id + " not found", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Get one data guest success", "data": guest})
}

func UpdateGuest(c *gin.Context) {
	id := c.Param("id")
	
	guest, err := services.GetGuestByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "guest with id " + id + " not found", "error": err.Error()})
		return
	}

	var updatedGuest models.Guest
	if err := c.ShouldBindJSON(&updatedGuest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&guest).Updates(&updatedGuest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Failed to update guest","error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message":"Updated guest successfully", "data": guest})
}

func DeleteGuest(c *gin.Context) {
	id := c.Param("id")
	guest, err := services.GetGuestByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "guest with id " + id + " not found", "error": err.Error()})
		return
	}

	if err := database.DB.Delete(&guest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed Delete Guest", "error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Guest deleted successfully","data": guest})
}