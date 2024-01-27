package sqlstorage

import (
	"context"
	"time"

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
	pool, err := pgxpool.Connect(ctx, s.cfg.BuildDBUrl())
	if err != nil {
		return err
	}
	s.conn = pool

	return nil
}

func (s *Storage) Close(_ context.Context) error {
	s.conn.Close()

	return nil
}

func (s *Storage) CreateEvent(evt domain.Event) error {
	q := `INSERT INTO
		events(title, description, start_time, duration, user_id, remind_for, notified)
		VALUES($1, $2, $3, $4, $5, $6, $7)`
	if _, err := s.conn.Exec(
		s.ctx, q, evt.Title, evt.Description, evt.StartTime, evt.Duration, evt.UserID, evt.NotifyBefore, evt.Notified,
	); err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteEvent(id domain.EventID) error {
	query := `DELETE FROM events WHERE id = $1`

	if _, err := s.conn.Exec(s.ctx, query, id); err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateEvent(id domain.EventID, evt domain.Event) error {
	q := `UPDATE events set
		title = $1,
		description = $2,
		start_time = $3,
		duration = $4,
		user_id = $5,
		remind_for = $6,
		notified = $7
	  	WHERE id = $8`
	if _, err := s.conn.Exec(
		s.ctx,
		q,
		evt.Title,
		evt.Description,
		evt.StartTime,
		evt.Duration,
		evt.UserID,
		evt.NotifyBefore,
		evt.Notified,
		id,
	); err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetEventByID(id domain.EventID) (domain.Event, error) {
	q := `SELECT
		title, description, start_time, duration, user_id, remind_for, notified
		FROM events
		WHERE id = $1`

	var evt domain.Event
	if err := s.conn.QueryRow(s.ctx, q, id).Scan(
		&evt.Title, &evt.Description, &evt.StartTime, &evt.Duration, &evt.UserID, &evt.NotifyBefore, &evt.Notified,
	); err != nil {
		return domain.Event{}, err
	}

	return evt, nil
}

func (s *Storage) GetAllEvents() ([]domain.Event, error) {
	return make([]domain.Event, 0), nil
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
