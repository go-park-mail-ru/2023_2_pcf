package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdGetAmountHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdUseCase := mock_entities.NewMockAdUseCaseInterface(ctrl)

	adRouter := AdRouter{
		Ad:     mockAdUseCase,
		logger: mockLogger,
	}

	fakeAds := []*entities.Ad{
		{Id: 1, Name: "Ad 1"},
		{Id: 2, Name: "Ad 2"},
		// Добавить дополнительные объявления при необходимости
	}

	userId := 1
	offset := 0
	amount := 2

	mockAdUseCase.EXPECT().AdReadList(userId).Return(fakeAds, nil)

	// Создание JSON запроса
	requestBody, _ := json.Marshal(map[string]int{
		"amount": amount,
		"offset": offset,
	})

	req, _ := http.NewRequest("POST", "/adGetAmount", bytes.NewBuffer(requestBody))

	// Создание контекста с пользовательским ID
	ctx := context.WithValue(req.Context(), "userid", userId)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	adRouter.AdGetAmountHandler(rr, req)

	// Проверка ответа
	assert.Equal(t, http.StatusOK, rr.Code)
	expectedResponseBody, _ := json.Marshal(fakeAds[offset:amount])
	assert.JSONEq(t, string(expectedResponseBody), rr.Body.String())
}
