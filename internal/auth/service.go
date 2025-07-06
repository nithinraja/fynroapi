package auth

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"fyrnoapi/config"
	"fyrnoapi/model"
	"fyrnoapi/pkg/token"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		DB: db,
	}
}

// RequestOTP generates and sends OTP to the given mobile number
func (s *AuthService) RequestOTP(mobile string) error {
	// Generate OTP
	code := otp.GenerateOTP(config.OTPLength)

	// Store OTP in database
	otpRecord := model.OTP{
		Mobile:    mobile,
		Code:      code,
		ExpiresAt: time.Now().Add(config.OTPExpiry),
	}

	if err := s.DB.Create(&otpRecord).Error; err != nil {
		return fmt.Errorf("failed to save OTP: %w", err)
	}

	// Send via Twilio
	return twilio.SendOTP(mobile, code)
}

// VerifyOTP checks the submitted OTP and issues JWT token
func (s *AuthService) VerifyOTP(mobile, code string) (string, error) {
	var otpRecord model.OTP

	err := s.DB.Where("mobile = ? AND code = ?", mobile, code).
		Order("created_at DESC").First(&otpRecord).Error

	if errors.Is(err, gorm.ErrRecordNotFound) || otpRecord.ExpiresAt.Before(time.Now()) {
		return "", errors.New(config.MessageOTPInvalid)
	}

	// Find or create user
	var user model.User
	err = s.DB.Where("mobile = ?", mobile).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = model.User{
			Mobile: mobile,
		}
		if err := s.DB.Create(&user).Error; err != nil {
			return "", fmt.Errorf("failed to create user: %w", err)
		}
	}

	// Generate JWT
	jwtToken, err := token.GenerateJWT(user.ID, config.JWTTokenExpiry)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return jwtToken, nil
}
