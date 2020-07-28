package integration_test

import (
	"os"
	"testing"
	"time"

	"github.com/wicaker/notification/internal/pkg/rmq"
	"github.com/wicaker/notification/internal/transport"
	"github.com/wicaker/notification/test/mock"
)

var (
	messageSent    mock.GomailMessage
	gomailMock     = mock.NewMockGomailDialer(&messageSent)
	timeoutContext = time.Duration(2) * time.Second
	listrmq        []rmq.Queue
	didAck         bool
)

func init() {
	os.Setenv("EMAIL_SENDING_FROM", "wicak_notification@mail.com")
}

func TestMain(m *testing.M) {
	// registering queue or channel
	registerMockQueue()

	// registering transport
	transport.Handler(listrmq, gomailMock)

	code := m.Run()

	os.Exit(code)
}

func registerMockQueue() {
	listrmq = append(listrmq, mock.NewMockQueueRMQ("notification.user.register", &didAck))
	listrmq = append(listrmq, mock.NewMockQueueRMQ("notification.user.change_password", &didAck))
	listrmq = append(listrmq, mock.NewMockQueueRMQ("notification.user.forgot_password", &didAck))
}

func getMessageSent() mock.GomailMessage {
	defer func() {
		messageSent = mock.GomailMessage{}
	}()
	return messageSent
}

func getDidAck() bool {
	defer func() {
		didAck = false
	}()
	return didAck
}
