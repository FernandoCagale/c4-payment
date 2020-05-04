package routers

import (
	"github.com/FernandoCagale/c4-payment/api/handlers"
	"github.com/gorilla/mux"
)

type SystemRoutes struct {
	healthHandler  *handlers.HealthHandler
	paymentHandler *handlers.PaymentHandler
}

func (routes *SystemRoutes) MakeHandlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", routes.healthHandler.Health).Methods("GET")
	r.HandleFunc("/payments", routes.paymentHandler.Create).Methods("POST")
	r.HandleFunc("/payments", routes.paymentHandler.FindAll).Methods("GET")
	r.HandleFunc("/payments/{id}", routes.paymentHandler.FindById).Methods("GET")
	r.HandleFunc("/payments/{id}", routes.paymentHandler.DeleteById).Methods("DELETE")

	return r
}

func NewSystem(healthHandler *handlers.HealthHandler, paymentHandler *handlers.PaymentHandler) *SystemRoutes {
	return &SystemRoutes{
		healthHandler:  healthHandler,
		paymentHandler: paymentHandler,
	}
}
