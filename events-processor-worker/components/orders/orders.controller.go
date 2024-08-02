package orders

import (
	"event-processor-worker/config/messageQueue"
	"event-processor-worker/utilities"
)

func StartOrdersConsumer() {
	ordersQueue, err := messageQueue.Channel.QueueDeclare(
		"orders",
		true,
		false,
		false,
		false,
		nil,
	)
	utilities.ErrorHandler(err, "Failed to connect to orders queue")
	ordersMessages, err := messageQueue.Channel.Consume(
		ordersQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utilities.ErrorHandler(err, "Failed to get orders messages")

	go func() {
		for d := range ordersMessages {
			PersistOrder(d)
		}
	}()
}
