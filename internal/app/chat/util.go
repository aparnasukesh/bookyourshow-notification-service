package chat

import "github.com/streadway/amqp"

type RabbitMQQueueChannel struct {
	ch *amqp.Channel
	*amqp.Queue
}

func RabbitMQQueue(conn *amqp.Connection, queueName string) (*RabbitMQQueueChannel, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &RabbitMQQueueChannel{
		ch:    ch,
		Queue: &q,
	}, nil
}

func MsgReceiver(queue *RabbitMQQueueChannel) (<-chan amqp.Delivery, error) {
	return queue.ch.Consume(
		queue.Name,
		"",
		false, // set to false for manual acknowledgement
		false,
		false,
		false,
		nil,
	)
}

func SendMessage(queue *RabbitMQQueueChannel, msg amqp.Delivery, response string) error {
	return queue.ch.Publish(
		"",
		msg.ReplyTo,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			Body:          []byte(response),
			CorrelationId: msg.CorrelationId,
		},
	)
}
