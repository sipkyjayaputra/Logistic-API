package database

import (
	"log"
	"logistic-api/models"
)

func initPaymenTerm() {
	var paymentTermCount int64
	DB.Model(&models.PaymentTerm{}).Count(&paymentTermCount)
	if paymentTermCount == 0 {
		paymentTerms := []models.PaymentTerm{{
			Name: "CIF",
		}, {
			Name: "CFR",
		}, {
			Name: "FOB",
		}, {
			Name: "CIP",
		}, {
			Name: "CPT",
		}, {
			Name: "DAP",
		}, {
			Name: "DAT",
		}, {
			Name: "DDP",
		}, {
			Name: "EXW",
		}, {
			Name: "FAS",
		}, {
			Name: "FCA",
		}}

		if err := DB.Create(&paymentTerms).Error; err != nil {
			log.Fatalf("Error inserting payment terms: %v", err)
		}
	}
}
