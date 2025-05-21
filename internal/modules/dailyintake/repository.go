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

func (repository *Repository) UpdateDailyIntake(id uint, params *DailyIntakeUpdateServiceDTO) (*DailyIntake, error) {
	// Handle MealItems more efficiently
	if params.MealItems != nil {
		// 1. Fetch existing meal items
		var existingItems []MealItem
		if err := repository.DB.Where("daily_intake_id = ?", id).Find(&existingItems).Error; err != nil {
			return nil, err
		}
		existingMap := make(map[uint]MealItem)
		for _, item := range existingItems {
			existingMap[item.ID] = item
		}

		// 2. Track IDs to keep
		idsToKeep := make(map[uint]bool)
		for _, item := range params.MealItems {
			if item.ID != 0 {
				idsToKeep[item.ID] = true
				// Update existing
				repository.DB.Model(&MealItem{}).Where("id = ?", item.ID).Updates(MealItem{
					Name:        item.Name,
					Measurement: item.Measurement,
					Quantity:    item.Quantity,
					Calorie:     item.Calorie,
				})
			} else {
				// Create new
				newItem := MealItem{
					Name:          item.Name,
					Measurement:   item.Measurement,
					Quantity:      item.Quantity,
					Calorie:       item.Calorie,
					DailyIntakeID: id,
				}
				repository.DB.Create(&newItem)
			}
		}

		// 3. Delete items not in the new list
		for _, oldItem := range existingItems {
			if !idsToKeep[oldItem.ID] {
				repository.DB.Delete(&MealItem{}, oldItem.ID)
			}
		}
	}
	var dailyIntake DailyIntake
	if err := repository.DB.Preload("MealItems").First(&dailyIntake, id).Error; err != nil {
		return nil, err
	}
	return &dailyIntake, nil
}

func (repository *Repository) GetDailyIntake(id uint) (*DailyIntake, error) {
	dailyIntake := DailyIntake{}
	if err := repository.DB.Preload("MealItems").First(&dailyIntake, id).Error; err != nil {
		return nil, err
	}
	return &dailyIntake, nil
}

func (repository *Repository) GetDailyIntakesList(params DailyIntakesListServiceDTO) (DailyIntakeListDTO, error) {
	var (
		dailyIntakes []DailyIntake
		totalCount   int64
	)

	// Build query with filters
	query := repository.DB.Model(&DailyIntake{}).Preload("MealItems")
	if params.UserID != nil {
		query = query.Where("user_id = ?", *params.UserID)
	}
	// Only Picks One Date
	// TODO: Use CreatedAt instead of Date and Range
	if params.Date != nil {
		query = query.Where("DATE(date) = ?", params.Date.Format("2006-01-02"))
	}
	if params.MealType != "" {
		query = query.Where("meal_type = ?", params.MealType)
	}

	// Count total items with filters
	if err := query.Count(&totalCount).Error; err != nil {
		return DailyIntakeListDTO{
			Items:      nil,
			Page:       params.Page,
			PageSize:   params.PageSize,
			TotalCount: 0,
		}, err
	}

	// Pagination: calculate offset
	offset := int((params.Page - 1) * params.PageSize)
	if offset < 0 {
		offset = 0
	}

	// Fetch paginated data
	if err := query.Offset(offset).Limit(int(params.PageSize)).Find(&dailyIntakes).Error; err != nil {
		return DailyIntakeListDTO{
			Items:      nil,
			Page:       params.Page,
			PageSize:   params.PageSize,
			TotalCount: uint(totalCount),
		}, err
	}

	return DailyIntakeListDTO{
		Items:      dailyIntakes,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalCount: uint(totalCount),
	}, nil
}
