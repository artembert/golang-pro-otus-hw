package app

import (
	"context"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/domain/storage"
)

type App struct {
	ctx    context.Context
	store  Storage
	logger Logger
}

type Logger interface { // TODO
}

type Storage interface {
	storage.Actions
}

func New(logger Logger, storage Storage) *App {
	return &App{}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
