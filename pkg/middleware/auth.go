package middleware

import (
	"AdHub/internal/pkg/entities"
	"AdHub/proto/api"
	"context"
	"fmt"
	"net/http"
	"time"
)

func Auth(ss interface{}, csrfUc entities.CsrfUseCaseInterface) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ss, ok := ss.(api.SessionClient)
			if !ok {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

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
			fmt.Print("12345678\n")
			fmt.Println(userId)
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

			csrfFromDb, err := csrfUc.GetByUserId(int(userId.Id))
			if err != nil {
				http.Error(w, "CSRF err1", http.StatusForbidden)
				return
			}

			if csrfFromDb.Token != csrfToken.Value {
				http.Error(w, "CSRF err2", http.StatusForbidden)
				return
			}

			csrfUc.CsrfRemove(csrfFromDb)
			newCsrf, err := csrfUc.CsrfCreate(int(userId.Id))
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
