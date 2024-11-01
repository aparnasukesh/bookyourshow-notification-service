package chat

import (
	"github.com/streadway/amqp"
)

type NewRabbitMqConsumer struct {
	conn *amqp.Connection
	svc  Service
}

func NewRabbitMQConsumer(svc Service, conn *amqp.Connection) *NewRabbitMqConsumer {
	return &NewRabbitMqConsumer{
		conn: conn,
		svc:  svc,
	}
}

func (c *NewRabbitMqConsumer) ChatConsumer() error {
	queue, err := RabbitMQQueue(c.conn, "chat_queue")
	if err != nil {
		return err
	}
	defer queue.ch.Close()
	if err := queue.ch.Qos(1, 0, false); err != nil {
		return err

	}
	msgs, err := MsgReceiver(queue)
	if err != nil {
		return err
	}

	done := make(chan struct{})
	for msg := range msgs {

		out, err := c.svc.chat(msg.Body, done)
		if err != nil {
			return err
		}

		if err := SendMessage(queue, msg, out); err != nil {
			return err
		}
		msg.Ack(false)
	}

	done <- struct{}{}
	return nil
}
