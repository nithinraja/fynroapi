// =================
// internal/services/question.go
// =================
package services

import (
    "fmt"

    "github.com/google/uuid"
    "fyrno.com/api/fyrnoapi/internal/database"
    "fyrno.com/api/fyrnoapi/internal/models"
)

type QuestionService struct {
    openaiService *OpenAIService
}

func NewQuestionService(openaiService *OpenAIService) *QuestionService {
    return &QuestionService{
        openaiService: openaiService,
    }
}

func (s *QuestionService) ProcessQuestion(req *models.QuestionRequest) (*models.QuestionResponse, error) {
    db := database.GetDB()

    // Generate UUID
    questionID := uuid.New().String()

    // Save question to database
    question := models.Question{
        QuestionID: questionID,
        Question:   req.Question,
        Username:   req.Username,
    }

    if err := db.Create(&question).Error; err != nil {
        return nil, fmt.Errorf("failed to save question: %w", err)
    }

    // Get response from OpenAI
    aiResponse, err := s.openaiService.GetResponse(req.Question)
    if err != nil {
        return nil, fmt.Errorf("failed to get AI response: %w", err)
    }

    // Save response to database
    response := models.Response{
        QuestionID: questionID,
        Response:   aiResponse,
    }

    if err := db.Create(&response).Error; err != nil {
        return nil, fmt.Errorf("failed to save response: %w", err)
    }

    return &models.QuestionResponse{
        ID:         question.ID,
        QuestionID: questionID,
        Question:   req.Question,
        Username:   req.Username,
        Answer:     aiResponse,
        CreatedAt:  question.CreatedAt,
    }, nil
}

func (s *QuestionService) GetQuestionByID(questionID string) (*models.QuestionResponse, error) {
    db := database.GetDB()

    var question models.Question
    if err := db.Where("questionid = ?", questionID).First(&question).Error; err != nil {
        return nil, fmt.Errorf("question not found: %w", err)
    }

    var response models.Response
    if err := db.Where("questionid = ?", questionID).First(&response).Error; err != nil {
        return nil, fmt.Errorf("response not found: %w", err)
    }

    return &models.QuestionResponse{
        ID:         question.ID,
        QuestionID: question.QuestionID,
        Question:   question.Question,
        Username:   question.Username,
        Answer:     response.Response,
        CreatedAt:  question.CreatedAt,
    }, nil
}

