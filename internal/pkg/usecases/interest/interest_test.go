package interest

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInterestUseCase_InterestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockInterestRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeInterest := &entities.Interest{
		Name: "Test Interest",
	}

	mockRepo.EXPECT().Create(gomock.Eq(fakeInterest)).Return(fakeInterest, nil)

	createdInterest, err := useCase.InterestCreate(fakeInterest)
	assert.NoError(t, err)
	assert.Equal(t, fakeInterest, createdInterest)
}

func TestInterestUseCase_InterestRemove(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockInterestRepoInterface(ctrl)

	useCase := New(mockRepo)

	mockRepo.EXPECT().Remove(1).Return(nil)

	err := useCase.InterestRemove(1)
	assert.NoError(t, err)
}
