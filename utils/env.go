package utils

import (
	"os"
)

// GetEnv retrieves the value of the environment variable named by the key.
// If the variable is not found, it returns the provided default value.
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
