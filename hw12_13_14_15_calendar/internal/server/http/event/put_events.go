package event

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app/event/command"
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

	request := command.UpdateEventRequest{
		ID:           evt.Id,
		Title:        evt.Title,
		StartTime:    evt.StartTime,
		Duration:     time.Duration(evt.Duration) * time.Minute,
		Description:  evt.Description,
		UserID:       r.Header.Get(app.HeaderUserId),
		NotifyBefore: time.Duration(evt.NotifyBefore) * time.Minute,
	}
	res, err := handler.app.UpdateEvent(request)
	if err != nil {
		handler.logger.Error("update event error: " + err.Error())
		handler.writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	handler.writeEvent(w, &res.Event)
}
