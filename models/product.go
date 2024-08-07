package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name                string  `json:"name"`
	NumPackages         int     `json:"num_packages"`
	NumPackagesName     string  `json:"num_packages_name"`
	HsCode              int     `json:"hs_code"`
	Quantity            float64 `json:"quantity"`
	QuantityName        string  `json:"quantity_name"`
	GrossWeight         float64 `json:"gross_weight"`
	NettWeight          float64 `json:"nett_weight"`
	CommercialInvoiceID uint    `json:"commercial_invoice_id"`
	PackageID           *uint   `json:"package_id"`
	Image               string  `json:"image"`
	UnitPrice           float64 `json:"unit_price"`
	UnitPriceName       string  `json:"unit_price_name"`
	TotalPrice          float64 `json:"total_price"`
	TotalPriceName      string  `json:"total_price_name"`
}
