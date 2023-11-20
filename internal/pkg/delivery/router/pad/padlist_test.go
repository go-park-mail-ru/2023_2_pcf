package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPadListHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPadUseCase := mock_entities.NewMockPadUseCaseInterface(ctrl)
	padRouter := PadRouter{
		Pad: mockPadUseCase,
	}

	userID := 1
	expectedPads := []*entities.Pad{
		{
			Id:       1,
			Name:     "Pad 1",
			Owner_id: userID,
		},
		{
			Id:       2,
			Name:     "Pad 2",
			Owner_id: userID,
		},
	}

	mockPadUseCase.EXPECT().PadReadList(userID).Return(expectedPads, nil)

	req := httptest.NewRequest("GET", "/pads/list", nil)
	req = req.WithContext(context.WithValue(req.Context(), "userId", userID))
	rec := httptest.NewRecorder()

	padRouter.PadListHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, rec.Code)
	}

	expectedBody, _ := json.Marshal(expectedPads)
	expectedBodyStr := string(expectedBody)
	expectedBodyStr = expectedBodyStr + "\n"
	if rec.Body.String() != expectedBodyStr {
		t.Errorf("Response body does not match the expected value.\nExpected: %s\nActual: %s", expectedBody, rec.Body.String())
	}
}
