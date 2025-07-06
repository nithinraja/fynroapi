package dto

// Request to update or register a user (e.g., after OTP verification or name submission)
type RegisterUserRequest struct {
	Name  string `json:"name" binding:"required,min=2,max=100"`
	Phone string `json:"phone" binding:"required,e164"` // Example: +919876543210
}

// Response after user registration or lookup
type UserResponse struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Token  string `json:"token,omitempty"` // JWT token (optional)
}

// Basic user profile response
type UserProfileResponse struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	JoinedAt    string `json:"joined_at"`
	IsPremium   bool   `json:"is_premium"`
	QuestionCount int  `json:"question_count"`
}

// Request to update profile (if needed later)
type UpdateProfileRequest struct {
	Name string `json:"name" binding:"omitempty,min=2,max=100"`
}

// Optional logout request (not mandatory for JWT but useful for tracking)
type LogoutRequest struct {
	UserID string `json:"user_id" binding:"required"`
}
