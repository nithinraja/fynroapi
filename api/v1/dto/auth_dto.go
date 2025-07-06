package dto

// Request to send OTP to a phone number
type SendOTPRequest struct {
	Phone string `json:"phone" binding:"required,e164"` // E.164 format recommended
}

// Response after sending OTP
type SendOTPResponse struct {
	Message string `json:"message"`
}

// Request to verify OTP
type VerifyOTPRequest struct {
	Phone string `json:"phone" binding:"required,e164"`
	Code  string `json:"code" binding:"required,len=6"`
}

// Response after successful OTP verification
type VerifyOTPResponse struct {
	Token   string `json:"token"`   // JWT token
	Message string `json:"message"` // e.g., "OTP verified successfully"
}

// JWT Claims (optional, if needed separately)
type JWTClaims struct {
	UserID string `json:"user_id"`
	Phone  string `json:"phone"`
	Name   string `json:"name,omitempty"`
}


