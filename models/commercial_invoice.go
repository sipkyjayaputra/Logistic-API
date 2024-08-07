package models

import (
	"gorm.io/gorm"
)

type CommercialInvoice struct {
	gorm.Model
	InvoiceNo       string    `json:"invoice_no"`
	TransportDetail string    `json:"transport_detail"`
	PaymentTermID   uint      `json:"payment_term_id"`
	Products        []Product `json:"products"`
}
