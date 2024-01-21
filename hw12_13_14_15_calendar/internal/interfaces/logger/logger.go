package logger

import (
	"errors"
	"net/http"
)

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Debug(args ...interface{})
	HTTPRequest(r *http.Request, args ...interface{})
}

var (
	ErrLoggerOutputFile  = errors.New("unable to create log file")
	ErrLoggerEmptyOutput = errors.New("output path is empty")
)
