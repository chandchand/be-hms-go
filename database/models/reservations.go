package models

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
    ReservationID  uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"reservation_id"`
    RoomID         uuid.UUID `gorm:"type:uuid;not null" json:"room_id"`
    StaffID        uuid.UUID `gorm:"type:uuid;not null" json:"staff_id"`
    GuestID        uuid.UUID `gorm:"type:uuid;not null" json:"guest_id"`
    ServiceID      uuid.UUID `gorm:"type:uuid" json:"service_id"`
    CheckInDate    time.Time `json:"check_in_date"`
    CheckOutDate   time.Time `json:"check_out_date"`
    Adult          int       `gorm:"not null" json:"adult"`
    Children       int       `json:"children"`
    NumberOfGuests int       `gorm:"not null" json:"number_of_guests"`
    ReservationStatus string `gorm:"type:role_enum;not null" json:"reservation_status"`
    TotalPrice     float64   `gorm:"not null" json:"total_price"`
    Deposit        float64   `gorm:"not null" json:"deposit"`
}
