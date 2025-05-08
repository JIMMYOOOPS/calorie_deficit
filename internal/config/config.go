package config

import (
	"calorie_deficit/internal/utils"
)

// Set the SERVER_PORT either from an environment variable or default to ":3000"
var SERVER_PORT = utils.GetEnv("SERVER_PORT", ":3000")

// Set the Database config
type DatabaseConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

var POSTGRES_DB_CONFIG = DatabaseConfig{
	Host:     utils.GetEnv("POSTGRES_HOST", "localhost"),
	Port:     utils.GetEnv("POSTGRES_PORT", "5432"),
	Database: utils.GetEnv("POSTGRES_DB", "calorie_deficit"),
	Username: utils.GetEnv("POSTGRES_USER", "default_user"),
	Password: utils.GetEnv("POSTGRES_PASSWORD", "default_password"),
}
