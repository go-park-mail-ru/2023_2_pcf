package logger

import (
	"net/http"
	"time"
)

type Logger interface {
	Info(message string)
	Error(message string)
	Fatal(message string)
	MW(message string, r *http.Request, duration time.Duration)
	SetLogLevel(level level)
}

type level struct {
	level LogrusLogger.level
}
