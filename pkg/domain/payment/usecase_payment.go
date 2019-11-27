package payment

import (
	"github.com/FernandoCagale/c4-payment/internal/errors"
	"github.com/FernandoCagale/c4-payment/internal/event"
	"github.com/FernandoCagale/c4-payment/pkg/entity"
)

type PaymentUseCase struct {
	repo  Repository
	event event.Event
}

func NewUseCase(repo Repository, event event.Event) *PaymentUseCase {
	return &PaymentUseCase{
		repo:  repo,
		event: event,
	}
}

func (usecase *PaymentUseCase) FindAll() (payments []*entity.Customer, err error) {
	return usecase.repo.FindAll()
}

func (usecase *PaymentUseCase) FindById(ID string) (payment *entity.Customer, err error) {
	return usecase.repo.FindById(ID)
}

func (usecase *PaymentUseCase) DeleteById(ID string) (err error) {
	return usecase.repo.DeleteById(ID)
}


func (usecase *PaymentUseCase) Create(e *entity.Ecommerce) error {
	err := e.Validate()
	if err != nil {
		return errors.ErrInvalidPayload
	}

	customer := e.ToCustomer()

	if err = usecase.repo.Create(&customer); err != nil {
		return err
	}

	//TODO notify

	return nil
}
