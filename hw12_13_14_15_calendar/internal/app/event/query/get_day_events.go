package query

import (
	"fmt"
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
)

type GetDayEventsRequest struct {
	Date time.Time
}

type GetDayEventsResponse struct {
	Events []*domain.Event
}

type GetDayEventsRequestHandler interface {
	Handle(request GetDayEventsRequest) (*GetDayEventsResponse, error)
}

func NewGetDayEventsRequestHandler(storage storage.EventsRepository) (GetDayEventsRequestHandler, error) {
	if storage == nil {
		return nil, fmt.Errorf("provided storage is nil")
	}
	return getDayEventsRequestHandler{storage: storage}, nil
}

type getDayEventsRequestHandler struct {
	storage storage.EventsRepository
}

func (h getDayEventsRequestHandler) Handle(
	request GetDayEventsRequest,
) (*GetDayEventsResponse, error) {
	events, err := h.storage.GetEventsByDate(request.Date)
	if err != nil {
		return nil, fmt.Errorf("storage GetDayEvents error: %w", err)
	}
	return &GetDayEventsResponse{Events: events}, nil
}
