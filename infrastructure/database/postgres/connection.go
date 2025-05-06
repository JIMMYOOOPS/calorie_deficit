package postgres

import (
	"calorie_deficit/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // DB is a global variable to hold the database connection

// ConnectDB establishes a connection to the PostgreSQL database using GORM
func ConnectPostgresDB() error {
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
		return err
	}

	log.Println("Connected to PostgreSQL database successfully")
	DB = db
	return nil
}
