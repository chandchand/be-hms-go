package utils

import (
	"reflect"

	"github.com/chandchand/be-hms-go/database/models"
)

// ValidateAndUpdateStaffAttributes melakukan validasi dan pembaruan atribut staff
func ValidateAndUpdateStaffAttributes(guest *models.Guest, updatedAttributes map[string]interface{}) error {
	guestValue := reflect.ValueOf(guest).Elem()

	for key, value := range updatedAttributes {
		if field := guestValue.FieldByName(key); field.IsValid() {
			if key == "FirstName" {
				if len(value.(string)) < 3 {
					return &CustomError{"FirstName harus memiliki panjang minimal 3 karakter"}
				}
			}

			field.Set(reflect.ValueOf(value))
		}
	}

	return nil
}