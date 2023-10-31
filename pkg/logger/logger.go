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

//go:generate /Users/bincom/go/bin/mockgen -source=logger.go -destination=mock_logger/mock.go
type Logger interface {
	Info(message string)
	Error(message string)
	Fatal(message string)
	MW(message string, r *http.Request, duration time.Duration)
	SetLogLevel(level string)
}
