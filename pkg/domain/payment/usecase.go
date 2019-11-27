package payment

import (
	"github.com/FernandoCagale/c4-payment/pkg/entity"
)

type UseCase interface {
	Create(ecommerce *entity.Ecommerce) (err error)
	FindAll() (payments []*entity.Customer, err error)
	FindById(ID string) (payment *entity.Customer, err error)
	DeleteById(ID string) (err error)
}
