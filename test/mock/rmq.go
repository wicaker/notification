package mock

import (
	"github.com/wicaker/notification/internal/pkg/rmq"

	"github.com/streadway/amqp"
)

type queue struct {
	name    string
	message chan *string
	didAck  *bool
}

// NewMockQueueRMQ /
func NewMockQueueRMQ(name string, didAck *bool) rmq.Queue {
	return &queue{
		name:    name,
		message: make(chan *string),
		didAck:  didAck,
	}
}

func (q *queue) Consume(consumer rmq.MsgCons) {
	go func() {
		for {
			msg := <-q.message
			delivery := amqp.Delivery{
				Body: []byte(*msg),
			}
			consumer(delivery)
		}
	}()
}

// Publish method for publishing message to rabbitmq
func (q *queue) Publish(message string, routingKey string, headers map[string]interface{}) error {
	q.message <- &message
	return nil
}

// GetQueueName /
func (q *queue) GetQueueName() string {
	return q.name
}

// Acknowledgement /
func (q *queue) Acknowledgement(err error, msg *amqp.Delivery) {
	if err != nil {
		*q.didAck = false
	} else {
		*q.didAck = true
	}
}
