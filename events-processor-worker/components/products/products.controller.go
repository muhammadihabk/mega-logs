package products

import (
	"event-processor-worker/config/messageQueue"
	"event-processor-worker/utilities"
)

func StartProductsConsumer() {
	productsQueue, err := messageQueue.Channel.QueueDeclare(
		"products",
		true,
		false,
		false,
		false,
		nil,
	)
	utilities.ErrorHandler(err, "Failed to connect to products queue")
	productsMessages, err := messageQueue.Channel.Consume(
		productsQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utilities.ErrorHandler(err, "Failed to get products messages")

	go func() {
		for d := range productsMessages {
			PersistProduct(d)
		}
	}()
}
