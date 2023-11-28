package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetBalanceHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock_entities.NewMockUserUseCaseInterface(ctrl)
	mockBalanceUseCase := mock_entities.NewMockBalanceUseCaseInterface(ctrl)
	mockSession := mock_entities.NewMockSessionUseCaseInterface(ctrl)

	balanceRouter := BalanceRouter{
		User:    mockUserUseCase,
		Balance: mockBalanceUseCase,
		Session: mockSession,
	}

	fakeBalance := &entities.Balance{
		Id:                1,
		Total_balance:     1000.0,
		Reserved_balance:  100.0,
		Available_balance: 900.0,
	}

	expectedUserID := 1
	mockBalanceUseCase.EXPECT().BalanceRead(expectedUserID).Return(fakeBalance, nil)

	req, _ := http.NewRequest("GET", "/balance", nil)
	ctx := context.WithValue(req.Context(), "userId", expectedUserID)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	balanceRouter.GetBalanceHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
