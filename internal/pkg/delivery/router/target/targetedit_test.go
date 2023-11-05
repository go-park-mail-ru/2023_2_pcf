package router

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateTargetHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTargetUseCase := mock_entities.NewMockTargetUseCaseInterface(ctrl)

	targetRouter := TargetRouter{
		Target: mockTargetUseCase,
	}

	// Задаём тестовые данные и ожидания
	fakeUserID := 1
	fakeTargetID := 123
	fakeTarget := &entities.Target{
		Id:       fakeTargetID,
		Owner_id: fakeUserID,
		// Другие поля...
	}

	updatedName := "Updated Name"
	updatedGender := "Female"
	updatedMinAge := 30
	updatedMaxAge := 45

	requestBody := `{
		"id": 123,
		"name": "Updated Name",
		"gender": "Female",
		"min_age": 30,
		"max_age": 45,
		"interests": "Interest1, Interest2",
		"tags": "Tag1, Tag2",
		"keys": "Key1, Key2",
		"regions": "Region1, Region2"
	}`

	mockTargetUseCase.EXPECT().
		TargetRead(gomock.Eq(fakeTargetID)).
		Return(fakeTarget, nil)

	mockTargetUseCase.EXPECT().
		TargetUpdate(gomock.Any()).
		DoAndReturn(func(target *entities.Target) error {
			assert.Equal(t, updatedName, target.Name)
			assert.Equal(t, updatedGender, target.Gender)
			assert.Equal(t, updatedMinAge, target.Min_age)
			assert.Equal(t, updatedMaxAge, target.Max_age)
			return nil
		})

	req, _ := http.NewRequest("POST", "/update-target", strings.NewReader(requestBody))
	req = req.WithContext(context.WithValue(req.Context(), "userid", fakeUserID))

	rr := httptest.NewRecorder()

	targetRouter.UpdateTargetHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
