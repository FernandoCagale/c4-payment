package handlers

import (
	"github.com/hellofresh/health-go"
	"gopkg.in/mgo.v2"
	"os"

	healthRabbitmq "github.com/hellofresh/health-go/checks/rabbitmq"

	"net/http"
	"time"
)

type HealthHandler struct {
	session *mgo.Session
}

func NewHealth(session *mgo.Session) *HealthHandler {
	return &HealthHandler{session}
}

func (handler *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	health.Register(health.Config{
		Name:    "rabbitmq",
		Timeout: time.Second * 5,
		Check: healthRabbitmq.New(healthRabbitmq.Config{
			DSN: os.Getenv("AMQP_URI"),
		}),
	})

	health.Register(health.Config{
		Name: "mongodb",
		Check: func() error {
			if err := handler.session.Ping(); err != nil {
				handler.session.Refresh()
				return err
			}
			return nil
		},
	})

	health.HandlerFunc(w, r)
}
