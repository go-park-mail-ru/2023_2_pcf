package middleware

import (
	"AdHub/pkg/logger"
	"fmt"
	"net/http"
	"runtime/debug"
)

func Recover(log logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					debug.PrintStack()
					log.Fatal(fmt.Sprintf("Panic: %v\n", err))

					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
