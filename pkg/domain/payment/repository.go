package payment

import "github.com/FernandoCagale/c4-payment/pkg/entity"

type Repository interface {
	Create(customer *entity.Customer) (err error)
	FindAll() (payments []*entity.Customer, err error)
	FindById(ID string) (payment *entity.Customer, err error)
	DeleteById(ID string) (err error)
}
