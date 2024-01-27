package loggerzap

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/logger"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/loglevel"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zap *zap.SugaredLogger
}

func createFileIfNotExists(path string) error {
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0o755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func Factory(level loglevel.LogLevel, outputPath string) (logger.Logger, error) {
	return New(level, outputPath)
}

func New(level loglevel.LogLevel, outputPath string) (*Logger, error) {
	if outputPath == "" {
		return nil, logger.ErrLoggerEmptyOutput
	}
	if err := createFileIfNotExists(outputPath); err != nil {
		return nil, errors.Wrap(logger.ErrLoggerOutputFile, err.Error())
	}

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
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.OutputPaths = []string{"stdout", outputPath}

	options := []zap.Option{zap.AddCallerSkip(1)}

	logg, err := cfg.Build(options...)
	if err != nil {
		return nil, err
	}

	return &Logger{
		zap: logg.Sugar(),
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

func (logger Logger) HTTPRequest(r *http.Request, data ...interface{}) {
	logger.zap.With("IP", r.RemoteAddr).Info(data)
}

// Compile-time check that Logger implements logger.Logger.
var (
	_ logger.Logger = &Logger{}
	_ logger.Logger = (*Logger)(nil)
)
