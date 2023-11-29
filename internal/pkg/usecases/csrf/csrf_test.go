package csrf

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCsrfUseCase_CsrfCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockCsrfRepoInterface(ctrl)

	useCase := New(mockRepo)

	userId := 1
	fakeCsrf := &entities.Csrf{
		Token:  "fakeCsrfToken",
		UserId: userId,
	}

	mockRepo.EXPECT().Create(gomock.Any()).Return(fakeCsrf, nil)

	createdCsrf, err := useCase.CsrfCreate(userId)
	assert.NoError(t, err)
	assert.Equal(t, fakeCsrf, createdCsrf)
}

func TestCsrfUseCase_CsrfRemove(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockCsrfRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeCsrf := &entities.Csrf{
		Token:  "fakeCsrfToken",
		UserId: 1,
	}

	mockRepo.EXPECT().Remove(gomock.Eq(fakeCsrf)).Return(nil)

	err := useCase.CsrfRemove(fakeCsrf)
	assert.NoError(t, err)
}

func TestCsrfUseCase_GetByUserId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockCsrfRepoInterface(ctrl)

	useCase := New(mockRepo)

	userId := 1
	fakeCsrf := &entities.Csrf{
		Token:  "fakeCsrfToken",
		UserId: userId,
	}

	mockRepo.EXPECT().Read(gomock.Eq(userId)).Return(fakeCsrf, nil)

	retrievedCsrf, err := useCase.GetByUserId(userId)
	assert.NoError(t, err)
	assert.Equal(t, fakeCsrf, retrievedCsrf)
}
