package rabbitmq

type MessageBroker interface {
	Connect(amqpURI string) error
	DeclareQueue(queueName string) (string, error)
	SendMessage(queueName, message string) error
	Close()
}
