package payment

import (
	"ai-financial-api/config"
	"ai-financial-api/models"
	"errors"
)

func MarkPaymentSuccess(paymentID string, status string) error {
	var payment models.Payment
	if err := config.DB.Where("razorpay_payment_id = ?", paymentID).First(&payment).Error; err != nil {
		return errors.New("payment not found")
	}

	payment.Status = status
	return config.DB.Save(&payment).Error
}
