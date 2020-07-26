package mail

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// NewDialer will return gomail.Dialer configuration
func NewDialer() *gomail.Dialer {
	configSMTPHOST := os.Getenv("CONFIG_SMTP_HOST")
	configSMTPPORT, _ := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	configEmail := os.Getenv("CONFIG_EMAIL")
	configPassword := os.Getenv("CONFIG_PASSWORD")

	dialer := gomail.NewDialer(
		configSMTPHOST,
		configSMTPPORT,
		configEmail,
		configPassword,
	)

	return dialer
}
