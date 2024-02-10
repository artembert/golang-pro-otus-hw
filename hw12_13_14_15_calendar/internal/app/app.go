package app

import (
	"context"
	"fmt"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app/event/command"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app/event/query"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/logger"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
)

type app struct {
	ctx                 context.Context
	store               Storage
	logger              Logger
	createHandler       command.CreateEventRequestHandler
	getDayEventsHandler query.GetDayEventsRequestHandler
}

type Application interface {
	CreateEvent(ctx context.Context, request command.CreateEventRequest) (*command.CreateEventResponse, error)
	GetDayEvents(ctx context.Context, request query.GetDayEventsRequest) (*query.GetDayEventsResponse, error)
	// UpdateEvent(ctx context.Context, request command.UpdateEventRequest) error
	// DeleteEvent(ctx context.Context, request command.DeleteEventRequest) error
	// GetWeekEvents(ctx context.Context, request query.GetWeekEventsRequest) (*query.GetWeekEventsResponse, error)
	// GetMonthEvents(ctx context.Context, request query.GetMonthEventsRequest) (*query.GetMonthEventsResponse, error)
}

type Logger interface {
	logger.Logger
}

type Storage interface {
	storage.EventsRepository
}

func New(ctx context.Context, logg Logger, storage Storage) (Application, error) {
	logg.Info("app created")
	createHandler, err := command.NewCreateEventRequestHandler(storage)
	if err != nil {
		logg.Error("create CreateEventRequestHandler error: %w", err)
		return nil, fmt.Errorf("create CreateEventRequestHandler error: %w", err)
	}
	getDayEventsHandler, err := query.NewGetDayEventsRequestHandler(storage)
	if err != nil {
		logg.Error("create GetDayEventsRequestHandler error: %w", err)
		return nil, fmt.Errorf("create GetDayEventsRequestHandler error: %w", err)
	}
	return &app{
		ctx: ctx, store: storage, logger: logg,
		createHandler:       createHandler,
		getDayEventsHandler: getDayEventsHandler,
	}, nil
}

func (a *app) CreateEvent(_ context.Context, request command.CreateEventRequest) (*command.CreateEventResponse, error) {
	return a.createHandler.Handle(request)
}

func (a *app) GetDayEvents(ctx context.Context, request query.GetDayEventsRequest) (*query.GetDayEventsResponse, error) {
	return a.getDayEventsHandler.Handle(request)
}
