package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"event-processor-worker/components/customers"
	"event-processor-worker/components/orderItems"
	"event-processor-worker/components/orders"
	"event-processor-worker/components/products"
	"event-processor-worker/components/sellers"
	"event-processor-worker/config/db"
	"event-processor-worker/config/messageQueue"
)

func main() {
	db.CreateTablesIfNotExist()

	go func() {
		customers.StartConsumers()
	}()

	go func() {
		products.StartConsumers()
	}()

	go func() {
		orders.StartConsumers()
	}()

	go func() {
		orderItems.StartConsumers()
	}()

	go func() {
		sellers.StartConsumers()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Clean up")
	db.CleanupOnExit()
	messageQueue.CleanupOnExit()
}
