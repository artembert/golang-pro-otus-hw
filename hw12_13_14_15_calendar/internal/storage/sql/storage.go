package sqlstorage

import (
	"context"
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

func (storage *Storage) Connect(ctx context.Context) error {
	pool, err := pgxpool.Connect(
		ctx,
		storage.cfg.GetDatabaseConnectionString(),
	)
	if err != nil {
		return err
	}
	storage.conn = pool

	return nil
}

func (storage *Storage) Close(ctx context.Context) error {
	storage.conn.Close()

	return nil
}
