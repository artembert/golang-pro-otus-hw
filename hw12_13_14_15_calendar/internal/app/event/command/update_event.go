package command

import (
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
)

type UpdateEventRequest struct {
	ID           string
	Title        string
	StartTime    time.Time
	Duration     time.Duration
	Description  string
	UserID       string
	NotifyBefore time.Duration
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
	if err := validateID(req.ID); err != nil {
		return nil, err
	}
	if err := validateTitle(req.Title); err != nil {
		return nil, err
	}
	if err := validateDuration(req.Duration); err != nil {
		return nil, err
	}
	if err := validateUserID(req.UserID); err != nil {
		return nil, err
	}
	evt := &domain.Event{
		ID:           domain.EventID(req.ID),
		Title:        req.Title,
		StartTime:    req.StartTime,
		Duration:     req.Duration,
		Description:  req.Description,
		UserID:       domain.UserID(req.UserID),
		NotifyBefore: req.NotifyBefore,
		Notified:     false,
	}
	updatedEvt, err := c.storage.UpdateEvent(evt)
	if err != nil {
		return nil, err
	}

	return &UpdateEventResponse{Event: *updatedEvt}, nil
}
