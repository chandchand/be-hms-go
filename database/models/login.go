package models

import (
	"github.com/google/uuid"
)

type Login struct {
    ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    Username string    `gorm:"not null;unique" json:"username" binding:"required"`
    Password string    `gorm:"not null" json:"password"`
    StaffID  uuid.UUID `gorm:"type:uuid;not null" json:"staff_id"`
    Role     string    `gorm:"type:role_enum;not null" json:"role"`
}
