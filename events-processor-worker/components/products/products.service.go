package products

import (
	"context"
	"encoding/json"
	"event-processor-worker/config/db"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Product struct {
	ProductKey               string      `json:"product_key"`
	ProductCategoryName      string      `json:"product_category_name"`
	ProductNameLength        json.Number `json:"product_name_length"`
	ProductDescriptionLength json.Number `json:"product_description_length"`
	ProductPhotosQty         json.Number `json:"product_photos_qty"`
	ProductWeightG           json.Number `json:"product_weight_g"`
	ProductLengthCm          json.Number `json:"product_length_cm"`
	ProductHeightCm          json.Number `json:"product_height_cm"`
	ProductWidthCm           json.Number `json:"product_width_cm"`
}

func processMessage(data []byte) error {
	db := db.GetDB()

	var product Product
	if err := json.Unmarshal(data, &product); err != nil {
		return err
	}
	productNameLength, _ := product.ProductNameLength.Int64()
	productDescriptionLength, _ := product.ProductDescriptionLength.Int64()
	productPhotosQty, _ := product.ProductPhotosQty.Int64()
	productWeightG, _ := product.ProductWeightG.Int64()
	productLengthCm, _ := product.ProductLengthCm.Int64()
	productHeightCm, _ := product.ProductHeightCm.Int64()
	productWidthCm, _ := product.ProductWidthCm.Int64()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := `
		INSERT INTO products (
			product_key,
			product_category_name,
			product_name_length,
			product_description_length,
			product_photos_qty,
			product_weight_g,
			product_length_cm,
			product_height_cm,
			product_width_cm
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.ExecContext(ctx, query,
		product.ProductKey,
		product.ProductCategoryName,
		productNameLength,
		productDescriptionLength,
		productPhotosQty,
		productWeightG,
		productLengthCm,
		productHeightCm,
		productWidthCm,
	)
	if err != nil {
		return err
	}
	return nil
}

func handleDlxMessages(delivery amqp.Delivery) {
	log.Printf("This is a placeholder for handling DLX messages, maybe alerting, depends on the business. Message:\n%s", delivery.Body)
	delivery.Ack(false)
}
