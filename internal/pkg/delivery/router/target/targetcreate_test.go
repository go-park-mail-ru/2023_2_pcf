package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreateTargetHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock dependencies
	mockTargetUseCase := mock_entities.NewMockTargetUseCaseInterface(ctrl)

	// Create handler
	tr := &TargetRouter{
		Target: mockTargetUseCase,
	}

	// Prepare request body
	requestBody, _ := json.Marshal(map[string]string{
		"name":      "SampleName",
		"gender":    "SampleGender",
		"min_age":   "18",
		"max_age":   "35",
		"interests": "Interest1, Interest2",
		"tags":      "Tag1, Tag2",
		"keys":      "Key1, Key2",
		"regions":   "Region1, Region2",
	})

	// Create test request
	req := httptest.NewRequest(http.MethodPost, "/target", bytes.NewBuffer(requestBody))
	req = req.WithContext(context.WithValue(req.Context(), "userId", 1)) // Set mock user ID in context
	rr := httptest.NewRecorder()

	// Set expectations on mocks
	newTarget := entities.Target{
		Name:      "SampleName",
		Owner_id:  1,
		Gender:    "SampleGender",
		Min_age:   18,
		Max_age:   35,
		Interests: []string{"Interest1", "Interest2"},
		Tags:      []string{"Tag1", "Tag2"},
		Keys:      []string{"Key1", "Key2"},
		Regions:   []string{"Region1", "Region2"},
	}
	mockTargetUseCase.EXPECT().TargetCreate(&newTarget).Return(&newTarget, nil)

	// Call the handler
	tr.CreateTargetHandler(rr, req)

	// Assert response
	assert.Equal(t, http.StatusCreated, rr.Code)
	var response entities.Target
	json.NewDecoder(rr.Body).Decode(&response)
	assert.Equal(t, newTarget, response)

}
