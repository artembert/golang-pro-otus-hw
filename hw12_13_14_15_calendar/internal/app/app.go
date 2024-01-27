package app

import (
	"context"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/logger"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
)

type App struct {
	ctx    context.Context
	store  Storage
	logger Logger
}

type Logger interface {
	logger.Logger
}

type Storage interface {
	storage.EventsRepository
}

func New(ctx context.Context, logger Logger, storage Storage) *App {
	logger.Info("app created")
	return &App{
		ctx, storage, logger,
	}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	_ = ctx
	_ = id
	_ = title
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
