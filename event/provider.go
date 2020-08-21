package event

import (
	"github.com/FernandoCagale/c4-payment/internal/broker/consumer"
	"github.com/FernandoCagale/c4-payment/internal/broker/producer"
	"github.com/FernandoCagale/c4-payment/pkg/domain/payment"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewPayment, payment.Set, consumer.Set, producer.Set)
