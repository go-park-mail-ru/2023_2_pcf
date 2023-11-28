package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAdGetHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdUseCase := mock_entities.NewMockAdUseCaseInterface(ctrl)

	adRouter := AdRouter{
		Ad: mockAdUseCase,
	}

	// Подготовка тестового объявления
	adID := 1
	expectedAd := &entities.Ad{
		Id:           adID,
		Name:         "Test Ad",
		Description:  "This is a test ad",
		Website_link: "https://example.com",
		Budget:       100,
		Owner_id:     1,
		// Заполните остальные поля как необходимо
	}

	// Настройка ожидаемого поведения моков
	mockAdUseCase.EXPECT().AdRead(adID).Return(expectedAd, nil)

	// Создание тестового запроса
	req := httptest.NewRequest("GET", fmt.Sprintf("/ad?adID=%d", adID), nil)
	rec := httptest.NewRecorder()

	adRouter.AdGetHandler(rec, req)

	// Проверка кода ответа
	if rec.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, rec.Code)
	}

	// Проверка тела ответа
	expectedBody, _ := json.Marshal(expectedAd)
	if strings.TrimSpace(rec.Body.String()) != strings.TrimSpace(string(expectedBody)) {
		t.Errorf("Response body does not match the expected value.\nExpected: %s\nActual: %s", string(expectedBody), rec.Body.String())
	}
}
