package fabric

import (
	"context"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/domain/storage"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/storage/sql"
)

func Init(ctx context.Context, cfg *config.Config) (storage.Actions, error) {
	switch storage.Type(cfg.Storage.Type) {
	case storage.Memory:
		return memorystorage.New(), nil
	case storage.SQL:
		store := sqlstorage.New(ctx, &cfg.DB)
		return store, nil
	}
	return nil, storage.ErrUnsupportedStorageType
}
