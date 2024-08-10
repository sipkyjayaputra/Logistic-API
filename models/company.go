package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}
