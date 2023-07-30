package main

import (
	rabbitmq "goeio/rabbitmq"
	"log"
	"os"
)

func consumeMessages(rmq *rabbitmq.RabbitMQBroker) {
	queueName := "goeio"

	msgs, err := rmq.Consume(queueName)
	if err != nil {
		log.Fatal("Failed to start consuming messages:", err)
	}

	log.Println("Consumer started. Waiting for messages...")

	for msg := range msgs {
		message := string(msg.Body)
		log.Printf("Received a message: %s\n", message)

		successful := processMessage(message)

		if successful {
			msg.Ack(false)
		} else {
			msg.Nack(false, true)
		}
	}
}

func processMessage(message string) bool {
	return true
}

func main() {
	amqpURI := os.Getenv("AMQP_URI")
	if amqpURI == "" {
		log.Printf("AMQP_URI environment variable is not set, using defaults.")
		amqpURI = "amqp://guest:guest@rabbitmq:5672/"
	}

	rmq := rabbitmq.NewRabbitMQBroker()
	err := rmq.Connect(amqpURI)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer rmq.Close()

	consumeMessages(rmq)
}
