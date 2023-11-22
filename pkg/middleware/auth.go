package middleware

import (
	"AdHub/internal/pkg/entities"
	"context"
	"net/http"
	"time"
)

func Auth(ss entities.SessionUseCaseInterface, csrfUc entities.CsrfUseCaseInterface) func(next http.Handler) http.Handler {
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

			userId, err := ss.GetUserId(sessionToken.Value)
			if err != nil {
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}

			//csrf logic

			csrfToken, err := r.Cookie("csrf_token")
			if err != nil {
				http.Error(w, "CSRF required", http.StatusUnauthorized)
				return
			}

			csrfFromDb, err := csrfUc.GetByUserId(userId)
			if err != nil {
				http.Error(w, "CSRF err", http.StatusForbidden)
				return
			}

			if csrfFromDb.Token != csrfToken.Value {
				http.Error(w, "CSRF err", http.StatusForbidden)
				return
			}

			csrfUc.CsrfRemove(csrfFromDb)
			newCsrf, err := csrfUc.CsrfCreate(userId)
			if err != nil {
				http.Error(w, "err", http.StatusInternalServerError)
				return
			}
			cookie := &http.Cookie{
				Name:     "csrf_token",
				Value:    newCsrf.Token,
				Expires:  time.Now().Add(24 * time.Hour),
				HttpOnly: true,
				Domain:   "127.0.0.1",
				Path:     "/",
			}
			http.SetCookie(w, cookie)

			ctx := context.WithValue(r.Context(), "userId", userId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

}
