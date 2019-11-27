package handlers

import (
	"github.com/FernandoCagale/c4-payment/api/render"
	"net/http"
)

type HealthHandler struct {
}

func (handler *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	render.Response(w, map[string]bool{"ok": true}, http.StatusOK)
}

func NewHealth() *HealthHandler {
	return &HealthHandler{}
}
