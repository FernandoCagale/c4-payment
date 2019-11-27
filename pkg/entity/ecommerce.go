package entity

import (
	"github.com/go-ozzo/ozzo-validation"
	"time"
)

type Ecommerce struct {
	ID       string            `json:"id"`
	Customer EcommerceCustomer `json:"customer"`
	Order    EcommerceOrder    `json:"order"`
}

func (e Ecommerce) Validate() error {
	return validation.ValidateStruct(&e,
	)
}

func (e Ecommerce) ToCustomer() Customer {
	var value float64

	for _, orderItem := range e.Order.Items {
		value += orderItem.Value * float64(orderItem.Quantity)
	}

	return Customer{
		Code: e.Customer.Code,
		Name: e.Customer.Name,
		Payments: []*Payment{
			{
				Code:  e.Order.Code,
				Data:  e.Order.Date,
				Value: value,
			},
		},
	}
}

type EcommerceCustomer struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type EcommerceOrder struct {
	Code  string           `json:"code"`
	Date  time.Time        `json:"date"`
	Items []*EcommerceItem `json:"items"`
}

type EcommerceItem struct {
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Value       float64 `json:"value"`
}
