package handler

import (
	"fyrnoapi/model"
	"fyrnoapi/pkg/database"
	"fyrnoapi/pkg/response"
	"fyrnoapi/pkg/token"
 
	"github.com/gin-gonic/gin"
)

// GetUserProfile handles GET /api/v1/user/profile
func GetUserProfile(c *gin.Context) {
	userID, exists := token.GetUserIDFromContext(c)
	if !exists {
		response.Unauthorized(c, "Unauthorized")
		return
	}

	var user model.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		response.InternalServerError(c, "User not found")
		return
	}

	response.Success(c, "User profile fetched", gin.H{
		"id":     user.ID,
		"name":   user.Name,
		"mobile": user.Mobile,
	})
}

// GetUserQuestions handles GET /api/v1/user/questions
// Returns a list of previously asked questions by the logged-in user
func GetUserQuestions(c *gin.Context) {
	userID, exists := token.GetUserIDFromContext(c)
	if !exists {
		response.Unauthorized(c, "Unauthorized")
		return
	}

	var questions []model.Question
	if err := database.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&questions).Error; err != nil {
		response.InternalServerError(c, "Failed to fetch questions")
		return
	}

	response.Success(c, "User questions fetched", questions)
}

// (Optional) UpdateUserName handles POST /api/v1/user/update-name
func UpdateUserName(c *gin.Context) {
	userID, exists := token.GetUserIDFromContext(c)
	if !exists {
		response.Unauthorized(c, "Unauthorized")
		return
	}

	var req struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Name is required")
		return
	}

	if err := database.DB.Model(&model.User{}).Where("id = ?", userID).Update("name", req.Name).Error; err != nil {
		response.InternalServerError(c, "Failed to update name")
		return
	}

	response.Success(c, "User name updated", gin.H{"name": req.Name})
}
