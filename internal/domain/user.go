package domain

// User represent structure of rabbitmq message from user service
type User struct {
	EmailDestination string `json:"email_destination"`
	Token            string `json:"token"`
}
