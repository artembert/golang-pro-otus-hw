package handler

import (
	"context"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app/event/command"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/server/grpc/handler/mapper"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/pkg/api/grpc/eventsservice"
)

func (handler *eventHandler) CreateEvent(ctx context.Context, req *eventsservice.CreateEventRequest) (*eventsservice.CreateEventResponse, error) {
	createEventRequest := command.CreateEventRequest{
		Event: &domain.Event{
			Title:        req.Event.GetTitle(),
			StartTime:    req.Event.GetStartTime().AsTime(),
			Duration:     req.Event.GetDuration().AsDuration(),
			Description:  req.Event.GetDescription(),
			UserID:       domain.UserID(req.Event.GetUserId()),
			NotifyBefore: req.Event.GetNotifyBefore().AsDuration(),
		},
	}
	res, err := handler.app.CreateEvent(createEventRequest)
	if err != nil {
		return nil, err
	}
	return &eventsservice.CreateEventResponse{Event: mapper.EventToProtoEvent(&res.Event)}, nil
}
