// =================
// internal/handlers/question.go
// =================
package handlers

import (
	"net/http"

	"fyrno.com/api/fyrnoapi/internal/models"
	"fyrno.com/api/fyrnoapi/internal/services"
	"fyrno.com/api/fyrnoapi/pkg/utils"
	"github.com/gin-gonic/gin"
)

type QuestionHandler struct {
    questionService *services.QuestionService
}

func NewQuestionHandler(questionService *services.QuestionService) *QuestionHandler {
    return &QuestionHandler{
        questionService: questionService,
    }
}

func (h *QuestionHandler) CreateQuestion(c *gin.Context) {
    var req models.QuestionRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
        return
    }

    response, err := h.questionService.ProcessQuestion(&req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to process question", err)
        return
    }

    utils.SuccessResponse(c, http.StatusCreated, "Question processed successfully", response)
}

func (h *QuestionHandler) GetQuestion(c *gin.Context) {
    questionID := c.Param("questionid")
    if questionID == "" {
        utils.ErrorResponse(c, http.StatusBadRequest, "Question ID is required", nil)
        return
    }

    response, err := h.questionService.GetQuestionByID(questionID)
    if err != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "Question not found", err)
        return
    }

    utils.SuccessResponse(c, http.StatusOK, "Question retrieved successfully", response)
}