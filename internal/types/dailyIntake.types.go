// internal/types/dailyIntake.types.go
package types

import (
	"calorie_deficit/internal/constants"
)

// DailyIntake contains the daily intake types

// MealItem is a struct for the meal item
type MealItem struct {
	Name            string                    `json:"name"`
	MeasurementType constants.MeasurementType `json:"measurementType"`
	Calories        float64                   `json:"calories"`
}
