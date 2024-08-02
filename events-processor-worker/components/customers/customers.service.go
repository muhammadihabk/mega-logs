package customers

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PersistCustomer(body amqp.Delivery) {
	log.Printf("Received a message: %s", body.Body)
}
