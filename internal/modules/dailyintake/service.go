// internal/modules/dailyintake/service.go
package dailyintake

import (
	"calorie_deficit/internal/constants"

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
	existingIntake, err := service.GetDailyIntakeByDateAndMealType(params.UserID, params.Date, params.MealType)
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

func (service *Service) GetDailyIntakeByDateAndMealType(userID uint, date string, mealType constants.MealType) (*DailyIntake, error) {
	// Call the repository method to get the daily intake by date and meal type
	dailyIntake, err := service.Repository.GetDailyIntakeByDateAndMealType(userID, date, mealType)
	if err != nil {
		return nil, err
	}
	return dailyIntake, nil
}

func (service *Service) GetDailyIntake(id uint) (*DailyIntake, error) {
	// Call the repository method to get the daily intake by ID
	repositoryResponse, repositoryError := service.Repository.GetDailyIntake(id)
	if repositoryError != nil {
		return nil, repositoryError
	}
	return repositoryResponse, nil
}
