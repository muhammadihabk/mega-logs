package sellers

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PersistSeller(body amqp.Delivery) {
	log.Printf("Received a message: %s", body.Body)
}
