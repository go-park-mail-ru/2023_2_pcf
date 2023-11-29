package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"context"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestPadCreateHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPadUseCase := mock_entities.NewMockPadUseCaseInterface(ctrl)

	padRouter := PadRouter{
		Pad: mockPadUseCase,
	}

	// Подготовка данных для тела запроса
	formData := url.Values{}
	formData.Set("name", "Test Pad")
	formData.Set("description", "This is a test pad")
	formData.Set("website_link", "https://examplepad.com")
	formData.Set("price", "200")
	formData.Set("target_id", "2")

	expectedPad := &entities.Pad{
		Name:         "Test Pad",
		Description:  "This is a test pad",
		Website_link: "https://examplepad.com",
		Price:        200,
		Owner_id:     1,
		Target_id:    2,
	}

	mockPadUseCase.EXPECT().PadCreate(gomock.Any()).Return(expectedPad, nil)

	// Создание тестового запроса
	userId := 1
	req := httptest.NewRequest("POST", "/pads/create", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(context.WithValue(req.Context(), "userId", userId))
	rec := httptest.NewRecorder()

	padRouter.PadCreateHandler(rec, req)

	// Проверка кода ответа
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusCreated, rec.Code)
	}

	// Проверка тела ответа
	expectedResponse := `{"id":0,"name":"Test Pad","description":"This is a test pad","website_link":"https://examplepad.com","price":200,"target_id":2,"Owner_id":1,"clicks":0,"balance":0}`
	if rec.Body.String() != expectedResponse {
		t.Errorf("Response body does not match the expected value.\nExpected: %s\nActual: %s", expectedResponse, rec.Body.String())
	}
}
