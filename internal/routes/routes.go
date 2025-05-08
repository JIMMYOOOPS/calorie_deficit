package routes

import (
	"calorie_deficit/internal/handler"
	"calorie_deficit/internal/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes initializes the Gin router and sets up the routes for the application
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// version 1 group
	v1 := router.Group("/v1")
	{
		// entry point for the API
		v1.GET("/", func(context *gin.Context) {
			context.JSON(200, types.SuccessResponse[string]{Message: "Welcome to the Calorie Deficit API!"})
		})

		// health check endpoint
		v1.GET("/healthz", handler.HealthCheckHandler(db))

		// other module routes can be added here...
	}
}
