package services

import (
	"github.com/chandchand/be-hms-go/database"
	"github.com/chandchand/be-hms-go/database/models"
	"github.com/gin-gonic/gin"
)

func GetBedByID(c *gin.Context, id string) (*models.BedType, error) {
	var bed models.BedType
	if err := database.DB.Where("id = ?", id).First(&bed).Error; err != nil {
		return nil, err
	}
	return &bed, nil
}

func GetRoomTypeByID(c *gin.Context, id string) (*models.RoomType, error) {
	var rt models.RoomType
	if err := database.DB.Where("id = ?", id).Preload("BedTypes").First(&rt).Error; err != nil {
		return nil, err
	}
	return &rt, nil
}


func GetRoomByID(c *gin.Context, id string) (*models.Room, error) {
	var room models.Room
	if err := database.DB.Where("id = ?", id).Preload("RoomType.BedTypes").First(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}
