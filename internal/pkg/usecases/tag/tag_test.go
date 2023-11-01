package tag

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTagUseCase_TagCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockTagRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeTag := &entities.Tag{
		Name: "Test Tag",
	}

	mockRepo.EXPECT().Create(gomock.Eq(fakeTag)).Return(fakeTag, nil)

	createdTag, err := useCase.TagCreate(fakeTag)
	assert.NoError(t, err)
	assert.Equal(t, fakeTag, createdTag)
}

func TestTagUseCase_TagRemove(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockTagRepoInterface(ctrl)

	useCase := New(mockRepo)

	mockRepo.EXPECT().Remove(1).Return(nil)

	err := useCase.TagRemove(1)
	assert.NoError(t, err)
}
