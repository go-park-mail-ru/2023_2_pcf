package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAdCreateHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdUseCase := mock_entities.NewMockAdUseCaseInterface(ctrl)
	mockFileUseCase := mock_entities.NewMockFileUseCaseInterface(ctrl)

	adRouter := AdRouter{
		Ad:   mockAdUseCase,
		File: mockFileUseCase,
	}

	fakeAd := &entities.Ad{
		Name:         "Test Ad",
		Description:  "Test Description",
		Website_link: "http://example.com",
		Budget:       100.0,
		Image_link:   "image.jpg",
		Owner_id:     1,
		Target_id:    1,
	}

	mockAdUseCase.EXPECT().AdCreate(gomock.Any()).Return(fakeAd, nil)
	mockFileUseCase.EXPECT().Save(gomock.Any(), gomock.Any()).Return("image.jpg", nil)

	adJSON, _ := json.Marshal(fakeAd)

	req, _ := http.NewRequest("POST", "/ad", strings.NewReader(string(adJSON)))
	rr := httptest.NewRecorder()

	// Добавление контекста с пользовательским ID
	ctx := context.WithValue(req.Context(), "userid", 1)
	req = req.WithContext(ctx)

	adRouter.AdCreateHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	responseAd := &entities.Ad{}
	if err := json.NewDecoder(rr.Body).Decode(responseAd); err != nil {
		t.Fatal(err)
	}

}
