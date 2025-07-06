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

// GenerateRecommendations - Call OpenAI and save recommendations
func GenerateRecommendations(c *gin.Context) {
	userCtx, exists := c.Get("user")
	if !exists {
		response.Unauthorized(c, "Unauthorized")
		return
	}
	claims := userCtx.(*token.JWTClaims)

	var req dto.RecommendationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input", err.Error())
		return
	}

	// Call OpenAI with prompt based on previous response
	result, err := utils.CallOpenAIRecommendations(claims.Name, req.PreviousResponse)
	if err != nil {
		response.InternalServerError(c, "OpenAI failed", err.Error())
		return
	}

	recommendation := model.Recommendation{
		ID:         uuid.New(),
		UserID:     claims.UserID,
		Prompt:     req.PreviousResponse,
		Response:   result,
		CreatedAt:  time.Now(),
	}

	if err := database.DB.Create(&recommendation).Error; err != nil {
		response.InternalServerError(c, "Failed to save recommendations", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Recommendations generated", gin.H{
		"recommendations": result,
	})
}

// GetRecommendations - List recommendations for a logged-in user
func GetRecommendations(c *gin.Context) {
	userCtx, exists := c.Get("user")
	if !exists {
		response.Unauthorized(c, "Unauthorized")
		return
	}
	claims := userCtx.(*token.JWTClaims)

	// Ensure user has paid
	var payment model.Payment
	err := database.DB.Where("user_id = ? AND status = ?", claims.UserID, "success").First(&payment).Error
	if err != nil {
		response.Forbidden(c, "Payment required to view recommendations")
		return
	}

	var recs []model.Recommendation
	if err := database.DB.Where("user_id = ?", claims.UserID).Order("created_at desc").Find(&recs).Error; err != nil {
		response.InternalServerError(c, "Failed to fetch recommendations", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Fetched recommendations", recs)
}
