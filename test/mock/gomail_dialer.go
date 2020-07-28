package mock

import (
	"reflect"
	"unsafe"

	"github.com/wicaker/notification/internal/domain"
	"gopkg.in/gomail.v2"
)

// GomailMessage /
type GomailMessage struct {
	Messages []Message
}

// Message /
type Message struct {
	From    []string
	Subject []string
	To      []string
}

type mockGomailDialer struct {
	GomailMsg *GomailMessage
}

// NewMockGomailDialer will return gomail.Dialer configuration
func NewMockGomailDialer(message *GomailMessage) domain.GomailDialer {
	return &mockGomailDialer{
		GomailMsg: message,
	}
}

func (mock *mockGomailDialer) DialAndSend(m ...*gomail.Message) error {
	msg := GomailMessage{}
	for i := range m {
		var header map[string][]string
		rs := reflect.ValueOf(*m[i])
		rs2 := reflect.New(rs.Type()).Elem()
		rs2.Set(rs)

		rf := rs2.Field(0)
		rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()

		ri := reflect.ValueOf(&header).Elem() // i, but writeable
		ri.Set(rf)

		msg.Messages = append(msg.Messages, Message{
			From:    header["From"],
			Subject: header["Subject"],
			To:      header["To"],
		})
	}
	*mock.GomailMsg = GomailMessage{
		Messages: msg.Messages,
	}
	return nil
}
