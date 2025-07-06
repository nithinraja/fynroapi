package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"fyrnoapi/api/v1/dto"
	"fyrnoapi/model"
	"fyrnoapi/pkg/database"
	"fyrnoapi/pkg/response" 
	"fyrnoapi/pkg/utils"
)

// AskQuestion allows anonymous user to submit a question
func AskQuestion(c *gin.Context) {
	var req dto.QuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input", err.Error())
		return
	}

	question := model.Question{
		ID:        uuid.New(),
		Text:      req.Question,
		UUID:      uuid.New().String(),
		CreatedAt: time.Now(),
	}

	if err := database.DB.Create(&question).Error; err != nil {
		response.InternalServerError(c, "Failed to save question", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Question submitted", gin.H{
		"question_id": question.ID,
		"uuid":        question.UUID,
	})
}

// AddNameToQuestion stores the name for a UUID-question and calls OpenAI
func AddNameToQuestion(c *gin.Context) {
	var req dto.QuestionNameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input", err.Error())
		return
	}

	var question model.Question
	if err := database.DB.Where("uuid = ?", req.UUID).First(&question).Error; err != nil {
		response.NotFound(c, "Question not found")
		return
	}

	question.Name = req.Name

	// Call OpenAI to get result (mock or real)
	openaiResp, err := utils.CallOpenAIWithName(req.Name, question.Text)
	if err != nil {
		response.InternalServerError(c, "OpenAI failed", err.Error())
		return
	}

	question.OpenAIResponse = openaiResp
	database.DB.Save(&question)

	response.Success(c, http.StatusOK, "Name saved and result generated", gin.H{
		"response": openaiResp,
	})
}

// GetQuestionsByUser fetches all questions of the authenticated user
func GetQuestionsByUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		response.Unauthorized(c, "User not found in context")
		return
	}

	claims := user.(*token.JWTClaims)

	var questions []model.Question
	if err := database.DB.Where("user_id = ?", claims.UserID).Order("created_at desc").Find(&questions).Error; err != nil {
		response.InternalServerError(c, "Failed to fetch questions", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User questions", questions)
}

// LinkQuestionsToUser updates all anonymous questions with the user ID after login
func LinkQuestionsToUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		response.Unauthorized(c, "User not found in context")
		return
	}
	claims := user.(*token.JWTClaims)

	var req dto.LinkQuestionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input", err.Error())
		return
	}

	if len(req.QuestionUUIDs) == 0 {
		response.BadRequest(c, "No questions to update", nil)
		return
	}

	if err := database.DB.Model(&model.Question{}).
		Where("uuid IN (?) AND user_id IS NULL", req.QuestionUUIDs).
		Update("user_id", claims.UserID).Error; err != nil {
		response.InternalServerError(c, "Failed to link questions", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Questions linked to user", nil)
}
