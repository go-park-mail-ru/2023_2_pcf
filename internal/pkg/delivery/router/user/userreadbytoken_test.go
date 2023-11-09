package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByTokenHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock_entities.NewMockUserUseCaseInterface(ctrl)
	mockSession := mock_entities.NewMockSessionUseCaseInterface(ctrl)

	userRouter := UserRouter{
		User:    mockUserUseCase,
		Session: mockSession,
	}

	// Задаём тестовые данные
	token := "test_token"
	userId := 1
	fakeUser := &entities.User{
		Id: userId,
		// Заполните другие поля, если необходимо
	}

	// Настройка ожидаемых вызовов
	mockUserUseCase.EXPECT().
		UserReadById(gomock.Eq(userId)).
		Return(fakeUser, nil)

	requestData := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	requestJSON, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/usergetbytoken", bytes.NewReader(requestJSON))
	rr := httptest.NewRecorder()

	// Установка контекста с user ID, как если бы middleware аутентификации уже было выполнено
	ctx := context.WithValue(req.Context(), "userId", userId)
	req = req.WithContext(ctx)

	userRouter.GetUserByTokenHandler(rr, req)

	// Проверяем код ответа
	assert.Equal(t, http.StatusOK, rr.Code)
}
