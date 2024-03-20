package models

import (
	"github.com/google/uuid"
)

type RoomType struct {
    ID              uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    Name            string     `gorm:"not null" json:"name"`
    PricePerNight   string     `gorm:"not null" json:"price_p_night"`
    PricePerPerson  string     `gorm:"not null" json:"price_p_person"`
    MaxAdult        string     `gorm:"not null" json:"max_adult"`
    MaxChildren     string     `gorm:"not null" json:"max_children"`
    BedTypes        []*BedType `gorm:"many2many:roomtype_bedtypes;" json:"bed_types"`
}

