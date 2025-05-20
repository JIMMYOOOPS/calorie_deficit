package utils

import (
	"calorie_deficit/internal/constants"
	"calorie_deficit/internal/pkg/logger"

	"strconv"
	"time"
)

// ParseStringToUint parses a string to uint and returns the value or error.
func ParseStringToUint(value string) (uint, error) {
	parsedInt, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		logger.Logger.Error(err.Error())
		return 0, err // Return 0 if the conversion fails
	}
	parsedValue := uint(parsedInt)
	return parsedValue, nil
}

// ParseStringToDate parses a string to date and returns the value or error.
func ParseStringToDate(value string) (time.Time, error) {
	parsedDate, err := time.Parse(constants.IsoFormatMilSec, value)
	if err != nil {
		logger.Logger.Error(err.Error())
		return time.Time{}, err // Return empty string if the conversion fails
	}
	return parsedDate, nil
}

// ParsePaginationParams parses the page and limit from the request and returns the values or error.
func ParsePaginationParams(page, pageSize string) (uint, uint) {
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	return uint(pageInt), uint(pageSizeInt)
}
