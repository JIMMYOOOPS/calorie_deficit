// internal/modules/dailyintake
package dailyintake

import (
	"calorie_deficit/internal/constants"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
func (repository *Repository) CreateDailyIntake(params *DailyIntakeCreateServiceDTO) (*DailyIntake, error) {
	// Map DTO to model
	dailyIntake := DailyIntake{
		UserID: params.UserID,
		Date: func() time.Time {
			parsedDate, _ := time.Parse("2006-01-02T15:04:05.000Z", params.Date)
			return parsedDate
		}(),
		MealType: params.MealType,
		MealItems: func() []MealItem {
			items := make([]MealItem, len(params.MealItems))
			for i, item := range params.MealItems {
				items[i] = MealItem{
					Name:        item.Name,
					Measurement: item.Measurement,
					Quantity:    item.Quantity,
					Calorie:     item.Calorie,
				}
			}
			return items
		}(),
	}

	if err := repository.DB.Create(&dailyIntake).Error; err != nil {
		return nil, err
	}
	if err := repository.DB.Preload("MealItems").First(&dailyIntake, dailyIntake.ID).Error; err != nil {
		return nil, err
	}

	return &dailyIntake, nil
}

func (repository *Repository) GetDailyIntakeByDateAndMealType(userID uint, date string, mealType constants.MealType) (*DailyIntake, error) {
	var dailyIntake DailyIntake
	// Parse the date string to time.Time
	parsedDate, err := time.Parse("2006-01-02T15:04:05.000Z", date)
	if err != nil {
		return nil, err
	}
	// Query the database for the daily intake record the date only needs to match the date part
	err = repository.DB.
		Where("user_id = ? AND DATE(date) = ? AND meal_type = ?", userID, parsedDate.Format("2006-01-02"), mealType).
		First(&dailyIntake).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No record found
		}
		return nil, err // Some other error occurred
	}
	if err := repository.DB.Preload("MealItems").First(&dailyIntake, dailyIntake.ID).Error; err != nil {
		return nil, err
	}
	return &dailyIntake, nil
}

func (repository *Repository) AddMealItemsToDailyIntake(dailyIntakeID uint, mealItems []MealItem) (*DailyIntake, error) {
	var updatedDailyIntake DailyIntake
	// Update the daily intake record with the new meal items
	if err := repository.DB.Model(&DailyIntake{ID: dailyIntakeID}).Association("MealItems").Append(&mealItems); err != nil {
		return nil, err
	}

	if err := repository.DB.Preload("MealItems").First(&updatedDailyIntake, dailyIntakeID).Error; err != nil {
		return nil, err
	}
	return &updatedDailyIntake, nil
}
