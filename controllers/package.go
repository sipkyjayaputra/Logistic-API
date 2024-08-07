package controllers

import (
	"logistic-api/database"
	"logistic-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPackages retrieves all packages
func GetPackages(c *gin.Context) {
	var packages []models.Package
	if err := database.DB.Find(&packages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving packages"})
		return
	}
	c.JSON(http.StatusOK, packages)
}

// CreatePackage creates a new package
func CreatePackage(c *gin.Context) {
	var pkg models.Package
	if err := c.ShouldBindJSON(&pkg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&pkg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating package"})
		return
	}
	c.JSON(http.StatusOK, pkg)
}

// GetPackage retrieves a package by ID
func GetPackage(c *gin.Context) {
	id := c.Param("id")
	var pkg models.Package
	if err := database.DB.First(&pkg, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found"})
		return
	}

	c.JSON(http.StatusOK, pkg)
}

// UpdatePackage updates a package by ID
func UpdatePackage(c *gin.Context) {
	id := c.Param("id")
	var pkg models.Package
	if err := database.DB.First(&pkg, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found"})
		return
	}
	if err := c.ShouldBindJSON(&pkg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Save(&pkg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating package"})
		return
	}
	c.JSON(http.StatusOK, pkg)
}

// DeletePackage deletes a package by ID
func DeletePackage(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Package{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Package deleted"})
}
