package event

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	fmt.Println("in the declareExchange func call...")
	return ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable?
		false,        // auto-deleted?
		false,        // internal
		false,        // no-wait?
		nil,          // arguments?
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		false, // durable?
		false, // delete when unused?
		true,  // exlusive?
		false, // no-wait?
		nil,   // args?
	)
}
