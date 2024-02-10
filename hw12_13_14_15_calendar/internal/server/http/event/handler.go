package event

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app/event/command"
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

func (handler *eventHandler) GetEvents(w http.ResponseWriter, r *http.Request, params openapi.GetEventsParams) {
	// TODO implement me
	panic("implement me")
}

func (handler *eventHandler) PutEvents(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (handler *eventHandler) DeleteEventsID(w http.ResponseWriter, r *http.Request, id string) {
	// TODO implement me
	panic("implement me")
}

func (handler *eventHandler) writeError(w http.ResponseWriter, err error) {
	if errors.Is(err, app.ErrEventNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var errValidate command.ErrValidate
	if errors.As(err, &errValidate) {
		w.WriteHeader(http.StatusBadRequest)
		errResponse := openapi.Error{Message: err.Error()}
		resBuf, errMarshal := json.Marshal(errResponse)
		if errMarshal != nil {
			handler.logger.Error("Event json marshal error: " + errMarshal.Error())
		}
		_, errWrite := w.Write(resBuf)
		if errWrite != nil {
			handler.logger.Error("Response write error: " + errWrite.Error())
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func (handler *eventHandler) writeEvent(w http.ResponseWriter, event domain.Event) {
	createdEvent := openapi.Event{
		Duration:     int(event.Duration / time.Minute),
		Id:           string(event.ID),
		NotifyBefore: int(event.NotifyBefore / time.Minute),
		StartTime:    event.StartTime, // TODO: Replace with time.MarshalText()
		Title:        event.Title,
	}
	resBuf, err := json.Marshal(createdEvent)
	if err != nil {
		handler.logger.Error("Event json marshal error: " + err.Error())
	}
	_, err = w.Write(resBuf)
	if err != nil {
		handler.logger.Error("Response write error: " + err.Error())
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}