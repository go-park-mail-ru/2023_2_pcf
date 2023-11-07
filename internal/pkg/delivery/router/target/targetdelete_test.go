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

func TestTargetDeleteHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTargetUseCase := mock_entities.NewMockTargetUseCaseInterface(ctrl)
	mockSession := mock_entities.NewMockSessionUseCaseInterface(ctrl)

	targetRouter := TargetRouter{
		Target:  mockTargetUseCase,
		Session: mockSession,
	}

	fakeTarget := &entities.Target{
		Id:       1,
		Name:     "Test Target",
		Owner_id: 1,
	}

	mockTargetUseCase.EXPECT().TargetRead(fakeTarget.Id).Return(fakeTarget, nil)
	mockTargetUseCase.EXPECT().TargetRemove(fakeTarget.Id).Return(nil)

	targetJSON, _ := json.Marshal(map[string]int{"id": fakeTarget.Id})

	req, _ := http.NewRequest("DELETE", "/target/delete", bytes.NewBuffer(targetJSON))
	rr := httptest.NewRecorder()

	// Создание контекста с имитацией пользователя (в вашем случае это ID = 1)
	ctx := context.WithValue(req.Context(), "userid", 1)
	req = req.WithContext(ctx)

	targetRouter.TargetDeleteHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
