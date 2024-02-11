package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/pkg/loggerzap"
	internalgrpc "github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/server/grpc"
	internalhttp "github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/server/http"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/storage/factory"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/config.yaml", "Path to configuration file")
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	cfg, err := config.New(configFile)
	if err != nil {
		log.Printf("failed to parse config %s: %s", configFile, err)
		os.Exit(1) //nolint:gocritic
	}

	logg, err := loggerzap.Factory(cfg.Logger.Level, cfg.Logger.OutputPath)
	if err != nil {
		log.Printf("failed to init logger: %s", err)
		os.Exit(1) //nolint:gocritic
	}

	store, err := storagefactory.Init(ctx, &cfg, logg)
	if err != nil {
		log.Printf("failed to init a storage %s, %s", cfg.Storage.Type, err)
		os.Exit(1) //nolint:gocritic
	}
	err = store.Connect(ctx)
	if err != nil {
		log.Printf("failed to connect to the storage %s, %s", cfg.Storage.Type, err)
		os.Exit(1) //nolint:gocritic
	}

	calendar, err := app.New(ctx, logg, store)
	if err != nil {
		log.Printf("failed to create an app %s", err)
		os.Exit(1) //nolint:gocritic
	}

	server := internalhttp.New(logg, calendar, internalhttp.Config{
		Host:              cfg.Server.Host,
		Port:              cfg.Server.Port,
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout,
		ReadTimeout:       cfg.Server.ReadTimeout,
		WriteTimeout:      cfg.Server.WriteTimeout,
	})
	grpcServer := internalgrpc.New(logg, calendar, internalgrpc.Config{
		Host: cfg.Server.Host,
		Port: cfg.Server.GRPCPort,
	})

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		<-ctx.Done()

		grpcServer.Stop()
		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	go func() {
		if err := server.Start(ctx); err != nil {
			logg.Error("failed to start http server: " + err.Error())
			cancel()

		}
		wg.Done()
	}()

	go func() {
		if err := grpcServer.Start(ctx); err != nil {
			logg.Error("failed to start grpc server: " + err.Error())
			cancel()
		}
		wg.Done()
	}()

	logg.Info("calendar is running...")
	wg.Wait()
}
