package storage

import (
	"context"
	"errors"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/domain/event"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/domain/storage"
	memorystorage "github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/storage/sql"
)

var (
	ErrEventNotFound          = errors.New("event by id '%s' not found")
	ErrEventAlreadyExist      = errors.New("event by id '%s' already exist")
	ErrUnsupportedStorageType = errors.New("storage type '%s' is wot supported")
)

type Storage interface {
	CreateEvent(evt event.Event) error
	DeleteEvent(evt event.Event) error
	UpdateEvent(evt event.Event) error
	GetEventByID(id string) (event.Event, error)
	GetAllEvents() ([]event.Event, error)
}

func Init(ctx context.Context, cfg *config.Config) (Storage, error) {
	switch storage.Type(cfg.Storage.Type) {
	case storage.Memory:
		return memorystorage.New(), nil
	case storage.SQL:
		store := sqlstorage.New(ctx, &cfg.DB)
		return store, nil
	}
	return nil, ErrUnsupportedStorageType
}
