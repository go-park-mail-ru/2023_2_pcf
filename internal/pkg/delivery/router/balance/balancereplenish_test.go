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

func TestBalanceReplenishHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBalanceUseCase := mock_entities.NewMockBalanceUseCaseInterface(ctrl)

	balanceRouter := BalanceRouter{
		Balance: mockBalanceUseCase,
	}

	currentBalance := &entities.Balance{
		Id:            1,
		Total_balance: 100,
	}

	// Настройка ожиданий на mock-объектах
	mockBalanceUseCase.EXPECT().
		BalanceRead(gomock.Any()). // Используйте конкретный ID, если это необходимо
		Return(currentBalance, nil)

	mockBalanceUseCase.EXPECT().
		BalanceUP(gomock.Eq(50.0), gomock.Eq(1)).
		Return(nil)

	replenishAmount := struct {
		Amount float64 `json:"amount"`
	}{
		Amount: 50.0,
	}

	body, _ := json.Marshal(replenishAmount)

	req, _ := http.NewRequest("POST", "/replenish", bytes.NewReader(body))
	req = req.WithContext(context.WithValue(req.Context(), "userid", 1)) // Устанавливаем userID в контекст запроса

	rr := httptest.NewRecorder()

	balanceRouter.BalanceReplenishHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
