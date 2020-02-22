package event

import (
	"encoding/json"
	"fmt"
	"github.com/FernandoCagale/c4-payment/internal/event"
	"github.com/FernandoCagale/c4-payment/pkg/domain/payment"
	"github.com/FernandoCagale/c4-payment/pkg/entity"
	"log"
)

const (
	EXCHANGE = "ecommerce"
	QUEUE    = "payment"
)

type PaymentEvent struct {
	usecase payment.UseCase
	event   event.Event
}

func NewPayment(usecase payment.UseCase, event event.Event) *PaymentEvent {
	return &PaymentEvent{
		usecase: usecase,
		event:   event,
	}
}

func (eventPayment *PaymentEvent) ProcessPayment() {
	messages, err := eventPayment.event.SubscribeExchange(EXCHANGE, QUEUE)
	if err != nil {
		fmt.Println(err.Error())
	}

	for msg := range messages {
		log.Printf("received message: %s, PAYMENT: %s", msg.UUID, string(msg.Payload))

		var ecommerce entity.Ecommerce

		if err := json.Unmarshal(msg.Payload, &ecommerce); err != nil {
			fmt.Println(err.Error())
			msg.Nacked()
		}

		if err = eventPayment.usecase.Create(&ecommerce); err != nil {
			fmt.Println(err.Error())
			msg.Nacked()
		}

		msg.Ack() //TODO x-dead-letter-exchange
	}
}
