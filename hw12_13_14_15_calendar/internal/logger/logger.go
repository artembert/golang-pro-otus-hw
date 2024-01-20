package logger

import (
	"fmt"
	abstractlogger "github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Logger interface {
	abstractlogger.Logger
	HTTPRequest(r *http.Request, args ...interface{})
}

type AppLogger struct {
	logger *logrus.Logger
}

// TODO: Choose logger implementation: logger OR logger_uber
func New(level string, output io.Writer) (Logger, error) {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	logger.SetLevel(logLevel)
	logger.SetOutput(output)

	return &AppLogger{logger}, nil
}

func (l *AppLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l AppLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l AppLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l AppLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l AppLogger) HTTPRequest(r *http.Request, args ...interface{}) {
	l.logger.WithFields(logrus.Fields{
		"IP": r.RemoteAddr,
	}).Info(args...)
}

func GetOutputFile(filePath string) (*os.File, error) {
	dirPath := filepath.Dir(filePath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0o700)
		if err != nil {
			return nil, fmt.Errorf("create dir %s: %w", dirPath, err)
		}
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return nil, err
	}

	return file, nil
}
