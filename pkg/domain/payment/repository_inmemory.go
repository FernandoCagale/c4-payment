package payment

import (
	"github.com/FernandoCagale/c4-payment/internal/errors"
	"github.com/FernandoCagale/c4-payment/pkg/entity"
)

type InMemoryRepository struct {
	m map[string]*entity.Customer
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{map[string]*entity.Customer{}}
}

func (repo *InMemoryRepository) FindAll() (payments []*entity.Customer, err error) {
	for _, payment := range repo.m {
		payments = append(payments, payment)
	}
	return payments, nil
}

func (repo *InMemoryRepository) FindById(ID string) (payment *entity.Customer, err error) {
	for _, payment := range repo.m {
		if payment.Code == ID {
			return payment, nil
		}
	}
	return nil, errors.ErrNotFound
}

func (repo *InMemoryRepository) DeleteById(ID string) (err error) {
	for _, payment := range repo.m {
		if payment.Code == ID {
			delete(repo.m, ID)
			return nil
		}
	}
	return errors.ErrNotFound
}

func (repo *InMemoryRepository) Create(e *entity.Customer) (err error) {
	customer := repo.m[e.Code]

	if customer == nil {
		repo.m[e.Code] = e
		return nil
	}

	for _, payment := range repo.m[e.Code].Payments {
		if payment.Code == e.Payments[0].Code {
			payment.Data = e.Payments[0].Data
			payment.Value = e.Payments[0].Value
			return nil
		}
	}
	customer.Payments = append(customer.Payments, e.Payments...)

	return nil
}
