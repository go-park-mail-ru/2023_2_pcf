package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestAdEditHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdUseCase := mock_entities.NewMockAdUseCaseInterface(ctrl)
	mockSession := mock_entities.NewMockSessionUseCaseInterface(ctrl)
	mockFileUseCase := mock_entities.NewMockFileUseCaseInterface(ctrl)

	adRouter := AdRouter{
		Ad:      mockAdUseCase,
		Session: mockSession,
		File:    mockFileUseCase,
	}

	// Задаём тестовые данные и ожидания
	fakeAdId := 123
	userID := 1
	fakeAd := &entities.Ad{
		Id:       fakeAdId,
		Owner_id: userID,
		// Заполните другие поля, если это необходимо
	}

	updatedName := "Updated Name"
	updateRequest := struct {
		AdId int     `json:"ad_id"`
		Name *string `json:"name"`
		// Остальные поля заполнены, если необходимо
	}{
		AdId: fakeAdId,
		Name: &updatedName,
	}

	mockAdUseCase.EXPECT().
		AdRead(gomock.Eq(fakeAdId)).
		Return(fakeAd, nil)

	mockAdUseCase.EXPECT().
		AdUpdate(gomock.Any()).
		DoAndReturn(func(ad *entities.Ad) error {
			assert.Equal(t, updatedName, ad.Name) // Убедитесь, что имя было обновлено
			return nil
		})

	adUpdateJSON, _ := json.Marshal(updateRequest)

	req, _ := http.NewRequest("POST", "/ad/update", bytes.NewBuffer(adUpdateJSON))
	req = mux.SetURLVars(req, map[string]string{"ad_id": strconv.Itoa(fakeAdId)}) // установка переменных для маршрута, если они используются
	rr := httptest.NewRecorder()

	// Установка контекста с user ID, как если бы middleware аутентификации уже было выполнено
	ctx := context.WithValue(req.Context(), "userId", userID)
	req = req.WithContext(ctx)

	adRouter.AdUpdateHandler(rr, req)

	// Проверяем код ответа
	assert.Equal(t, http.StatusOK, rr.Code)
}
