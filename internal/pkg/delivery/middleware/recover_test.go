package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("Test panic")
	})

	req := httptest.NewRequest("GET", "http://example.com", nil)
	rr := httptest.NewRecorder()

	recoverMiddleware := Recover(handler)

	recoverMiddleware.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
