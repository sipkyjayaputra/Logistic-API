package models

import "gorm.io/gorm"

type Container struct {
	gorm.Model
	Name         string        `json:"name"`
	PackingLists []PackingList `json:"packing_lists"`
}
