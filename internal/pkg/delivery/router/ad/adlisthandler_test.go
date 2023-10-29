package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"AdHub/pkg/auth/mock_session"
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
	mockSessionStorage := mock_session.NewMockSessionStorageInterface(ctrl)

	adRouter := AdRouter{
		Ad: mockAdUseCase,
	}

	fakeToken := "fake-token"
	fakeUserID := 1

	req := httptest.NewRequest("GET", "/ad?token="+fakeToken, nil)
	rr := httptest.NewRecorder()

	mockSessionStorage.EXPECT().GetUserId(fakeToken).Return(fakeUserID, nil)

	fakeAds := []*entities.Ad{
		{
			Id:          1,
			Name:        "Ad 1",
			Description: "Description for Ad 1",
			Sector:      "Sector 1",
			Owner_id:    1,
		},
		{
			Id:          2,
			Name:        "Ad 2",
			Description: "Description for Ad 2",
			Sector:      "Sector 2",
			Owner_id:    1,
		},
	}

	mockAdUseCase.EXPECT().AdReadList(fakeUserID).Return(fakeAds, nil)

	adRouter.AdListHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var responseAds []*entities.Ad
	if err := json.NewDecoder(rr.Body).Decode(&responseAds); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, fakeAds, responseAds)
}
