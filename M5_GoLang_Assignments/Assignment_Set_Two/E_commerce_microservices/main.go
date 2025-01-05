package main

import (
	"fmt"
	"log"
	"E_commerce_microservices/database"
	"E_commerce_microservices/controllers"
	"E_commerce_microservices/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Create the router
	r := gin.Default()

	// Apply middleware globally
	r.Use(middleware.RequestLogger())

	// Public routes
	r.POST("/product", controllers.CreateProduct)

	// Protected routes (JWT Authentication)
	r.GET("/product/:id", middleware.JWTMiddleware(), controllers.GetProductByID)
	r.PUT("/product/:id", middleware.JWTMiddleware(), controllers.UpdateProduct)
	r.DELETE("/product/:id", middleware.JWTMiddleware(), controllers.DeleteProduct)

	// Start the server
	fmt.Println("Server running on http://localhost:8000")
	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
