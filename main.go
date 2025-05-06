package main

import (
	// Import utils package early to load environment variables
	"calorie_deficit/config" // Import config package early to set up environment variables
	"calorie_deficit/infrastructure/database/postgres"
	"calorie_deficit/routes"

	// Import utils package early to load environment variables
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize connection to PostgreSQL database
	if err := postgres.ConnectPostgresDB(); err != nil {
		panic(err)
	}

	// Create a new Gin router
	router := gin.Default()
	// Register the routes with the router
	routes.RegisterRoutes(router)

	// Start the server on port 8080
	router.Run(config.SERVER_PORT)
}
