package transport

import (
	"time"

	"github.com/wicaker/notification/internal/helper"
	"github.com/wicaker/notification/internal/pkg/rmq"
	"github.com/wicaker/notification/internal/usecase"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// LogMqError to log amqp error
func LogMqError(err error, msg *amqp.Delivery, message string) {
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		}).Printf("%s. Message from routingKey: %s rejected...", err, msg.RoutingKey)
		msg.Reject(false)
	} else {
		msg.Ack(true)
	}
}

// Handler /
func Handler(rmqQueue []rmq.Queue) {
	timeoutContext := time.Duration(2) * time.Second

	emailSendingHelper := helper.NewEmailSendingHelper(timeoutContext)
	emailSendingUsecase := usecase.NewEmailSendingUsecase(timeoutContext, emailSendingHelper)
	NewEmailSendingHandler(rmqQueue, emailSendingUsecase)
}
