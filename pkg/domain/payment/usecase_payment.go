package payment

import (
	"github.com/FernandoCagale/c4-payment/internal/broker/producer"
	"github.com/FernandoCagale/c4-payment/internal/errors"
	"github.com/FernandoCagale/c4-payment/pkg/entity"
)

const TOPIC = "payment.registered"

type PaymentUseCase struct {
	repo     Repository
	producer producer.Producer
}

func NewUseCase(repo Repository, producer producer.Producer) *PaymentUseCase {
	return &PaymentUseCase{
		repo:     repo,
		producer: producer,
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

	if err := usecase.producer.Producer(TOPIC, customer); err != nil {
		return err
	}

	return nil
}
