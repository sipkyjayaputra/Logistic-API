package controllers

import (
	"logistic-api/database"
	"logistic-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCommercialInvoices retrieves all commercial invoices
func GetCommercialInvoices(c *gin.Context) {
	var invoices []models.CommercialInvoice
	if err := database.DB.Find(&invoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving commercial invoices"})
		return
	}
	c.JSON(http.StatusOK, invoices)
}

// CreateCommercialInvoice creates a new commercial invoice
func CreateCommercialInvoice(c *gin.Context) {
	var invoice models.CommercialInvoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating commercial invoice"})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

// GetCommercialInvoice retrieves a commercial invoice by ID
func GetCommercialInvoice(c *gin.Context) {
	id := c.Param("id")
	var invoice models.CommercialInvoice
	if err := database.DB.Preload("Products").First(&invoice, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Commercial invoice not found"})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

// UpdateCommercialInvoice updates an existing commercial invoice
func UpdateCommercialInvoice(c *gin.Context) {
	id := c.Param("id")
	var invoice models.CommercialInvoice
	if err := database.DB.First(&invoice, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Commercial invoice not found"})
		return
	}
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Save(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating commercial invoice"})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

// DeleteCommercialInvoice deletes a commercial invoice by ID
func DeleteCommercialInvoice(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.CommercialInvoice{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Commercial invoice not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Commercial invoice deleted"})
}
