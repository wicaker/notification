package config

import (
	"log"
	"os"

	"github.com/streadway/amqp"

	"github.com/wicaker/notification/internal/pkg/rmq"
)

// MqConfig collects all of necessary field for connection configuration to rabbitmq
type MqConfig struct {
	IsClose        bool
	AmqpConnection *amqp.Connection
	ErrorChannel   chan *amqp.Error
	rmqConn        *rmq.Connection
	Queue          []rmq.Queue
}

// NewRabbitmq to start a new configuration to rabbitmq
func NewRabbitmq() *MqConfig {
	config := new(MqConfig)

	config.rmqConn = rmq.NewConnection(os.Getenv("RABBITMQ_SERVER"))
	config.AmqpConnection = config.rmqConn.Dial()
	config.ErrorChannel = config.rmqConn.ErrorChannel(config.AmqpConnection)
	config.rmqConn.NotifyClose(config.AmqpConnection, config.ErrorChannel)
	config.registerQueue()

	return config
}

// Close to close amqp connection
func (c *MqConfig) Close() error {
	log.Println("closing rabbitmq connection...")
	c.IsClose = true
	return c.AmqpConnection.Close()
}

// Reconnect to reconnecting connection
func (c *MqConfig) Reconnect(err error) {
	log.Printf("%s: %s", "Reconnecting after connection closed", err)
	c.AmqpConnection = c.rmqConn.Dial()
	c.ErrorChannel = c.rmqConn.ErrorChannel(c.AmqpConnection)
	c.rmqConn.NotifyClose(c.AmqpConnection, c.ErrorChannel)
	c.registerQueue()
}

func (c *MqConfig) registerQueue() {
	exchange := rmq.Exchange{
		ExcName: "events",
		ExcType: rmq.TOPIC,
	}

	registerChannel := rmq.NewQueue("notification.user.register", c.AmqpConnection, exchange, []string{
		"user.register",
	}, true)
	c.Queue = append(c.Queue, registerChannel)

	changePassworChannel := rmq.NewQueue("notification.user.change_password", c.AmqpConnection, exchange, []string{
		"user.change_password",
	}, true)
	c.Queue = append(c.Queue, changePassworChannel)

	forgotPassworChannel := rmq.NewQueue("notification.user.forgot_password", c.AmqpConnection, exchange, []string{
		"user.forgot_password",
	}, true)
	c.Queue = append(c.Queue, forgotPassworChannel)
}
