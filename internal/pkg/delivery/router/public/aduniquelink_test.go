package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAdBannerHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdUseCase := mock_entities.NewMockAdUseCaseInterface(ctrl)

	adRouter := &AdRouter{
		Ad:   mockAdUseCase,
		addr: "http://example.com",
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ad-banner?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	expectedAdID := 1

	fakeAd := &entities.Ad{
		Image_link: "fake_image.png",
	}

	mockAdUseCase.EXPECT().AdRead(expectedAdID).Return(fakeAd, nil)

	adRouter.AdBannerHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
