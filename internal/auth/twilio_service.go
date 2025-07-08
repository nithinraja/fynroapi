package auth

import (
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(to, body string) error {
    client := twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: os.Getenv("TWILIO_ACCOUNT_SID"),
        Password: os.Getenv("TWILIO_AUTH_TOKEN"),
    })

    params := &openapi.CreateMessageParams{}
    params.SetTo(to)
    params.SetFrom(os.Getenv("TWILIO_FROM_NUMBER"))
    params.SetBody(body)

    _, err := client.Api.CreateMessage(params)
    return err
}
