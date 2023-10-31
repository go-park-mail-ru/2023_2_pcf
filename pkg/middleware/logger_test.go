package middleware

import (
	"AdHub/pkg/logger/mock_logger"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestLoggerMiddleware(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := mock_logger.NewMockLogger(ctrl)

	mockLogger.EXPECT().MW(gomock.Any(), gomock.Any(), gomock.Any())

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	middleware := Logger(mockLogger)

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	handler.ServeHTTP(w, req)
}
