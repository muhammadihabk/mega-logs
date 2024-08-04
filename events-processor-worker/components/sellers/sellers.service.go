package sellers

import (
	"context"
	"encoding/json"
	"event-processor-worker/config/db"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Seller struct {
	SellerKey           string `json:"seller_key"`
	SellerZipCodePrefix string `json:"seller_zip_code_prefix"`
	SellerCity          string `json:"seller_city"`
	SellerState         string `json:"seller_state"`
}

func PersistSeller(delivery amqp.Delivery) {
	err := processMessage(delivery.Body)
	if err != nil {
		log.Printf("Failed to persist seller.\n%s Message requeued.\n%v", delivery.Body, err)
		delivery.Nack(false, true)
		return
	}

	delivery.Ack(false)
	log.Printf("Message acknowledged.\n%v ", delivery.Body)
}

func processMessage(data []byte) error {
	db := db.GetDB()

	var seller Seller
	if err := json.Unmarshal(data, &seller); err != nil {
		return err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := `
		INSERT INTO sellers (
			seller_key,
			seller_zip_code_prefix,
			seller_city,
			seller_state
		)
		VALUES (?, ?, ?, ?)
	`

	_, err := db.ExecContext(ctx, query,
		seller.SellerKey,
		seller.SellerZipCodePrefix,
		seller.SellerCity,
		seller.SellerState,
	)
	if err != nil {
		return err
	}
	return nil
}

func HandleDlxMessages(delivery amqp.Delivery) {
	log.Printf("This is a placeholder for handling DLX messages, maybe alerting, depends on the business. Message:\n%s", delivery.Body)
	delivery.Ack(false)
}
