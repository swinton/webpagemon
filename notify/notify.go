package notify

import (
	"os"

	twilio "github.com/sfreiberg/gotwilio"
)

// Notify sends notifications
func Notify(to string, message string) (err error) {
	client := twilio.NewTwilioClient(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"))
	_, _, err = client.SendSMS(os.Getenv("TWILIO_SMS_SENDER"), to, message, "", "")
	return err
}
