package event

import (
	eventImp "github.com/FernandoCagale/c4-payment/internal/event"
	"github.com/FernandoCagale/c4-payment/pkg/domain/payment"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewPayment, payment.Set, eventImp.Set)
