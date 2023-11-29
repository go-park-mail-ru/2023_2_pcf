package router

import (
	mock_entities2 "AdHub/auth/pkg/entities/mock_entities"
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"AdHub/proto/api"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserCreateHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock_entities.NewMockUserUseCaseInterface(ctrl)
	mockBalanceUseCase := mock_entities.NewMockBalanceUseCaseInterface(ctrl)
	mockSession := mock_entities2.NewMockSessionUseCaseInterface(ctrl)

	userRouter := UserRouter{
		User:    mockUserUseCase,
		Session: mockSession,
		Balance: mockBalanceUseCase,
	}

	fakeUser := &entities.User{
		Id:        1,
		Login:     "testuser@mail.ru",
		Password:  "test312",
		FName:     "test",
		LName:     "test	",
		BalanceId: 1,
	}

	//fakeSession := &entities.Session{
	//	Token:  "test",
	//	UserId: 1,
	//}

	fakeBalance := &entities.Balance{
		Id:                1,
		Available_balance: 0,
		Reserved_balance:  0,
	}

	userJSON, _ := json.Marshal(fakeUser)

	req, err := http.NewRequest("POST", "/user", strings.NewReader(string(userJSON)))
	if err != nil {
		t.Fatal(err)
	}

	mockBalanceUseCase.EXPECT().
		BalanceCreate(gomock.Any()).
		Return(fakeBalance, nil)

	mockUserUseCase.EXPECT().
		UserCreate(gomock.Any()).
		Return(fakeUser, nil)

	mockSession.EXPECT().
		Auth(context.Background(), gomock.Any()).
		Return(&api.AuthResponse{}, nil)

	mockUserUseCase.EXPECT().
		UserUpdate(gomock.Any()).
		Return(nil)

	rr := httptest.NewRecorder()

	userRouter.UserCreateHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	responseJSON := rr.Body.String()
	assert.JSONEq(t, string(userJSON), responseJSON)
}
