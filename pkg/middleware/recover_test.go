package middleware

import (
	"AdHub/pkg/logger/mock_logger"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestRecoverMiddleware(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := mock_logger.NewMockLogger(ctrl)

	mockLogger.EXPECT().Fatal(gomock.Any())

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	middleware := Recover(mockLogger)

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("Simulated panic")
	}))

	handler.ServeHTTP(w, req)
}

func TestRecoverMiddlewareNoPanic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := mock_logger.NewMockLogger(ctrl)

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	middleware := Recover(mockLogger)

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	handler.ServeHTTP(w, req)
}
