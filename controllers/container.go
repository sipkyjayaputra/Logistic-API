package controllers

import (
	"logistic-api/database"
	"logistic-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetContainers(c *gin.Context) {
	var containers []models.Container
	database.DB.Find(&containers)
	c.JSON(http.StatusOK, containers)
}

func CreateContainer(c *gin.Context) {
	var container models.Container
	if err := c.ShouldBindJSON(&container); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&container)
	c.JSON(http.StatusOK, container)
}

func GetContainer(c *gin.Context) {
	id := c.Param("id")
	var container models.Container
	if err := database.DB.First(&container, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Container not found"})
		return
	}

	// Fetch packing lists for the container
	var packingLists []models.PackingList
	if err := database.DB.Where("container_id = ?", id).Find(&packingLists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving packing lists"})
		return
	}

	for _, pl := range packingLists {
		var products []models.Product
		if err := database.DB.Where("packing_list_id = ?", pl.Model.ID).Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving products"})
			return
		}
		pl.Products = products

		container.PackingLists = append(container.PackingLists, pl)
	}

	c.JSON(http.StatusOK, container)
}

func UpdateContainer(c *gin.Context) {
	id := c.Param("id")
	var container models.Container
	if err := database.DB.First(&container, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Container not found"})
		return
	}
	if err := c.ShouldBindJSON(&container); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&container)
	c.JSON(http.StatusOK, container)
}

func DeleteContainer(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Container{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Container not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Container deleted"})
}
