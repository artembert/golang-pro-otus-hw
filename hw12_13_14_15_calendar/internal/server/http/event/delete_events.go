package event

import (
	"net/http"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app/event/command"
)

func (handler *eventHandler) DeleteEventsID(w http.ResponseWriter, r *http.Request, id string) {
	err := handler.app.DeleteEvent(command.DeleteEventRequest{ID: id})
	if err != nil {
		handler.logger.Error("failed to delete event by ID: [" + id + "]:" + err.Error())
		handler.writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
