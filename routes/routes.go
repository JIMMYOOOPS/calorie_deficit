package routes

import (
	"calorie_deficit/types"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes initializes the Gin router and sets up the routes for the application
func RegisterRoutes(router *gin.Engine) {
	// version 1 group
	v1 := router.Group("/v1")
	{
		// entry point for the API
		v1.GET("/", func(context *gin.Context) {
			context.JSON(200, types.Response{Message: "Welcome to the Calorie Deficit API!"})
		})

		// health check endpoint
		v1.GET("/health", func(context *gin.Context) {
			context.JSON(200, types.Response{Message: "OK"})
		})

		// other module routes can be added here...
	}
}
