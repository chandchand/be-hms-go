package services

import (
	"github.com/chandchand/be-hms-go/database"
	"github.com/chandchand/be-hms-go/database/models"
	"github.com/google/uuid"
)

func GetUserByID(userID uuid.UUID) (*models.Staff, error) {
	var user models.Staff
	if err := database.DB.Preload("Login").First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}