package storage

import (
	"errors"
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
)

type Type string

const (
	Memory Type = "memory"
	SQL    Type = "sql"
)

type EventsRepository interface {
	CreateEvent(evt *domain.Event) (*domain.Event, error)
	DeleteEvent(id domain.EventID) error
	UpdateEvent(id domain.EventID, evt *domain.Event) error
	GetEventsByDate(date time.Time) ([]*domain.Event, error)
	GetEventsByWeek(startOfWeek time.Time) ([]*domain.Event, error)
	GetEventsByMonth(startOfMonth time.Time) ([]*domain.Event, error)
}

var (
	ErrEventNotFound          = errors.New("event by id '%s' not found")
	ErrUnsupportedStorageType = errors.New("storage type '%s' is wot supported")
)
