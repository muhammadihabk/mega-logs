package customers

import (
	"context"
	"encoding/json"
	"event-processor-worker/config/db"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Customer struct {
	CustomerKey           string `json:"customer_key"`
	CustomerZipCodePrefix string `json:"customer_zip_code_prefix"`
	CustomerCity          string `json:"customer_city"`
	CustomerState         string `json:"customer_state"`
}

func processMessage(data []byte) error {
	var customer Customer
	if err := json.Unmarshal(data, &customer); err != nil {
		return err
	}

	db := db.GetDB()
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := `
		INSERT INTO customers (
			customer_key,
			customer_zip_code_prefix,
			customer_city,
			customer_state
		)
		VALUES (?, ?, ?, ?)
	`
	_, err := db.ExecContext(ctx, query, customer.CustomerKey,
		customer.CustomerZipCodePrefix,
		customer.CustomerCity,
		customer.CustomerState)
	if err != nil {
		return err
	}
	return nil
}

func handleDlxMessages(delivery amqp.Delivery) {
	log.Printf("This is a placeholder for handling DLX messages, maybe alerting, depends on the business. Message:\n%s", delivery.Body)
	delivery.Ack(false)
}
