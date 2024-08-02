package customers

import (
	"event-processor-worker/config/messageQueue"
	"event-processor-worker/utilities"
)

func StartCustomerConsumer() {
	customersQueue, err := messageQueue.Channel.QueueDeclare(
		"customers",
		true,
		false,
		false,
		false,
		nil,
	)
	utilities.ErrorHandler(err, "Failed to connect to customers queue")
	customersMessages, err := messageQueue.Channel.Consume(
		customersQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utilities.ErrorHandler(err, "Failed to get customers messages")

	go func() {
		for d := range customersMessages {
			PersistCustomer(d)
		}
	}()
}
