package dailyintake

import (
	dailyintakeapplication "calorie_deficit/internal/application/dailyintake"
	"calorie_deficit/internal/infrastructure/mcp"
	"calorie_deficit/internal/modules/dailyintake"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterDailyIntakeRoutes(router *gin.RouterGroup, db *gorm.DB, llmClient mcp.LLMClient) {
	// Initialize the repository and service
	repository := dailyintake.NewRepository(db)
	service := dailyintake.NewService(repository)
	dailyIntakeGroup := router.Group("/daily-intake")
	mcpDailyIntakeService := dailyintakeapplication.NewMCPDailyIntakeService(llmClient)
	handler := dailyintake.NewHandler(service, mcpDailyIntakeService)
	{
		dailyIntakeGroup.POST("/", handler.CreateDailyIntake)
		dailyIntakeGroup.PATCH("/:id", handler.UpdateDailyIntake)
		dailyIntakeGroup.GET("/:id", handler.GetDailyIntake)
		dailyIntakeGroup.GET("/", handler.GetDailyIntakesList)
	}
}
