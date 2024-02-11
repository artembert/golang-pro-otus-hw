package internalgrpc

import (
	"context"
	"net"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/logger"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/server/grpc/handler"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/pkg/api/grpc/eventsservice"
	errorspkg "github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Config struct {
	Host string
	Port string
}

type Server struct {
	eventsservice.UnimplementedCalendarServer
	application app.Application
	grpcServer  *grpc.Server
	logger      logger.Logger
	addr        string
}

type Logger interface {
	logger.Logger
}

func New(logger Logger, app app.Application, cfg Config) *Server {
	server := grpc.NewServer() // TODO: add logg interceptor
	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	eventsservice.RegisterCalendarServer(server, handler.New(app, logger))

	return &Server{
		grpcServer: server,
		logger:     logger,
		addr:       addr,
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("starting GRPC server at", s.addr)
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return errorspkg.Wrap(err, "error while starting gRPC server")
	}
	err = s.grpcServer.Serve(listener)
	if err != nil {
		return errorspkg.Wrap(err, "error while starting gRPC server")
	}

	<-ctx.Done()

	return nil
}

func (s *Server) Stop() {
	s.logger.Info("stopping gRPC server")
	s.grpcServer.GracefulStop()
}
