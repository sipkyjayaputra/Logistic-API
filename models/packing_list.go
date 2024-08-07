package models

import (
	"time"

	"logistic-api/utils"

	"gorm.io/gorm"
)

type PackingList struct {
	gorm.Model
	Issuer            string       `json:"issuer"`
	To                string       `json:"to"`
	NoSC              string       `json:"no_sc"`
	Date              time.Time    `json:"date"`
	TotalPackages     int          `json:"total_packages"`
	TotalGrossWeight  float64      `json:"total_gross_weight"`
	TotalNettWeight   float64      `json:"total_nett_weight"`
	ContainerID       *uint        `json:"container_id"`
	Products          []Product    `json:"products"`
	PackingListStatus utils.Status `json:"packing_list_status"`
}
