package models

import (
	"gorm.io/gorm"
)

type PackingList struct {
	gorm.Model
	PackingListNo       *string `json:"packing_list_no"`
	TotalPackages       int     `json:"total_packages"`
	TotalPackagesName   string  `json:"total_packages_name"`
	CommercialInvoiceID uint    `json:"commercial_invoice_id"`
}
