// internal/modules/dailyintake/service.go
package dailyintake

import (
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (service *Service) CreateDailyIntake(params DailyIntakeCreateServiceDTO) (DailyIntakeCreateResponseDTO, error) {
	testResponse := DailyIntakeCreateResponseDTO{
		UserID:    params.UserID,
		Date:      params.Date,
		MealType:  params.MealType,
		MealItems: make([]MealItemResponseDTO, len(params.MealItems)),
		CreatedAt: "2025-05-18T02:55:46.844Z",
		UpdatedAt: "2025-05-18T02:55:46.844Z",
	}
	return testResponse, nil
}
