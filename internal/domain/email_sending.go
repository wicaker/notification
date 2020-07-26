package domain

import (
	"context"
)

// EmailSending represent structure of email sender
type EmailSending struct {
	EmailFrom   string
	EmailTo     []string
	Subject     string
	HTMLMessage string
	Attach      string
}

// EmailSendingUsecase represent emailsending's usecase contract
type EmailSendingUsecase interface {
	RegisterNotification(context.Context, *EmailSending) error
	ChangePasswordNotification(context.Context, *EmailSending) error
}

// EmailSendingHelper represent emailsending's helper contract
type EmailSendingHelper interface {
	Send(context.Context, *EmailSending) error
}
