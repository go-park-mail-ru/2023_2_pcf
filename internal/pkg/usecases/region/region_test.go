package region

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegionUseCase_RegionCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockRegionRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeRegion := &entities.Region{
		Name: "Test Region",
	}

	mockRepo.EXPECT().Create(gomock.Eq(fakeRegion)).Return(fakeRegion, nil)

	createdRegion, err := useCase.RegionCreate(fakeRegion)
	assert.NoError(t, err)
	assert.Equal(t, fakeRegion, createdRegion)
}

func TestRegionUseCase_RegionRemove(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockRegionRepoInterface(ctrl)

	useCase := New(mockRepo)

	mockRepo.EXPECT().Remove(1).Return(nil)

	err := useCase.RegionRemove(1)
	assert.NoError(t, err)
}
