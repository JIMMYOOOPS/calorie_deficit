// infrastructure/database/postgres/migrations/migrate.go
package migrations

import (
	"calorie_deficit/internal/constants"
	"calorie_deficit/internal/modules/dailyintake"
	"calorie_deficit/internal/pkg/logger"

	"gorm.io/gorm"
)

// Initialize the constants for the migration
var (
	MigrationSuccess = constants.LogMessages.Database.PostgresMigration.MigrationSuccess
	MigrationFailed  = constants.LogMessages.Database.PostgresMigration.MigrationFailed
)

func AutoMigrate(db *gorm.DB) error {
	// Migrate the database
	if err := db.AutoMigrate(
		&dailyintake.DailyIntake{},
		&dailyintake.MealItem{}); err != nil {
		logger.Logger.Error(MigrationFailed, err)
		return err
	}
	logger.Logger.Info(MigrationSuccess)
	return nil
}
