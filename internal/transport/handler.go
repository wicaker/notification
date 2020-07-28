package transport

import (
	"time"

	"github.com/wicaker/notification/internal/domain"
	"github.com/wicaker/notification/internal/helper"
	"github.com/wicaker/notification/internal/pkg/rmq"
	"github.com/wicaker/notification/internal/usecase"
)

// Handler /
func Handler(rmqQueue []rmq.Queue, dialer domain.GomailDialer) {
	timeoutContext := time.Duration(2) * time.Second

	emailSendingHelper := helper.NewEmailSendingHelper(timeoutContext, dialer)
	emailSendingUsecase := usecase.NewEmailSendingUsecase(timeoutContext, emailSendingHelper)
	NewEmailSendingHandler(rmqQueue, emailSendingUsecase)
}
