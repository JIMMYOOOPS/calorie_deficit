package clients

import (
	"context"

	openRouter "github.com/revrost/go-openrouter"
)

type OpenRouterClient struct {
	client *openRouter.Client
}

func NewOpenRouterClient(apiKey string) *OpenRouterClient {
	return &OpenRouterClient{
		client: openRouter.NewClient(apiKey),
	}
}

func (openRouterClient *OpenRouterClient) CreateChatCompletion(context context.Context, input string) (string, error) {
	response, err := openRouterClient.client.CreateChatCompletion(
		context,
		openRouter.ChatCompletionRequest{
			Model: "qwen/qwen3-0.6b-04-28:free",
			Messages: []openRouter.ChatCompletionMessage{
				{
					Role:    openRouter.ChatMessageRoleUser,
					Content: openRouter.Content{Text: input},
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Message.Content.Text, nil
}
