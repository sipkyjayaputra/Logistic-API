package controllers

import (
	"logistic-api/database"
	"logistic-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCompanies retrieves all companies from the database
func GetCompanies(c *gin.Context) {
	var companies []models.Company
	database.DB.Find(&companies)
	c.JSON(http.StatusOK, companies)
}

// CreateCompany handles the creation of a new company record
func CreateCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&company)
	c.JSON(http.StatusOK, company)
}

// GetCompany retrieves a specific company by its ID
func GetCompany(c *gin.Context) {
	id := c.Param("id")
	var company models.Company
	if err := database.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	c.JSON(http.StatusOK, company)
}

// UpdateCompany handles updating an existing company record
func UpdateCompany(c *gin.Context) {
	id := c.Param("id")
	var company models.Company
	if err := database.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&company)
	c.JSON(http.StatusOK, company)
}

// DeleteCompany handles the deletion of a company record
func DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Company{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
}
