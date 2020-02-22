package event

import "github.com/ThreeDotsLabs/watermill/message"

type Event interface {
	SubscribeExchange(topic, queue string) (<-chan *message.Message, error)
	SubscribeQueue(topic string) (<-chan *message.Message, error)
	PublishExchange(topic string, payload interface{}) (error)
	PublishQueue(topic string, payload interface{}) (error)
}
