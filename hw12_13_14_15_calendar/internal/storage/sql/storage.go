package sqlstorage

import (
	"context"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type config interface {
	BuildDBUrl() string
}

type Storage struct {
	ctx  context.Context
	cfg  config
	conn *pgxpool.Pool
}

func New(ctx context.Context, cfg config) *Storage {
	return &Storage{ctx: ctx, cfg: cfg}
}

func (s *Storage) Connect(ctx context.Context) error {
	pool, err := pgxpool.Connect(
		ctx,
		s.cfg.BuildDBUrl(),
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

func (s *Storage) CreateEvent(evt domain.Event) error {
	return nil
}

func (s *Storage) DeleteEvent(evt domain.Event) error {
	return nil
}

func (s *Storage) UpdateEvent(evt domain.Event) error {
	return nil
}

func (s *Storage) GetEventByID(id string) (domain.Event, error) {
	return domain.Event{}, nil
}

func (s *Storage) GetAllEvents() ([]domain.Event, error) {
	return make([]domain.Event, 0), nil
}

// Compile-time check that Storage implements storage.Storage
var _ storage.Actions = &Storage{}
var _ storage.Actions = (*Storage)(nil)
