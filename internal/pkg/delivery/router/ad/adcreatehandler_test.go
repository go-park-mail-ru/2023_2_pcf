package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAdCreateHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdUseCase := mock_entities.NewMockAdUseCaseInterface(ctrl)
	//mockFileUseCase := mock_entities.NewMockFileUseCaseInterface(ctrl)

	adRouter := AdRouter{
		Ad: mockAdUseCase,
	}

	// Prepare the request payload
	payload := struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		WebsiteLink string  `json:"website_link"`
		Budget      float64 `json:"budget"`
		TargetId    int     `json:"target_id"`
	}{
		Name:        "Test Ad",
		Description: "This is a test ad",
		WebsiteLink: "https://example.com",
		Budget:      100,
		TargetId:    1,
	}

	// Mock the FileUseCaseInterface.Save method to return a fake image link
	//mockFileUseCase.EXPECT().Save(gomock.Any(), gomock.Any()).Return("fake_image_link", nil)

	expectedAd := &entities.Ad{
		Name:         payload.Name,
		Description:  payload.Description,
		Website_link: payload.WebsiteLink,
		Budget:       payload.Budget,
		Image_link:   "fake_image_link",
		Owner_id:     1,
		Target_id:    payload.TargetId,
	}

	mockAdUseCase.EXPECT().AdCreate(gomock.Any()).Return(expectedAd, nil)

	// Create a test request
	userId := 1
	reqBody, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/ads/create", bytes.NewReader(reqBody))
	req = req.WithContext(context.WithValue(req.Context(), "userid", userId))
	rec := httptest.NewRecorder()

	adRouter.AdCreateHandler(rec, req)

	// Check the response status code
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusCreated, rec.Code)
	}

	// Check the response body
	expectedResponse := `{"id":0,"name":"Test Ad","description":"This is a test ad","website_link":"https://example.com","budget":100,"target_id":1,"image_link":"fake_image_link","Owner_id":1}`
	if rec.Body.String() != expectedResponse {
		t.Errorf("Response body does not match the expected value.\nExpected: %s\nActual: %s", expectedResponse, rec.Body.String())
	}
}
