package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/wicaker/notification/internal/domain"
	"github.com/wicaker/notification/internal/pkg/rmq"

	"github.com/streadway/amqp"
)

// emailSendingHandler represent the email sending handler
type emailSendingHandler struct {
	EmailSendingUsecase     domain.EmailSendingUsecase
	queueUserRegister       rmq.Queue
	queueUserChangePassword rmq.Queue
	queueUserForgotPassword rmq.Queue
}

// NewEmailSendingHandler will initialize handler
func NewEmailSendingHandler(rmqQueue []rmq.Queue, u domain.EmailSendingUsecase) {
	handler := new(emailSendingHandler)
	handler.EmailSendingUsecase = u
	ctx := context.Background()

	for _, rmqQ := range rmqQueue {
		switch name := rmqQ.GetQueueName(); name {
		case "notification.user.register":
			handler.queueUserRegister = rmqQ
		case "notification.user.change_password":
			handler.queueUserChangePassword = rmqQ
		case "notification.user.forgot_password":
			handler.queueUserForgotPassword = rmqQ
		}
	}

	handler.queueUserRegister.Consume(func(msg amqp.Delivery) {
		err := handler.Register(ctx, msg.Body)
		LogMqError(err, &msg, "")
		return
	})

	handler.queueUserChangePassword.Consume(func(msg amqp.Delivery) {
		err := handler.ChangePassword(ctx, msg.Body)
		LogMqError(err, &msg, "")
		return
	})

	handler.queueUserForgotPassword.Consume(func(msg amqp.Delivery) {
		err := handler.ForgotPassword(ctx, msg.Body)
		LogMqError(err, &msg, "")
		return
	})
}

func (e *emailSendingHandler) Register(ctx context.Context, msg []byte) error {
	userMessage, err := e.validateUserMessage(msg)
	if err != nil {
		return fmt.Errorf("Validation error... %s", err)
	}

	err = e.EmailSendingUsecase.RegisterNotif(ctx, userMessage)
	if err != nil {
		return fmt.Errorf("Failed send email... %s", err)
	}

	return nil
}

func (e *emailSendingHandler) ChangePassword(ctx context.Context, msg []byte) error {
	userMessage, err := e.validateUserMessage(msg)
	if err != nil {
		return fmt.Errorf("Validation error... %s", err)
	}

	err = e.EmailSendingUsecase.ChangePasswordNotif(ctx, userMessage)
	if err != nil {
		return fmt.Errorf("Failed send email... %s", err)
	}

	return nil
}

func (e *emailSendingHandler) ForgotPassword(ctx context.Context, msg []byte) error {
	userMessage, err := e.validateUserMessage(msg)
	if err != nil {
		return fmt.Errorf("Validation error... %s", err)
	}

	err = e.EmailSendingUsecase.ChangePasswordNotif(ctx, userMessage)
	if err != nil {
		return fmt.Errorf("Failed send email... %s", err)
	}

	return nil
}

func (e *emailSendingHandler) validateUserMessage(msg []byte) (*domain.User, error) {
	var message domain.User
	if err := json.Unmarshal(msg, &message); err != nil {
		return nil, err
	}

	// message validation
	switch {
	case message.EmailDestination == "":
		return nil, errors.New("`emailDestination` was missing")
	case message.Token == "":
		return nil, errors.New("`token` was missing")
	}

	return &message, nil
}
