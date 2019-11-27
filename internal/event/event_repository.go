package event

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"os"
)

type EventRepository struct {
	uri string
}

func New() *EventRepository {
	return &EventRepository{
		uri: os.Getenv("AMQP_URI"),
	}
}

func (o *EventRepository) getConfig() (config amqp.Config) {
	return amqp.NewDurablePubSubConfig(o.uri, nil)
}

func (o *EventRepository) getConfigQueue(queue string) (config amqp.Config) {
	queueConstant := amqp.GenerateQueueNameConstant(queue)
	return amqp.NewDurablePubSubConfig(o.uri, queueConstant)
}

func (o *EventRepository) message(payload interface{}) (*message.Message, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return message.NewMessage(watermill.NewUUID(), body), nil
}

func (o *EventRepository) Subscribe(topic, queue string) (<-chan *message.Message, error) {
	subscriber, err := amqp.NewSubscriber(o.getConfigQueue(queue), watermill.NewStdLogger(false, false))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return subscriber.Subscribe(context.Background(), topic)
}

func (o *EventRepository) Publish(topic string, payload interface{}) (error) {
	publisher, err := amqp.NewPublisher(o.getConfig(), watermill.NewStdLogger(false, false), )
	if err != nil {
		return err
	}

	msg, err := o.message(payload)
	if err != nil {
		return err
	}

	err = publisher.Publish(topic, msg)
	if err != nil {
		return err
	}
	return nil
}
