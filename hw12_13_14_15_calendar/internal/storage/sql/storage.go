package sqlstorage

import (
	"context"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Config interface {
	DriverName() string
	GetDatabaseConnectionString() string
	MigrationDir() string
}

type Logger interface {
	Error(msg string)
}

type Storage struct {
	conn *pgxpool.Pool
	log  Logger
	cfg  Config
}

func New(conn *pgxpool.Pool, cfg Config, log Logger) *Storage {
	return &Storage{conn, log, cfg}
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

func (s *Storage) CreateEvent(evt storage.Event) error {
	return nil
}

func (s *Storage) DeleteEvent(evt storage.Event) error {
	return nil
}

func (s *Storage) UpdateEvent(evt storage.Event) error {
	return nil
}

func (s *Storage) GetEventByID(id string) (storage.Event, error) {
	return storage.Event{}, nil
}

func (s *Storage) GetAllEvents() ([]storage.Event, error) {
	return make([]storage.Event, 0), nil
}

// Compile-time check that Storage implements storage.EventStorage
var _ storage.EventStorage = &Storage{}
var _ storage.EventStorage = (*Storage)(nil)
