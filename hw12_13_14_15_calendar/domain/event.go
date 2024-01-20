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
	UserId       UserId
	NotifyBefore *time.Duration
	Notified     bool
}

type UserId string
