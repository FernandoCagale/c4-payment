package handlers

import (
	"encoding/json"
	"github.com/FernandoCagale/c4-payment/api/render"
	"github.com/FernandoCagale/c4-payment/internal/errors"
	"github.com/FernandoCagale/c4-payment/pkg/domain/payment"
	"github.com/FernandoCagale/c4-payment/pkg/entity"
	"github.com/gorilla/mux"
	"net/http"
)

type PaymentHandler struct {
	usecase payment.UseCase
}

func NewPayment(usecase payment.UseCase) *PaymentHandler {
	return &PaymentHandler{
		usecase: usecase,
	}
}

func (handler *PaymentHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	payments, err := handler.usecase.FindAll()
	if err != nil {
		render.ResponseError(w, err, http.StatusInternalServerError)
		return
	}

	render.Response(w, payments, http.StatusOK)
}

func (handler *PaymentHandler) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ID := vars["id"]

	payment, err := handler.usecase.FindById(ID)
	if err != nil {
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, payment, http.StatusOK)
}

func (handler *PaymentHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ID := vars["id"]

	err := handler.usecase.DeleteById(ID)
	if err != nil {
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, nil, http.StatusNoContent)
}

func (handler *PaymentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var ecommerce *entity.Ecommerce

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ecommerce); err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := handler.usecase.Create(ecommerce); err != nil {
		switch err {
		case errors.ErrInvalidPayload:
			render.ResponseError(w, err, http.StatusBadRequest)
		case errors.ErrConflict:
			render.ResponseError(w, err, http.StatusConflict)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, nil, http.StatusCreated)
}
