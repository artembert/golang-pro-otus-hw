package query

import (
	"fmt"
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
)

type GetWeekEventsRequest struct {
	Date time.Time
}

type GetWeekEventsResponse struct {
	Events []*domain.Event
}

type GetWeekEventsRequestHandler interface {
	Handle(request GetWeekEventsRequest) (*GetWeekEventsResponse, error)
}

func NewGetWeekEventsRequestHandler(storage storage.EventsRepository) (GetWeekEventsRequestHandler, error) {
	if storage == nil {
		return nil, fmt.Errorf("provided storage is nil")
	}
	return getWeekEventsRequestHandler{storage: storage}, nil
}

type getWeekEventsRequestHandler struct {
	storage storage.EventsRepository
}

func (h getWeekEventsRequestHandler) Handle(
	request GetWeekEventsRequest,
) (*GetWeekEventsResponse, error) {
	events, err := h.storage.GetEventsByWeek(request.Date)
	if err != nil {
		return nil, fmt.Errorf("storage GetWeekEvents error: %w", err)
	}
	return &GetWeekEventsResponse{Events: events}, nil
}
