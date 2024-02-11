package mapper

import (
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/pkg/api/grpc/eventsservice"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func EventToProtoEvent(event *domain.Event) *eventsservice.Event {
	evt := eventsservice.Event{
		Title:        event.Title,
		Description:  event.Description,
		UserId:       string(event.UserID),
		StartTime:    timestamppb.New(event.StartTime),
		Duration:     durationpb.New(event.Duration),
		NotifyBefore: durationpb.New(event.NotifyBefore),
		Id:           string(event.ID),
		Notified:     false,
	}

	return &evt
}
