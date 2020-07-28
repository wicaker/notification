package config

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// NewGomailDialer will return gomail.Dialer configuration
func NewGomailDialer() *gomail.Dialer {
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
