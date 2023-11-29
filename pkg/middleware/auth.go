package middleware

import (
	"AdHub/internal/pkg/entities"
	"AdHub/proto/api"
	"context"
	"net/http"
)

func Auth(ss api.SessionClient, csrfUc entities.CsrfUseCaseInterface) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api/v1/auth" {
				next.ServeHTTP(w, r)
				return
			}
			if r.URL.Path == "/api/v1/user" {
				next.ServeHTTP(w, r)
				return
			}

			sessionToken, err := r.Cookie("session_token")
			if err != nil {
				http.Error(w, "Not authorized", http.StatusUnauthorized)
				return
			}

			userId, err := ss.GetUserId(context.Background(), &api.GetRequest{Token: sessionToken.Value})
			if err != nil {
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "userId", int(userId.Id))

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

}
