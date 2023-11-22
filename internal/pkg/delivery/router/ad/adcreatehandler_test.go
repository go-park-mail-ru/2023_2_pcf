package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAdCreateHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdUseCase := mock_entities.NewMockAdUseCaseInterface(ctrl)

	adRouter := AdRouter{
		Ad: mockAdUseCase,
		// Возможно, вам также потребуется мок логгера и других зависимостей
	}

	// Подготовка тела запроса в виде формы
	formData := url.Values{}
	formData.Set("name", "Test Ad")
	formData.Set("description", "This is a test ad")
	formData.Set("website_link", "https://example.com")
	formData.Set("budget", "100")
	formData.Set("target_id", "1")

	expectedAd := &entities.Ad{
		Name:         "Test Ad",
		Description:  "This is a test ad",
		Website_link: "https://example.com",
		Budget:       100,
		Image_link:   "fake_image_link",
		Owner_id:     1,
		Target_id:    1,
	}

	mockAdUseCase.EXPECT().AdCreate(gomock.Any()).Return(expectedAd, nil)

	// Создание тестового запроса с данными формы
	userId := 1
	req := httptest.NewRequest("POST", "/ads/create", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(context.WithValue(req.Context(), "userId", userId))
	rec := httptest.NewRecorder()

	adRouter.AdCreateHandler(rec, req)

	// Проверка кода ответа
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusCreated, rec.Code)
	}

	// Проверка тела ответа
	expectedResponse := `{"id":0,"name":"Test Ad","description":"This is a test ad","website_link":"https://example.com","budget":100,"target_id":1,"image_link":"fake_image_link","Owner_id":1}`
	if rec.Body.String() != expectedResponse {
		t.Errorf("Response body does not match the expected value.\nExpected: %s\nActual: %s", expectedResponse, rec.Body.String())
	}
}
