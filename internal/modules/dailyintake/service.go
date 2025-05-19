// internal/modules/dailyintake/service.go
package dailyintake

import (
	"errors"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (service *Service) CreateDailyIntake(params DailyIntakeCreateServiceDTO) (*DailyIntake, error) {
	// Check if there is already a record for the same date and meal type
	existingIntake, err := service.Repository.GetDailyIntakeByDateAndMealType(params.UserID, params.Date, params.MealType)
	if err != nil {
		return nil, err
	}
	// if a record exists for the same date and meal type we want to add the meal items to the existing record
	if existingIntake != nil {
		// Update the existing record with the service method AddMealItemsToDailyIntake
		addMealItemResponse, err := service.AddMealItemsToDailyIntake(existingIntake.ID, params.MealItems)
		if err != nil {
			return nil, err
		}
		return addMealItemResponse, nil
	}

	createResponse, err := service.Repository.CreateDailyIntake(&params)
	if err != nil {
		return nil, err
	}
	return createResponse, nil
}

func (service *Service) AddMealItemsToDailyIntake(dailyIntakeID uint, mealItems []MealItemDTO) (*DailyIntake, error) {
	if len(mealItems) == 0 {
		return nil, errors.New("no meal items provided")
	}
	// Map the meal items from the DTO to the existing intake
	updatedMealItems := make([]MealItem, len(mealItems))
	for i, item := range mealItems {
		updatedMealItems[i] = MealItem{
			Name:        item.Name,
			Measurement: item.Measurement,
			Quantity:    item.Quantity,
			Calorie:     item.Calorie,
		}
	}
	// Update the existing record
	addMealItemResponse, err := service.Repository.AddMealItemsToDailyIntake(dailyIntakeID, updatedMealItems)
	if err != nil {
		return nil, err
	}
	return addMealItemResponse, nil
}
