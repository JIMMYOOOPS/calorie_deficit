package handler

import (
	"calorie_deficit/internal/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HealthCheckHandler(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Check Postgres connection
		postgresqlDB, err := db.DB()
		if err != nil || postgresqlDB.Ping() != nil {
			context.JSON(500, types.ErrorResponse{Code: 500, Message: "Database connection failed"})
			return
		}

		context.JSON(200, types.SuccessResponse[string]{Message: "OK"})
	}
}
