package controllers

import (
	"logistic-api/database"
	"logistic-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPackingLists(c *gin.Context) {
	var packingLists []models.PackingList
	database.DB.Find(&packingLists)
	c.JSON(http.StatusOK, packingLists)
}

func CreatePackingList(c *gin.Context) {
	var packingList models.PackingList
	if err := c.ShouldBindJSON(&packingList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&packingList)
	c.JSON(http.StatusOK, packingList)
}

func GetPackingList(c *gin.Context) {
	id := c.Param("id")
	var packingList models.PackingList
	if err := database.DB.Preload("CommercialInvoice.Products").Preload("CommercialInvoice").First(&packingList, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Packing List not found"})
		return
	}

	c.JSON(http.StatusOK, packingList)
}

func UpdatePackingList(c *gin.Context) {
	id := c.Param("id")
	var packingList models.PackingList
	if err := database.DB.First(&packingList, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Packing List not found"})
		return
	}
	if err := c.ShouldBindJSON(&packingList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&packingList)
	c.JSON(http.StatusOK, packingList)
}

func DeletePackingList(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.PackingList{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Packing List not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Packing List deleted"})
}
