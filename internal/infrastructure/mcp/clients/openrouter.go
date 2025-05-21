// internal/infrastructure/mcp/clients/openrouter.go

package clients

import (
	"calorie_deficit/internal/constants"
	"calorie_deficit/internal/pkg/logger"

	"context"

	openRouter "github.com/revrost/go-openrouter"
)

type OpenRouterClient struct {
	client *openRouter.Client
}

// NewOpenRouterClient creates a new OpenRouterClient with the provided API key
func NewOpenRouterClient(apiKey string) *OpenRouterClient {
	return &OpenRouterClient{
		client: openRouter.NewClient(apiKey),
	}
}

// CreateChatCompletion sends a chat completion request to the OpenRouter API and returns the json response as a string
func (openRouterClient *OpenRouterClient) CreateChatCompletion(ctx context.Context, systemRoleInput, userRoleInput string) (string, error) {
	response, err := openRouterClient.client.CreateChatCompletion(
		ctx,
		openRouter.ChatCompletionRequest{
			Model: constants.LLMs.OpenRouter.Models.Google[0], // current free model can be switched to any other model
			Messages: []openRouter.ChatCompletionMessage{
				{
					Role:    openRouter.ChatMessageRoleSystem,
					Content: openRouter.Content{Text: systemRoleInput},
				},
				{
					Role:    openRouter.ChatMessageRoleUser,
					Content: openRouter.Content{Text: userRoleInput},
				},
			},
			ResponseFormat: &openRouter.ChatCompletionResponseFormat{
				Type: openRouter.ChatCompletionResponseFormatTypeJSONObject,
			},
		},
	)
	if err != nil {
		logger.Logger.Errorf("Error calling OpenRouter API: %v", err)
		return "", err
	}
	return response.Choices[0].Message.Content.Text, nil
}
