// internal/constants/dailyIntake.go
package constants

// DailyIntake contains the daily intake constants

// MealType is an enum for the meal type
type MealType string

const (
	Breakfast      MealType = "breakfast"
	Brunch         MealType = "brunch"
	Lunch          MealType = "lunch"
	AfternoonSnack MealType = "afternoonSnack"
	Dinner         MealType = "dinner"
	LateNightSnack MealType = "lateNightSnack"
	Drinks         MealType = "drinks" // definitely designed for Taiwanese people
	Other          MealType = "other"
)

// MeasurementType is an enum for the measurement type
type MeasurementType string

const (
	Grams MeasurementType = "grams"
	ML    MeasurementType = "ml"
)
