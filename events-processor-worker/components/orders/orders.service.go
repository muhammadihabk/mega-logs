package orders

import (
	"context"
	"encoding/json"
	"event-processor-worker/config/db"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Order struct {
	OrderKey                   string `json:"order_key"`
	CustomerID                 int    `json:"customer_id"`
	OrderStatus                string `json:"order_status"`
	OrderPurchaseTimestamp     string `json:"order_purchase_timestamp"`
	OrderApprovedAt            string `json:"order_approved_at"`
	OrderDeliveredCarrierDate  string `json:"order_delivered_carrier_date"`
	OrderDeliveredCustomerDate string `json:"order_delivered_customer_date"`
	OrderEstimatedDeliveryDate string `json:"order_estimated_delivery_date"`
}

func processMessage(data []byte) error {
	db := db.GetDB()

	var order Order
	if err := json.Unmarshal(data, &order); err != nil {
		return err
	}
	layout := "2006-01-02 15:04:05"
	orderPurchaseTimestamp, _ := time.Parse(layout, order.OrderPurchaseTimestamp)
	orderApprovedAt, _ := time.Parse(layout, order.OrderApprovedAt)
	orderDeliveredCarrierDate, _ := time.Parse(layout, order.OrderDeliveredCarrierDate)
	orderDeliveredCustomerDate, _ := time.Parse(layout, order.OrderDeliveredCustomerDate)
	orderEstimatedDeliveryDate, _ := time.Parse(layout, order.OrderEstimatedDeliveryDate)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := `
		INSERT INTO orders (
			order_key,
			customer_id,
			order_status,
			order_purchase_timestamp,
			order_approved_at,
			order_delivered_carrier_date,
			order_delivered_customer_date,
			order_estimated_delivery_date
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.ExecContext(ctx, query,
		order.OrderKey,
		order.CustomerID,
		order.OrderStatus,
		orderPurchaseTimestamp,
		orderApprovedAt,
		orderDeliveredCarrierDate,
		orderDeliveredCustomerDate,
		orderEstimatedDeliveryDate,
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
