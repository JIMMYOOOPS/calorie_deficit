package mcp

import (
	"calorie_deficit/internal/config"
	"calorie_deficit/internal/constants"
	"calorie_deficit/internal/infrastructure/mcp/clients"
	"calorie_deficit/internal/pkg/logger"

	"context"
	"errors"
)

// LLMClient defines the interface for all LLM clients
type LLMClient interface {
	CreateChatCompletion(ctx context.Context, input string) (string, error)
}

func InitializeClient(llm string) (LLMClient, error) {
	switch llm {
	case "openai":
		apiKey := config.LLM_CONFIG.OpenAIConfig.APIKey
		client := clients.NewOpenAIClient(apiKey)
		logger.Logger.Infof(constants.LogMessages.MCP.Client.ClientInitialized, llm)
		return client, nil
	case "openRouter":
		apiKey := config.LLM_CONFIG.OpenRouterAIConfig.APIKey
		client := clients.NewOpenRouterClient(apiKey)
		logger.Logger.Infof(constants.LogMessages.MCP.Client.ClientInitialized, llm)
		return client, nil
	default:
		return nil, errors.New("invalid LLM specified")
	}
}
