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

func TestUserDeleteHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock_entities.NewMockUserUseCaseInterface(ctrl)

	userRouter := UserRouter{
		User: mockUserUseCase,
	}

	// Задаём тестовые данные
	userId := 1
	login := "testuser"
	userFromDb := &entities.User{
		Id:    userId,
		Login: login,
		// Другие поля, если необходимо
	}

	// Настройка ожидаемых вызовов
	mockUserUseCase.EXPECT().
		UserReadByLogin(gomock.Eq(login)).
		Return(userFromDb, nil)

	mockUserUseCase.EXPECT().
		UserDelete(gomock.Eq(login)).
		Return(nil)

	req, _ := http.NewRequest("DELETE", "/user?login="+login, nil)

	rr := httptest.NewRecorder()

	// Установка контекста с user ID, как если бы middleware аутентификации уже было выполнено
	ctx := context.WithValue(req.Context(), "userid", userId)
	req = req.WithContext(ctx)

	userRouter.UserDeleteHandler(rr, req)

	// Проверяем код ответа
	assert.Equal(t, http.StatusNoContent, rr.Code)
}
