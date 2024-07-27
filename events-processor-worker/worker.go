package main

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	var connection *amqp.Connection
	var err error

	for i := 0; i < 5; i++ {
		connection, err = amqp.Dial("amqp://guest:guest@messageQueue:5672/")
		if err == nil {
			break
		}
		log.Printf("Failed to connect to RabbitMQ, retrying in 5 seconds... (%d/5)\n", i+1)
		time.Sleep(5 * time.Second)
	}
	errorHandler(err, "Failed to connect to the message queue")
	defer connection.Close()

	channel, err := connection.Channel()
	errorHandler(err, "Failed to create a channel")
	defer channel.Close()

	customersQueue, err := channel.QueueDeclare(
		"customers",
		true,
		false,
		false,
		false,
		nil,
	)
	errorHandler(err, "Failed to connect to customers queue")

	productsQueue, err := channel.QueueDeclare(
		"products",
		true,
		false,
		false,
		false,
		nil,
	)
	errorHandler(err, "Failed to connect to products queue")

	ordersQueue, err := channel.QueueDeclare(
		"orders",
		true,
		false,
		false,
		false,
		nil,
	)
	errorHandler(err, "Failed to connect to orders queue")

	orderItemsQueue, err := channel.QueueDeclare(
		"orderItems",
		true,
		false,
		false,
		false,
		nil,
	)
	errorHandler(err, "Failed to connect to orderItems queue")

	var keepRunning chan struct{}

	customersMessages, err := channel.Consume(
		customersQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	errorHandler(err, "Failed to get customers messages")

	go func() {
		for d := range customersMessages {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	productsMessages, err := channel.Consume(
		productsQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	errorHandler(err, "Failed to get products messages")

	go func() {
		for d := range productsMessages {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	ordersMessages, err := channel.Consume(
		ordersQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	errorHandler(err, "Failed to get orders messages")

	go func() {
		for d := range ordersMessages {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	orderItemsMessages, err := channel.Consume(
		orderItemsQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	errorHandler(err, "Failed to get orderItems messages")

	go func() {
		for d := range orderItemsMessages {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	<-keepRunning
}

func errorHandler(err error, message string) {
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
}
