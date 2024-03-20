package services

import (
	"github.com/chandchand/be-hms-go/database"
	"github.com/chandchand/be-hms-go/database/models"
	"github.com/gin-gonic/gin"
)

func GetGuestByID(c *gin.Context, id string) (*models.Guest, error) {
	var guest models.Guest
	if err := database.DB.Where("id = ?", id).First(&guest).Error; err != nil {
		return nil, err
	}
	return &guest, nil
}
