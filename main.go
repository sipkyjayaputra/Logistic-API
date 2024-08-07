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

	authorized.GET("/payment-terms", controllers.GetPaymentTerms)
	authorized.POST("/payment-terms", controllers.CreatePaymentTerm)
	authorized.GET("/payment-terms/:id", controllers.GetPaymentTerm)
	authorized.PUT("/payment-terms/:id", controllers.UpdatePaymentTerm)
	authorized.DELETE("/payment-terms/:id", controllers.DeletePaymentTerm)

	authorized.GET("/commercial-invoices", controllers.GetCommercialInvoices)
	authorized.POST("/commercial-invoices", controllers.CreateCommercialInvoice)
	authorized.GET("/commercial-invoices/:id", controllers.GetCommercialInvoice)
	authorized.PUT("/commercial-invoices/:id", controllers.UpdateCommercialInvoice)
	authorized.DELETE("/commercial-invoices/:id", controllers.DeleteCommercialInvoice)

	authorized.GET("/packages", controllers.GetPackages)
	authorized.POST("/packages", controllers.CreatePackage)
	authorized.GET("/packages/:id", controllers.GetPackage)
	authorized.PUT("/packages/:id", controllers.UpdatePackage)
	authorized.DELETE("/packages/:id", controllers.DeletePackage)

	authorized.GET("/products", controllers.GetProducts)
	authorized.POST("/products", controllers.CreateProduct)
	authorized.GET("/products/:id", controllers.GetProduct)
	authorized.PUT("/products/:id", controllers.UpdateProduct)
	authorized.DELETE("/products/:id", controllers.DeleteProduct)

	authorized.GET("/packing-lists", controllers.GetPackingLists)
	authorized.POST("/packing-lists", controllers.CreatePackingList)
	authorized.GET("/packing-lists/:id", controllers.GetPackingList)
	authorized.PUT("/packing-lists/:id", controllers.UpdatePackingList)
	authorized.DELETE("/packing-lists/:id", controllers.DeletePackingList)

	r.Run()
}
