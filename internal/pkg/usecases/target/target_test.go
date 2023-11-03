package target

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTargetUseCase_TargetCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockTargetRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeTarget := &entities.Target{
		Id:        1,
		Name:      "Test Target",
		Owner_id:  1,
		Gender:    "Male",
		Min_age:   18,
		Max_age:   50,
		Interests: []string{"interest1", "interest2"},
		Tags:      []string{"tag1", "tag2"},
		Keys:      []string{"key1", "key2"},
		Regions:   []string{"region1", "region2"},
	}

	mockRepo.EXPECT().Create(gomock.Eq(fakeTarget)).Return(fakeTarget, nil)

	createdTarget, err := useCase.TargetCreate(fakeTarget)
	assert.NoError(t, err)
	assert.Equal(t, fakeTarget, createdTarget)
}

func TestTargetUseCase_TargetRemove(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockTargetRepoInterface(ctrl)

	useCase := New(mockRepo)

	targetID := 1

	mockRepo.EXPECT().Remove(targetID).Return(nil)

	err := useCase.TargetRemove(targetID)
	assert.NoError(t, err)
}

func TestTargetUseCase_TargetRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockTargetRepoInterface(ctrl)

	useCase := New(mockRepo)

	targetID := 1

	fakeTarget := &entities.Target{
		Id:        targetID,
		Name:      "Test Target",
		Owner_id:  1,
		Gender:    "Male",
		Min_age:   18,
		Max_age:   50,
		Interests: []string{"interest1", "interest2"},
		Tags:      []string{"tag1", "tag2"},
		Keys:      []string{"key1", "key2"},
		Regions:   []string{"region1", "region2"},
	}

	mockRepo.EXPECT().Read(targetID).Return(fakeTarget, nil)

	target, err := useCase.TargetRead(targetID)
	assert.NoError(t, err)
	assert.Equal(t, fakeTarget, target)
}

func TestTargetUseCase_TargetReadList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockTargetRepoInterface(ctrl)

	useCase := New(mockRepo)

	ownerID := 1

	fakeTargets := []*entities.Target{
		{
			Id:        1,
			Name:      "Test Target 1",
			Owner_id:  ownerID,
			Gender:    "Male",
			Min_age:   18,
			Max_age:   50,
			Interests: []string{"interest1", "interest2"},
			Tags:      []string{"tag1", "tag2"},
			Keys:      []string{"key1", "key2"},
			Regions:   []string{"region1", "region2"},
		},
		{
			Id:        2,
			Name:      "Test Target 2",
			Owner_id:  ownerID,
			Gender:    "Female",
			Min_age:   20,
			Max_age:   60,
			Interests: []string{"interest3", "interest4"},
			Tags:      []string{"tag3", "tag4"},
			Keys:      []string{"key3", "key4"},
			Regions:   []string{"region3", "region4"},
		},
	}

	mockRepo.EXPECT().ReadList(ownerID).Return(fakeTargets, nil)

	targets, err := useCase.TargetReadList(ownerID)
	assert.NoError(t, err)
	assert.Len(t, targets, len(fakeTargets))
}
