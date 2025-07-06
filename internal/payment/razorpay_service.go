package payment

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/razorpay/razorpay-go"
)

type RazorpayService struct {
	Client      *razorpay.Client
	KeySecret   string
	CallbackURL string
}

func NewRazorpayService() *RazorpayService {
	keyID := os.Getenv("RAZORPAY_KEY_ID")
	keySecret := os.Getenv("RAZORPAY_KEY_SECRET")
	callbackURL := os.Getenv("RAZORPAY_CALLBACK_URL")

	if keyID == "" || keySecret == "" {
		panic("RAZORPAY_KEY_ID or RAZORPAY_KEY_SECRET is not set in environment")
	}

	client := razorpay.NewClient(keyID, keySecret)

	return &RazorpayService{
		Client:      client,
		KeySecret:   keySecret,
		CallbackURL: callbackURL,
	}
}

// CreateOrder creates a new payment order in Razorpay
func (r *RazorpayService) CreateOrder(amount int, currency, receipt, notes string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"amount":          amount * 100, // amount in paise
		"currency":        currency,
		"receipt":         receipt,
		"payment_capture": 1,
		"notes": map[string]string{
			"description": notes,
		},
	}

	order, err := r.Client.Order.Create(data, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create razorpay order: %v", err)
	}

	return order, nil
}

// VerifySignature checks if the payment signature is valid
func (r *RazorpayService) VerifySignature(orderID, paymentID, providedSignature string) bool {
	data := orderID + "|" + paymentID
	h := hmac.New(sha256.New, []byte(r.KeySecret))
	h.Write([]byte(data))
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	return hmac.Equal([]byte(expectedSignature), []byte(providedSignature))
}
