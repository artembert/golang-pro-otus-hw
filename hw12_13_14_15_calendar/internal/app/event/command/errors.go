package command

import (
	//	"context"
	"errors"
	//	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/logger"
	//	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
)

var ErrStorageIsUndefined = errors.New("storage is undefined")

//
//type Service interface {
//	CreateEvent(ctx context.Context, request command.CreateEventRequest) (*command.CreateEventResponse, error)
//	UpdateEvent(ctx context.Context, request command.UpdateEventRequest) error
//	DeleteEvent(ctx context.Context, request command.DeleteEventRequest) error
//	GetDayEvents(ctx context.Context, request query.GetDayEventsRequest) (*query.GetDayEventsResponse, error)
//	GetWeekEvents(ctx context.Context, request query.GetWeekEventsRequest) (*query.GetWeekEventsResponse, error)
//	GetMonthEvents(ctx context.Context, request query.GetMonthEventsRequest) (*query.GetMonthEventsResponse, error)
//}
//
//type command struct {
//	logger *logger.Logger
//	repo   storage.EventsRepository
//}
//
//func NewService(logger *logger.Logger, repo storage.EventsRepository) *Service {
//	return &command{
//		logger: logger,
//		repo:   repo,
//	}
//}
//
////func (command *command) CreateEvent(evt *domain.Event) (domain.EventID, error) {
////	return command.repo.CreateEvent(evt)
////}
