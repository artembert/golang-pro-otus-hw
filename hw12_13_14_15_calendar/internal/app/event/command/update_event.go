package command

import (
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
)

type UpdateEventRequest struct {
	Event *domain.Event
}

type UpdateEventResponse struct {
	Event domain.Event
}

type UpdateEventRequestHandler interface {
	Handle(request UpdateEventRequest) (*UpdateEventResponse, error)
}

type updateEventRequestHandler struct {
	storage storage.EventsRepository
}

func NewUpdateEventRequestHandler(storage storage.EventsRepository) (UpdateEventRequestHandler, error) {
	if storage == nil {
		return nil, ErrStorageIsUndefined
	}
	return &updateEventRequestHandler{storage: storage}, nil
}

func (c *updateEventRequestHandler) Handle(
	req UpdateEventRequest,
) (*UpdateEventResponse, error) {
	evt := req.Event
	if err := validateID(string(evt.ID)); err != nil {
		return nil, err
	}
	if err := validateTitle(evt.Title); err != nil {
		return nil, err
	}
	if err := validateDuration(evt.Duration); err != nil {
		return nil, err
	}
	if err := validateUserID(evt.UserID); err != nil {
		return nil, err
	}
	updatedEvt, err := c.storage.UpdateEvent(evt)
	if err != nil {
		return nil, err
	}

	return &UpdateEventResponse{Event: *updatedEvt}, nil
}
