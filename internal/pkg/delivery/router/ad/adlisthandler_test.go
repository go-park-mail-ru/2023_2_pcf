package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAdListHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdUseCase := mock_entities.NewMockAdUseCaseInterface(ctrl)
	mockSession := mock_entities.NewMockSessionUseCaseInterface(ctrl)

	adRouter := AdRouter{
		Ad:      mockAdUseCase,
		Session: mockSession,
	}

	token := "fakeToken"
	userID := 1

	fakeAds := []*entities.Ad{
		{
			Id:          1,
			Name:        "Fake Ad 1",
			Description: "This is a fake ad 1",
			Target_id:   1,
			Owner_id:    1,
		},
		{
			Id:          2,
			Name:        "Fake Ad 2",
			Description: "This is a fake ad 2",
			Target_id:   2,
			Owner_id:    2,
		},
		{
			Id:          3,
			Name:        "Fake Ad 3",
			Description: "This is a fake ad 3",
			Target_id:   1,
			Owner_id:    1,
		},
	}

	mockSession.EXPECT().GetUserId(token).Return(userID, nil)
	mockAdUseCase.EXPECT().AdReadList(userID).Return(fakeAds, nil)

	req := httptest.NewRequest("GET", "/ad-list?token="+token, nil)
	rr := httptest.NewRecorder()

	adRouter.AdListHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var responseAds []*entities.Ad
	if err := json.Unmarshal(rr.Body.Bytes(), &responseAds); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	assert.Equal(t, fakeAds, responseAds)
}
