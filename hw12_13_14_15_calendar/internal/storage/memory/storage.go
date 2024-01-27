package memorystorage

import (
	"sync"
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/storage/timeutils"
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

func (s *Storage) CreateEvent(evt domain.Event) (domain.EventID, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	uuidV4 := uuid.Must(uuid.NewV4())
	id := domain.EventID(uuidV4.String())
	evt.ID = id
	s.events[id] = evt

	return id, nil
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
	startDate := timeutils.BeginningOfDay(date)
	endDate := timeutils.EndOfDay(date)
	return s.getEventsForPeriod(startDate, endDate)
}

func (s *Storage) GetEventsByWeek(startOfWeek time.Time) ([]domain.Event, error) {
	startDate := timeutils.BeginningOfDay(startOfWeek)
	endDate := startOfWeek.AddDate(0, 0, timeutils.DaysInWeek)
	return s.getEventsForPeriod(startDate, endDate)
}

func (s *Storage) GetEventsByMonth(startOfMonth time.Time) ([]domain.Event, error) {
	startDate := timeutils.BeginningOfDay(startOfMonth)
	endDate := startOfMonth.AddDate(0, 0, timeutils.DaysInMonth)
	return s.getEventsForPeriod(startDate, endDate)
}

func (s *Storage) getEventsForPeriod(startDate time.Time, endDate time.Time) ([]domain.Event, error) {
	events := make([]domain.Event, 0)
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, evt := range s.events {
		if evt.StartTime.After(startDate) && evt.StartTime.Before(endDate) {
			events = append(events, evt)
		}
	}
	return events, nil
}

// Compile-time check that Storage implements storage.Storage.
var (
	_ storage.EventsRepository = &Storage{}
	_ storage.EventsRepository = (*Storage)(nil)
)
