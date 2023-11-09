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

func TestRedirectHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdUseCase := mock_entities.NewMockAdUseCaseInterface(ctrl)

	PublicRouter := &PublicRouter{
		Ad: mockAdUseCase,
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/redirect?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	expectedURL := "example.com/ad/1"
	fakeAd := &entities.Ad{
		Website_link: expectedURL,
	}

	mockAdUseCase.EXPECT().AdRead(1).Return(fakeAd, nil)

	PublicRouter.RedirectHandler(rr, req)

	assert.Equal(t, http.StatusSeeOther, rr.Code)

	expectedLocation := rr.Header().Get("Location")
	assert.Equal(t, "http://"+expectedURL, expectedLocation)
}
