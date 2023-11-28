package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBalanceReplenishHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock_entities.NewMockUserUseCaseInterface(ctrl)
	mockBalanceUseCase := mock_entities.NewMockBalanceUseCaseInterface(ctrl)

	balanceRouter := BalanceRouter{
		User:    mockUserUseCase,
		Balance: mockBalanceUseCase,
	}

	// Подготовка тестовых данных
	replenishAmount := 100.0
	userId := 1
	balanceId := 1
	currentBalance := &entities.Balance{
		Id:                balanceId,
		Available_balance: 500.0, // Пример текущего баланса
		// Дополнительные поля по необходимости
	}

	// Настройка ожидаемого поведения моков
	mockUserUseCase.EXPECT().UserReadById(userId).Return(&entities.User{BalanceId: balanceId}, nil)
	mockBalanceUseCase.EXPECT().BalanceRead(balanceId).Return(currentBalance, nil)
	mockBalanceUseCase.EXPECT().BalanceUP(gomock.Any(), balanceId).Return(nil)

	// Создание тестового запроса
	requestBody, _ := json.Marshal(map[string]string{"amount": fmt.Sprintf("%.2f", replenishAmount)})
	req := httptest.NewRequest("POST", "/balance/replenish", bytes.NewReader(requestBody))
	req = req.WithContext(context.WithValue(req.Context(), "userId", userId))
	rec := httptest.NewRecorder()

	balanceRouter.BalanceReplenishHandler(rec, req)

	// Проверка кода ответа
	if rec.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, rec.Code)
	}

	// Проверка тела ответа
	// В данном случае предполагается, что в ответе возвращается текущий баланс
	expectedBody, _ := json.Marshal(currentBalance)
	if strings.TrimSpace(rec.Body.String()) != strings.TrimSpace(string(expectedBody)) {
		t.Errorf("Response body does not match the expected value.\nExpected: %s\nActual: %s", string(expectedBody), rec.Body.String())
	}
}
