// internal/modules/dailyintake/dto.go
package dailyintake

import (
	"calorie_deficit/internal/constants"
	"calorie_deficit/internal/types"

	"time"
)

type DailyIntakeCreateRequestDTO struct {
	UserID    uint                 `json:"user_id" binding:"required"`
	Date      string               `json:"date" binding:"required,datetime=2006-01-02T15:04:05.000Z"`       // Date in ISO 8601
	MealType  constants.MealType   `json:"meal_type" binding:"required,oneof=breakfast lunch dinner snack"` // MealType is an enum for the meal type
	MealItems []MealItemRequestDTO `json:"meal_items" binding:"required,dive"`                              // MealItems is a slice of MealItem
}

type MealItemRequestDTO struct {
	Name        string                    `json:"name" binding:"required"`                       // Name of the meal item
	Measurement constants.MeasurementType `json:"measurement" binding:"required,oneof=grams ml"` // Measurement is an enum for the measurement type
	Quantity    float64                   `json:"quantity" binding:"required"`                   // Quantity of the meal item
	Calorie     float64                   `json:"calorie" binding:"required"`                    // Calorie of the meal item
}

type DailyIntakeCreateResponseDTO struct {
	ID        uint                  `json:"id"`         // ID is the primary key of the daily intake table
	UserID    uint                  `json:"user_id"`    // UserID is the foreign key to the user table
	Date      string                `json:"date"`       // Date in ISO 8601 format 2025-05-18T02:55:46.844Z
	MealType  constants.MealType    `json:"meal_type"`  // MealType is an enum for the meal type
	MealItems []MealItemResponseDTO `json:"meal_items"` // MealItems is a slice of MealItem
	CreatedAt string                `json:"created_at"` // CreatedAt in ISO 8601 format 2025-05-18T02:55:46.844Z
	UpdatedAt string                `json:"updated_at"` // UpdatedAt in ISO 8601 format 2025-05-18T02:55:46.844Z
}

type MealItemResponseDTO struct {
	ID            uint                      `json:"id"`              // ID is the primary key of the meal item table
	DailyIntakeID uint                      `json:"daily_intake_id"` // DailyIntakeID is the foreign key to the daily intake table
	Name          string                    `json:"name"`            // Name of the meal item
	Measurement   constants.MeasurementType `json:"measurement"`     // Measurement is an enum for the measurement type
	Quantity      float64                   `json:"quantity"`        // Quantity of the meal item
	Calorie       float64                   `json:"calorie"`         // Calorie of the meal item
}

type DailyIntakeCreateServiceDTO struct {
	UserID    uint               `json:"userId"`    // UserID is the foreign key to the user table
	Date      string             `json:"date"`      // Date in ISO 8601 format 2025-05-18T02:55:46.844Z
	MealType  constants.MealType `json:"mealType"`  // MealType is an enum for the meal type
	MealItems []MealItemDTO      `json:"mealItems"` // MealItems is a slice of MealItem
}

type MealItemDTO struct {
	Name        string                    `json:"name"`        // Name of the meal item
	Measurement constants.MeasurementType `json:"measurement"` // Measurement is an enum for the measurement type
	Quantity    float64                   `json:"quantity"`    // Quantity of the meal item
	Calorie     float64                   `json:"calorie"`     // Calorie of the meal item
}

type DailyIntakesListServiceDTO struct {
	UserID   *uint      `json:"userId,omitempty"`   // UserID is the foreign key to the user table
	Date     *time.Time `json:"date,omitempty"`     // Date in ISO 8601 format 2025-05-18T02:55:46.844Z
	MealType string     `json:"mealType,omitempty"` // MealType is an enum for the meal type
	Page     uint       `json:"page"`               // Page number for pagination
	PageSize uint       `json:"pageSize"`           // Page size for pagination
}

type DailyIntakeListDTO struct {
	Items      []DailyIntake `json:"items"`      // Items is a slice of DailyIntakeCreateResponseDTO
	Page       uint          `json:"page"`       // Page number for pagination
	PageSize   uint          `json:"pageSize"`   // Page size for pagination
	TotalCount uint          `json:"totalCount"` // Total number of items
}

type DailyIntakeListResponseDTO struct {
	Items []DailyIntake `json:"items"` // Items is a slice of DailyIntakeCreateResponseDTO
	Meta  *types.Meta   `json:"meta"`  // Meta is a pointer to the Meta struct
}
