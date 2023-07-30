package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQBroker struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQBroker() *RabbitMQBroker {
	return &RabbitMQBroker{}
}

func (r *RabbitMQBroker) Connect(amqpURI string) error {
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return err
	}
	r.conn = conn

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	r.ch = ch

	return nil
}

func (r *RabbitMQBroker) Close() {
	r.ch.Close()
	r.conn.Close()
}

func (r *RabbitMQBroker) DeclareQueue(queueName string) (string, error) {
	q, err := r.ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return "", err
	}

	return q.Name, nil
}

func (r *RabbitMQBroker) SendMessage(queueName, message string) error {
	err := r.ch.PublishWithContext(
		context.Background(),
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	return err
}

func (r *RabbitMQBroker) Consume(queueName string) (<-chan amqp.Delivery, error) {
	return r.ch.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}
