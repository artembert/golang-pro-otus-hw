package internalhttp

import (
	"context"
	"net"
	"net/http"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/logger"
	"github.com/pkg/errors"
)

type Config struct {
	Host string
	Port string
}

type Server struct {
	server *http.Server
	logger logger.Logger
}

type Logger interface {
	logger.Logger
}

type Application interface { // TODO
}

type Handler struct {
	App Application
}

func New(logger Logger, app Application, cfg Config) *Server {
	serveMux := initRoutes(app, logger)
	addr := net.JoinHostPort(cfg.Host, cfg.Port)

	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: serveMux,
		},
		logger: logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("starting server")
	if err := s.server.ListenAndServe(); err != nil {
		return errors.Wrap(err, "error while starting server")
	}

	<-ctx.Done()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("stopping server")
	if err := s.server.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "error while stopping server")
	}

	return nil
}

func initRoutes(app Application, logger Logger) *http.ServeMux {
	handler := Handler{App: app}
	requestHandler := http.NewServeMux()
	requestHandler.HandleFunc("/hello", loggingMiddleware(http.HandlerFunc(handler.Hello), logger))

	return requestHandler
}

func (h *Handler) Hello(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Hello, it's calendar!"))
}
