package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAdDeleteHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdUseCase := mock_entities.NewMockAdUseCaseInterface(ctrl)
	mockUserUseCase := mock_entities.NewMockUserUseCaseInterface(ctrl)
	mockBalanceUseCase := mock_entities.NewMockBalanceUseCaseInterface(ctrl)

	adRouter := AdRouter{
		Ad:      mockAdUseCase,
		User:    mockUserUseCase,
		Balance: mockBalanceUseCase,
	}

	testAd := entities.Ad{
		Id:       1,
		Owner_id: 123,
	}

	expectedUser := &entities.User{
		Id:          1,
		Login:       "testuser",
		Password:    "test",
		FName:       "test",
		LName:       "test",
		CompanyName: "Yandex",
		Avatar:      "test.jpg",
	}

	// Setting up the expected calls and returns for the mock object
	mockAdUseCase.EXPECT().
		AdRead(gomock.Eq(testAd.Id)).
		Return(&testAd, nil)

	mockAdUseCase.EXPECT().
		AdRemove(gomock.Eq(testAd.Id)).
		Return(nil)

	mockUserUseCase.EXPECT().UserReadById(123).Return(expectedUser, nil)

	mockBalanceUseCase.EXPECT().BalanceUP(0.0, 0).Return(nil)

	requestBody, _ := json.Marshal(map[string]int{
		"ad_id": testAd.Id,
	})

	req, _ := http.NewRequest("DELETE", "/ad", strings.NewReader(string(requestBody)))

	// Вставляем userId в контекст запроса
	ctx := context.WithValue(req.Context(), "userId", testAd.Owner_id)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	adRouter.AdDeleteHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
