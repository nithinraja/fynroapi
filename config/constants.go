package config

import "time"

const (
	// App context keys
	ContextUserIDKey = "userID"
	ContextUUIDKey   = "uuid"

	// Token & Security
	JWTTokenExpiry   = time.Hour * 72
	OTPExpiry        = time.Minute * 5
	OTPLength        = 6

	// Roles (if needed in future)
	RoleUser  = "user"
	RoleAdmin = "admin"

	// Question Status
	QuestionPending   = "pending"
	QuestionCompleted = "completed"

	// Payment Status
	PaymentInitiated = "initiated"
	PaymentSuccess   = "success"
	PaymentFailed    = "failed"

	// Recommendation Access
	AccessFree     = "free"
	AccessPremium  = "premium"

	// Messages
	MessageOTPInvalid       = "Invalid or expired OTP"
	MessageUnauthorized     = "Unauthorized access"
	MessagePaymentFailed    = "Payment failed"
	MessageQuestionNotFound = "Question not found"
)
