package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPadDeleteHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPadUseCase := mock_entities.NewMockPadUseCaseInterface(ctrl)
	padRouter := PadRouter{
		Pad: mockPadUseCase,
	}

	padID := 1
	userID := 1
	padToDelete := &entities.Pad{
		Id:       padID,
		Owner_id: userID,
	}

	mockPadUseCase.EXPECT().PadRead(padID).Return(padToDelete, nil)
	mockPadUseCase.EXPECT().PadRemove(padID).Return(nil)

	reqBody, _ := json.Marshal(map[string]int{"pad_id": padID})
	req := httptest.NewRequest("POST", "/pad/delete", bytes.NewReader(reqBody))
	req = req.WithContext(context.WithValue(req.Context(), "userId", userID))
	rec := httptest.NewRecorder()

	padRouter.PadDeleteHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, rec.Code)
	}
	expectedResponse := `{"message": "Pad deleted successfully"}`
	if rec.Body.String() != expectedResponse {
		t.Errorf("Response body does not match the expected value.\nExpected: %s\nActual: %s", expectedResponse, rec.Body.String())
	}
}
