package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string  `json:"name"`
	NumPackages   int     `json:"num_packages"`
	HsCode        string  `json:"hs_code"`
	Quantity      float64 `json:"quantity"`
	QuantityUnit  string  `json:"quantity_unit"`
	GrossWeight   float64 `json:"gross_weight"`
	NettWeight    float64 `json:"nett_weight"`
	PackingListID uint    `json:"packing_list_id"`
	Image         string  `json:"image"`
}
