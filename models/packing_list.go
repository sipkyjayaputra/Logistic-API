package models

import (
	"time"

	"gorm.io/gorm"
)

type PackingList struct {
	gorm.Model
	PackingListNo       *string   `json:"packing_list_no"`
	IssuerName          string    `json:"issuer_name"`
	IssuerAddress       string    `json:"issuer_address"`
	IssuerEmail         string    `json:"issuer_email"`
	IssuerPhone         string    `json:"issuer_phone"`
	ReceiverName        string    `json:"receiver_name"`
	ReceiverAddress     string    `json:"receiver_address"`
	ReceiverEmail       string    `json:"receiver_email"`
	ReceiverPhone       string    `json:"receiver_phone"`
	NoSC                string    `json:"no_sc"`
	Date                time.Time `json:"date"`
	TotalPackages       int       `json:"total_packages"`
	TotalPackagesName   string    `json:"total_packages_name"`
	TotalGrossWeight    float64   `json:"total_gross_weight"`
	TotalNettWeight     float64   `json:"total_nett_weight"`
	CommercialInvoiceID uint      `json:"commercial_invoice_id"`
}
