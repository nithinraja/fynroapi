package auth

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioService struct {
	Client     *twilio.RestClient
	FromNumber string
}

func NewTwilioService() *TwilioService {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	fromNumber := os.Getenv("TWILIO_PHONE_NUMBER")
	if fromNumber == "" {
		panic("TWILIO_PHONE_NUMBER is not set in environment variables")
	}

	return &TwilioService{
		Client:     client,
		FromNumber: fromNumber,
	}
}

// SendOTP sends an OTP code to the provided phone number
func (s *TwilioService) SendOTP(toPhone string, otpCode string) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(toPhone)
	params.SetFrom(s.FromNumber)
	params.SetBody(fmt.Sprintf("Your OTP code is: %s", otpCode))

	_, err := s.Client.Api.CreateMessage(params)
	return err
}
