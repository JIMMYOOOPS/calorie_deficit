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
	// Init constants for readability
	openAIName := constants.LLMs.OpenAI.Name
	openAIAPIkey := config.LLM_CONFIG.OpenAIConfig.APIKey
	openRouterAIName := constants.LLMs.OpenRouterAI.Name
	openRouterAPIkey := config.LLM_CONFIG.OpenRouterAIConfig.APIKey
	clientInitMessage := constants.LogMessages.MCP.Client.ClientInitialized
	invalidModelSpecified := constants.LogMessages.MCP.Client.InvalidModelSpecified

	switch llm {
	case openAIName:
		apiKey := openAIAPIkey
		client := clients.NewOpenAIClient(apiKey)
		logger.Logger.Infof(clientInitMessage, llm)
		return client, nil
	case openRouterAIName:
		apiKey := openRouterAPIkey
		client := clients.NewOpenRouterClient(apiKey)
		logger.Logger.Infof(clientInitMessage, llm)
		return client, nil
	default:
		return nil, errors.New(invalidModelSpecified)
	}
}
