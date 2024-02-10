package query

import (
	"fmt"
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
)

type GetMonthEventsRequest struct {
	Date time.Time
}

type GetMonthEventsResponse struct {
	Events []*domain.Event
}

type GetMonthEventsRequestHandler interface {
	Handle(request GetMonthEventsRequest) (*GetMonthEventsResponse, error)
}

func NewGetMonthEventsRequestHandler(storage storage.EventsRepository) (GetMonthEventsRequestHandler, error) {
	if storage == nil {
		return nil, fmt.Errorf("provided storage is nil")
	}
	return getMonthEventsRequestHandler{storage: storage}, nil
}

type getMonthEventsRequestHandler struct {
	storage storage.EventsRepository
}

func (h getMonthEventsRequestHandler) Handle(
	request GetMonthEventsRequest,
) (*GetMonthEventsResponse, error) {
	events, err := h.storage.GetEventsByMonth(request.Date)
	if err != nil {
		return nil, fmt.Errorf("storage GetMonthEvents error: %w", err)
	}
	return &GetMonthEventsResponse{Events: events}, nil
}
