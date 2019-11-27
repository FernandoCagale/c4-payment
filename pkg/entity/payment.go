package entity

import (
	"time"
)

type Customer struct {
	Code     string     `json:"code" bson:"_id"`
	Name     string     `json:"name"`
	Payments []*Payment `json:"payments"`
}

type Payment struct {
	Code  string    `json:"code"`
	Data  time.Time `json:"data"`
	Value float64   `json:"value"`
}
