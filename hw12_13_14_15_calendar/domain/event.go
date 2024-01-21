package domain

import (
	"time"
)

type Event struct {
	ID           string
	Title        string
	StartTime    time.Time
	Duration     time.Duration
	Description  *string
	UserID       UserID
	NotifyBefore *time.Duration
	Notified     bool
}

type UserID string
