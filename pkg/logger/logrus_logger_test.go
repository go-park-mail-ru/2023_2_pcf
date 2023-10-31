package logger

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogrusLogger_Info(t *testing.T) {
	var buf bytes.Buffer
	logger := logrus.New()
	logger.SetOutput(&buf)

	logrusLogger := &LogrusLogger{log: logger}

	message := "Info message"
	logrusLogger.Info(message)

	assert.Contains(t, buf.String(), message)
}

func TestLogrusLogger_Error(t *testing.T) {
	var buf bytes.Buffer
	logger := logrus.New()
	logger.SetOutput(&buf)

	logrusLogger := &LogrusLogger{log: logger}

	message := "Error message"
	logrusLogger.Error(message)

	assert.Contains(t, buf.String(), message)
}

func TestLogrusLogger_MW(t *testing.T) {
	var buf bytes.Buffer
	logger := logrus.New()
	logger.SetOutput(&buf)

	logrusLogger := &LogrusLogger{log: logger}

	message := "MW message"
	request, _ := http.NewRequest("GET", "http://example.com", nil)
	duration := time.Second
	logrusLogger.MW(message, request, duration)

	assert.Contains(t, buf.String(), message)
}

func TestLogrusLogger_SetLogLevel(t *testing.T) {
	logger := logrus.New()
	logrusLogger := &LogrusLogger{log: logger}

	logrusLogger.SetLogLevel(DebugLevel)
	assert.Equal(t, logrus.DebugLevel, logger.GetLevel())

	logrusLogger.SetLogLevel(InfoLevel)
	assert.Equal(t, logrus.InfoLevel, logger.GetLevel())

	logrusLogger.SetLogLevel(ErrorLevel)
	assert.Equal(t, logrus.ErrorLevel, logger.GetLevel())

	logrusLogger.SetLogLevel(FatalLevel)
	assert.Equal(t, logrus.FatalLevel, logger.GetLevel())
}
