package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdDeleteHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdUseCase := mock_entities.NewMockAdUseCaseInterface(ctrl)
	mockSession := mock_entities.NewMockSessionUseCaseInterface(ctrl)

	adRouter := AdRouter{
		Ad:      mockAdUseCase,
		Session: mockSession,
	}

	payload := struct {
		Token       string `json:"token"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Sector      string `json:"sector"`
		TargetId    int    `json:"target_id"`
	}{
		Token:       "fakeToken",
		Name:        "Test Ad",
		Description: "This is a test ad",
		Sector:      "Technology",
	}

	payloadJSON, err := json.Marshal(payload)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/addelete", bytes.NewReader(payloadJSON))
	rr := httptest.NewRecorder()

	mockSession.EXPECT().GetUserId("fakeToken").Return(1, nil)

	fakeAd := &entities.Ad{
		Id:          1,
		Name:        payload.Name,
		Description: payload.Description,
		Target_id:   payload.TargetId,
		Owner_id:    1,
	}
	mockAdUseCase.EXPECT().AdCreate(gomock.Any()).Return(fakeAd, nil)

	adRouter.AdCreateHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response struct {
		Id int `json:"id"`
	}

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, 1, response.Id)
}
