package payment

import (
	"errors"
	"fmt"
	"time"

	"fyrnoapi/model"
	"fyrnoapi/pkg/database"
)

type PaymentService struct {
	Razorpay *RazorpayService
	DB       *database.Database
}

func NewPaymentService(db *database.Database) *PaymentService {
	return &PaymentService{
		Razorpay: NewRazorpayService(),
		DB:       db,
	}
}

// InitiatePayment creates an order in Razorpay and stores the record in DB
func (s *PaymentService) InitiatePayment(userID uint, amount int, currency string, notes string) (map[string]interface{}, error) {
	receipt := fmt.Sprintf("receipt_%d_%d", userID, time.Now().Unix())

	order, err := s.Razorpay.CreateOrder(amount, currency, receipt, notes)
	if err != nil {
		return nil, err
	}

	// Store in DB
	payment := &model.Payment{
		UserID:         userID,
		Amount:         amount,
		Currency:       currency,
		PaymentStatus:  "created",
		RazorpayOrderID: order["id"].(string),
		ReceiptID:       receipt,
		Notes:           notes,
	}

	if err := s.DB.DB.Create(payment).Error; err != nil {
		return nil, err
	}

	return order, nil
}

// ConfirmPayment verifies Razorpay signature and updates DB
func (s *PaymentService) ConfirmPayment(userID uint, orderID, paymentID, signature string) error {
	// Verify signature
	if !s.Razorpay.VerifySignature(orderID, paymentID, signature) {
		return errors.New("invalid Razorpay signature")
	}

	var payment model.Payment
	if err := s.DB.DB.Where("razorpay_order_id = ? AND user_id = ?", orderID, userID).First(&payment).Error; err != nil {
		return err
	}

	payment.PaymentID = paymentID
	payment.Signature = signature
	payment.PaymentStatus = "paid"
	payment.PaidAt = time.Now()

	return s.DB.DB.Save(&payment).Error
}

// GetUserPaymentStatus checks if the user has paid
func (s *PaymentService) GetUserPaymentStatus(userID uint) (bool, error) {
	var payment model.Payment
	err := s.DB.DB.Where("user_id = ? AND payment_status = ?", userID, "paid").First(&payment).Error
	if err != nil {
		return false, nil
	}
	return true, nil
}
