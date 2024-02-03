package command

import (
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
)

type CreateEventRequest struct {
	Title        string
	StartTime    time.Time
	Duration     time.Duration
	Description  string
	UserID       string
	NotifyBefore time.Duration
}

type CreateEventResponse struct {
	EventID string
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
	evt := &domain.Event{
		Title:        req.Title,
		StartTime:    req.StartTime,
		Duration:     req.Duration,
		Description:  req.Description,
		UserID:       domain.UserID(req.UserID),
		NotifyBefore: req.NotifyBefore,
		Notified:     false,
	}
	id, err := c.storage.CreateEvent(evt)
	if err != nil {
		return nil, err
	}
	return &CreateEventResponse{EventID: string(id)}, nil
}
