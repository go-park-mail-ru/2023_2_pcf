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

func TestCreateTargetHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTargetUseCase := mock_entities.NewMockTargetUseCaseInterface(ctrl)
	mockSession := mock_entities.NewMockSessionUseCaseInterface(ctrl)

	targetRouter := TargetRouter{
		Target:  mockTargetUseCase,
		Session: mockSession,
	}

	fakeTarget := &entities.Target{
		Name:   "TargetName",
		Gender: "male",
	}

	userId := 1 // Предполагаемый ID пользователя

	// Подготовка мока для создания таргета
	mockTargetUseCase.EXPECT().
		TargetCreate(gomock.Any()).
		DoAndReturn(func(target *entities.Target) (*entities.Target, error) {
			target.Id = 123 // Предполагаем, что созданный таргет получает ID 123
			return target, nil
		})

	targetJSON, _ := json.Marshal(fakeTarget)

	req, _ := http.NewRequest("POST", "/create-target", bytes.NewReader(targetJSON))
	req = req.WithContext(context.WithValue(req.Context(), "userid", userId))

	rr := httptest.NewRecorder()

	targetRouter.CreateTargetHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}
