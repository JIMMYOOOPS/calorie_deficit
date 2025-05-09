package postgres

import (
	"calorie_deficit/internal/config"
	"calorie_deficit/internal/constants"
	"calorie_deficit/internal/pkg/logger"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB establishes a connection to the PostgreSQL database using GORM
func ConnectPostgresDB() (*gorm.DB, error) {
	// Construct the PostgreSQL connection string
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.POSTGRES_DB_CONFIG.Host,
		config.POSTGRES_DB_CONFIG.Port,
		config.POSTGRES_DB_CONFIG.Username,
		config.POSTGRES_DB_CONFIG.Password,
		config.POSTGRES_DB_CONFIG.Database,
	)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		logger.Logger.Error(constants.PostgresConnectionFailed, err)
		return nil, err
	}

	logger.Logger.Info(constants.PostgresConnectionSuccess)
	return db, nil
}
