package handler

import (
	"context"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/logger"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/pkg/api/grpc/eventsservice"
)

type eventHandler struct {
	eventsservice.UnimplementedCalendarServer
	app    app.Application
	logger logger.Logger
}

func (handler *eventHandler) UpdateEvent(ctx context.Context, request *eventsservice.UpdateEventRequest) (*eventsservice.UpdateEventResponse, error) {
	// TODO implement me
	panic("implement me")
}

func New(app app.Application, logger logger.Logger) eventsservice.CalendarServer {
	return &eventHandler{
		app:    app,
		logger: logger,
	}
}
