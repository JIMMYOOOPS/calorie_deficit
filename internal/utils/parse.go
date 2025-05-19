package utils

import (
	"calorie_deficit/internal/pkg/logger"
	"strconv"
)

// ParseStringToUint parses a string to uint and returns the value or error.
func ParseStringToUint(value string) (uint, error) {
	uint64, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		logger.Logger.Error(err.Error())
		return 0, err // Return 0 if the conversion fails
	}
	parsedValue := uint(uint64)
	return parsedValue, nil
}
