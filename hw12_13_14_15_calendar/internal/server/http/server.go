package internalhttp

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/server/http/event"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/pkg/api/openapi"
	"github.com/go-chi/chi/v5"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/logger"
	errorspkg "github.com/pkg/errors"
)

type Config struct {
	Host              string
	Port              string
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
}

type Server struct {
	server *http.Server
	logger logger.Logger
}

type Logger interface {
	logger.Logger
}

func New(logger Logger, app app.Application, cfg Config) *Server {
	handler := event.NewEventHandler(app, logger)
	router := chi.NewRouter()
	router.Use(loggingMiddleware(logger))
	openapi.HandlerFromMux(handler, router)
	addr := net.JoinHostPort(cfg.Host, cfg.Port)

	return &Server{
		server: &http.Server{
			Addr:              addr,
			Handler:           router,
			ReadHeaderTimeout: cfg.ReadHeaderTimeout,
			ReadTimeout:       cfg.ReadTimeout,
			WriteTimeout:      cfg.WriteTimeout,
		},
		logger: logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("starting server at", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			s.logger.Info("http server stopped gracefully")
			return nil
		}
		return errorspkg.Wrap(err, "error while starting server")
	}

	<-ctx.Done()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("stopping server")
	if err := s.server.Shutdown(ctx); err != nil {
		return errorspkg.Wrap(err, "error while stopping server")
	}

	return nil
}
