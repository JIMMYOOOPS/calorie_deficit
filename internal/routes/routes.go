package routes

import (
	"calorie_deficit/internal/handler"
	"calorie_deficit/internal/infrastructure/mcp"
	"calorie_deficit/internal/routes/dailyintake"
	"calorie_deficit/internal/types"

	_ "calorie_deficit/docs" // Import the generated swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// RegisterRoutes initializes the Gin router and sets up the routes for the application
func RegisterRoutes(router *gin.Engine, db *gorm.DB, llmClient mcp.LLMClient) {
	// api version 1
	api := router.Group("/api")
	// version 1 group
	v1 := api.Group("/v1")
	// swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	{
		// entry point for the API
		v1.GET("/", func(context *gin.Context) {
			context.JSON(200, types.SuccessResponse[string]{Message: "Welcome to the Calorie Deficit API!"})
		})

		// health check endpoint
		v1.GET("/healthz", handler.HealthCheckHandler(db))

		// import daily intake routes
		dailyintake.RegisterDailyIntakeRoutes(v1, db, llmClient)

	}
}
