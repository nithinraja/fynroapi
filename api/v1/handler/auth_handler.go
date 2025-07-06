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
	"fyrnoapi/pkg/token"
	"fyrnoapi/pkg/utils"
)

// SendOTP handles generating OTP and sending it to the user
func SendOTP(c *gin.Context) {
	var req dto.SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input", err.Error())
		return
	}

	if req.Phone == ""  {
		response.BadRequest(c, "Phone and Name are required", nil)
		return
	}

	// Check if user already exists
	var user model.User
	result := database.DB.Where("phone = ?", req.Phone).First(&user)

	if result.Error != nil {
		// Create new user if not exists
		user = model.User{
			ID:    uuid.New(), 
			Phone: req.Phone,
		}
		if err := database.DB.Create(&user).Error; err != nil {
			response.InternalServerError(c, "Failed to create user", err.Error())
			return
		}
	}

	// Generate OTP
	otp := utils.GenerateOTP(6)

	// Send OTP via Twilio
	err := utils.SendOTPViaTwilio(req.Phone, otp)
	if err != nil {
		response.InternalServerError(c, "Failed to send OTP", err.Error())
		return
	}

	// Store OTP in session table
	session := model.Session{
		ID:        uuid.New(),
		UserID:    user.ID,
		Phone:     req.Phone,
		OTP:       otp,
		IsUsed:    false,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	if err := database.DB.Create(&session).Error; err != nil {
		response.InternalServerError(c, "Failed to store session", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "OTP sent successfully", gin.H{
		"user_id": user.ID,
	})
}

// VerifyOTP handles OTP validation and JWT token generation
func VerifyOTP(c *gin.Context) {
	var req dto.VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input", err.Error())
		return
	}

	var session model.Session
	err := database.DB.
		Where("phone = ? AND otp = ? AND is_used = false AND expires_at >= ?", req.Phone, req.OTP, time.Now()).
		First(&session).Error

	if err != nil {
		response.Unauthorized(c, "Invalid or expired OTP")
		return
	}

	// Mark session as used
	session.IsUsed = true
	database.DB.Save(&session)

	// Generate JWT
	tokenStr, err := token.GenerateToken(session.UserID.String(), session.Phone, time.Hour*24)
	if err != nil {
		response.InternalServerError(c, "Failed to generate token", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "OTP verified", gin.H{
		"token": tokenStr,
	})
}
