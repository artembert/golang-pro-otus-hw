package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/pkg/loggerzap"
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

	logg, err := loggerzap.New(cfg.Logger.Level, cfg.Logger.OutputPath)
	if err != nil {
		log.Printf("failed to init logger: %s", err)
		os.Exit(1) //nolint:gocritic
	}

	store, err := factory.Init(ctx, &cfg)
	if err != nil {
		log.Printf("failed to connect to storage %s, %s", cfg.Storage.Type, err)
		os.Exit(1) //nolint:gocritic
	}

	calendar := app.New(ctx, logg, store)

	server := internalhttp.New(logg, calendar, internalhttp.Config(cfg.Server))

	go func() {
		<-ctx.Done()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
