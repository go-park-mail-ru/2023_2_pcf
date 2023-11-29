package router

import (
	mock_entities2 "AdHub/auth/pkg/entities/mock_entities"
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetTargetHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTargetUseCase := mock_entities.NewMockTargetUseCaseInterface(ctrl)
	mockSession := mock_entities2.NewMockSessionUseCaseInterface(ctrl)

	targetRouter := TargetRouter{
		Target:  mockTargetUseCase,
		Session: mockSession,
	}

	fakeTarget := &entities.Target{
		Id:       1,
		Name:     "Test Target",
		Owner_id: 1,
	}

	// Настройка ожидаемых вызовов
	targetId := 1
	mockTargetUseCase.EXPECT().TargetRead(gomock.Eq(targetId)).Return(fakeTarget, nil)

	// Симуляция контекста с пользовательским ID
	userId := 1
	ctx := context.WithValue(context.Background(), "userid", userId)

	// Создание запроса
	req, _ := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("/target/%d", targetId), nil)
	rr := httptest.NewRecorder()

	// Создание роутера с поддержкой переменных пути
	r := mux.NewRouter()
	r.HandleFunc("/target/{id}", targetRouter.GetTargetHandler)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
