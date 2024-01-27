package domain

import (
	"time"
)

type Event struct {
	ID           EventID
	Title        string
	StartTime    time.Time
	Duration     time.Duration
	Description  string
	UserID       UserID
	NotifyBefore time.Duration
	Notified     bool
}

type UserID string

type EventID string
