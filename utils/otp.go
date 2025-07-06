package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"
)

// OTPStore is an in-memory store for OTPs.
// In production, this should be backed by Redis or a persistent cache.
var otpStore = make(map[string]OTPData)

// OTPData holds the OTP and its expiration time.
type OTPData struct {
	OTP       string
	ExpiresAt time.Time
}

// GenerateOTP generates a 6-digit OTP and stores it with an expiration.
func GenerateOTP(mobile string) (string, error) {
	otp, err := generateRandomDigits(6)
	if err != nil {
		return "", err
	}

	otpStore[mobile] = OTPData{
		OTP:       otp,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	return otp, nil
}

// ValidateOTP checks if the OTP is valid and not expired.
func ValidateOTP(mobile, inputOTP string) error {
	data, exists := otpStore[mobile]
	if !exists {
		return errors.New("OTP not found")
	}

	if time.Now().After(data.ExpiresAt) {
		delete(otpStore, mobile)
		return errors.New("OTP expired")
	}

	if data.OTP != inputOTP {
		return errors.New("Invalid OTP")
	}

	delete(otpStore, mobile) // remove on success
	return nil
}

// generateRandomDigits generates a random numeric OTP of specified length.
func generateRandomDigits(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	digits := ""
	for _, v := range b {
		digits += fmt.Sprintf("%d", (v % 10))
	}

	return digits[:length], nil
}

// GenerateTokenID generates a base64 token ID (for future use if needed).
func GenerateTokenID() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
