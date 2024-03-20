package models

import "github.com/google/uuid"

type Room struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	RoomNumber     string    `gorm:"not null" json:"room_number"`
	Availability   string    `gorm:"type:avail_enum;default:Available" json:"availability"`
	RoomTypeID      uuid.UUID `gorm:"type:uuid" json:"room_type_id"`
	RoomType        RoomType   `gorm:"foreignKey:RoomTypeID" json:"room_type"`
}