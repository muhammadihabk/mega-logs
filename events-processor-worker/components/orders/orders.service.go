package orders

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PersistOrder(body amqp.Delivery) {
	log.Printf("Received a message: %s", body.Body)
}
