package controllers

import (
	"logistic-api/database"
	"logistic-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPaymentTerms retrieves all payment terms
func GetPaymentTerms(c *gin.Context) {
	var paymentTerms []models.PaymentTerm
	if err := database.DB.Find(&paymentTerms).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving payment terms"})
		return
	}
	c.JSON(http.StatusOK, paymentTerms)
}

// CreatePaymentTerm creates a new payment term
func CreatePaymentTerm(c *gin.Context) {
	var paymentTerm models.PaymentTerm
	if err := c.ShouldBindJSON(&paymentTerm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&paymentTerm).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating payment term"})
		return
	}
	c.JSON(http.StatusOK, paymentTerm)
}

// GetPaymentTerm retrieves a payment term by ID
func GetPaymentTerm(c *gin.Context) {
	id := c.Param("id")
	var paymentTerm models.PaymentTerm
	if err := database.DB.First(&paymentTerm, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment term not found"})
		return
	}
	c.JSON(http.StatusOK, paymentTerm)
}

// UpdatePaymentTerm updates an existing payment term
func UpdatePaymentTerm(c *gin.Context) {
	id := c.Param("id")
	var paymentTerm models.PaymentTerm
	if err := database.DB.First(&paymentTerm, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment term not found"})
		return
	}
	if err := c.ShouldBindJSON(&paymentTerm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Save(&paymentTerm).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating payment term"})
		return
	}
	c.JSON(http.StatusOK, paymentTerm)
}

// DeletePaymentTerm deletes a payment term by ID
func DeletePaymentTerm(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.PaymentTerm{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment term not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Payment term deleted"})
}
