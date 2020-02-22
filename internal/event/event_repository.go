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

func (o *EventRepository) message(payload interface{}) (*message.Message, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return message.NewMessage(watermill.NewUUID(), body), nil
}

func (o *EventRepository) SubscribeExchange(topic, queue string) (<-chan *message.Message, error) {
	amqpConfig := amqp.NewDurablePubSubConfig(o.uri, amqp.GenerateQueueNameConstant(queue))
	subscriber, err := amqp.NewSubscriber(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return subscriber.Subscribe(context.Background(), topic)
}

func (o *EventRepository) SubscribeQueue(topic string) (<-chan *message.Message, error) {
	amqpConfig := amqp.NewDurableQueueConfig(o.uri)
	subscriber, err := amqp.NewSubscriber(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return subscriber.Subscribe(context.Background(), topic)
}

func (o *EventRepository) PublishExchange(topic string, payload interface{}) (error) {
	amqpConfig := amqp.NewDurablePubSubConfig(o.uri, nil)
	publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false), )
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

func (o *EventRepository) PublishQueue(topic string, payload interface{}) (error) {
	amqpConfig := amqp.NewDurableQueueConfig(o.uri)
	publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false), )
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