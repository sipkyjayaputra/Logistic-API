package models

import "gorm.io/gorm"

type Package struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}
