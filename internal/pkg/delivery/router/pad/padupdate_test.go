package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"context"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestPadUpdateHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPadUseCase := mock_entities.NewMockPadUseCaseInterface(ctrl)

	padRouter := PadRouter{
		Pad: mockPadUseCase,
	}

	padID := 1
	userID := 1
	updatedName := "Updated Pad Name"
	updatedPrice := "300.50"
	updatedTargetId := "3"

	padToUpdate := &entities.Pad{
		Id:        padID,
		Owner_id:  userID,
		Name:      "Original Pad Name",
		Price:     200,
		Target_id: 2,
	}

	updatedPad := &entities.Pad{
		Id:        padID,
		Owner_id:  userID,
		Name:      updatedName,
		Price:     300.50,
		Target_id: 3,
	}

	mockPadUseCase.EXPECT().PadRead(padID).Return(padToUpdate, nil)
	mockPadUseCase.EXPECT().PadUpdate(updatedPad).Return(nil)

	formData := url.Values{}
	formData.Set("id", strconv.Itoa(padID))
	formData.Set("name", updatedName)
	formData.Set("price", updatedPrice)
	formData.Set("target_id", updatedTargetId)

	req := httptest.NewRequest("POST", "/pad/update", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(context.WithValue(req.Context(), "userId", userID))
	rec := httptest.NewRecorder()

	padRouter.PadUpdateHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, rec.Code)
	}

}
