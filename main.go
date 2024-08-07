package main

import (
	"logistic-api/controllers"
	"logistic-api/database"
	"logistic-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())

	authorized.GET("/containers", controllers.GetContainers)
	authorized.POST("/containers", controllers.CreateContainer)
	authorized.GET("/containers/:id", controllers.GetContainer)
	authorized.PUT("/containers/:id", controllers.UpdateContainer)
	authorized.DELETE("/containers/:id", controllers.DeleteContainer)

	authorized.GET("/packinglists", controllers.GetPackingLists)
	authorized.POST("/packinglists", controllers.CreatePackingList)
	authorized.GET("/packinglists/:id", controllers.GetPackingList)
	authorized.PUT("/packinglists/:id", controllers.UpdatePackingList)
	authorized.DELETE("/packinglists/:id", controllers.DeletePackingList)

	authorized.GET("/products", controllers.GetProducts)
	authorized.POST("/products", controllers.CreateProduct)
	authorized.GET("/products/:id", controllers.GetProduct)
	authorized.PUT("/products/:id", controllers.UpdateProduct)
	authorized.DELETE("/products/:id", controllers.DeleteProduct)

	r.Run()
}
