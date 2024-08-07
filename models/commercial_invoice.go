package models

import (
	"time"

	"gorm.io/gorm"
)

type CommercialInvoice struct {
	gorm.Model
	InvoiceNo       string    `json:"invoice_no"`
	IssuerName      string    `json:"issuer_name"`
	IssuerAddress   string    `json:"issuer_address"`
	IssuerEmail     string    `json:"issuer_email"`
	IssuerPhone     string    `json:"issuer_phone"`
	ReceiverName    string    `json:"receiver_name"`
	ReceiverAddress string    `json:"receiver_address"`
	ReceiverEmail   string    `json:"receiver_email"`
	ReceiverPhone   string    `json:"receiver_phone"`
	NotifyParty     bool      `json:"notify_party"`
	PartyName       string    `json:"party_name"`
	PartyAddress    string    `json:"party_address"`
	PartyEmail      string    `json:"party_email"`
	PartyPhone      string    `json:"party_phone"`
	NoSC            *string   `json:"no_sc"`
	Date            time.Time `json:"date"`
	TransportDetail string    `json:"transport_detail"`
	PaymentTermID   uint      `json:"payment_term_id"`
	Products        []Product `json:"products"`
}
