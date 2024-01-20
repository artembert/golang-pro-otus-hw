package loggerzap

import (
	"fmt"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/loglevel"
	"go.uber.org/zap"
)

type Logger struct {
	zap *zap.SugaredLogger
}

func New(level loglevel.LogLevel, outputPath string) (*Logger, error) {
	l := zap.NewAtomicLevel()
	switch level {
	case loglevel.Debug:
		l.SetLevel(zap.DebugLevel)
	case loglevel.Error:
		l.SetLevel(zap.ErrorLevel)
	case loglevel.Warn:
		l.SetLevel(zap.WarnLevel)
	case loglevel.Info:
		l.SetLevel(zap.InfoLevel)
	default:
		return nil, fmt.Errorf("unsupported log level: %s", level)
	}
	cfg := zap.NewProductionConfig()
	cfg.Level = l
	cfg.OutputPaths = []string{"stdout", outputPath}
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	sugar := logger.Sugar()
	return &Logger{
		zap: sugar,
	}, nil
}

func (logger Logger) Info(data ...interface{}) {
	logger.zap.Info(data)
}

func (logger Logger) Error(data ...interface{}) {
	logger.zap.Error(data)
}

func (logger Logger) Debug(data ...interface{}) {
	logger.zap.Debug(data)
}

func (logger Logger) Warn(data ...interface{}) {
	logger.zap.Warn(data)
}
