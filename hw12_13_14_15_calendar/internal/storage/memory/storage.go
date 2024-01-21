package memorystorage

import (
	"sync"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
)

type eventsCollection map[string]domain.Event

type Logger interface {
	Error(msg string)
}

type Storage struct {
	mu     sync.RWMutex
	events eventsCollection
	log    *Logger
}

func New() *Storage {
	return &Storage{
		events: make(eventsCollection),
	}
}

func (s *Storage) CreateEvent(evt domain.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.events[evt.ID] = evt

	return nil
}

func (s *Storage) DeleteEvent(evt domain.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.events, evt.ID)

	return nil
}

func (s *Storage) UpdateEvent(evt domain.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.events[evt.ID] = evt

	return nil
}

func (s *Storage) GetEventByID(id string) (domain.Event, error) {
	evt, ok := s.events[id]
	if !ok {
		return domain.Event{}, storage.ErrEventNotFound
	}

	return evt, nil
}

func (s *Storage) GetAllEvents() ([]domain.Event, error) {
	events := make([]domain.Event, 0)

	for _, evt := range s.events {
		events = append(events, evt)
	}

	return events, nil
}

// Compile-time check that Storage implements storage.Storage.
var (
	_ storage.Actions = &Storage{}
	_ storage.Actions = (*Storage)(nil)
)
