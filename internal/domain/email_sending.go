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
	RegisterNotif(context.Context, *User) error
	ChangePasswordNotif(context.Context, *User) error
	ForgotPasswordNotif(context.Context, *User) error
}

// EmailSendingHelper represent emailsending's helper contract
type EmailSendingHelper interface {
	Send(context.Context, *EmailSending) error
}
