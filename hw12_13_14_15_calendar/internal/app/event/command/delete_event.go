package command

import (
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
)

type DeleteEventRequest struct {
	ID string
}

type DeleteEventRequestHandler interface {
	Handle(request DeleteEventRequest) error
}

type deleteEventRequestHandler struct {
	storage storage.EventsRepository
}

func NewDeleteEventRequestHandler(storage storage.EventsRepository) (DeleteEventRequestHandler, error) {
	if storage == nil {
		return nil, ErrStorageIsUndefined
	}
	return &deleteEventRequestHandler{storage: storage}, nil
}

func (c *deleteEventRequestHandler) Handle(
	req DeleteEventRequest,
) error {
	err := c.storage.DeleteEvent(domain.EventID(req.ID))
	if err != nil {
		return err
	}
	return nil
}
