package producer

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"os"
	"strings"
)

type ProducerKafka struct {
	address string
}

func NewProducerKafka() *ProducerKafka {
	return &ProducerKafka{
		address: os.Getenv("ADDRESS_KAFKA"),
	}
}

func (e *ProducerKafka) Producer(topic string, payload interface{}) error {
	produce := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  strings.Split(e.address, ","),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	defer produce.Close()

	value, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return produce.WriteMessages(context.Background(),
		kafka.Message{
			Value: value,
		},
	)
}