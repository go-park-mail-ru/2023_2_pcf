package logger

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLogger(level string) *LogrusLogger {
	log := logrus.New()

	l := &LogrusLogger{log: log}
	l.SetLevel(level)

	return l
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

func (l *LogrusLogger) SetLevel(level string) {
	if level == "debug" {
		l.log.SetLevel(logrus.DebugLevel)
	} else if level == "info" {
		l.log.SetLevel(logrus.InfoLevel)
	} else if level == "error" {
		l.log.SetLevel(logrus.ErrorLevel)
	} else if level == "fatal" {
		l.log.SetLevel(logrus.FatalLevel)
	}
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
