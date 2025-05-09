package utils

import (
	"calorie_deficit/internal/constants"

	"log"
	"os"

	"github.com/joho/godotenv"
)

// init function is used to load environment variables from a .env file before the main function is called.
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print(constants.EnvFileLoadFailed, err)
	} else {
		log.Println(constants.EnvFileLoaded)
	}
}

// GetEnv retrieves the value of the environment variable named by the key.
// If the variable is not found, it returns the provided default value.
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}
