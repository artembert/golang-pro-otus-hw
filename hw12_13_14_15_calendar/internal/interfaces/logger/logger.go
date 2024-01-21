package logger

import "net/http"

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Debug(args ...interface{})
	HTTPRequest(r *http.Request, args ...interface{})
}
