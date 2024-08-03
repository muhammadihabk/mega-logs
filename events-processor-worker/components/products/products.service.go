package products

import (
	"context"
	"event-processor-worker/config/db"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PersistProduct(delivery amqp.Delivery) {
	err := processMessage(delivery.Body)
	if err != nil {
		log.Printf("Failed to persist product.\n%s Message requeued.\n%v", delivery.Body, err)
		delivery.Nack(false, true)
		return
	}

	delivery.Ack(false)
	log.Printf("Message acknowledged.\n%v ", delivery.Body)
}

func processMessage(data []byte) error {
	db := db.GetDB()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := ""
	_, err := db.ExecContext(ctx, query, string(data))
	if err != nil {
		return err
	}
	return nil
}

func HandleDlxMessages(delivery amqp.Delivery) {
	log.Printf("This is a placeholder for handling DLX messages, maybe alerting, depends on the business. Message:\n%s", delivery.Body)
	delivery.Ack(false)
}
