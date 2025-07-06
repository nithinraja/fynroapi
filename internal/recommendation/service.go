package recommendation

import (
	"fmt"

	"fyrnoapi/internal/openai"
	"fyrnoapi/model"
	"fyrnoapi/pkg/database"
)

type RecommendationService struct {
	DB           *database.Database
	OpenAIClient *openai.OpenAIService
}

func NewRecommendationService(db *database.Database, openaiClient *openai.OpenAIService) *RecommendationService {
	return &RecommendationService{
		DB:           db,
		OpenAIClient: openaiClient,
	}
}

// GetRecommendations generates a list of suggestions from OpenAI for the logged-in user
func (s *RecommendationService) GetRecommendations(userID uint, userName string) ([]string, error) {
	var questions []model.Question
	if err := s.DB.DB.Where("user_id = ?", userID).Find(&questions).Error; err != nil {
		return nil, err
	}

	if len(questions) == 0 {
		return nil, fmt.Errorf("no questions found for user")
	}

	var prompts []string
	for _, q := range questions {
		prompts = append(prompts, q.Question)
	}

	// Use OpenAI service to get recommendation string
	recommendationText, err := s.OpenAIClient.GetSuggestions(userName, prompts)
	if err != nil {
		return nil, err
	}

	// Optionally split string into slice based on formatting
	recommendations := parseSuggestionsToList(recommendationText)

	return recommendations, nil
}

// Helper to split OpenAI text into a slice, depending on format
func parseSuggestionsToList(text string) []string {
	var list []string
	lines := splitLines(text)
	for _, line := range lines {
		if line != "" {
			list = append(list, line)
		}
	}
	return list
}

func splitLines(text string) []string {
	var lines []string
	start := 0
	for i, r := range text {
		if r == '\n' || r == '\r' {
			lines = append(lines, text[start:i])
			start = i + 1
		}
	}
	if start < len(text) {
		lines = append(lines, text[start:])
	}
	return lines
}
