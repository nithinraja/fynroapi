package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/razorpay/razorpay-go"

	"fyrnoapi/api/v1/dto"
	"fyrnoapi/model"
	"fyrnoapi/pkg/database"
	"fyrnoapi/pkg/response"
)

// CreateRazorpayOrder initiates a payment and creates a Razorpay order
func CreateRazorpayOrder(c *gin.Context) {
	var req dto.PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	// Razorpay client
	client := razorpay.NewClient(os.Getenv("RAZORPAY_KEY"), os.Getenv("RAZORPAY_SECRET"))

	// Create Razorpay order
	data := map[string]interface{}{
		"amount":          req.Amount * 100, // Razorpay expects amount in paise
		"currency":        "INR",
		"receipt":         uuid.New().String(),
		"payment_capture": 1,
	}

	order, err := client.Order.Create(data, nil)
	if err != nil {
		response.InternalServerError(c, "Failed to create Razorpay order", err.Error())
		return
	}

	// Store in DB
	payment := model.Payment{
		ID:         uuid.New(),
		UserID:     req.UserID,
		Amount:     float64(req.Amount),
		Status:     "created",
		OrderID:    order["id"].(string),
		CreatedAt:  time.Now(),
	}
	if err := database.DB.Create(&payment).Error; err != nil {
		response.InternalServerError(c, "Failed to store payment", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Order created", gin.H{
		"order_id": order["id"],
		"amount":   req.Amount,
		"currency": "INR",
		"razorpay_key": os.Getenv("RAZORPAY_KEY"),
	})
}

// VerifyPayment handles payment verification (basic post-payment confirmation)
// This endpoint can be secured by Razorpay webhook or callback from frontend
func VerifyPayment(c *gin.Context) {
	var req dto.PaymentVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	var payment model.Payment
	err := database.DB.Where("order_id = ?", req.OrderID).First(&payment).Error
	if err != nil {
		response.NotFound(c, "Payment record not found")
		return
	}

	payment.Status = "success"
	payment.PaymentID = req.PaymentID
	payment.UpdatedAt = time.Now()

	if err := database.DB.Save(&payment).Error; err != nil {
		response.InternalServerError(c, "Failed to update payment", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Payment verified", gin.H{
		"unlocked_features": true,
	})
}
