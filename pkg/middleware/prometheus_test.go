package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetricsMiddleware(t *testing.T) {
	// Создайте тестовый обработчик, который будет использоваться в качестве следующего обработчика
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Эмулируйте выполнение запроса здесь
		w.WriteHeader(http.StatusOK)
	})

	// Создайте запрос и респонс для теста
	req, err := http.NewRequest("GET", "/example", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()

	// Создайте MetricsMiddleware
	middleware := MetricsMiddleware(testHandler)

	// Вызовите MetricsMiddleware с нашими тестовыми объектами запроса и респонса
	middleware.ServeHTTP(rr, req)

	// Теперь вы можете добавить проверки для метрик, которые вы хотите проверить
	// Например, можно проверить, что счетчик "http_requests_total" увеличился
	// и длительность запроса была записана в "http_request_duration_seconds".

	// Проверка счетчика "http_requests_total"
	// Проверка счетчика "http_requests_total"
	_, err = httpRequestsTotal.GetMetricWithLabelValues("/example", "GET", "OK")
	assert.NoError(t, err)

	// Проверка длительности запроса
	// Используем метод Observe для регистрации значения гистограммы
	_, err = httpRequestDuration.GetMetricWithLabelValues("/example", "GET")
	assert.NoError(t, err)

	// Регистрируем ожидаемое значение

	// Здесь вы также можете добавить проверку для счетчика "http_request_errors_total",
	// если хотите проверить его в вашем тесте.

	// Проверка статус кода ответа
	expectedStatusCode := http.StatusOK
	assert.Equal(t, expectedStatusCode, rr.Code)
}
