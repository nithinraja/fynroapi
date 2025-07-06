package question

import (
	"fmt"
	"time"

	"fyrnoapi/internal/openai"
	"fyrnoapi/model"
	"fyrnoapi/pkg/database"
)

type QuestionService struct {
	DB           *database.Database
	OpenAIClient *openai.OpenAIService
}

func NewQuestionService(db *database.Database, openaiClient *openai.OpenAIService) *QuestionService {
	return &QuestionService{
		DB:           db,
		OpenAIClient: openaiClient,
	}
}

// AskQuestion stores the question and returns OpenAI's initial response
func (s *QuestionService) AskQuestion(sessionUUID, questionText string) (string, error) {
	question := &model.Question{
		SessionUUID: sessionUUID,
		Question:    questionText,
		CreatedAt:   time.Now(),
	}

	if err := s.DB.DB.Create(question).Error; err != nil {
		return "", err
	}

	// Initial anonymous OpenAI response (name will be empty)
	response, err := s.OpenAIClient.GetInitialResponse("User", questionText)
	if err != nil {
		return "", err
	}

	question.InitialResponse = response
	s.DB.DB.Save(question)

	return response, nil
}

// AddUserDetails links a name to questions in the session and generates detailed response
func (s *QuestionService) AddUserDetails(sessionUUID, name string) (string, error) {
	var questions []model.Question
	if err := s.DB.DB.Where("session_uuid = ?", sessionUUID).Find(&questions).Error; err != nil {
		return "", err
	}

	if len(questions) == 0 {
		return "", fmt.Errorf("no questions found for session")
	}

	// Update name for existing questions
	for i := range questions {
		questions[i].Name = name
		s.DB.DB.Save(&questions[i])
	}

	// Get detailed response for the last question
	latest := questions[len(questions)-1]
	resp, err := s.OpenAIClient.GetDetailedAnalysis(name, latest.Question)
	if err != nil {
		return "", err
	}

	latest.DetailedResponse = resp
	s.DB.DB.Save(&latest)

	return resp, nil
}

// LinkQuestionsToUser connects all session questions to a logged-in user
func (s *QuestionService) LinkQuestionsToUser(sessionUUID string, userID uint) error {
	return s.DB.DB.Model(&model.Question{}).
		Where("session_uuid = ?", sessionUUID).
		Update("user_id", userID).Error
}

// GetUserSuggestions gets recommendations from OpenAI based on user's history
func (s *QuestionService) GetUserSuggestions(userID uint, userName string) ([]string, error) {
	var questions []model.Question
	if err := s.DB.DB.Where("user_id = ?", userID).Find(&questions).Error; err != nil {
		return nil, err
	}

	if len(questions) == 0 {
		return nil, fmt.Errorf("no questions found for user")
	}

	var questionTexts []string
	for _, q := range questions {
		questionTexts = append(questionTexts, q.Question)
	}

	response, err := s.OpenAIClient.GetSuggestions(userName, questionTexts)
	if err != nil {
		return nil, err
	}

	// You may choose to parse this string response into list format if needed.
	// For now, we return a single-item list containing the full response.
	return []string{response}, nil
}
