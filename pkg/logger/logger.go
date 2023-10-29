package logger

import (
	"net/http"
	"time"
)

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	ErrorLevel = "error"
	FatalLevel = "fatal"
)

type Logger interface {
	Info(message string)
	Error(message string)
	Fatal(message string)
	MW(message string, r *http.Request, duration time.Duration)
	SetLogLevel(level string)
}
