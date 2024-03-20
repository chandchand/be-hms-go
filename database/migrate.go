package database

import (
	"github.com/chandchand/be-hms-go/database/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Staff{}, 
		&models.Login{}, 
		&models.Guest{}, 
		&models.BedType{}, 
		&models.RoomType{}, 
		&models.Room{},
		&models.Reservation{})
}