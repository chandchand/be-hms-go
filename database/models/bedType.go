package models

import (
	"github.com/google/uuid"
)

type BedType struct {
    ID           uuid.UUID   `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    Name         string      `gorm:"not null" json:"name"`
    MaxPerson    string      `gorm:"not null" json:"max_person"`
    Width        string      `gorm:"not null" json:"width"`
    Length       string      `gorm:"not null" json:"length"`
    ExtraCharge  string      `gorm:"not null" json:"extraCharge"`
	RoomTypes  []*RoomType `gorm:"many2many:roomtype_bedtypes;" json:"room_types"`
}