package storage

import (
	"errors"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/domain/event"
)

var (
	ErrEventNotFound     = errors.New("event by id '%s' not found")
	ErrEventAlreadyExist = errors.New("event by id '%s' already exist")
)

type Storage interface {
	CreateEvent(evt event.Event) error
	DeleteEvent(evt event.Event) error
	UpdateEvent(evt event.Event) error
	GetEventByID(id string) (event.Event, error)
	GetAllEvents() ([]event.Event, error)
}
