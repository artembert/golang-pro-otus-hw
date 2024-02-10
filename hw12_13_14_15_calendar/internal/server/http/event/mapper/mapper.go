package mapper

import (
	"errors"
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/pkg/api/openapi"
)

var (
	ErrorDurationParse     = errors.New("error parsing duration")
	ErrorNotifyBeforeParse = errors.New("error parsing notifyBefore")
)

func EventToOpenapi(event *domain.Event) *openapi.Event {
	evt := openapi.Event{
		Description:  event.Description,
		Duration:     event.Duration.String(),
		Id:           string(event.ID),
		NotifyBefore: event.NotifyBefore.String(),
		StartTime:    event.StartTime,
		Title:        event.Title,
	}

	return &evt
}

func OpenapiNewEventToEvent(event *openapi.NewEvent, userId string) (*domain.Event, error) {
	duration, err := time.ParseDuration(event.Duration)
	if err != nil {
		return nil, ErrorDurationParse
	}
	notifyBefore, err := time.ParseDuration(event.NotifyBefore)
	if err != nil {
		return nil, ErrorNotifyBeforeParse
	}

	evt := domain.Event{
		Title:        event.Title,
		StartTime:    event.StartTime,
		Duration:     duration,
		Description:  event.Description,
		UserID:       domain.UserID(userId),
		NotifyBefore: notifyBefore,
	}

	return &evt, nil
}

func OpenapiEventToEvent(event *openapi.Event, userId string) (*domain.Event, error) {
	newEvent := openapi.NewEvent{
		Description:  event.Description,
		Duration:     event.Duration,
		NotifyBefore: event.NotifyBefore,
		StartTime:    event.StartTime,
		Title:        event.Title,
	}
	evt, err := OpenapiNewEventToEvent(&newEvent, userId)
	if err != nil {
		return nil, err
	}

	evt.ID = domain.EventID(event.Id)

	return evt, nil
}
