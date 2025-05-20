package utils

import (
	"calorie_deficit/internal/constants"
	"calorie_deficit/internal/types"
	"errors"

	"gorm.io/gorm"
)

// HandleNotFoundError wraps gorm.ErrRecordNotFound with your custom AppError
func HandleNotFoundError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return types.NewAppError(constants.ErrNotFound.Code, constants.ErrNotFound.Message)
	}
	return err
}
