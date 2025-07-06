// =================
// internal/services/openai.go
// =================
package services

import (
    "context"
    "fmt"

    "github.com/sashabaranov/go-openai"
)

type OpenAIService struct {
    client *openai.Client
}

func NewOpenAIService(apiKey string) *OpenAIService {
    return &OpenAIService{
        client: openai.NewClient(apiKey),
    }
}

func (s *OpenAIService) GetResponse(question string) (string, error) {
    ctx := context.Background()

    req := openai.ChatCompletionRequest{
        Model: openai.GPT3Dot5Turbo,
        Messages: []openai.ChatCompletionMessage{
            {
                Role:    openai.ChatMessageRoleUser,
                Content: question,
            },
        },
        MaxTokens: 1000,
    }

    resp, err := s.client.CreateChatCompletion(ctx, req)
    if err != nil {
        return "", fmt.Errorf("failed to get OpenAI response: %w", err)
    }

    if len(resp.Choices) == 0 {
        return "", fmt.Errorf("no response from OpenAI")
    }

    return resp.Choices[0].Message.Content, nil
}

