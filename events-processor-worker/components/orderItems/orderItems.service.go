package orderItems

import (
	"context"
	"encoding/json"
	"event-processor-worker/config/db"
	"event-processor-worker/utilities"
	"log"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderItem struct {
	OrderID           int         `json:"order_id"`
	ProductID         int         `json:"product_id"`
	OrderItemNum      json.Number `json:"order_item_num"`
	SellerKey         string      `json:"seller_key"`
	ShippingLimitDate string      `json:"shipping_limit_date"`
	Price             string      `json:"price"`
	FreightValue      string      `json:"freight_value"`
}

func processMessage(data []byte) error {
	db := db.GetDB()

	var orderItem OrderItem
	if err := json.Unmarshal(data, &orderItem); err != nil {
		return err
	}
	orderItemNum, _ := orderItem.OrderItemNum.Int64()

	layout := "2006-01-02 15:04:05"
	shippingLimitDate, _ := time.Parse(layout, orderItem.ShippingLimitDate)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := `
		INSERT INTO orderItems (
			order_id,
			product_id,
			order_item_num,
			seller_key,
			shipping_limit_date,
			price,
			freight_value
		)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	price, err := strconv.ParseFloat(orderItem.Price, 64)
	utilities.ErrorHandler(err, "Failed to parse price")
	freightValue, err := strconv.ParseFloat(orderItem.FreightValue, 64)
	utilities.ErrorHandler(err, "Failed to parse freightValue")

	_, err = db.ExecContext(ctx, query,
		orderItem.OrderID,
		orderItem.ProductID,
		orderItemNum,
		orderItem.SellerKey,
		shippingLimitDate,
		price,
		freightValue,
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
