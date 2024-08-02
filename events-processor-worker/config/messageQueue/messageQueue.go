package messageQueue

import (
	"event-processor-worker/utilities"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	connection *amqp.Connection
	Channel    *amqp.Channel
)

func init() {
	var err error

	for i := 0; i < 5; i++ {
		connection, err = amqp.Dial(os.Getenv("amqpUri"))
		if err == nil {
			break
		}
		log.Printf("Failed to connect to RabbitMQ, retrying in 5 seconds... (%d/5)\n", i+1)
		time.Sleep(5 * time.Second)
	}
	utilities.ErrorHandler(err, "Failed to connect to the message queue")

	channel, err := connection.Channel()
	utilities.ErrorHandler(err, "Failed to create a channel")

	Channel = channel
}

func CleanupOnExit() {
	defer connection.Close()
	defer Channel.Close()
}
