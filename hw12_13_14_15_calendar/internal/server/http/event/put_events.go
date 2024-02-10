package event

import (
	"encoding/json"
	"net/http"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app/event/command"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/server/http/event/mapper"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/pkg/api/openapi"
)

func (handler *eventHandler) PutEvents(w http.ResponseWriter, r *http.Request) {
	var evt openapi.Event
	err := json.NewDecoder(r.Body).Decode(&evt)
	if err != nil {
		handler.logger.Error("openapi.NewEvent json decode error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	evtToUpdate, err := mapper.OpenapiEventToEvent(&evt, r.Header.Get(app.HeaderUserId))
	if err != nil {
		handler.logger.Error("openapi.Event to Event error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := handler.app.UpdateEvent(command.UpdateEventRequest{
		Event: evtToUpdate,
	})
	if err != nil {
		handler.logger.Error("update event error: " + err.Error())
		handler.writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	handler.writeEvent(w, &res.Event)
}
