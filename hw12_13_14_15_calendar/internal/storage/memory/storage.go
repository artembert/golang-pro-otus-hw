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
	mu     sync.RWMutex
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
	s.mu.Lock()
	defer s.mu.Unlock()

	s.events[evt.ID] = evt

	return nil
}

func (s *Storage) DeleteEvent(evt storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.events, evt.ID)

	return nil
}

func (s *Storage) UpdateEvent(evt storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.events[evt.ID] = evt

	return nil
}

func (s *Storage) GetEventByID(id string) (storage.Event, error) {
	evt, ok := s.events[id]
	if !ok {
		return storage.Event{}, storage.ErrEventNotFound
	}

	return evt, nil
}

func (s *Storage) GetAllEvents() ([]storage.Event, error) {
	events := make([]storage.Event, 0)

	for _, evt := range s.events {
		events = append(events, evt)
	}

	return events, nil
}

// Compile-time check that Storage implements storage.EventStorage
var _ storage.EventStorage = &Storage{}
var _ storage.EventStorage = (*Storage)(nil)
