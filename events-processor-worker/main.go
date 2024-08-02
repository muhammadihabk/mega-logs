package main

import (
	"event-processor-worker/components/customers"
	"event-processor-worker/components/orderItems"
	"event-processor-worker/components/orders"
	"event-processor-worker/components/products"
	"event-processor-worker/components/sellers"
	"log"
	"os"
	"os/signal"
	"syscall"

	"event-processor-worker/config/db"
	"event-processor-worker/config/messageQueue"
)

func main() {
	db.CreateTablesIfNotExist()

	customers.StartCustomerConsumer()
	products.StartProductsConsumer()
	orders.StartOrdersConsumer()
	orderItems.StartOrderItemsConsumer()
	sellers.StartSellersConsumer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gracefully...")
	db.CleanupOnExit()
	messageQueue.CleanupOnExit()
}
