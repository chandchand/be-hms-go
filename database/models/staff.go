package models

import (
	"github.com/google/uuid"
)

type Staff struct {
    ID            uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    FirstName     string    `gorm:"not null" json:"first_name"`
    LastName      string    `gorm:"not null" json:"last_name"`
    ContactNumber string    `json:"contact_number"`
    Email         string    `gorm:"not null" json:"email"`
    Position      string    `gorm:"not null" json:"position"`
    Login         Login     `gorm:"foreignKey:StaffID;onDelete:CASCADE" json:"login" binding:"required"`
}
