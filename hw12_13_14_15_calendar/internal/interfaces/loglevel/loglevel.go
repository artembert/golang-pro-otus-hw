package loglevel

type LogLevel string

const (
	Debug LogLevel = "debug"
	Error LogLevel = "error"
	Warn  LogLevel = "warn"
	Info  LogLevel = "info"
)
