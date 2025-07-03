package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/lucasfeitozas/golang-clean-arch/pkg/events"
	"github.com/streadway/amqp"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	// Declara a fila se ela não existir
	_, err := rabbitMQChannel.QueueDeclare(
		"orders", // nome da fila
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		panic(err)
	}

	// Bind da fila ao exchange
	err = rabbitMQChannel.QueueBind(
		"orders",        // nome da fila
		"order.created", // routing key
		"amq.direct",    // exchange
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		panic(err)
	}

	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order created: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct",    // exchange
		"order.created", // routing key
		false,           // mandatory
		false,           // immediate
		msgRabbitmq,     // message to publish
	)
}
