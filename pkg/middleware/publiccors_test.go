package middleware

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPub_CORS(t *testing.T) {
	// Mock handler that does nothing but is necessary for the middleware
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// Wrap the mock handler with the Pub_CORS middleware
	handlerWithCORS := Pub_CORS(mockHandler)

	// Test scenarios
	tests := []struct {
		name           string
		method         string
		wantStatusCode int
		wantHeaders    map[string]string
	}{
		{
			name:           "OPTIONS method",
			method:         http.MethodOptions,
			wantStatusCode: http.StatusOK,
			wantHeaders: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Methods":     "GET, POST, OPTIONS, DELETE",
				"Access-Control-Allow-Headers":     "Content-Type",
				"Access-Control-Allow-Credentials": "true",
			},
		},
		{
			name:           "GET method",
			method:         http.MethodGet,
			wantStatusCode: http.StatusOK, // Assuming next handler returns StatusOK
			wantHeaders: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Credentials": "true",
			},
		},
		// Add other HTTP methods if necessary
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request with the specified method
			req, err := http.NewRequest(tc.method, "/", nil)
			assert.NoError(t, err)

			// Record the response
			rr := httptest.NewRecorder()
			handlerWithCORS.ServeHTTP(rr, req)

			// Assert status code
			assert.Equal(t, tc.wantStatusCode, rr.Code)

			// Assert headers
			for key, wantValue := range tc.wantHeaders {
				assert.Equal(t, wantValue, rr.Header().Get(key))
			}
		})
	}
}
