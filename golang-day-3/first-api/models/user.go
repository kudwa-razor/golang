package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json"phone"`
	Address string `json:address"`
}
