package config

import (
	"os"
	"strconv"

	"github.com/wicaker/notification/internal/domain"

	"gopkg.in/gomail.v2"
)

type gomailDialer struct {
	dialer domain.GomailDialer
}

// NewGomailDialer will return gomail.Dialer configuration
func NewGomailDialer() domain.GomailDialer {
	configSMTPHOST := os.Getenv("GOMAIL_SMTP_HOST")
	configSMTPPORT, _ := strconv.Atoi(os.Getenv("GOMAIL_SMTP_PORT"))
	configEmail := os.Getenv("GOMAIL_EMAIL")
	configPassword := os.Getenv("GOMAIL_PASSWORD")

	dialer := gomail.NewDialer(
		configSMTPHOST,
		configSMTPPORT,
		configEmail,
		configPassword,
	)

	return &gomailDialer{
		dialer: dialer,
	}
}

func (g *gomailDialer) DialAndSend(m ...*gomail.Message) error {
	return g.dialer.DialAndSend(m[:]...)
}
