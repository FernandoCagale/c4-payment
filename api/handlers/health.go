package handlers

import (
	"github.com/hellofresh/health-go"
	"gopkg.in/mgo.v2"
	"net/http"
)

type HealthHandler struct {
	session *mgo.Session
}

func NewHealth(session *mgo.Session) *HealthHandler {
	return &HealthHandler{session}
}

func (handler *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	health.HandlerFunc(w, r)
}
