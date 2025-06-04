package rabbitmq

import (
	"github.com/streadway/amqp"
)

// func NewRabbitMQConnection() (*amqp.Connection, error) {
// 	return amqp.Dial("amqp://guest:guest@localhost:5672/")
// }

func NewRabbitMQConnection() (*amqp.Connection, error) {
	return amqp.Dial("amqp://admin:admin123@rabbitmq:5672/") // Use updated credentials
}
