package sellers

import (
	"event-processor-worker/config/messageQueue"
	"event-processor-worker/utilities"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var dlxName = "dlx_exchange"
var dlxRoutingKey = "dlx_routing_key"

const numWorkers = 10

var taskChannel = make(chan amqp.Delivery)

func StartConsumers() {
	startSellerConsumer()
	startDlxConsumer()
}

func startSellerConsumer() {
	args := amqp.Table{
		"x-queue-type":              "quorum",
		"x-delivery-limit":          5,
		"x-dead-letter-exchange":    dlxName,
		"x-dead-letter-routing-key": dlxRoutingKey,
	}
	queueName := "sellers"
	_, err := messageQueue.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		args,
	)
	utilities.ErrorHandler(err, "Failed to declare sellers queue")

	sellersMessages, err := messageQueue.Channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utilities.ErrorHandler(err, "Failed to get sellers messages")

	for i := 0; i < numWorkers; i++ {
		go worker()
	}

	func() {
		for d := range sellersMessages {
			taskChannel <- d
		}
	}()
	close(taskChannel)
}

func worker() {
	for delivery := range taskChannel {
		err := processMessage(delivery.Body)
		if err != nil {
			log.Printf("Failed to persist seller.\n%s Message requeued.\n%v", delivery.Body, err)
			delivery.Nack(false, true)
		} else {
			delivery.Ack(false)
			log.Printf("Message acknowledged.\n%s ", delivery.Body)
		}
	}
}

func startDlxConsumer() {
	dlxArgs := amqp.Table{
		"x-queue-type": "quorum",
	}
	queueName := "dlx_queue"

	err := messageQueue.Channel.ExchangeDeclare(
		dlxName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	utilities.ErrorHandler(err, "Failed to declare seller DLX exchange")

	_, err = messageQueue.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		dlxArgs,
	)
	utilities.ErrorHandler(err, "Failed to declare seller DLX queue")

	messageQueue.Channel.QueueBind(
		queueName,
		dlxRoutingKey,
		dlxName,
		false,
		nil,
	)

	dlxMessages, err := messageQueue.Channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utilities.ErrorHandler(err, "Failed to get DLX messages")

	go func() {
		for d := range dlxMessages {
			handleDlxMessages(d)
		}
	}()
}
