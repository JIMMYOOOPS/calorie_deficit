package handler

import (
	"calorie_deficit/internal/types"

	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HealthCheckHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		sqlDB, err := db.DB()
		log.Printf("Health check executed: %v", err)
		if err != nil || sqlDB.Ping() != nil {
			context.JSON(500, types.ErrorResponse{Code: 500, Message: "Database connection failed"})
			return
		}

		context.JSON(200, types.SuccessResponse[string]{Message: "OK"})
	}
}
