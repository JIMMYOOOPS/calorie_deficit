// internal/modules/dailyintake/model.go
package dailyintake

import (
	"calorie_deficit/internal/constants"

	"time"
)

// DailyIntake contains the daily intake model
type DailyIntake struct {
	ID        uint               `gorm:"primaryKey"`
	UserID    uint               `gorm:"index"` // UserID is the foreign key to the user table
	Date      time.Time          `gorm:"index"`
	MealType  constants.MealType `gorm:"index"`                    // MealType is an enum for the meal type
	MealItems []MealItem         `gorm:"foreignKey:DailyIntakeID"` // MealItems is a slice of MealItem
	CreatedAt time.Time          `gorm:"autoCreateTime"`
	UpdatedAt time.Time          `gorm:"autoUpdateTime"`
}

// MealItem contains the meal item model
type MealItem struct {
	ID            uint `gorm:"primaryKey"`
	DailyIntakeID uint `gorm:"index"`
	Name          string
	Measurement   constants.MeasurementType
	Quantity      float64
	Calorie       float64
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

// TableName returns the table name for the DailyIntake model
func (DailyIntake) TableName() string {
	return "daily_intake"
}

// TableName returns the table name for the MealItem model
func (MealItem) TableName() string {
	return "meal_item"
}
