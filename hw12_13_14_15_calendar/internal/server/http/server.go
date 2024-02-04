package internalhttp

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

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
			Addr:              addr,
			Handler:           serveMux,
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

func initRoutes(app Application, logger Logger) *http.ServeMux {
	handler := Handler{App: app}
	requestHandler := http.NewServeMux()
	requestHandler.HandleFunc("/hello", loggingMiddleware(http.HandlerFunc(handler.Hello), logger))

	return requestHandler
}

func (h *Handler) Hello(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Hello, it's calendar!"))
}
