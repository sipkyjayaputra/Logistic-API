package models

import "gorm.io/gorm"

type PaymentTerm struct {
	gorm.Model
	Name string `json:"name"`
}
