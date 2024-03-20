package models

import (
	"github.com/google/uuid"
)

type Guest struct {
    ID       	uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    FirstName 	string    `gorm:"not null" json:"first_name"`
    LastName 	string    `gorm:"not null" json:"last_name"`
    Gender     	string    `gorm:"type:gender_enum;not null" json:"gender"`
	Address 	string 	  `gorm:"not null" json:"address"`
	Phone 		string 	  `gorm:"not null;unique" json:"phone"`
	Email 		string 	  `gorm:"not null" json:"email"`
}
