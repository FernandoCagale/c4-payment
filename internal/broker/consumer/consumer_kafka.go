package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/FernandoCagale/c4-payment/pkg/domain/payment"
	"github.com/FernandoCagale/c4-payment/pkg/entity"
	"github.com/segmentio/kafka-go"
	"os"
	"strings"
	"time"
)

const (
	NOTIFY_ECOMMERCE_TOPIC = "notify.ecommerce"
)

type ConsumerKafka struct {
	address string
	usecase payment.UseCase
}

func NewConsumerKafka(usecase payment.UseCase, ) *ConsumerKafka {
	return &ConsumerKafka{
		usecase: usecase,
		address: os.Getenv("ADDRESS_KAFKA"),
	}
}

func (e *ConsumerKafka) Consumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        strings.Split(e.address, ","),
		Topic:          NOTIFY_ECOMMERCE_TOPIC,
		GroupID:        "c4-payment",
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: 10 * time.Second,
	})

	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err.Error()) //TODO
		}

		var ecommerce entity.Ecommerce

		if err := json.Unmarshal(m.Value, &ecommerce); err != nil {
			fmt.Println(err.Error()) //TODO
		}

		if err := e.usecase.Create(&ecommerce); err != nil {
			fmt.Println(err.Error()) //TODO
		}

		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
