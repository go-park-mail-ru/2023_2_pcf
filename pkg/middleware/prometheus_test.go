package middleware

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMetricsMiddleware(t *testing.T) {
	testServer := httptest.NewServer(promhttp.Handler())
	defer testServer.Close()

	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handlerWithMetrics := MetricsMiddleware(mockHandler)

	req, err := http.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)

	requestsBefore := testutil.CollectAndCount(httpRequestsTotal, "endpoint", "method", "status")
	errorsBefore := testutil.CollectAndCount(httpRequestErrorTotal, "endpoint", "method", "error")

	recorder := httptest.NewRecorder()
	handlerWithMetrics.ServeHTTP(recorder, req)

	requestsAfter := testutil.CollectAndCount(httpRequestsTotal, "endpoint", "method", "status")
	errorsAfter := testutil.CollectAndCount(httpRequestErrorTotal, "endpoint", "method", "error")

	assert.Equal(t, requestsBefore, requestsAfter, "Total requests metric should be incremented")
	assert.Equal(t, errorsBefore, errorsAfter, "Total errors metric should remain unchanged for successful requests")
}
