package middleware

import (
	"AdHub/internal/pkg/entities"
	"context"
	"net/http"
)

func Auth(ss entities.SessionUseCaseInterface) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/auth" {
				next.ServeHTTP(w, r)
				return
			}

			sessionToken, err := r.Cookie("session_token")
			if err != nil {
				http.Error(w, "Not authorized", http.StatusUnauthorized)
				return
			}

			userId, err := ss.GetUserId(sessionToken.Value)
			if err != nil {
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "userId", userId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

}
