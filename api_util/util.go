package api_util

import (
	"github.com/autom8ter/api/go/api"
	"os"
)

func FromEnv(debug bool, index api.CustomerIndex) *api.Config {
	return &api.Config{
		Debug:         debug,
		TwilioAccount: os.Getenv("TWILIO_ACCOUNT"),
		TwilioKey:     os.Getenv("TWILIO_KEY"),
		SendgridKey:   os.Getenv("SENDGRID_KEY"),
		SlackKey:      os.Getenv("SLACK_KEY"),
		StripeKey:     os.Getenv("STRIPE_KEY"),
		CustomerIndex:         index,
		EmailAddress: &api.EmailAddress{
			Address: os.Getenv("EMAIL_ADDRESS"),
			Name:    os.Getenv("EMAIL_NAME"),
		},
		LogConfig: &api.LogConfig{
			Username: os.Getenv("SLACK_LOG_USERNAME"),
			Channel:  os.Getenv("SLACK_LOG_CHANNEL"),
		},
	}
}
