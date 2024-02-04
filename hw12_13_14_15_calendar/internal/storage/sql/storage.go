package sqlstorage

import (
	"context"
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/logger"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/storage/timeutils"
	"github.com/jackc/pgx/v4/pgxpool"
)

type config interface {
	BuildDBUrl() string
}

type Storage struct {
	ctx  context.Context
	cfg  config
	conn *pgxpool.Pool
	logg logger.Logger
}

func New(ctx context.Context, cfg config, logg logger.Logger) *Storage {
	return &Storage{ctx: ctx, cfg: cfg, logg: logg}
}

func (s *Storage) Connect(ctx context.Context) error {
	pool, err := pgxpool.Connect(ctx, s.cfg.BuildDBUrl())
	if err != nil {
		return err
	}
	s.conn = pool

	s.logg.Info("Connected to SQL database at " + s.cfg.BuildDBUrl())
	return nil
}

func (s *Storage) Close(_ context.Context) error {
	s.conn.Close()

	return nil
}

func (s *Storage) CreateEvent(evt *domain.Event) (*domain.Event, error) {
	q := `INSERT INTO
		events(title, description, start_time, duration, user_id, remind_for, notified)
		VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`
	var id domain.EventID
	if err := s.conn.QueryRow(
		s.ctx, q, evt.Title, evt.Description, evt.StartTime, evt.Duration, evt.UserID, evt.NotifyBefore, evt.Notified,
	).Scan(&id); err != nil {
		var newEvt *domain.Event
		return newEvt, err
	}

	return s.GetEventByID(id)
}

func (s *Storage) DeleteEvent(id domain.EventID) error {
	query := `DELETE FROM events WHERE id = $1`

	if _, err := s.conn.Exec(s.ctx, query, id); err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateEvent(id domain.EventID, evt *domain.Event) error {
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

func (s *Storage) GetEventByID(id domain.EventID) (*domain.Event, error) {
	q := `SELECT
		title, description, start_time, duration, user_id, remind_for, notified
		FROM events
		WHERE id = $1`

	var evt domain.Event
	if err := s.conn.QueryRow(s.ctx, q, id).Scan(
		&evt.Title, &evt.Description, &evt.StartTime, &evt.Duration, &evt.UserID, &evt.NotifyBefore, &evt.Notified,
	); err != nil {
		return &domain.Event{}, err
	}

	return &evt, nil
}

func (s *Storage) GetAllEvents() ([]domain.Event, error) {
	return make([]domain.Event, 0), nil
}

func (s *Storage) GetEventsByDate(date time.Time) ([]*domain.Event, error) {
	startDate := timeutils.BeginningOfDay(date)
	endDate := timeutils.EndOfDay(date)
	return s.getEventsForPeriod(startDate, endDate)
}

func (s *Storage) GetEventsByWeek(startOfWeek time.Time) ([]*domain.Event, error) {
	startDate := timeutils.BeginningOfDay(startOfWeek)
	endDate := startOfWeek.AddDate(0, 0, timeutils.DaysInWeek)
	return s.getEventsForPeriod(startDate, endDate)
}

func (s *Storage) GetEventsByMonth(startOfMonth time.Time) ([]*domain.Event, error) {
	startDate := timeutils.BeginningOfDay(startOfMonth)
	endDate := startOfMonth.AddDate(0, 0, timeutils.DaysInMonth)
	return s.getEventsForPeriod(startDate, endDate)
}

func (s *Storage) getEventsForPeriod(startDate time.Time, endDate time.Time) ([]*domain.Event, error) {
	q := `SELECT 
    	title, description, start_time, duration, user_id, remind_for, notified 
	FROM
		events 
	WHERE 
	    start_time >= $1 AND start_time < $2 
	ORDER BY 
	    start_time
	`
	events := make([]*domain.Event, 0)
	rows, err := s.conn.Query(s.ctx, q, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var evt domain.Event
		if err := rows.Scan(
			&evt.Title, &evt.Description, &evt.StartTime, &evt.Duration, &evt.UserID, &evt.NotifyBefore, &evt.Notified,
		); err != nil {
			s.logg.Warn("failed to scan row: " + err.Error())
		} else {
			events = append(events, &evt)
		}
	}
	if err := rows.Err(); err != nil {
		s.logg.Error("failed to scan rows: " + err.Error())
		return events, err
	}

	return events, nil
}

// Compile-time check that Storage implements storage.Storage.
var (
	_ storage.EventsRepository = &Storage{}
	_ storage.EventsRepository = (*Storage)(nil)
)
