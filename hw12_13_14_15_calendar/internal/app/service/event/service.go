package event

import (
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/domain"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/logger"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/storage"
)

type Service struct {
	logger *logger.Logger
	repo   storage.EventsRepository
}

func NewService(logger *logger.Logger, repo storage.EventsRepository) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
	}
}

func (service *Service) CreateEvent(evt *domain.Event) (domain.EventID, error) {
	return service.repo.CreateEvent(evt)
}
