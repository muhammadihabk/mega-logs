package orderItems

import (
	"event-processor-worker/config/messageQueue"
	"event-processor-worker/utilities"
)

func StartOrderItemsConsumer() {
	orderItemsQueue, err := messageQueue.Channel.QueueDeclare(
		"orderItems",
		true,
		false,
		false,
		false,
		nil,
	)
	utilities.ErrorHandler(err, "Failed to connect to orderItems queue")
	orderItemsMessages, err := messageQueue.Channel.Consume(
		orderItemsQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utilities.ErrorHandler(err, "Failed to get orderItems messages")

	go func() {
		for d := range orderItemsMessages {
			PersistOrderItems(d)
		}
	}()
}
