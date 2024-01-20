package storage

import (
	"errors"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
)

type Type string

const (
	Memory Type = "memory"
	SQL    Type = "sql"
)

type Actions interface {
	CreateEvent(evt domain.Event) error
	DeleteEvent(evt domain.Event) error
	UpdateEvent(evt domain.Event) error
	GetEventByID(id string) (domain.Event, error)
	GetAllEvents() ([]domain.Event, error)
}

var (
	ErrEventNotFound          = errors.New("event by id '%s' not found")
	ErrEventAlreadyExist      = errors.New("event by id '%s' already exist")
	ErrUnsupportedStorageType = errors.New("storage type '%s' is wot supported")
)
