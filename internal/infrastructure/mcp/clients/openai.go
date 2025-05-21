package clients

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
	client *openai.Client
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{
		client: openai.NewClient(apiKey),
	}
}

func (oac *OpenAIClient) CreateChatCompletion(ctx context.Context, systemRoleInput, userRoleInput string) (string, error) {
	response, err := oac.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: "gpt-3.5-turbo",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemRoleInput,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userRoleInput,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Message.Content, nil
}
