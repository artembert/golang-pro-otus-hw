package event

import (
	"encoding/json"
	"net/http"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app/event/command"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/server/http/event/mapper"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/pkg/api/openapi"
)

func (handler *eventHandler) PostEvents(w http.ResponseWriter, r *http.Request) {
	var newEvent openapi.NewEvent
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		handler.logger.Error("openapi.NewEvent json decode error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newEvt, err := mapper.OpenapiNewEventToEvent(&newEvent, r.Header.Get(app.HeaderUserId))
	if err != nil {
		handler.logger.Error("openapi.NewEvent to Event error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := handler.app.CreateEvent(command.CreateEventRequest{
		Event: newEvt,
	})
	if err != nil {
		handler.logger.Error("create event error: " + err.Error())
		handler.writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	handler.writeEvent(w, &res.Event)
}
