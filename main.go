package main

import (
	"logistic-api/controllers"
	"logistic-api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	// r.GET("/containers", controllers.GetContainers)
	// r.POST("/containers", controllers.CreateContainer)
	// r.GET("/containers/:id", controllers.GetContainer)
	// r.PUT("/containers/:id", controllers.UpdateContainer)
	// r.DELETE("/containers/:id", controllers.DeleteContainer)

	r.GET("/packinglists", controllers.GetPackingLists)
	r.POST("/packinglists", controllers.CreatePackingList)
	r.GET("/packinglists/:id", controllers.GetPackingList)
	r.PUT("/packinglists/:id", controllers.UpdatePackingList)
	r.DELETE("/packinglists/:id", controllers.DeletePackingList)

	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id", controllers.GetProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	r.Run()
}
