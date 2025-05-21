package main

import (
	"calorie_deficit/internal/config"    // Import config package early to set up environment variables
	"calorie_deficit/internal/constants" // Import constants package early to load environment variables
	"calorie_deficit/internal/infrastructure/database/postgres"
	"calorie_deficit/internal/infrastructure/database/postgres/migrations"
	"calorie_deficit/internal/infrastructure/mcp"
	"calorie_deficit/internal/pkg/logger"
	"calorie_deficit/internal/routes"

	// Import utils package early to load environment variables
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize logger
	logger.InitLogger()
	// Initialize connection to PostgreSQL database
	db, databaseError := postgres.ConnectPostgresDB()
	if databaseError != nil {
		panic(databaseError)
	}
	// Migrate the database
	if migrateError := migrations.AutoMigrate(db); migrateError != nil {
		panic(migrateError)
	}
	// Initialize MCP client
	client, mcpClientError := mcp.InitializeClient(constants.LLMs.OpenRouter.Name)
	if mcpClientError != nil {
		panic(mcpClientError)
	}
	if mcpClientError != nil {
		logger.Logger.Errorf(constants.LogMessages.MCP.Client.ClientInitFailed, constants.LLMs.OpenRouter.Name)
		panic(mcpClientError)
	}
	// Create a new Gin router
	router := gin.Default()
	// Register the routes with the router
	routes.RegisterRoutes(router, db, client)

	// Start the server on port 8080
	router.Run(config.SERVER_PORT)
}
