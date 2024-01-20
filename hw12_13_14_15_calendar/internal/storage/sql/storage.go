package sqlstorage

import (
	"context"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/domain/event"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Config interface {
	DriverName() string
	GetDatabaseConnectionString() string
	MigrationDir() string
}

type Storage struct {
	conn *pgxpool.Pool
	cfg  Config
}

func New(conn *pgxpool.Pool, cfg Config) *Storage {
	return &Storage{conn, cfg}
}

func (s *Storage) Connect(ctx context.Context) error {
	pool, err := pgxpool.Connect(
		ctx,
		s.cfg.GetDatabaseConnectionString(),
	)
	if err != nil {
		return err
	}
	s.conn = pool

	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	s.conn.Close()

	return nil
}

func (s *Storage) CreateEvent(evt event.Event) error {
	return nil
}

func (s *Storage) DeleteEvent(evt event.Event) error {
	return nil
}

func (s *Storage) UpdateEvent(evt event.Event) error {
	return nil
}

func (s *Storage) GetEventByID(id string) (event.Event, error) {
	return event.Event{}, nil
}

func (s *Storage) GetAllEvents() ([]event.Event, error) {
	return make([]event.Event, 0), nil
}

// Compile-time check that Storage implements storage.Storage
var _ storage.Storage = &Storage{}
var _ storage.Storage = (*Storage)(nil)
