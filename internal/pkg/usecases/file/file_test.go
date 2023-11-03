package file

import (
	"AdHub/internal/pkg/entities/mock_entities"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFileUseCase_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockFileRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeFileName := "unique_filename"
	fakeFileData := []byte("file_data")

	mockRepo.EXPECT().Save(gomock.Eq(fakeFileData), gomock.Eq(fakeFileName)).Return(fakeFileName, nil)

	savedFileName, err := useCase.Save(fakeFileData, fakeFileName)
	assert.NoError(t, err)
	assert.Equal(t, fakeFileName, savedFileName)
}

func TestFileUseCase_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockFileRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeFileName := "unique_filename"
	fakeFileData := []byte("file_data")

	mockRepo.EXPECT().Get(gomock.Eq(fakeFileName)).Return(fakeFileData, nil)

	fileData, err := useCase.Get(fakeFileName)
	assert.NoError(t, err)
	assert.Equal(t, fakeFileData, fileData)
}

func TestFileUseCase_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockFileRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeFileName := "unique_filename"

	mockRepo.EXPECT().Delete(gomock.Eq(fakeFileName)).Return(nil)

	err := useCase.Delete(fakeFileName)
	assert.NoError(t, err)
}
