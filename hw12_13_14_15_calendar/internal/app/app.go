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
	ctx                   context.Context
	store                 Storage
	logger                Logger
	createHandler         command.CreateEventRequestHandler
	getDayEventsHandler   query.GetDayEventsRequestHandler
	getWeekEventsHandler  query.GetWeekEventsRequestHandler
	getMonthEventsHandler query.GetMonthEventsRequestHandler
	deleteEventHandler    command.DeleteEventRequestHandler
}

type Application interface {
	CreateEvent(request command.CreateEventRequest) (*command.CreateEventResponse, error)
	GetDayEvents(request query.GetDayEventsRequest) (*query.GetDayEventsResponse, error)
	GetWeekEvents(request query.GetWeekEventsRequest) (*query.GetWeekEventsResponse, error)
	GetMonthEvents(request query.GetMonthEventsRequest) (*query.GetMonthEventsResponse, error)
	DeleteEvent(request command.DeleteEventRequest) error
	// UpdateEvent(request command.UpdateEventRequest) error
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
	deleteEventHandler, err := command.NewDeleteEventRequestHandler(storage)
	if err != nil {
		logg.Error("create NewDeleteEventRequestHandler error: %w", err)
		return nil, fmt.Errorf("create NewDeleteEventRequestHandler error: %w", err)
	}
	getDayEventsHandler, err := query.NewGetDayEventsRequestHandler(storage)
	if err != nil {
		logg.Error("create GetDayEventsRequestHandler error: %w", err)
		return nil, fmt.Errorf("create GetDayEventsRequestHandler error: %w", err)
	}
	getWeekEventsHandler, err := query.NewGetWeekEventsRequestHandler(storage)
	if err != nil {
		logg.Error("create GetWeekEventsRequestHandler error: %w", err)
		return nil, fmt.Errorf("create GetWeekEventsRequestHandler error: %w", err)
	}
	getMonthEventsHandler, err := query.NewGetMonthEventsRequestHandler(storage)
	if err != nil {
		logg.Error("create GetMonthEventsRequestHandler error: %w", err)
		return nil, fmt.Errorf("create GetMonthEventsRequestHandler error: %w", err)
	}

	return &app{
		ctx: ctx, store: storage, logger: logg,
		createHandler:         createHandler,
		getDayEventsHandler:   getDayEventsHandler,
		getWeekEventsHandler:  getWeekEventsHandler,
		getMonthEventsHandler: getMonthEventsHandler,
		deleteEventHandler:    deleteEventHandler,
	}, nil
}

func (a *app) CreateEvent(request command.CreateEventRequest) (*command.CreateEventResponse, error) {
	return a.createHandler.Handle(request)
}

func (a *app) DeleteEvent(request command.DeleteEventRequest) error {
	return a.deleteEventHandler.Handle(request)
}

func (a *app) GetDayEvents(request query.GetDayEventsRequest) (*query.GetDayEventsResponse, error) {
	return a.getDayEventsHandler.Handle(request)
}

func (a *app) GetWeekEvents(request query.GetWeekEventsRequest) (*query.GetWeekEventsResponse, error) {
	return a.getWeekEventsHandler.Handle(request)
}

func (a *app) GetMonthEvents(request query.GetMonthEventsRequest) (*query.GetMonthEventsResponse, error) {
	return a.getMonthEventsHandler.Handle(request)
}
