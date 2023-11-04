package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdGetHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdUseCase := mock_entities.NewMockAdUseCaseInterface(ctrl)

	adRouter := AdRouter{
		router:  nil, // В данном случае не нужен для теста
		Ad:      mockAdUseCase,
		Session: nil, // В данном случае не нужен для теста
	}

	expectedAd := &entities.Ad{
		Id:           1,
		Name:         "Test Ad",
		Description:  "Test Description",
		Website_link: "https://example.com",
		Budget:       100.0,
		Target_id:    1,
		Image_link:   "https://example.com/image.jpg",
		Owner_id:     1,
	}

	mockAdUseCase.EXPECT().
		AdRead(gomock.Eq(expectedAd.Id)).
		Return(expectedAd, nil)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/ad/%d", expectedAd.Id), nil)
	rr := httptest.NewRecorder()

	// Создание роутера с поддельным обработчиком для тестирования
	r := mux.NewRouter()
	r.HandleFunc("/ad/{adID}", adRouter.AdGetHandler)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Проверяем, что тело ответа содержит информацию об объявлении
	adJSON, _ := json.Marshal(expectedAd)
	assert.JSONEq(t, string(adJSON), rr.Body.String())
}
