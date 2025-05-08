package postgres

import (
	"calorie_deficit/internal/config"
	"fmt"
	"log"

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
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	log.Println("Connected to PostgreSQL database successfully")
	return db, nil
}
