package storage

import "time"

type Event struct {
	ID           string
	Title        string
	StartTime    time.Time
	Duration     time.Duration
	Description  *string
	UserId       UserId
	NotifyBefore *time.Duration
}

type EventStorage interface {
	CreateEvent(evt Event) error
	DeleteEvent(evt Event) error
	UpdateEvent(evt Event) error
	GetEventByID(id string) error
	GetAllEvents() ([]Event, error)
}

type UserId string
