package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBalanceHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBalanceUseCase := mock_entities.NewMockBalanceUseCaseInterface(ctrl)
	mockSession := mock_entities.NewMockSessionUseCaseInterface(ctrl)

	balanceRouter := BalanceRouter{
		Balance: mockBalanceUseCase,
		Session: mockSession,
		// предполагаем, что logger уже инициализирован
	}

	fakeBalance := &entities.Balance{
		Id:                1,
		Total_balance:     1000.0,
		Reserved_balance:  100.0,
		Available_balance: 900.0,
	}

	expectedUserID := 1 // Идентификатор пользователя для теста
	mockBalanceUseCase.EXPECT().BalanceRead(expectedUserID).Return(fakeBalance, nil)

	req, _ := http.NewRequest("GET", "/balance", nil)
	// Добавление идентификатора пользователя в контекст запроса
	ctx := context.WithValue(req.Context(), "userid", expectedUserID)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	balanceRouter.GetBalanceHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
