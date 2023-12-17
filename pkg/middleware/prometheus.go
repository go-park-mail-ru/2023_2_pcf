package middleware

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests.",
		},
		[]string{"endpoint", "method", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"endpoint", "method"},
	)

	httpRequestErrorTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_errors_total",
			Help: "Total number of HTTP request errors.",
		},
		[]string{"endpoint", "method", "error"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal, httpRequestDuration, httpRequestErrorTotal)
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		method := r.Method
		recorder := &StatusRecorder{ResponseWriter: w, Status: 200}

		startTime := time.Now()
		next.ServeHTTP(recorder, r)
		duration := time.Since(startTime)

		statusCode := recorder.Status

		// Increment the total requests counter
		httpRequestsTotal.WithLabelValues(path, method, http.StatusText(statusCode)).Inc()

		// Record the duration of the request
		httpRequestDuration.WithLabelValues(path, method).Observe(duration.Seconds())

		// If there was an error, increment the error counter
		if statusCode >= 400 {
			httpRequestErrorTotal.WithLabelValues(path, method, http.StatusText(statusCode)).Inc()
		}
	})
}

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (rec *StatusRecorder) WriteHeader(code int) {
	rec.Status = code
	rec.ResponseWriter.WriteHeader(code)
}
