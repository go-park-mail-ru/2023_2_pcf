package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "http://example.com", nil)
	rr := httptest.NewRecorder()

	corsMiddleware := CORS(handler)

	corsMiddleware.ServeHTTP(rr, req)

	assert.Equal(t, "http://127.0.0.1:8081", rr.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "true", rr.Header().Get("Access-Control-Allow-Credentials"))
	assert.Equal(t, http.StatusOK, rr.Code)
}
