package products

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PersistProduct(body amqp.Delivery) {
	log.Printf("Received a message: %s", body.Body)
}
