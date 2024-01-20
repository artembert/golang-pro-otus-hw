package main

import (
	"context"
	"flag"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/logger_uber"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/storage/fabric"
	"log"
	"os/signal"
	"syscall"
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
		log.Fatalf("failed to parse config %s: %s", configFile, err)
	}

	//logFile, err := logger.GetOutputFile(cfg.Logger.OutputPath)
	//if err != nil {
	//	log.Fatalf("failed to resolve log output %s: %s", cfg.Logger.OutputPath, err)
	//}
	//logg, err := logger.New(cfg.Logger.Level, logFile)
	logg, err := logger_uber.New(cfg.Logger.Level, cfg.Logger.OutputPath)

	store, err := fabric.Init(ctx, &cfg)
	if err != nil {
		log.Fatalf("failed to connect to srorage %s, %s", cfg.Storage.Type, err)
	}

	calendar := app.New(ctx, logg, store)
	_ = calendar
	//server := internalhttp.NewServer(logg, calendar)

	//go func() {
	//	<-ctx.Done()
	//
	//	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	//	defer cancel()
	//
	//	if err := server.Stop(ctx); err != nil {
	//		logg.Error("failed to stop http server: " + err.Error())
	//	}
	//}()

	//logg.Info("calendar is running...")

	//if err := server.Start(ctx); err != nil {
	//logg.Error("failed to start http server: " + err.Error())
	//cancel()
	//os.Exit(1) //nolint:gocritic
	//}
}
