package app

import (
	"context"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app/service/event"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/logger"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
)

type App struct {
	ctx          context.Context
	store        Storage
	logger       Logger
	eventService *event.Service
}

type EventService interface {
	event.Service
}

type Logger interface {
	logger.Logger
}

type Storage interface {
	storage.EventsRepository
}

func New(ctx context.Context, logg Logger, storage Storage) *App {
	logg.Info("app created")
	eventService := event.NewService((*logger.Logger)(&logg), storage)
	return &App{
		ctx, storage, logg, eventService,
	}
}
