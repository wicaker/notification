package helper

import (
	"context"
	"time"

	"github.com/wicaker/notification/internal/domain"

	"gopkg.in/gomail.v2"
)

type emailSendingHelper struct {
	contextTimeout time.Duration
	dialer         domain.GomailDialer
}

// NewEmailSendingHelper will create new an emailSendingHelper object representation of domain.EmailSendingHelper interface
func NewEmailSendingHelper(timeout time.Duration, dialer domain.GomailDialer) domain.EmailSendingHelper {
	return &emailSendingHelper{
		contextTimeout: timeout,
		dialer:         dialer,
	}
}

func (eu *emailSendingHelper) Send(ctx context.Context, email *domain.EmailSending) error {
	m := gomail.NewMessage()
	m.SetHeader("From", email.EmailFrom)
	m.SetHeader("To", email.EmailTo[0])
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", email.HTMLMessage)

	if email.Attach != "" {
		m.Attach(email.Attach)
	}

	err := eu.dialer.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
