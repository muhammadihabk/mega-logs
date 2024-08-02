package sellers

import (
	"event-processor-worker/config/messageQueue"
	"event-processor-worker/utilities"
)

func StartSellersConsumer() {
	sellersQueue, err := messageQueue.Channel.QueueDeclare(
		"sellers",
		true,
		false,
		false,
		false,
		nil,
	)
	utilities.ErrorHandler(err, "Failed to connect to sellers queue")
	sellersMessages, err := messageQueue.Channel.Consume(
		sellersQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utilities.ErrorHandler(err, "Failed to get sellers messages")

	go func() {
		for d := range sellersMessages {
			PersistSeller(d)
		}
	}()
}
