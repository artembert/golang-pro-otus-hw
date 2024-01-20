package memorystorage

import (
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/domain/event"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/storage"
	"sync"
)

type eventsCollection map[string]event.Event

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

func (s *Storage) CreateEvent(evt event.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.events[evt.ID] = evt

	return nil
}

func (s *Storage) DeleteEvent(evt event.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.events, evt.ID)

	return nil
}

func (s *Storage) UpdateEvent(evt event.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.events[evt.ID] = evt

	return nil
}

func (s *Storage) GetEventByID(id string) (event.Event, error) {
	evt, ok := s.events[id]
	if !ok {
		return event.Event{}, storage.ErrEventNotFound
	}

	return evt, nil
}

func (s *Storage) GetAllEvents() ([]event.Event, error) {
	events := make([]event.Event, 0)

	for _, evt := range s.events {
		events = append(events, evt)
	}

	return events, nil
}

// Compile-time check that Storage implements storage.Storage
var _ storage.Storage = &Storage{}
var _ storage.Storage = (*Storage)(nil)
