package event

import (
	"github.com/FernandoCagale/c4-payment/internal/broker/consumer"
)

type PaymentEvent struct {
	consumer consumer.Consumer
}

func NewPayment(consumer consumer.Consumer) *PaymentEvent {
	return &PaymentEvent{
		consumer: consumer,
	}
}

func (event *PaymentEvent) MakeEvents() {
	go event.consumer.Consumer()
}
