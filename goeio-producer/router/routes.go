package router

import (
	rabbitmq "goeio/goeio-rabbitmq"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter(rmq *rabbitmq.RabbitMQBroker) *gin.Engine {
	router := gin.Default()

	router.Use(LoggerMiddleware())

	queueName := "goeio"
	queue, err := rmq.DeclareQueue(queueName)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	router.POST("/send", func(c *gin.Context) {
		var request struct {
			Message string `json:"message"`
		}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}

		err = rmq.SendMessage(queue, request.Message)
		if err != nil {
			log.Println("Failed to send a message:", err)
			c.JSON(500, gin.H{"error": "failed to send message"})
			return
		}

		c.JSON(200, gin.H{"status": "message sent"})
	})

	return router
}
