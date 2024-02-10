package event

import (
	"net/http"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app/event/query"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/pkg/api/openapi"
)

func (handler *eventHandler) GetEvents(w http.ResponseWriter, r *http.Request, params openapi.GetEventsParams) {
	var events []*domain.Event
	switch params.Period {
	case openapi.Day:
		response, err := handler.app.GetDayEvents(query.GetDayEventsRequest{Date: params.Date})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		events = response.Events
	}
	w.WriteHeader(http.StatusOK)
	handler.writeEventsList(w, events)
}
