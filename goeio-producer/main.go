package main

import (
	"goeio/goeio-producer/router"
	rabbitmq "goeio/goeio-rabbitmq"
	"log"
	"os"
)

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

	router := router.SetupRouter(rmq)

	router.Run(":8080")
}
