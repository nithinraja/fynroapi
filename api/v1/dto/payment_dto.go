package dto

type PaymentRequest struct {
	UserID string  `json:"user_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

type PaymentVerifyRequest struct {
	OrderID   string `json:"order_id" binding:"required"`
	PaymentID string `json:"payment_id" binding:"required"`
}

// Request to initiate Razorpay payment order
type CreatePaymentRequest struct {
	Amount int64  `json:"amount" binding:"required,min=100"` // Amount in paise (e.g., ₹100 = 10000)
	Currency string `json:"currency" binding:"required"`     // e.g., "INR"
}

// Response to return order details from Razorpay
type CreatePaymentResponse struct {
	OrderID string `json:"order_id"`
	Currency string `json:"currency"`
	Amount int64    `json:"amount"`
	Key     string  `json:"key"` // Razorpay public key for frontend
}

// Request to confirm and verify payment after success
type ConfirmPaymentRequest struct {
	OrderID   string `json:"order_id" binding:"required"`
	PaymentID string `json:"payment_id" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

// Response after successful payment confirmation
type ConfirmPaymentResponse struct {
	Message   string `json:"message"`
	Status    string `json:"status"` // "success" or "failed"
	Access    string `json:"access"` // e.g., "full"
	PaymentID string `json:"payment_id"`
}
