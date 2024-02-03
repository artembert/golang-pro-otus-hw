package event

import (
	"net/http"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/logger"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/pkg/api/openapi"
)

type eventHandler struct {
	app    app.Application
	logger logger.Logger
}

func NewEventHandler(app app.Application, logger logger.Logger) openapi.ServerInterface {
	return &eventHandler{
		app:    app,
		logger: logger,
	}
}

func (e *eventHandler) GetEvents(w http.ResponseWriter, r *http.Request, params openapi.GetEventsParams) {
	// TODO implement me
	panic("implement me")
}

func (e *eventHandler) PostEvents(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (e *eventHandler) PutEvents(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (e *eventHandler) DeleteEventsID(w http.ResponseWriter, r *http.Request, id string) {
	// TODO implement me
	panic("implement me")
}
