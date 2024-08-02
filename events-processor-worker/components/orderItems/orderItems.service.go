package orderItems

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PersistOrderItems(body amqp.Delivery) {
	log.Printf("Received a message: %s", body.Body)
}
