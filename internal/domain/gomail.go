package domain

import (
	"gopkg.in/gomail.v2"
)

// GomailDialer represent gomail dialer abstraction
type GomailDialer interface {
	DialAndSend(m ...*gomail.Message) error
}
