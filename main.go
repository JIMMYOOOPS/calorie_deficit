package main

import (
	"calorie_deficit/internal/config"                           // Import config package early to set up environment variables
	"calorie_deficit/internal/infrastructure/database/postgres" // Import mcp package early to initialize the client
	"calorie_deficit/internal/pkg/logger"                       // Import logger package early to set up logging
	"calorie_deficit/internal/routes"

	// Import utils package early to load environment variables
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize logger
	logger.InitLogger()
	// Initialize connection to PostgreSQL database
	db, err := postgres.ConnectPostgresDB()
	if err != nil {
		panic(err)
	}

	// Create a new Gin router
	router := gin.Default()
	// Register the routes with the router
	routes.RegisterRoutes(router, db)

	// Start the server on port 8080
	router.Run(config.SERVER_PORT)
}
