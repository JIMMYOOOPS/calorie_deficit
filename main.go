package main

import (
	"calorie_deficit/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a simple GET endpoint
	routes.RegisterRoutes(router)

	// Start the server on port 8080
	router.Run(":3000")
}
