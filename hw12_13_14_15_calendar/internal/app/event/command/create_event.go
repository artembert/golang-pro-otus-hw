package command

import (
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
)

type CreateEventRequest struct {
	Event *domain.Event
}

type CreateEventResponse struct {
	Event domain.Event
}

type CreateEventRequestHandler interface {
	Handle(request CreateEventRequest) (*CreateEventResponse, error)
}

type createEventRequestHandler struct {
	storage storage.EventsRepository
}

func NewCreateEventRequestHandler(storage storage.EventsRepository) (CreateEventRequestHandler, error) {
	if storage == nil {
		return nil, ErrStorageIsUndefined
	}
	return &createEventRequestHandler{storage: storage}, nil
}

func (c *createEventRequestHandler) Handle(
	req CreateEventRequest,
) (*CreateEventResponse, error) {
	evt := req.Event
	if err := validateTitle(evt.Title); err != nil {
		return nil, err
	}
	if err := validateStartTime(evt.StartTime); err != nil {
		return nil, err
	}
	if err := validateDuration(evt.Duration); err != nil {
		return nil, err
	}
	if err := validateUserID(evt.UserID); err != nil {
		return nil, err
	}

	newEvent, err := c.storage.CreateEvent(evt)
	if err != nil {
		return nil, err
	}
	return &CreateEventResponse{Event: *newEvent}, nil
}
