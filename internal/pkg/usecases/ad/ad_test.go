package ad

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAdUseCase_AdCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockAdRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeAd := &entities.Ad{
		Name:        "Test Ad",
		Description: "This is a test ad",
		Sector:      "Technology",
		Owner_id:    1,
	}

	mockRepo.EXPECT().Create(gomock.Eq(fakeAd)).Return(fakeAd, nil)

	createdAd, err := useCase.AdCreate(fakeAd)
	assert.NoError(t, err)
	assert.Equal(t, fakeAd, createdAd)
}

func TestAdUseCase_AdReadList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockAdRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeAds := []*entities.Ad{
		{
			Id:          1,
			Name:        "Ad 1",
			Description: "Description 1",
			Sector:      "Sector 1",
			Owner_id:    1,
		},
		{
			Id:          2,
			Name:        "Ad 2",
			Description: "Description 2",
			Sector:      "Sector 2",
			Owner_id:    2,
		},
	}

	mockRepo.EXPECT().Read(1).Return(fakeAds, nil)

	ads, err := useCase.AdReadList(1)
	assert.NoError(t, err)
	assert.Equal(t, fakeAds, ads)
}
