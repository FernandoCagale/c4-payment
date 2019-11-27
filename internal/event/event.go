package event

import "github.com/ThreeDotsLabs/watermill/message"

type Event interface {
	Subscribe(topic, queue string) (<-chan *message.Message, error)
	Publish(topic string, payload interface{}) (error)
}
