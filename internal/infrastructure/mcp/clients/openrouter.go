// internal/infrastructure/mcp/clients/openrouter.go

package clients

import (
	"calorie_deficit/internal/constants"
	"calorie_deficit/internal/pkg/logger"

	"context"
	"encoding/json"
	"fmt"

	openRouter "github.com/revrost/go-openrouter"
	// Import the package where MealItemDTO is defined
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
//
//	Respond ONLY with a JSON array of meal items, each with fields: name, measurement, quantity, calories. Do not include any explanation or extra text.
func (openRouterClient *OpenRouterClient) CreateChatCompletion(ctx context.Context, userRoleInput string, schema any) (string, error) {
	// Chat completion request
	// model := constants.LLMs.OpenRouter.Models.Google[0] // current free model can be switched to any other model
	model := constants.LLMs.OpenRouter.Models.Meta[0] // current free model can be switched to any other model"
	messages := []openRouter.ChatCompletionMessage{
		{
			Role:    openRouter.ChatMessageRoleUser,
			Content: openRouter.Content{Text: userRoleInput},
		},
	}

	// Convert the schema to JSON bytes
	if schema == nil {
		// If schema is nil, use an empty JSON object
		schema = make(map[string]interface{})
	}
	schemaBytes, jsonParseError := json.Marshal(schema)
	if jsonParseError != nil {
		logger.Logger.Errorf("Error marshalling schema to JSON: %v", jsonParseError)
		return "", jsonParseError
	}
	/*
		the response schema is set to MealItemDTO type
	*/
	responseFormat := &openRouter.ChatCompletionResponseFormat{
		Type: openRouter.ChatCompletionResponseFormatTypeJSONSchema,
		// The schema is set to the MealItemDTO type
		JSONSchema: &openRouter.ChatCompletionResponseFormatJSONSchema{
			Name:   "MealItems",
			Strict: true,
			Schema: json.RawMessage(schemaBytes),
		},
	}

	chatCompletionRequest := openRouter.ChatCompletionRequest{
		Model:          model,
		Messages:       messages,
		ResponseFormat: responseFormat,
	}

	fmt.Println("ChatCompletionRequest:", chatCompletionRequest)
	reqBytes, _ := json.MarshalIndent(chatCompletionRequest, "", "  ")
	fmt.Println("Outgoing request JSON:", string(reqBytes))

	response, err := openRouterClient.client.CreateChatCompletion(
		ctx,
		chatCompletionRequest,
	)
	if err != nil {
		// Print out the json response
		responseBytes, _ := json.MarshalIndent(err, "", "  ")
		fmt.Println("Response JSON:", string(responseBytes))
		logger.Logger.Errorf("Error calling OpenRouter API: %v", err)
		return "", err
	}
	// Print out the json response
	responseBytes, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println("Response JSON:", string(responseBytes))
	return response.Choices[0].Message.Content.Text, nil
}
