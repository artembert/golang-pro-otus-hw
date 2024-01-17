package memorystorage

import (
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/storage"
	"sync"
)

type eventsCollection map[string]storage.Event

type Logger interface {
	Error(msg string)
}

type Storage struct {
	// TODO
	mu     sync.RWMutex //nolint:unused
	events eventsCollection
	log    Logger
}

func New(log Logger) *Storage {
	return &Storage{
		events: make(eventsCollection),
		log:    log,
	}
}

func (s *Storage) CreateEvent(evt storage.Event) error {
	return nil
}

func (s *Storage) DeleteEvent(evt storage.Event) error {
	return nil
}

func (s *Storage) UpdateEvent(evt storage.Event) error {
	return nil
}

func (s *Storage) GetEventByID(id string) error {
	return nil
}

func (s *Storage) GetAllEvents() ([]storage.Event, error) {
	return make([]storage.Event, 0), nil
}
