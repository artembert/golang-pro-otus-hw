package mapper

import (
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/pkg/api/openapi"
)

func EventToOpenapi(event *domain.Event) *openapi.Event {
	evt := openapi.Event{
		Description:  event.Description,
		Duration:     int(event.Duration / time.Minute),
		Id:           string(event.ID),
		NotifyBefore: int(event.NotifyBefore / time.Minute),
		StartTime:    event.StartTime, // TODO: Replace with time.MarshalText()
		Title:        event.Title,
	}

	return &evt
}
