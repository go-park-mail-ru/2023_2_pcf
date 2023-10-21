package logger

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	log := logrus.New()
	return &LogrusLogger{log: log}
}

func (l *LogrusLogger) Info(message string) {
	l.log.Info(message)
}

func (l *LogrusLogger) Error(message string) {
	l.log.Error(message)
}

func (l *LogrusLogger) Fatal(message string) {
	l.log.Fatal(message)
}

func (l *LogrusLogger) MW(message string, r *http.Request, duration time.Duration) {
	log := l.log.WithFields(logrus.Fields{
		"Method":     r.Method,
		"URL":        r.RequestURI,
		"RemoteAddr": r.RemoteAddr,
		"Duration":   duration,
	})
	log.Info(message)
}
