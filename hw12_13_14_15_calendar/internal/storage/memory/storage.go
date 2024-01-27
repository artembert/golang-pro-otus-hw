package memorystorage

import (
	"sync"
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
	"github.com/gofrs/uuid"
)

type eventsCollection map[domain.EventID]domain.Event

type Logger interface {
	Error(msg string)
}

type Storage struct {
	mu     sync.RWMutex
	events eventsCollection
	_      *Logger
}

func New() *Storage {
	return &Storage{
		events: make(eventsCollection),
	}
}

func (s *Storage) CreateEvent(evt domain.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := domain.EventID(uuid.UUID{}.String())
	evt.ID = id
	s.events[id] = evt

	return nil
}

func (s *Storage) DeleteEvent(id domain.EventID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[id]; !ok {
		return storage.ErrEventNotFound
	}
	delete(s.events, id)

	return nil
}

func (s *Storage) UpdateEvent(id domain.EventID, evt domain.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.events[id] = evt

	return nil
}

func (s *Storage) GetEventByID(id domain.EventID) (domain.Event, error) {
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

func (s *Storage) GetEventsByDate(date time.Time) ([]domain.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) GetEventsByWeek(startOfWeek time.Time) ([]domain.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) GetEventsByMonth(startOfMonth time.Time) ([]domain.Event, error) {
	//TODO implement me
	panic("implement me")
}

// Compile-time check that Storage implements storage.Storage.
var (
	_ storage.Actions = &Storage{}
	_ storage.Actions = (*Storage)(nil)
)
