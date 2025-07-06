package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// TwilioConfig stores credentials and settings for Twilio
type TwilioConfig struct {
	AccountSID string
	AuthToken  string
	FromNumber string
}

// LoadTwilioConfig loads Twilio credentials from environment variables
func LoadTwilioConfig() (*TwilioConfig, error) {
	sid := os.Getenv("TWILIO_ACCOUNT_SID")
	token := os.Getenv("TWILIO_AUTH_TOKEN")
	from := os.Getenv("TWILIO_PHONE_NUMBER")

	if sid == "" || token == "" || from == "" {
		return nil, errors.New("Twilio credentials not set in environment")
	}

	return &TwilioConfig{
		AccountSID: sid,
		AuthToken:  token,
		FromNumber: from,
	}, nil
}

// SendOTP sends an OTP SMS using Twilio API
func SendOTP(mobile, otp string) error {
	cfg, err := LoadTwilioConfig()
	if err != nil {
		return err
	}

	// Format phone number with country code, example: "+91" for India
	to := formatPhoneNumber(mobile)

	message := fmt.Sprintf("Your OTP is: %s. It will expire in 5 minutes.", otp)

	data := url.Values{}
	data.Set("To", to)
	data.Set("From", cfg.FromNumber)
	data.Set("Body", message)

	reqURL := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", cfg.AccountSID)
	req, err := http.NewRequest("POST", reqURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}

	req.SetBasicAuth(cfg.AccountSID, cfg.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	var resBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&resBody)

	return fmt.Errorf("Twilio error: %v", resBody["message"])
}

// formatPhoneNumber formats a raw mobile number to E.164 format.
// You can modify this logic depending on your target country.
func formatPhoneNumber(mobile string) string {
	if len(mobile) == 10 {
		return "+91" + mobile // default to India if not passed
	}
	return mobile
}
