package usecase

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/wicaker/notification/internal/domain"
)

type emailSendingUsecase struct {
	contextTimeout time.Duration
	helper         domain.EmailSendingHelper
}

// NewEmailSendingUsecase will create new an emailSendingUsecase object representation of domain.EmailSendingUsecase interface
func NewEmailSendingUsecase(timeout time.Duration, helper domain.EmailSendingHelper) domain.EmailSendingUsecase {
	return &emailSendingUsecase{
		contextTimeout: timeout,
		helper:         helper,
	}
}

func (eu *emailSendingUsecase) RegisterNotif(ctx context.Context, data domain.EmailSending, user domain.User) error {
	// preparing to send email
	emailSending := domain.EmailSending{
		EmailFrom:   os.Getenv("EMAIL_SENDING_FROM"),
		EmailTo:     data.EmailTo,
		Subject:     "Please confirm your email address",
		HTMLMessage: fmt.Sprintf(`<p>%s</p>`, user.Token),
	}

	// send email
	return eu.helper.Send(ctx, &emailSending)
}

func (eu *emailSendingUsecase) ChangePasswordNotif(ctx context.Context, data domain.EmailSending, user domain.User) error {
	// preparing to send email
	emailSending := domain.EmailSending{
		EmailFrom:   os.Getenv("EMAIL_SENDING_FROM"),
		EmailTo:     data.EmailTo,
		Subject:     "Change Password Confirmation",
		HTMLMessage: fmt.Sprintf(`<p>%s</p>`, user.Token),
	}

	// send email
	return eu.helper.Send(ctx, &emailSending)
}

func (eu *emailSendingUsecase) ForgotPasswordNotif(ctx context.Context, data domain.EmailSending, user domain.User) error {
	// preparing to send email
	emailSending := domain.EmailSending{
		EmailFrom:   os.Getenv("EMAIL_SENDING_FROM"),
		EmailTo:     data.EmailTo,
		Subject:     "Forgot Password Confirmation",
		HTMLMessage: fmt.Sprintf(`<p>%s</p>`, user.Token),
	}

	// send email
	return eu.helper.Send(ctx, &emailSending)
}
