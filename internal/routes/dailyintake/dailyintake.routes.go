package dailyintake

import (
	"calorie_deficit/internal/modules/dailyintake"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterDailyIntakeRoutes(router *gin.RouterGroup, db *gorm.DB) {
	service := dailyintake.NewService(db)
	handler := dailyintake.NewHandler(service)
	dailyIntakeGroup := router.Group("/daily-intake")
	{
		dailyIntakeGroup.POST("/", handler.CreateDailyIntakeHandler)
	}
}
