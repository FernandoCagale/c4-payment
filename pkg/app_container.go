package pkg

import (
	"github.com/FernandoCagale/c4-payment/api/handlers"
	"github.com/FernandoCagale/c4-payment/api/routers"
	"github.com/FernandoCagale/c4-payment/internal/broker/producer"
	"github.com/FernandoCagale/c4-payment/pkg/domain/payment"
	"github.com/google/wire"
)

var Container = wire.NewSet(payment.Set, handlers.Set, routers.Set, producer.Set)
