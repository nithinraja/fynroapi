package payment

import (
	"errors"
	"fmt"
	"os"

	razorpay "github.com/razorpay/razorpay-go"
)

func Create(questionUUID string, amount float64) (string, error) {
    client := razorpay.NewClient(os.Getenv("RAZORPAY_KEY_ID"), os.Getenv("RAZORPAY_KEY_SECRET"))

    orderData := map[string]interface{}{
        "amount":   int(amount * 100), // Razorpay accepts paise
        "currency": "INR",
        "receipt":  questionUUID,
    }

    order, err := client.Order.Create(orderData, nil)
    if err != nil {
        return "", err
    }

    if order["id"] == nil {
        return "", errors.New("order creation failed")
    }

    paymentLink := fmt.Sprintf("https://checkout.razorpay.com/v1/checkout.js?order_id=%s", order["id"].(string))
    return paymentLink, nil
}
