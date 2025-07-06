package openai

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIService struct {
	Client         *openai.Client
	PromptBuilder  *PromptBuilder
	DefaultModel   string
	MaxTokenLength int
}

// NewOpenAIService initializes and returns a new OpenAI service
func NewOpenAIService() *OpenAIService {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY not set in environment")
	}

	client := openai.NewClient(apiKey)

	return &OpenAIService{
		Client:         client,
		PromptBuilder:  NewPromptBuilder(),
		DefaultModel:   openai.GPT3Dot5Turbo, // or "gpt-4"
		MaxTokenLength: 1000,
	}
}

// GetInitialResponse gets a response for the user's question (pre-login)
func (s *OpenAIService) GetInitialResponse(name, question string) (string, error) {
	prompt := s.PromptBuilder.BuildInitialPrompt(name, question)
	return s.callOpenAI(prompt)
}

// GetDetailedAnalysis provides an in-depth answer after login
func (s *OpenAIService) GetDetailedAnalysis(name, question string) (string, error) {
	prompt := s.PromptBuilder.BuildAnalysisPrompt(name, question)
	return s.callOpenAI(prompt)
}

// GetSuggestions returns personalized recommendations based on past questions
func (s *OpenAIService) GetSuggestions(userName string, questions []string) (string, error) {
	prompt := s.PromptBuilder.BuildSuggestionPrompt(userName, questions)
	return s.callOpenAI(prompt)
}

// callOpenAI sends the constructed prompt to OpenAI and returns the response
func (s *OpenAIService) callOpenAI(prompt string) (string, error) {
	resp, err := s.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: s.DefaultModel,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful financial assistant providing expert advice in simple terms.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: s.MaxTokenLength,
		},
	)

	if err != nil {
		return "", fmt.Errorf("OpenAI API error: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}
