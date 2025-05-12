package config

import (
	"calorie_deficit/internal/utils"
)

// Set the SERVER_PORT either from an environment variable or default to ":3000"
var SERVER_PORT = utils.GetEnv("SERVER_PORT", ":8080")

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

// Set the OpenAI API key
type OpenAIConfig struct {
	APIKey string
}

type OpenRouterAIConfig struct {
	APIKey string
}

var LLM_CONFIG = struct {
	OpenAIConfig       OpenAIConfig
	OpenRouterAIConfig OpenRouterAIConfig
}{
	OpenAIConfig: OpenAIConfig{
		APIKey: utils.GetEnv("OPENAI_API_KEY", "default_openai_api_key"),
	},
	OpenRouterAIConfig: OpenRouterAIConfig{
		APIKey: utils.GetEnv("OPEN_ROUTER_AI_API_KEY", "default_openrouterai_api_key"),
	},
}
