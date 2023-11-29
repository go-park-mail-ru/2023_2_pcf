package router

import (
	mock_entities2 "AdHub/auth/pkg/entities/mock_entities"
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdUpdateHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock dependencies
	mockAdUseCase := mock_entities.NewMockAdUseCaseInterface(ctrl)
	mockSession := mock_entities2.NewMockSessionUseCaseInterface(ctrl)

	// Create handler
	mr := &AdRouter{
		Ad:      mockAdUseCase,
		Session: mockSession,
	}

	// Sample ad data
	currentAd := &entities.Ad{
		Owner_id:     1,
		Name:         "Old Ad",
		Description:  "Old Description",
		Website_link: "http://oldwebsite.com",
		Budget:       50.0,
		Target_id:    2,
		Click_cost:   1.0,
		Image_link:   "oldimage.jpg",
	}

	// Prepare request body with multipart form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("ad_id", "1")
	_ = writer.WriteField("name", "Updated Ad")
	_ = writer.WriteField("description", "Updated Description")
	_ = writer.WriteField("website_link", "http://updatedwebsite.com")
	_ = writer.WriteField("budget", "100.50")
	_ = writer.WriteField("target_id", "3")
	_ = writer.WriteField("click_cost", "1.5")
	// Add other fields as needed
	writer.Close()

	// Create test request
	req := httptest.NewRequest(http.MethodPost, "/ad/update", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	userID := 1 // Sample user ID
	req = req.WithContext(context.WithValue(req.Context(), "userId", userID))
	rr := httptest.NewRecorder()

	// Set expectations on mocks
	mockAdUseCase.EXPECT().AdRead(gomock.Any()).Return(currentAd, nil)
	mockAdUseCase.EXPECT().AdUpdate(gomock.Any()).Return(nil)

	// Call the handler
	mr.AdUpdateHandler(rr, req)

	// Assert response
	assert.Equal(t, http.StatusOK, rr.Code)

}
