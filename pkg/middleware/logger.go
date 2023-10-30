package middleware

import (
	"AdHub/pkg/logger"
	"net/http"
	"time"
)

func Logger(log logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			next.ServeHTTP(w, r)

			duration := time.Since(startTime)

			log.MW("Request handled", r, duration)
		})
	}
}
