package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodOptions, "http://example.com", nil)

	rr := httptest.NewRecorder()

	corsHandler := CORS(nextHandler)

	corsHandler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "http://127.0.0.1:8081", rr.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "GET, POST, OPTIONS, DELETE", rr.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Content-Type", rr.Header().Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "true", rr.Header().Get("Access-Control-Allow-Credentials"))

	req = httptest.NewRequest(http.MethodGet, "http://example.com", nil)

	rr = httptest.NewRecorder()

	corsHandler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "http://127.0.0.1:8081", rr.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "true", rr.Header().Get("Access-Control-Allow-Credentials"))
}
