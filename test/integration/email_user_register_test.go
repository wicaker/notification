package integration_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserRegisterSuccess(t *testing.T) {
	err := listrmq[0].Publish(`{"email_destination":"test@mail.com","token":"12345.qwerty.54321"}`, "user.register", map[string]interface{}{})
	assert.NoError(t, err)

	time.Sleep(5 * time.Millisecond)
	msg := getMessageSent()
	ack := getDidAck()
	assert.Equal(t, "test@mail.com", msg.Messages[0].To[0])
	assert.Equal(t, "wicak_notification@mail.com", msg.Messages[0].From[0])
	assert.Equal(t, "Please confirm your email address", msg.Messages[0].Subject[0])
	assert.True(t, ack)
}

func TestUserRegisterFailed(t *testing.T) {
	t.Run("Not contain email", func(t *testing.T) {
		err := listrmq[0].Publish(`{"email_destination":"","token":"12345.qwerty.54321"}`, "user.register", map[string]interface{}{})
		assert.NoError(t, err)

		time.Sleep(5 * time.Millisecond)
		msg := getMessageSent()
		ack := getDidAck()
		assert.Empty(t, msg)
		assert.False(t, ack)
	})

	t.Run("Not contain token", func(t *testing.T) {
		err := listrmq[0].Publish(`{"email_destination":"test@mail.com"}`, "user.register", map[string]interface{}{})
		assert.NoError(t, err)

		time.Sleep(5 * time.Millisecond)
		msg := getMessageSent()
		ack := getDidAck()
		assert.Empty(t, msg)
		assert.False(t, ack)
	})
}
