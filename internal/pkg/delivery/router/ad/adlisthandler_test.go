package router

import (
	mock_entities2 "AdHub/auth/pkg/entities/mock_entities"
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"context"
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
	mockSession := mock_entities2.NewMockSessionUseCaseInterface(ctrl)

	adRouter := AdRouter{
		Ad:      mockAdUseCase,
		Session: mockSession,
	}

	userId := 1 // Предполагаемый ID пользователя

	// Подготовка мока для функции AdReadList
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

	mockAdUseCase.EXPECT().AdReadList(userId).Return(fakeAds, nil)

	req, _ := http.NewRequest("GET", "/ad/list", nil)
	req = req.WithContext(context.WithValue(req.Context(), "userId", userId))

	rr := httptest.NewRecorder()

	adRouter.AdListHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Дополнительные проверки на данные возвращенного JSON-ответа, если необходимо
}
