package storage

import (
	"errors"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/domain/event"
)

type Type string

const (
	Memory Type = "memory"
	SQL    Type = "sql"
)

type Actions interface {
	CreateEvent(evt event.Event) error
	DeleteEvent(evt event.Event) error
	UpdateEvent(evt event.Event) error
	GetEventByID(id string) (event.Event, error)
	GetAllEvents() ([]event.Event, error)
}

var (
	ErrEventNotFound          = errors.New("event by id '%s' not found")
	ErrEventAlreadyExist      = errors.New("event by id '%s' already exist")
	ErrUnsupportedStorageType = errors.New("storage type '%s' is wot supported")
)
